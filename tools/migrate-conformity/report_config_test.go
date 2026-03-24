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
	if configs[0].ChecksFilter == nil || len(configs[0].ChecksFilter.Tags) != 2 {
		t.Fatalf("expected merged tags on Report A")
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

func boolPtr(val bool) *bool {
	return &val
}
