package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func toStringSlice(value interface{}) []string {
	switch v := value.(type) {
	case []interface{}:
		out := make([]string, 0, len(v))
		for _, raw := range v {
			if s, ok := raw.(string); ok {
				out = append(out, s)
			}
		}
		return out
	case []string:
		return v
	default:
		return nil
	}
}

func uniqueStrings(values []string) []string {
	seen := map[string]struct{}{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func isValueSetType(settingType string) bool {
	switch settingType {
	case "multiple-string-values", "multiple-ip-values", "multiple-aws-account-values", "multiple-number-values", "regions", "ignored-regions", "tags", "countries":
		return true
	default:
		return false
	}
}

func isNumericType(settingType string) bool {
	switch settingType {
	case "ttl", "single-number-value", "multiple-number-values":
		return true
	default:
		return false
	}
}

func toBoolValue(value interface{}, fallback bool) bool {
	if raw, ok := value.(bool); ok {
		return raw
	}
	return fallback
}

func toIntValue(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	case float32:
		return int(v)
	default:
		return 0
	}
}

func toStringValue(value interface{}) string {
	if value == nil {
		return ""
	}
	if s, ok := value.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", value)
}

func toTerraformName(name string, counter map[string]int, fallback string) string {
	normalized := normalizeName(name)
	if normalized == "" {
		normalized = fallback
	}

	counter[normalized]++
	if counter[normalized] == 1 {
		return normalized
	}

	return fmt.Sprintf("%s_%d", normalized, counter[normalized])
}

var invalidIdentifier = regexp.MustCompile(`[^a-z0-9_]+`)
var invalidIdentifierPreserve = regexp.MustCompile(`[^A-Za-z0-9_]+`)

func normalizeName(input string) string {
	value := strings.ToLower(strings.TrimSpace(input))
	value = strings.ReplaceAll(value, "-", "_")
	value = invalidIdentifier.ReplaceAllString(value, "_")
	value = strings.Trim(value, "_")
	if value == "" {
		return value
	}
	if value[0] >= '0' && value[0] <= '9' {
		return "group_" + value
	}
	return value
}

func sanitizeIdentifier(input, fallback string) string {
	value := strings.TrimSpace(input)
	if value == "" {
		value = fallback
	}
	value = invalidIdentifierPreserve.ReplaceAllString(value, "_")
	value = strings.Trim(value, "_")
	if value == "" {
		value = fallback
	}
	if len(value) > 0 && value[0] >= '0' && value[0] <= '9' {
		value = "res_" + value
	}
	return value
}

func resourceNameFromState(res stateResource, inst stateInstance, fallback string) string {
	name := sanitizeIdentifier(res.Name, fallback)
	if inst.IndexKey == nil {
		return name
	}
	suffix := sanitizeIdentifier(fmt.Sprintf("%v", inst.IndexKey), "index")
	return fmt.Sprintf("%s_%s", name, suffix)
}

func uniqueResourceName(name string, counter map[string]int) string {
	if name == "" {
		return ""
	}
	counter[name]++
	if counter[name] == 1 {
		return name
	}
	return fmt.Sprintf("%s_%d", name, counter[name])
}

func formatTags(tags []string) string {
	escaped := make([]string, 0, len(tags))
	for _, tag := range tags {
		escaped = append(escaped, fmt.Sprintf("\"%s\"", escapeHCL(tag)))
	}
	return strings.Join(escaped, ", ")
}

func escapeHCL(value string) string {
	value = strings.ReplaceAll(value, "\\", "\\\\")
	value = strings.ReplaceAll(value, "\"", "\\\"")
	return value
}

func escapeHCLMultiline(value string) string {
	value = strings.ReplaceAll(value, "\r\n", "\n")
	value = strings.ReplaceAll(value, "\r", "\n")
	value = escapeHCL(value)
	value = strings.ReplaceAll(value, "\n", "\\n")
	return value
}

func formatImportLine(resourceType, resourceName, resourceID string, dryRun bool) string {
	if resourceID == "" {
		resourceID = "<resource_id>"
	}
	line := fmt.Sprintf("terraform import %s.%s %s", resourceType, resourceName, resourceID)

	return line
}

func formatMappingLine(sourceType, sourceName, targetType, targetName string) string {
	return fmt.Sprintf("- %s.%s -> %s.%s", sourceType, sourceName, targetType, targetName)
}

func formatAttributeMappingLine(sourceAttribute, targetAttribute string) string {
	return fmt.Sprintf("  - %s -> %s", sourceAttribute, targetAttribute)
}

func shouldForceDryRun(outputDir string) (bool, error) {
	path := filepath.Join(outputDir, ".terraform")
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("check .terraform: %w", err)
}

func writeMainTF(outputDir string, lines []string) (string, error) {
	if outputDir == "" {
		return "", fmt.Errorf("output directory is empty")
	}
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return "", fmt.Errorf("create output directory: %w", err)
	}
	path := filepath.Join(outputDir, "main.tf")
	content := strings.Join(lines, "\n") + "\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return "", fmt.Errorf("write main.tf: %w", err)
	}
	return path, nil
}

func runTerraformImports(outputDir string, importLines []string) error {
	if len(importLines) == 0 {
		return nil
	}
	if err := runTerraformInit(outputDir); err != nil {
		return err
	}
	for _, line := range importLines {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "terraform import ") {
			continue
		}
		parts := strings.Fields(trimmed)
		if len(parts) < 4 {
			return fmt.Errorf("invalid import command: %s", trimmed)
		}
		cmd := exec.Command(parts[0], parts[1], parts[2], parts[3])
		cmd.Dir = outputDir
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("import failed (%s): %s", trimmed, strings.TrimSpace(string(output)))
		}
	}
	return nil
}

func runTerraformInit(outputDir string) error {
	cmd := exec.Command("terraform", "init", "-input=false")
	cmd.Dir = outputDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("terraform init failed: %s", strings.TrimSpace(string(output)))
	}
	return nil
}
