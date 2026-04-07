package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadCheckSuppressionsFromState(t *testing.T) {
	statePath := filepath.Join("testdata", "check_suppression_state.json")

	items, err := loadCheckSuppressionsFromState(statePath)
	if err != nil {
		t.Fatalf("loadCheckSuppressionsFromState error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 check suppressions, got %d", len(items))
	}

	if items[0].RuleID != "EC2-074" || items[0].Service != "EC2" || items[0].ResourceName != "critical" {
		t.Fatalf("unexpected first suppression: %+v", items[0])
	}
	if items[1].RuleID != "S3-001" || items[1].Service != "S3" || items[1].ResourceName != "global" {
		t.Fatalf("unexpected second suppression: %+v", items[1])
	}
}

func TestAppendCheckSuppressionHCL(t *testing.T) {
	item := checkSuppressionItem{
		ID:         "ccc:account-1:EC2-074:EC2:us-east-1:sg-123",
		AccountID:  "account-1",
		Service:    "EC2",
		RuleID:     "EC2-074",
		Region:     "us-east-1",
		ResourceID: "sg-123",
		Note:       "maintenance",
	}

	var lines []string
	lines = appendCheckSuppressionHCL(lines, item, "critical")

	expected := strings.Join([]string{
		"resource \"visionone_crm_check_suppression\" \"critical\" {",
		"  account_id = \"account-1\"",
		"  service = \"EC2\"",
		"  rule_id = \"EC2-074\"",
		"  region = \"us-east-1\"",
		"  resource_id = \"sg-123\"",
		"  note = \"maintenance\"",
		"}",
		"",
	}, "\n") + "\n"

	if got := strings.Join(lines, "\n") + "\n"; got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}

func TestDeriveServiceFromRuleID(t *testing.T) {
	if got := deriveServiceFromRuleID("EC2-074"); got != "EC2" {
		t.Fatalf("unexpected service: %s", got)
	}
	if got := deriveServiceFromRuleID(""); got != "" {
		t.Fatalf("expected empty service, got: %s", got)
	}
}
