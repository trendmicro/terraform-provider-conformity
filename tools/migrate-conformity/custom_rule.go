package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type customRuleItem struct {
	ID               string
	Name             string
	ResourceName     string
	Description      string
	CloudProvider    string
	RemediationNotes string
	Service          string
	ResourceType     string
	Severity         string
	Enabled          bool
	Categories       []string
	Attributes       []customRuleAttributeItem
	Rules            []customRuleEventRuleItem
}

type customRuleAttributeItem struct {
	Name     string
	Path     string
	Required bool
}

type customRuleEventRuleItem struct {
	Description string
	Operator    string
	Conditions  []customRuleConditionItem
}

type customRuleConditionItem struct {
	Operator string
	Fact     string
	Path     string
	Value    interface{}
}

func loadCustomRulesFromState(path string) ([]customRuleItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read state: %w", err)
	}

	var state terraformState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("parse state: %w", err)
	}

	var rules []customRuleItem
	for _, res := range state.Resources {
		if res.Mode != "managed" || res.Type != "conformity_custom_rule" {
			continue
		}
		for _, inst := range res.Instances {
			attrs := inst.Attributes
			if attrs == nil {
				continue
			}
			item := parseCustomRuleAttributes(attrs)
			if item.Name == "" {
				continue
			}
			item.ResourceName = resourceNameFromState(res, inst, "custom_rule")
			rules = append(rules, item)
		}
	}

	sort.Slice(rules, func(i, j int) bool {
		return rules[i].ResourceName < rules[j].ResourceName
	})

	return rules, nil
}

func parseCustomRuleAttributes(attrs map[string]interface{}) customRuleItem {
	item := customRuleItem{}
	item.ID = strings.TrimSpace(toStringValue(attrs["id"]))
	item.Name = strings.TrimSpace(toStringValue(attrs["name"]))
	item.Description = strings.TrimSpace(toStringValue(attrs["description"]))
	item.CloudProvider = strings.TrimSpace(toStringValue(attrs["cloud_provider"]))
	item.RemediationNotes = strings.TrimSpace(toStringValue(attrs["remediation_notes"]))
	item.Service = strings.TrimSpace(toStringValue(attrs["service"]))
	item.ResourceType = strings.TrimSpace(toStringValue(attrs["resource_type"]))
	item.Severity = strings.TrimSpace(toStringValue(attrs["severity"]))
	item.Enabled = toBoolValue(attrs["enabled"], true)
	item.Categories = toStringSlice(attrs["categories"])
	if len(item.Categories) > 1 {
		sortCustomRuleCategories(item.Categories)
	}
	item.Attributes = parseCustomRuleAttributesList(attrs["attributes"])
	item.Rules = parseCustomRuleRules(attrs["rules"])
	return item
}

func sortCustomRuleCategories(categories []string) {
	order := map[string]int{
		"security":               0,
		"cost-optimisation":      1,
		"reliability":            2,
		"performance-efficiency": 3,
		"operational-excellence": 4,
		"sustainability":         5,
	}

	sort.SliceStable(categories, func(i, j int) bool {
		iOrder, iOk := order[categories[i]]
		jOrder, jOk := order[categories[j]]
		if iOk && jOk {
			if iOrder == jOrder {
				return categories[i] < categories[j]
			}
			return iOrder < jOrder
		}
		if iOk {
			return true
		}
		if jOk {
			return false
		}
		return categories[i] < categories[j]
	})
}

func parseCustomRuleAttributesList(value interface{}) []customRuleAttributeItem {
	entries := toMapSlice(value)
	if len(entries) == 0 {
		return nil
	}

	attrs := make([]customRuleAttributeItem, 0, len(entries))
	for _, entry := range entries {
		name := strings.TrimSpace(toStringValue(entry["name"]))
		path := strings.TrimSpace(toStringValue(entry["path"]))
		required := toBoolValue(entry["required"], false)
		if name == "" || path == "" {
			continue
		}
		attrs = append(attrs, customRuleAttributeItem{
			Name:     name,
			Path:     path,
			Required: required,
		})
	}

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Name < attrs[j].Name
	})

	return attrs
}

func parseCustomRuleRules(value interface{}) []customRuleEventRuleItem {
	entries := toMapSlice(value)
	if len(entries) == 0 {
		return nil
	}

	rules := make([]customRuleEventRuleItem, 0, len(entries))
	for _, entry := range entries {
		eventType := strings.TrimSpace(toStringValue(entry["event_type"]))
		description := eventType
		if description == "" {
			description = "Migrated custom rule"
		}
		operation := strings.TrimSpace(toStringValue(entry["operation"]))

		rules = append(rules, customRuleEventRuleItem{
			Description: description,
			Operator:    operation,
			Conditions:  parseCustomRuleConditions(entry["conditions"]),
		})
	}

	return rules
}

func parseCustomRuleConditions(value interface{}) []customRuleConditionItem {
	entries := toMapSlice(value)
	if len(entries) == 0 {
		return nil
	}

	conditions := make([]customRuleConditionItem, 0, len(entries))
	for _, entry := range entries {
		fact := strings.TrimSpace(toStringValue(entry["fact"]))
		operator := strings.TrimSpace(toStringValue(entry["operator"]))
		path := strings.TrimSpace(toStringValue(entry["path"]))
		value := normalizeCustomRuleConditionValue(entry["value"])
		if fact == "" || operator == "" {
			continue
		}
		conditions = append(conditions, customRuleConditionItem{
			Operator: operator,
			Fact:     fact,
			Path:     path,
			Value:    value,
		})
	}

	return conditions
}

func normalizeCustomRuleConditionValue(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	if str, ok := value.(string); ok {
		return strings.TrimSpace(str)
	}
	return value
}

func appendCustomRuleHCL(lines []string, item customRuleItem, resourceName string) []string {
	lines = append(lines, fmt.Sprintf("resource \"visionone_crm_custom_rule\" \"%s\" {", resourceName))
	lines = append(lines, fmt.Sprintf("  name = \"%s\"", escapeHCL(item.Name)))
	if item.Description != "" {
		lines = append(lines, fmt.Sprintf("  description = \"%s\"", escapeHCL(item.Description)))
	}
	if item.Severity != "" {
		lines = append(lines, fmt.Sprintf("  risk_level = \"%s\"", escapeHCL(item.Severity)))
	}
	if item.CloudProvider != "" {
		lines = append(lines, fmt.Sprintf("  cloud_provider = \"%s\"", escapeHCL(item.CloudProvider)))
	}
	if item.Service != "" {
		lines = append(lines, fmt.Sprintf("  service = \"%s\"", escapeHCL(item.Service)))
	}
	if item.ResourceType != "" {
		lines = append(lines, fmt.Sprintf("  resource_type = \"%s\"", escapeHCL(item.ResourceType)))
	}
	lines = append(lines, fmt.Sprintf("  enabled = %t", item.Enabled))
	if len(item.Categories) > 0 {
		lines = append(lines, fmt.Sprintf("  categories = [%s]", formatQuotedList(item.Categories)))
	}
	if item.RemediationNotes != "" {
		lines = append(lines, fmt.Sprintf("  remediation_note = \"%s\"", escapeHCLMultiline(item.RemediationNotes)))
	}

	for _, attr := range item.Attributes {
		lines = append(lines, "  attribute {")
		lines = append(lines, fmt.Sprintf("    name = \"%s\"", escapeHCL(attr.Name)))
		lines = append(lines, fmt.Sprintf("    path = \"%s\"", escapeHCL(attr.Path)))
		lines = append(lines, fmt.Sprintf("    required = %t", attr.Required))
		lines = append(lines, "  }")
	}

	for _, rule := range item.Rules {
		lines = append(lines, "  event_rule {")
		lines = append(lines, fmt.Sprintf("    description = \"%s\"", escapeHCL(rule.Description)))
		if rule.Operator != "" || len(rule.Conditions) > 0 {
			lines = append(lines, "    conditions {")
			if rule.Operator != "" {
				lines = append(lines, fmt.Sprintf("      operator = \"%s\"", escapeHCL(rule.Operator)))
			}
			for _, condition := range rule.Conditions {
				lines = append(lines, "      condition {")
				lines = append(lines, fmt.Sprintf("        operator = \"%s\"", escapeHCL(condition.Operator)))
				lines = append(lines, fmt.Sprintf("        fact = \"%s\"", escapeHCL(condition.Fact)))
				if condition.Path != "" {
					lines = append(lines, fmt.Sprintf("        path = \"%s\"", escapeHCL(condition.Path)))
				}
				lines = append(lines, fmt.Sprintf("        value = %s", formatCustomRuleValue(condition.Value)))
				lines = append(lines, "      }")
			}
			lines = append(lines, "    }")
		}
		lines = append(lines, "  }")
	}

	lines = append(lines, "}")
	lines = append(lines, "")
	return lines
}

func appendCustomRuleMappingLines(mappingLines []string, item customRuleItem, targetName string) []string {
	if mappingLines == nil {
		return mappingLines
	}

	sourceName := item.ResourceName
	if sourceName == "" {
		sourceName = targetName
	}
	sourceType := "conformity_custom_rule"
	targetType := "visionone_crm_custom_rule"
	mappingLines = append(mappingLines, formatMappingLine(sourceType, sourceName, targetType, targetName))

	seen := map[string]struct{}{}
	appendUnique := func(line string) {
		if _, ok := seen[line]; ok {
			return
		}
		seen[line] = struct{}{}
		mappingLines = append(mappingLines, line)
	}
	mapField := func(sourceAttribute, targetAttribute string) {
		appendUnique(formatAttributeMappingLine(sourceAttribute, targetAttribute))
	}

	mapField("name", "name")
	if item.Description != "" {
		mapField("description", "description")
	}
	if item.CloudProvider != "" {
		mapField("cloud_provider", "cloud_provider")
	}
	if item.Service != "" {
		mapField("service", "service")
	}
	if item.ResourceType != "" {
		mapField("resource_type", "resource_type")
	}
	mapField("enabled", "enabled")
	if item.Severity != "" {
		mapField("severity", "risk_level")
	}
	if len(item.Categories) > 0 {
		mapField("categories", "categories")
	}
	if item.RemediationNotes != "" {
		mapField("remediation_notes", "remediation_note")
	}
	if len(item.Attributes) > 0 {
		mapField("attributes", "attribute")
		mapField("attributes.name", "attribute.name")
		mapField("attributes.path", "attribute.path")
		mapField("attributes.required", "attribute.required")
	}
	if len(item.Rules) > 0 {
		mapField("rules", "event_rule")
		mapField("rules.event_type", "event_rule.description")
		mapField("rules.operation", "event_rule.conditions.operator")
		mapField("rules.conditions", "event_rule.conditions.condition")
		mapField("rules.conditions.operator", "event_rule.conditions.condition.operator")
		mapField("rules.conditions.fact", "event_rule.conditions.condition.fact")
		mapField("rules.conditions.path", "event_rule.conditions.condition.path")
		mapField("rules.conditions.value", "event_rule.conditions.condition.value")
	}

	return mappingLines
}

func formatCustomRuleValue(raw interface{}) string {
	if raw == nil {
		return "\"\""
	}

	switch value := raw.(type) {
	case bool:
		if value {
			return "true"
		}
		return "false"
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%v", value)
	case float32, float64:
		return fmt.Sprintf("%v", value)
	case string:
		value = strings.TrimSpace(value)
		if value == "" {
			return "\"\""
		}

		lower := strings.ToLower(value)
		if lower == "null" {
			return "\"\""
		}
		if lower == "true" || lower == "false" {
			return lower
		}
		if _, err := strconv.ParseInt(value, 10, 64); err == nil {
			return value
		}
		if _, err := strconv.ParseFloat(value, 64); err == nil {
			return value
		}

		if strings.HasPrefix(value, "{") || strings.HasPrefix(value, "[") {
			var decoded interface{}
			if err := json.Unmarshal([]byte(value), &decoded); err == nil {
				return fmt.Sprintf("jsonencode(%s)", value)
			}
		}

		return fmt.Sprintf("\"%s\"", escapeHCL(value))
	default:
		return fmt.Sprintf("\"%s\"", escapeHCL(strings.TrimSpace(toStringValue(value))))
	}
}

func toMapSlice(value interface{}) []map[string]interface{} {
	list, ok := value.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}

	items := make([]map[string]interface{}, 0, len(list))
	for _, raw := range list {
		entry, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		items = append(items, entry)
	}
	return items
}
