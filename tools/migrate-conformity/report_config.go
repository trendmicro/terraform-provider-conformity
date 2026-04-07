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
	IgnoredTags           []string
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
	SuppressedReviewNote  bool
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
	filterTags := toStringSlice(filterTagsValue)
	ignoredTags := toStringSlice(entry["tags"])

	filter := &reportFilterItem{
		Categories:            toStringSlice(entry["categories"]),
		ComplianceStandardIds: toStringSlice(entry["compliance_standards"]),
		Tags:                  uniqueStrings(filterTags),
		IgnoredTags:           uniqueStrings(ignoredTags),
		Description:           parseReportFilterMessage(entry["message"]),
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
			if suppressed {
				filter.SuppressedReviewNote = true
			}
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

func parseReportFilterMessage(value interface{}) string {
	message, ok := value.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(message)
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
	if item.AccountID == "" {
		*lines = append(*lines, "  # @TODO review manually `include_account_names`: Conformity state is inconsistent for this field")
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
			*lines = append(*lines, "  # @TODO review manually `applied_compliance_standard_id`: set this for compliance report")
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
	if len(filter.IgnoredTags) > 0 {
		*lines = append(*lines, fmt.Sprintf("    # @TODO review manually `checks_filter.tags`: source filter.tags was ignored: [%s]", formatQuotedList(filter.IgnoredTags)))
	}
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
	if filter.SuppressedReviewNote {
		*lines = append(*lines, "    # @TODO review manually `suppressed`: suppressed=true in v2 may come from omitted source field in Conformity state")
	}
	if filter.Suppressed != nil {
		*lines = append(*lines, fmt.Sprintf("    suppressed = %t", *filter.Suppressed))
	}
	*lines = append(*lines, "  }")
}

func appendReportConfigMappingLines(mappingLines *[]string, item reportConfigItem, targetName string) {
	if mappingLines == nil {
		return
	}

	sourceName := item.ResourceName
	if sourceName == "" {
		sourceName = targetName
	}
	sourceType := "conformity_report_config"
	targetType := "visionone_crm_report_config"
	*mappingLines = append(*mappingLines, formatMappingLine(sourceType, sourceName, targetType, targetName))

	seen := map[string]struct{}{}
	appendUnique := func(line string) {
		if _, ok := seen[line]; ok {
			return
		}
		seen[line] = struct{}{}
		*mappingLines = append(*mappingLines, line)
	}
	mapField := func(sourceAttribute, targetAttribute string) {
		appendUnique(formatAttributeMappingLine(sourceAttribute, targetAttribute))
	}

	if item.AccountID != "" {
		mapField("account_id", "account_id")
	}
	if item.GroupID != "" {
		mapField("group_id", "group_id")
	}
	mapField("configuration.title", "report_title")
	if item.ReportType != "" {
		mapField("configuration.generate_report_type", "report_type")
	}
	if item.IncludeChecks {
		mapField("configuration.include_checks", "include_checks")
	}
	if item.EmailRecipientsSet {
		mapField("configuration.emails", "email_recipients")
	}
	if item.ReportFormatsInEmailSet {
		mapField("configuration.should_email_include_pdf/csv", "report_formats_in_email")
	}
	if item.Schedule != nil {
		if item.Schedule.Enabled != nil {
			mapField("configuration.scheduled", "schedule.enabled")
		}
		if item.Schedule.Frequency != "" {
			mapField("configuration.frequency", "schedule.frequency")
		}
		if item.Schedule.Timezone != "" {
			mapField("configuration.tz", "schedule.timezone")
		}
	}
	if item.ReportType == "COMPLIANCE-STANDARD" {
		mapField("filter.report_compliance_standard_id", "applied_compliance_standard_id")
		mapField("filter.with_checks/without_checks", "controls_type")
	}
	if item.ChecksFilter != nil && hasReportFilterValues(item.ChecksFilter) {
		mapField("filter.categories", "checks_filter.categories")
		if item.ReportType == "GENERIC" {
			mapField("filter.compliance_standards", "checks_filter.compliance_standard_ids")
		}
		mapField("filter.filter_tags", "checks_filter.tags")
		mapField("filter.message", "checks_filter.description")
		mapField("filter.newer_than_days", "checks_filter.newer_than_days")
		mapField("filter.older_than_days", "checks_filter.older_than_days")
		mapField("filter.providers", "checks_filter.providers")
		mapField("filter.regions", "checks_filter.regions")
		mapField("filter.resource", "checks_filter.resource_id")
		mapField("filter.resource_search_mode", "checks_filter.resource_search_mode")
		mapField("filter.resource_types", "checks_filter.resource_types")
		mapField("filter.risk_levels", "checks_filter.risk_levels")
		mapField("filter.rule_ids", "checks_filter.rule_ids")
		mapField("filter.services", "checks_filter.services")
		mapField("filter.statuses", "checks_filter.statuses")
		mapField("filter.suppressed", "checks_filter.suppressed")
	}
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
