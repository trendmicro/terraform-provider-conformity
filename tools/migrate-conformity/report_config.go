package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type reportConfigItem struct {
	ID                          string
	ResourceName                string
	AccountID                   string
	GroupID                     string
	ReportTitle                 string
	ReportType                  string
	IncludeAccountNames         *bool
	IncludeChecks               bool
	EmailRecipients             []string
	EmailRecipientsSet          bool
	ReportFormatsInEmail        []string
	ReportFormatsInEmailSet     bool
	Schedule                    *reportScheduleItem
	ChecksFilter                *reportFilterItem
	AppliedComplianceStandardID string
	ControlsType                string
}

type reportScheduleItem struct {
	Enabled   *bool
	Frequency string
	Timezone  string
}

type reportFilterItem struct {
	Categories            []string
	ComplianceStandardIds []string
	Tags                  []string
	Description           string
	NewerThanDays         int
	OlderThanDays         int
	Providers             []string
	Regions               []string
	ResourceID            string
	ResourceSearchMode    string
	ResourceTypes         []string
	RiskLevels            []string
	RuleIds               []string
	Services              []string
	Statuses              []string
	Suppressed            *bool
}

func loadReportConfigsFromState(path string) ([]reportConfigItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read state: %w", err)
	}

	var state terraformState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("parse state: %w", err)
	}

	var configs []reportConfigItem
	for _, res := range state.Resources {
		if res.Mode != "managed" || res.Type != "conformity_report_config" {
			continue
		}
		for _, inst := range res.Instances {
			attrs := inst.Attributes
			if attrs == nil {
				continue
			}
			item := parseReportConfigAttributes(attrs)
			if item.ReportTitle == "" {
				continue
			}
			item.ResourceName = resourceNameFromState(res, inst, "report_config")
			configs = append(configs, item)
		}
	}

	sort.Slice(configs, func(i, j int) bool {
		return configs[i].ResourceName < configs[j].ResourceName
	})

	return configs, nil
}

func parseReportConfigAttributes(attrs map[string]interface{}) reportConfigItem {
	item := reportConfigItem{}
	item.ID = strings.TrimSpace(toStringValue(attrs["id"]))

	item.AccountID = strings.TrimSpace(toStringValue(attrs["account_id"]))
	item.GroupID = strings.TrimSpace(toStringValue(attrs["group_id"]))

	config := firstListEntry(attrs["configuration"])
	if config == nil {
		return item
	}
	item.ReportTitle = strings.TrimSpace(toStringValue(config["title"]))
	item.ReportType = strings.TrimSpace(toStringValue(config["generate_report_type"]))
	if item.ReportType == "" {
		item.ReportType = "GENERIC"
	}

	if item.AccountID == "" {
		if includeAccountRaw, ok := config["include_account_names"]; ok {
			includeAccountNames := toBoolValue(includeAccountRaw, true)
			item.IncludeAccountNames = &includeAccountNames
		} else {
			includeAccountNames := true
			item.IncludeAccountNames = &includeAccountNames
		}
	}

	item.IncludeChecks = toBoolValue(config["include_checks"], false)

	sendEmail := toBoolValue(config["send_email"], false)
	emails := toStringSlice(config["emails"])
	if len(emails) > 0 {
		item.EmailRecipients = emails
		item.EmailRecipientsSet = true
	} else if sendEmail {
		item.EmailRecipientsSet = true
	}
	if _, ok := config["should_email_include_pdf"]; ok {
		item.ReportFormatsInEmailSet = true
		item.ReportFormatsInEmail = deriveReportFormats(
			toBoolValue(config["should_email_include_pdf"], false),
			toBoolValue(config["should_email_include_csv"], false),
		)
	} else if _, ok := config["should_email_include_csv"]; ok {
		item.ReportFormatsInEmailSet = true
		item.ReportFormatsInEmail = deriveReportFormats(
			toBoolValue(config["should_email_include_pdf"], false),
			toBoolValue(config["should_email_include_csv"], false),
		)
	}

	frequency := strings.TrimSpace(toStringValue(config["frequency"]))
	timezone := strings.TrimSpace(toStringValue(config["tz"]))
	if scheduledRaw, ok := config["scheduled"]; ok {
		enabled := toBoolValue(scheduledRaw, false)
		item.Schedule = &reportScheduleItem{
			Enabled:   &enabled,
			Frequency: frequency,
			Timezone:  timezone,
		}
	} else if frequency != "" || timezone != "" {
		item.Schedule = &reportScheduleItem{
			Frequency: frequency,
			Timezone:  timezone,
		}
	}

	filterEntry := firstListEntry(attrs["filter"])
	filter := parseReportConfigFilter(filterEntry, item.ReportType)
	if filter != nil && hasReportFilterValues(filter) {
		item.ChecksFilter = filter
	}

	if item.ReportType == "COMPLIANCE-STANDARD" {
		if filterEntry != nil {
			item.AppliedComplianceStandardID = strings.TrimSpace(toStringValue(filterEntry["report_compliance_standard_id"]))
		}
		if item.AppliedComplianceStandardID == "" && filter != nil && len(filter.ComplianceStandardIds) == 1 {
			item.AppliedComplianceStandardID = filter.ComplianceStandardIds[0]
		}

		withChecks := false
		withoutChecks := false
		if filterEntry != nil {
			withChecks = toBoolValue(filterEntry["with_checks"], false)
			withoutChecks = toBoolValue(filterEntry["without_checks"], false)
		}
		item.ControlsType = deriveControlsType(withChecks, withoutChecks)
	}

	return item
}

func parseReportConfigFilter(entry map[string]interface{}, reportType string) *reportFilterItem {
	if entry == nil {
		return nil
	}

	filterTagsValue := entry["filter_tags"]
	tagsValue := entry["tags"]
	filterTags := toStringSlice(filterTagsValue)
	tags := toStringSlice(tagsValue)
	mergedTags := append([]string{}, tags...)
	mergedTags = append(mergedTags, filterTags...)
	mergedTags = uniqueStrings(mergedTags)

	filter := &reportFilterItem{
		Categories:            toStringSlice(entry["categories"]),
		ComplianceStandardIds: toStringSlice(entry["compliance_standards"]),
		Tags:                  mergedTags,
		NewerThanDays:         toIntValue(entry["newer_than_days"]),
		OlderThanDays:         toIntValue(entry["older_than_days"]),
		Providers:             toStringSlice(entry["providers"]),
		Regions:               toStringSlice(entry["regions"]),
		ResourceID:            strings.TrimSpace(toStringValue(entry["resource"])),
		ResourceSearchMode:    strings.TrimSpace(toStringValue(entry["resource_search_mode"])),
		ResourceTypes:         toStringSlice(entry["resource_types"]),
		RiskLevels:            toStringSlice(entry["risk_levels"]),
		RuleIds:               toStringSlice(entry["rule_ids"]),
		Services:              toStringSlice(entry["services"]),
		Statuses:              toStringSlice(entry["statuses"]),
	}

	mode := strings.TrimSpace(toStringValue(entry["suppressed_filter_mode"]))
	if suppressedRaw, ok := entry["suppressed"]; ok {
		suppressed := toBoolValue(suppressedRaw, true)
		switch mode {
		case "v2":
			filter.Suppressed = &suppressed
		case "v1":
			if !suppressed {
				filter.Suppressed = &suppressed
			}
		default:
			if !suppressed {
				filter.Suppressed = &suppressed
			}
		}
	}

	if reportType == "COMPLIANCE-STANDARD" {
		filter.ComplianceStandardIds = nil
	}

	return filter
}

func hasReportFilterValues(filter *reportFilterItem) bool {
	if filter == nil {
		return false
	}
	return len(filter.Categories) > 0 ||
		len(filter.ComplianceStandardIds) > 0 ||
		len(filter.Tags) > 0 ||
		filter.Description != "" ||
		filter.NewerThanDays > 0 ||
		filter.OlderThanDays > 0 ||
		len(filter.Providers) > 0 ||
		len(filter.Regions) > 0 ||
		filter.ResourceID != "" ||
		filter.ResourceSearchMode != "" ||
		len(filter.ResourceTypes) > 0 ||
		len(filter.RiskLevels) > 0 ||
		len(filter.RuleIds) > 0 ||
		len(filter.Services) > 0 ||
		len(filter.Statuses) > 0 ||
		filter.Suppressed != nil
}

func deriveReportFormats(includePDF, includeCSV bool) []string {
	switch {
	case includePDF && includeCSV:
		return []string{"all"}
	case includePDF:
		return []string{"PDF"}
	case includeCSV:
		return []string{"CSV"}
	default:
		return nil
	}
}

func deriveControlsType(withChecks, withoutChecks bool) string {
	switch {
	case withChecks && !withoutChecks:
		return "withChecksOnly"
	case withoutChecks && !withChecks:
		return "noChecksOnly"
	case withChecks && withoutChecks:
		return "all"
	}
	return "all"
}

func appendReportConfigHCL(lines *[]string, item reportConfigItem, resourceName string) {
	*lines = append(*lines, fmt.Sprintf("resource \"visionone_crm_report_config\" \"%s\" {", resourceName))
	if item.AccountID != "" {
		*lines = append(*lines, fmt.Sprintf("  account_id = \"%s\"", escapeHCL(item.AccountID)))
	}
	if item.GroupID != "" {
		*lines = append(*lines, fmt.Sprintf("  group_id = \"%s\"", escapeHCL(item.GroupID)))
	}
	*lines = append(*lines, fmt.Sprintf("  report_title = \"%s\"", escapeHCL(item.ReportTitle)))
	if item.ReportType != "" {
		*lines = append(*lines, fmt.Sprintf("  report_type = \"%s\"", escapeHCL(item.ReportType)))
	}
	if item.IncludeAccountNames != nil {
		*lines = append(*lines, fmt.Sprintf("  include_account_names = %t", *item.IncludeAccountNames))
	}
	if item.IncludeChecks {
		*lines = append(*lines, fmt.Sprintf("  include_checks = %t", item.IncludeChecks))
	}
	if item.EmailRecipientsSet {
		if len(item.EmailRecipients) == 0 {
			*lines = append(*lines, "  email_recipients = []")
		} else {
			*lines = append(*lines, fmt.Sprintf("  email_recipients = [%s]", formatQuotedList(item.EmailRecipients)))
		}
	}
	if item.ReportFormatsInEmailSet && len(item.ReportFormatsInEmail) > 0 {
		*lines = append(*lines, fmt.Sprintf("  report_formats_in_email = [%s]", formatQuotedList(item.ReportFormatsInEmail)))
	}
	if item.ReportType == "COMPLIANCE-STANDARD" {
		if item.AppliedComplianceStandardID != "" {
			*lines = append(*lines, fmt.Sprintf("  applied_compliance_standard_id = \"%s\"", escapeHCL(item.AppliedComplianceStandardID)))
		} else {
			*lines = append(*lines, "  # TODO: set applied_compliance_standard_id for compliance report")
		}
		if item.ControlsType != "" {
			*lines = append(*lines, fmt.Sprintf("  controls_type = \"%s\"", escapeHCL(item.ControlsType)))
		}
	}

	if item.Schedule != nil {
		*lines = append(*lines, "  schedule {")
		if item.Schedule.Enabled != nil {
			*lines = append(*lines, fmt.Sprintf("    enabled = %t", *item.Schedule.Enabled))
		}
		if item.Schedule.Frequency != "" {
			*lines = append(*lines, fmt.Sprintf("    frequency = \"%s\"", escapeHCL(item.Schedule.Frequency)))
		}
		if item.Schedule.Timezone != "" {
			*lines = append(*lines, fmt.Sprintf("    timezone = \"%s\"", escapeHCL(item.Schedule.Timezone)))
		}
		*lines = append(*lines, "  }")
	}

	if item.ChecksFilter != nil && hasReportFilterValues(item.ChecksFilter) {
		appendReportFilterHCL(lines, item)
	}

	*lines = append(*lines, "}")
	*lines = append(*lines, "")
}

func appendReportFilterHCL(lines *[]string, item reportConfigItem) {
	filter := item.ChecksFilter
	*lines = append(*lines, "  checks_filter {")
	if len(filter.Categories) > 0 {
		*lines = append(*lines, fmt.Sprintf("    categories = [%s]", formatQuotedList(filter.Categories)))
	}
	if item.ReportType == "GENERIC" && len(filter.ComplianceStandardIds) > 0 {
		*lines = append(*lines, fmt.Sprintf("    compliance_standard_ids = [%s]", formatQuotedList(filter.ComplianceStandardIds)))
	}
	if len(filter.Tags) > 0 {
		*lines = append(*lines, fmt.Sprintf("    tags = [%s]", formatQuotedList(filter.Tags)))
	}
	if filter.Description != "" {
		*lines = append(*lines, fmt.Sprintf("    description = \"%s\"", escapeHCL(filter.Description)))
	}
	if filter.NewerThanDays > 0 {
		*lines = append(*lines, fmt.Sprintf("    newer_than_days = %d", filter.NewerThanDays))
	}
	if filter.OlderThanDays > 0 {
		*lines = append(*lines, fmt.Sprintf("    older_than_days = %d", filter.OlderThanDays))
	}
	if len(filter.Providers) > 0 {
		*lines = append(*lines, fmt.Sprintf("    providers = [%s]", formatQuotedList(filter.Providers)))
	}
	if len(filter.Regions) > 0 {
		*lines = append(*lines, fmt.Sprintf("    regions = [%s]", formatQuotedList(filter.Regions)))
	}
	if filter.ResourceID != "" {
		*lines = append(*lines, fmt.Sprintf("    resource_id = \"%s\"", escapeHCL(filter.ResourceID)))
	}
	if filter.ResourceSearchMode != "" {
		*lines = append(*lines, fmt.Sprintf("    resource_search_mode = \"%s\"", escapeHCL(filter.ResourceSearchMode)))
	}
	if len(filter.ResourceTypes) > 0 {
		*lines = append(*lines, fmt.Sprintf("    resource_types = [%s]", formatQuotedList(filter.ResourceTypes)))
	}
	if len(filter.RiskLevels) > 0 {
		*lines = append(*lines, fmt.Sprintf("    risk_levels = [%s]", formatQuotedList(filter.RiskLevels)))
	}
	if len(filter.RuleIds) > 0 {
		*lines = append(*lines, fmt.Sprintf("    rule_ids = [%s]", formatQuotedList(filter.RuleIds)))
	}
	if len(filter.Services) > 0 {
		*lines = append(*lines, fmt.Sprintf("    services = [%s]", formatQuotedList(filter.Services)))
	}
	if len(filter.Statuses) > 0 {
		*lines = append(*lines, fmt.Sprintf("    statuses = [%s]", formatQuotedList(filter.Statuses)))
	}
	if filter.Suppressed != nil {
		*lines = append(*lines, fmt.Sprintf("    suppressed = %t", *filter.Suppressed))
	}
	*lines = append(*lines, "  }")
}

func firstListEntry(value interface{}) map[string]interface{} {
	list, ok := value.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}
	entry, ok := list[0].(map[string]interface{})
	if !ok {
		return nil
	}
	return entry
}
