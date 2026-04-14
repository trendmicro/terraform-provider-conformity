package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type applyProfileItem struct {
	ID           string
	ResourceName string
	ProfileID    string
	AccountIDs   []string
	Mode         string
	Notes        string
	Include      *applyProfileIncludeItem
}

type applyProfileIncludeItem struct {
	Exceptions *bool
}

func loadApplyProfilesFromState(path string) ([]applyProfileItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read state: %w", err)
	}

	var state terraformState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("parse state: %w", err)
	}

	var items []applyProfileItem
	for _, res := range state.Resources {
		if res.Type != "conformity_apply_profile" {
			continue
		}
		if res.Mode != "data" && res.Mode != "managed" {
			continue
		}
		for _, inst := range res.Instances {
			attrs := inst.Attributes
			if attrs == nil {
				continue
			}
			item := parseApplyProfileAttributes(attrs)
			if item.ProfileID == "" {
				continue
			}
			item.ResourceName = resourceNameFromState(res, inst, "apply_profile")
			items = append(items, item)
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].ResourceName < items[j].ResourceName
	})

	return items, nil
}

func parseApplyProfileAttributes(attrs map[string]interface{}) applyProfileItem {
	item := applyProfileItem{}
	item.ID = strings.TrimSpace(toStringValue(attrs["id"]))
	item.ProfileID = strings.TrimSpace(toStringValue(attrs["profile_id"]))
	item.AccountIDs = toStringSlice(attrs["account_ids"])
	item.Mode = strings.TrimSpace(toStringValue(attrs["mode"]))
	item.Notes = strings.TrimSpace(toStringValue(attrs["notes"]))
	item.Include = parseApplyProfileInclude(attrs["include"])

	sort.Strings(item.AccountIDs)

	return item
}

func parseApplyProfileInclude(value interface{}) *applyProfileIncludeItem {
	entry := firstSetEntry(value)
	if entry == nil {
		return nil
	}

	if raw, ok := entry["exceptions"].(bool); ok {
		return &applyProfileIncludeItem{Exceptions: &raw}
	}

	return nil
}

func appendApplyProfileHCL(lines []string, item applyProfileItem, resourceName string) []string {
	lines = append(lines, fmt.Sprintf("data \"visionone_crm_apply_profile\" \"%s\" {", resourceName))
	lines = append(lines, fmt.Sprintf("  profile_id = \"%s\"", escapeHCL(item.ProfileID)))
	if len(item.AccountIDs) > 0 {
		lines = append(lines, fmt.Sprintf("  account_ids = [%s]", formatQuotedList(item.AccountIDs)))
	}
	if item.Mode != "" {
		lines = append(lines, fmt.Sprintf("  mode = \"%s\"", escapeHCL(item.Mode)))
	}
	if item.Notes != "" {
		lines = append(lines, fmt.Sprintf("  notes = \"%s\"", escapeHCL(item.Notes)))
	}
	if item.Include != nil && item.Include.Exceptions != nil {
		lines = append(lines, "  include = {")
		lines = append(lines, fmt.Sprintf("    exceptions = %t", *item.Include.Exceptions))
		lines = append(lines, "  }")
	}
	lines = append(lines, "}")
	lines = append(lines, "")
	return lines
}

func appendApplyProfileMappingLines(mappingLines []string, item applyProfileItem, targetName string) []string {
	if mappingLines == nil {
		return mappingLines
	}

	sourceName := item.ResourceName
	if sourceName == "" {
		sourceName = targetName
	}
	sourceType := "conformity_apply_profile"
	targetType := "visionone_crm_apply_profile"
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

	mapField("profile_id", "profile_id")
	if len(item.AccountIDs) > 0 {
		mapField("account_ids", "account_ids")
	}
	if item.Mode != "" {
		mapField("mode", "mode")
	}
	if item.Notes != "" {
		mapField("notes", "notes")
	}
	if item.Include != nil && item.Include.Exceptions != nil {
		mapField("include.exceptions", "include.exceptions")
	}

	return mappingLines
}
