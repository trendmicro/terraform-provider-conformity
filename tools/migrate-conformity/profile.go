package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type profileItem struct {
	ID           string
	Name         string
	ResourceName string
	Description  string
	ScanRules    []scanRuleItem
}

type scanRuleItem struct {
	ID            string
	Provider      string
	Enabled       bool
	RiskLevel     string
	Exceptions    *ruleExceptionsItem
	ExtraSettings []extraSettingItem
}

type ruleExceptionsItem struct {
	FilterTags  []string
	Tags        []string
	ResourceIds []string
}

type extraSettingItem struct {
	Name     string
	Type     string
	Value    string
	ValueSet []string
	Values   []valueItem
}

type valueItem struct {
	Value          string
	Enabled        *bool
	CustomizedTags []string
	CustomRisk     string
	IgnoredSettings []string
}

func loadProfilesFromState(path string) ([]profileItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read state: %w", err)
	}

	var state terraformState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("parse state: %w", err)
	}

	var profiles []profileItem
	for _, res := range state.Resources {
		if res.Mode != "managed" || res.Type != "conformity_profile" {
			continue
		}
		for _, inst := range res.Instances {
			attrs := inst.Attributes
			if attrs == nil {
				continue
			}
			name, _ := attrs["name"].(string)
			if name == "" {
				continue
			}
			profiles = append(profiles, profileItem{
				ID:           toStringValue(attrs["id"]),
				Name:         name,
				ResourceName: resourceNameFromState(res, inst, "profile"),
				Description:  toStringValue(attrs["description"]),
				ScanRules:    parseScanRules(attrs["included"]),
			})
		}
	}

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].ResourceName < profiles[j].ResourceName
	})

	return profiles, nil
}

func parseScanRules(value interface{}) []scanRuleItem {
	items, ok := value.([]interface{})
	if !ok {
		return nil
	}

	var rules []scanRuleItem
	for _, raw := range items {
		entry, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		ruleID := toStringValue(entry["id"])
		if ruleID == "" {
			continue
		}
		provider := toStringValue(entry["provider"])
		riskLevel := toStringValue(entry["risk_level"])
		enabled := toBoolValue(entry["enabled"], true)

		rule := scanRuleItem{
			ID:        ruleID,
			Provider:  provider,
			Enabled:   enabled,
			RiskLevel: riskLevel,
		}

		exceptions := parseExceptions(entry["exceptions"])
		if exceptions != nil {
			rule.Exceptions = exceptions
		}

		rule.ExtraSettings = parseExtraSettings(entry["extra_settings"])
		rules = append(rules, rule)
	}

	sort.Slice(rules, func(i, j int) bool {
		return rules[i].ID < rules[j].ID
	})

	return rules
}

func parseExceptions(value interface{}) *ruleExceptionsItem {
	items, ok := value.([]interface{})
	if !ok || len(items) == 0 {
		return nil
	}

	filterTags := []string{}
	tags := []string{}
	resourceIds := []string{}
	for _, raw := range items {
		entry, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		filterTags = append(filterTags, toStringSlice(entry["filter_tags"])...)
		tags = append(tags, toStringSlice(entry["tags"])...)
		resourceIds = append(resourceIds, toStringSlice(entry["resources"])...)
	}

	filterTags = uniqueStrings(filterTags)
	tags = uniqueStrings(tags)
	resourceIds = uniqueStrings(resourceIds)
	if len(filterTags) == 0 && len(tags) == 0 && len(resourceIds) == 0 {
		return nil
	}

	return &ruleExceptionsItem{
		FilterTags:  filterTags,
		Tags:        tags,
		ResourceIds: resourceIds,
	}
}

func parseExtraSettings(value interface{}) []extraSettingItem {
	items, ok := value.([]interface{})
	if !ok {
		return nil
	}

	var settings []extraSettingItem
	for _, raw := range items {
		entry, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		name := toStringValue(entry["name"])
		typeValue := toStringValue(entry["type"])
		if name == "" || typeValue == "" {
			continue
		}

		valuesArray := toStringSlice(entry["values_array"])
		valuesList := parseValues(entry["values"])
		if typeValue == "choice-multiple-value" && hasCustomizedTags(valuesList) {
			typeValue = "choice-multiple-value-with-tags"
		}
		if typeValue == "choice-multiple-value" && hasCustomizedRiskLevel(valuesList) {
			typeValue = "choice-multiple-value-with-risk-level"
		}

		setting := extraSettingItem{
			Name:  name,
			Type:  typeValue,
			Value: toStringValue(entry["value"]),
		}
		if len(valuesArray) > 0 {
			setting.ValueSet = valuesArray
		} else if isValueSetType(typeValue) && len(valuesList) > 0 {
			for _, val := range valuesList {
				if val.Value != "" {
					setting.ValueSet = append(setting.ValueSet, val.Value)
				}
			}
		} else {
			setting.Values = valuesList
		}

		settings = append(settings, setting)
	}

	sort.Slice(settings, func(i, j int) bool {
		return settings[i].Name < settings[j].Name
	})

	return settings
}

func parseValues(value interface{}) []valueItem {
	items, ok := value.([]interface{})
	if !ok {
		return nil
	}

	values := make([]valueItem, 0, len(items))
	for _, raw := range items {
		entry, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		settings := entry["settings"]
		val := valueItem{
			Value:          toStringValue(entry["value"]),
			CustomizedTags: parseCustomizedTags(settings),
			CustomRisk:     parseCustomizedRiskLevel(settings),
			IgnoredSettings: parseIgnoredSettings(settings),
		}
		if enabled, ok := entry["enabled"].(bool); ok {
			val.Enabled = &enabled
		}
		values = append(values, val)
	}

	return values
}

func parseCustomizedTags(value interface{}) []string {
	items, ok := value.([]interface{})
	if !ok {
		return nil
	}

	var tags []string
	for _, raw := range items {
		entry, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		if toStringValue(entry["name"]) != "tags-override" {
			continue
		}
		tags = append(tags, toStringSlice(entry["values_array"])...)
		if values, ok := entry["values"].([]interface{}); ok {
			for _, rawValue := range values {
				valueEntry, ok := rawValue.(map[string]interface{})
				if !ok {
					continue
				}
				val := toStringValue(valueEntry["value"])
				if val != "" {
					tags = append(tags, val)
				}
			}
		}
	}

	return uniqueStrings(tags)
}

func parseCustomizedRiskLevel(value interface{}) string {
	items, ok := value.([]interface{})
	if !ok {
		return ""
	}

	for _, raw := range items {
		entry, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		name := toStringValue(entry["name"])
		if name != "selectedSeverity" && name != "selected_severity" && name != "customizedRiskLevel" && name != "customized_risk_level" {
			continue
		}
		values := toStringSlice(entry["values_array"])
		if len(values) == 0 {
			if rawValues, ok := entry["values"].([]interface{}); ok {
				for _, rawValue := range rawValues {
					valueEntry, ok := rawValue.(map[string]interface{})
					if !ok {
						continue
					}
					val := toStringValue(valueEntry["value"])
					if val != "" {
						values = append(values, val)
					}
				}
			}
		}
		if len(values) > 0 {
			return values[0]
		}
	}

	return ""
}

func parseIgnoredSettings(value interface{}) []string {
	items, ok := value.([]interface{})
	if !ok {
		return nil
	}

	ignored := []string{}
	for _, raw := range items {
		entry, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		name := toStringValue(entry["name"])
		if name == "" {
			continue
		}
		switch name {
		case "tags-override":
			continue
		case "selectedSeverity", "selected_severity", "customizedRiskLevel", "customized_risk_level":
			continue
		default:
			ignored = append(ignored, name)
		}
	}

	return uniqueStrings(ignored)
}

func hasCustomizedTags(values []valueItem) bool {
	for _, value := range values {
		if len(value.CustomizedTags) > 0 {
			return true
		}
	}
	return false
}

func hasCustomizedRiskLevel(values []valueItem) bool {
	for _, value := range values {
		if value.CustomRisk != "" {
			return true
		}
	}
	return false
}

func appendScanRuleHCL(lines *[]string, rule scanRuleItem) {
	*lines = append(*lines, "")
	*lines = append(*lines, "  scan_rule {")
	*lines = append(*lines, fmt.Sprintf("    id = \"%s\"", escapeHCL(rule.ID)))
	if rule.Provider != "" {
		*lines = append(*lines, fmt.Sprintf("    provider = \"%s\"", escapeHCL(rule.Provider)))
	}
	*lines = append(*lines, fmt.Sprintf("    enabled = %t", rule.Enabled))
	if rule.RiskLevel != "" {
		*lines = append(*lines, fmt.Sprintf("    risk_level = \"%s\"", escapeHCL(rule.RiskLevel)))
	}

	if rule.Exceptions != nil {
		if len(rule.Exceptions.Tags) > 0 {
			*lines = append(*lines, fmt.Sprintf("    # tags: [%s] // This field was not mapped because it has no effect. Did you mean to use filter_tags here?", formatQuotedList(rule.Exceptions.Tags)))
		}
		if len(rule.Exceptions.FilterTags) > 0 || len(rule.Exceptions.ResourceIds) > 0 {
			*lines = append(*lines, "    exceptions {")
			if len(rule.Exceptions.FilterTags) > 0 {
				*lines = append(*lines, fmt.Sprintf("      filter_tags = [%s]", formatQuotedList(rule.Exceptions.FilterTags)))
			}
			if len(rule.Exceptions.ResourceIds) > 0 {
				*lines = append(*lines, fmt.Sprintf("      resource_ids = [%s]", formatQuotedList(rule.Exceptions.ResourceIds)))
			}
			*lines = append(*lines, "    }")
		}
	}

	for _, setting := range rule.ExtraSettings {
		*lines = append(*lines, "    extra_settings {")
		*lines = append(*lines, fmt.Sprintf("      name = \"%s\"", escapeHCL(setting.Name)))
		*lines = append(*lines, fmt.Sprintf("      type = \"%s\"", escapeHCL(setting.Type)))
		if setting.Value != "" {
			*lines = append(*lines, fmt.Sprintf("      value = %s", formatValue(setting.Value, setting.Type)))
		}
		if len(setting.ValueSet) > 0 {
			*lines = append(*lines, fmt.Sprintf("      value_set = [%s]", formatValueList(setting.ValueSet, setting.Type)))
		}
		for _, val := range setting.Values {
			*lines = append(*lines, "      values {")
			if val.Value != "" {
				*lines = append(*lines, fmt.Sprintf("        value = %s", formatValue(val.Value, setting.Type)))
			}
			if val.Enabled != nil {
				*lines = append(*lines, fmt.Sprintf("        enabled = %t", *val.Enabled))
			}
			if len(val.CustomizedTags) > 0 {
				*lines = append(*lines, fmt.Sprintf("        customized_tags = [%s]", formatQuotedList(val.CustomizedTags)))
			}
			if val.CustomRisk != "" {
				*lines = append(*lines, fmt.Sprintf("        customized_risk_level = \"%s\"", escapeHCL(val.CustomRisk)))
			}
			if len(val.IgnoredSettings) > 0 {
				*lines = append(*lines, fmt.Sprintf("        # Unmapped settings: %s", strings.Join(val.IgnoredSettings, ", ")))
			}
			*lines = append(*lines, "      }")
		}
		*lines = append(*lines, "    }")
	}

	*lines = append(*lines, "  }")
}

func formatValue(raw string, settingType string) string {
	if !isNumericType(settingType) {
		return fmt.Sprintf("\"%s\"", escapeHCL(raw))
	}
	if _, err := strconv.ParseInt(raw, 10, 64); err == nil {
		return raw
	}
	if _, err := strconv.ParseFloat(raw, 64); err == nil {
		return raw
	}
	return fmt.Sprintf("\"%s\"", escapeHCL(raw))
}

func formatValueList(values []string, settingType string) string {
	items := make([]string, 0, len(values))
	for _, value := range values {
		items = append(items, formatValue(value, settingType))
	}
	return strings.Join(items, ", ")
}

func formatQuotedList(values []string) string {
	items := make([]string, 0, len(values))
	for _, value := range values {
		items = append(items, fmt.Sprintf("\"%s\"", escapeHCL(value)))
	}
	return strings.Join(items, ", ")
}
