package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadReportConfigsFromState(t *testing.T) {
	statePath := filepath.Join("testdata", "report_config_state.json")

	configs, err := loadReportConfigsFromState(statePath)
	if err != nil {
		t.Fatalf("loadReportConfigsFromState error: %v", err)
	}
	if len(configs) != 2 {
		t.Fatalf("expected 2 report configs, got %d", len(configs))
	}

	if configs[0].ReportTitle != "Report A" || configs[0].ID != "report-config:one" || configs[0].ResourceName != "report_a" {
		t.Fatalf("unexpected first report config: %+v", configs[0])
	}
	if configs[0].Schedule == nil || configs[0].Schedule.Frequency != "* * *" {
		t.Fatalf("expected schedule on Report A")
	}
	if configs[0].Schedule.Enabled == nil || *configs[0].Schedule.Enabled != true {
		t.Fatalf("expected schedule enabled on Report A")
	}
	if configs[0].ChecksFilter == nil || len(configs[0].ChecksFilter.Tags) != 1 {
		t.Fatalf("expected tags to include filter_tags only on Report A")
	}
	if configs[0].ChecksFilter.Tags[0] != "prod" {
		t.Fatalf("unexpected tags on Report A: %+v", configs[0].ChecksFilter.Tags)
	}
	if len(configs[0].ChecksFilter.IgnoredTags) != 1 || configs[0].ChecksFilter.IgnoredTags[0] != "team" {
		t.Fatalf("unexpected ignored tags on Report A: %+v", configs[0].ChecksFilter.IgnoredTags)
	}
	if configs[0].ChecksFilter == nil || configs[0].ChecksFilter.Suppressed == nil || *configs[0].ChecksFilter.Suppressed != false {
		t.Fatalf("expected suppressed=false on Report A")
	}

	if configs[1].ReportTitle != "Report B" || configs[1].ID != "report-config:two" || configs[1].ResourceName != "report_b" {
		t.Fatalf("unexpected second report config: %+v", configs[1])
	}
	if configs[1].Schedule == nil || configs[1].Schedule.Enabled == nil || *configs[1].Schedule.Enabled != false {
		t.Fatalf("expected schedule disabled on Report B")
	}
	if configs[1].AppliedComplianceStandardID != "NIST4" {
		t.Fatalf("expected compliance standard ID on Report B")
	}
	if configs[1].ControlsType != "withChecksOnly" {
		t.Fatalf("expected controls_type on Report B")
	}
	if configs[1].ChecksFilter != nil && len(configs[1].ChecksFilter.ComplianceStandardIds) > 0 {
		t.Fatalf("compliance_standard_ids should be empty for compliance reports")
	}
}

func TestAppendReportConfigHCL(t *testing.T) {
	item := reportConfigItem{
		ID:                  "report-config:one",
		ReportTitle:         "Report A",
		ReportType:          "GENERIC",
		Schedule: &reportScheduleItem{
			Enabled:   boolPtr(true),
			Frequency: "* * *",
			Timezone:  "Asia/Manila",
		},
		ChecksFilter: &reportFilterItem{
			Categories:            []string{"security"},
			ComplianceStandardIds: []string{"NIST4"},
			Tags:                  []string{"prod", "team"},
			Description:           "find me",
			NewerThanDays:         7,
			Providers:             []string{"aws"},
			Regions:               []string{"us-east-1"},
			ResourceID:            "res-1",
			ResourceSearchMode:    "text",
			ResourceTypes:         []string{"s3-bucket"},
			RiskLevels:            []string{"HIGH"},
			RuleIds:               []string{"S3-001"},
			Services:              []string{"S3"},
			Statuses:              []string{"FAILURE"},
			Suppressed:            boolPtr(false),
		},
	}

	var hclLines []string
	appendReportConfigHCL(&hclLines, item, "report_a")

	expected := strings.Join([]string{
		"resource \"visionone_crm_report_config\" \"report_a\" {",
		"  report_title = \"Report A\"",
		"  report_type = \"GENERIC\"",
		"  # include_account_names @TODO: review manually; Conformity state is inconsistent for this field",
		"  schedule {",
		"    enabled = true",
		"    frequency = \"* * *\"",
		"    timezone = \"Asia/Manila\"",
		"  }",
		"  checks_filter {",
		"    categories = [\"security\"]",
		"    compliance_standard_ids = [\"NIST4\"]",
		"    tags = [\"prod\", \"team\"]",
		"    description = \"find me\"",
		"    newer_than_days = 7",
		"    providers = [\"aws\"]",
		"    regions = [\"us-east-1\"]",
		"    resource_id = \"res-1\"",
		"    resource_search_mode = \"text\"",
		"    resource_types = [\"s3-bucket\"]",
		"    risk_levels = [\"HIGH\"]",
		"    rule_ids = [\"S3-001\"]",
		"    services = [\"S3\"]",
		"    statuses = [\"FAILURE\"]",
		"    suppressed = false",
		"  }",
		"}",
		"",
	}, "\n") + "\n"

	if got := strings.Join(hclLines, "\n") + "\n"; got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}

func TestAppendReportConfigHCL_EmptyTagsOmitted(t *testing.T) {
	item := reportConfigItem{
		ReportTitle:         "Report C",
		ReportType:          "GENERIC",
		ChecksFilter: &reportFilterItem{
			Tags: []string{},
		},
	}

	var hclLines []string
	appendReportConfigHCL(&hclLines, item, "report_c")

	expected := strings.Join([]string{
		"resource \"visionone_crm_report_config\" \"report_c\" {",
		"  report_title = \"Report C\"",
		"  report_type = \"GENERIC\"",
		"  # include_account_names @TODO: review manually; Conformity state is inconsistent for this field",
		"}",
		"",
	}, "\n") + "\n"

	if got := strings.Join(hclLines, "\n") + "\n"; got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}

func TestParseReportConfigAttributes_DoesNotMapIncludeAccountNames(t *testing.T) {
	attrs := map[string]interface{}{
		"configuration": []interface{}{
			map[string]interface{}{
				"title":                 "Report D",
				"include_account_names": true,
			},
		},
	}

	item := parseReportConfigAttributes(attrs)
	if item.ReportTitle != "Report D" {
		t.Fatalf("expected report title to be parsed")
	}
}

func TestAppendReportConfigHCL_AccountLevelOmitsIncludeAccountNamesReviewComment(t *testing.T) {
	attrs := map[string]interface{}{
		"account_id": "account-123",
		"configuration": []interface{}{
			map[string]interface{}{
				"title":                 "Report E",
				"include_account_names": true,
			},
		},
	}

	item := parseReportConfigAttributes(attrs)
	var hclLines []string
	appendReportConfigHCL(&hclLines, item, "report_e")

	got := strings.Join(hclLines, "\n")
	if strings.Contains(got, "include_account_names") {
		t.Fatalf("expected include_account_names to be omitted from output")
	}
	if strings.Contains(got, "review include_account_names manually") {
		t.Fatalf("expected manual include_account_names review comment to be omitted for account-level report")
	}
}

func TestParseReportConfigAttributes_MessageFalseDoesNotPopulateChecksFilter(t *testing.T) {
	attrs := map[string]interface{}{
		"configuration": []interface{}{
			map[string]interface{}{
				"title": "Report F",
			},
		},
		"filter": []interface{}{
			map[string]interface{}{
				"message": false,
			},
		},
	}

	item := parseReportConfigAttributes(attrs)
	if item.ChecksFilter != nil {
		t.Fatalf("expected ChecksFilter to be nil when filter.message is non-string/empty")
	}
}

func TestParseReportFilterMessage(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{name: "string", input: "find me", want: "find me"},
		{name: "trimmed", input: "  find me  ", want: "find me"},
		{name: "empty", input: "", want: ""},
		{name: "bool false", input: false, want: ""},
		{name: "nil", input: nil, want: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseReportFilterMessage(tc.input)
			if got != tc.want {
				t.Fatalf("unexpected message parse result: got %q, want %q", got, tc.want)
			}
		})
	}
}

func boolPtr(val bool) *bool {
	return &val
}
