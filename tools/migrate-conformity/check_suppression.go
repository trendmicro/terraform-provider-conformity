package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type checkSuppressionItem struct {
	ID           string
	ResourceName string
	AccountID    string
	Service      string
	RuleID       string
	Region       string
	ResourceID   string
	Note         string
}

func loadCheckSuppressionsFromState(path string) ([]checkSuppressionItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read state: %w", err)
	}

	var state terraformState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("parse state: %w", err)
	}

	var suppressions []checkSuppressionItem
	for _, res := range state.Resources {
		if res.Mode != "managed" || res.Type != "conformity_check_suppression" {
			continue
		}
		for _, inst := range res.Instances {
			attrs := inst.Attributes
			if attrs == nil {
				continue
			}
			item := parseCheckSuppressionAttributes(attrs)
			if item.RuleID == "" || item.AccountID == "" || item.Region == "" || item.ResourceID == "" {
				continue
			}
			item.ResourceName = resourceNameFromState(res, inst, "check_suppression")
			suppressions = append(suppressions, item)
		}
	}

	sort.Slice(suppressions, func(i, j int) bool {
		return suppressions[i].ResourceName < suppressions[j].ResourceName
	})

	return suppressions, nil
}

func parseCheckSuppressionAttributes(attrs map[string]interface{}) checkSuppressionItem {
	item := checkSuppressionItem{}
	item.ID = strings.TrimSpace(toStringValue(attrs["id"]))
	item.AccountID = strings.TrimSpace(toStringValue(attrs["account_id"]))
	item.RuleID = strings.TrimSpace(toStringValue(attrs["rule_id"]))
	item.Service = deriveServiceFromRuleID(item.RuleID)
	item.Region = strings.TrimSpace(toStringValue(attrs["region"]))
	item.ResourceID = strings.TrimSpace(toStringValue(attrs["resource_id"]))
	item.Note = strings.TrimSpace(toStringValue(attrs["note"]))
	return item
}

func deriveServiceFromRuleID(ruleID string) string {
	parts := strings.SplitN(ruleID, "-", 2)
	if len(parts) == 0 {
		return ""
	}
	return strings.TrimSpace(parts[0])
}

func appendCheckSuppressionHCL(lines []string, item checkSuppressionItem, resourceName string) []string {
	lines = append(lines, fmt.Sprintf("resource \"visionone_crm_check_suppression\" \"%s\" {", resourceName))
	lines = append(lines, fmt.Sprintf("  account_id = \"%s\"", escapeHCL(item.AccountID)))
	lines = append(lines, fmt.Sprintf("  service = \"%s\"", escapeHCL(item.Service)))
	lines = append(lines, fmt.Sprintf("  rule_id = \"%s\"", escapeHCL(item.RuleID)))
	lines = append(lines, fmt.Sprintf("  region = \"%s\"", escapeHCL(item.Region)))
	lines = append(lines, fmt.Sprintf("  resource_id = \"%s\"", escapeHCL(item.ResourceID)))
	lines = append(lines, fmt.Sprintf("  note = \"%s\"", escapeHCL(item.Note)))
	lines = append(lines, "}")
	lines = append(lines, "")
	return lines
}

func appendCheckSuppressionMappingLines(mappingLines []string, item checkSuppressionItem, targetName string) []string {
	if mappingLines == nil {
		return mappingLines
	}

	sourceName := item.ResourceName
	if sourceName == "" {
		sourceName = targetName
	}
	sourceType := "conformity_check_suppression"
	targetType := "visionone_crm_check_suppression"
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

	mapField("account_id", "account_id")
	mapField("rule_id", "rule_id")
	mapField("rule_id (prefix)", "service")
	mapField("region", "region")
	mapField("resource_id", "resource_id")
	if item.Note != "" {
		mapField("note", "note")
	}

	return mappingLines
}
