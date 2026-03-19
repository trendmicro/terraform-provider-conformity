package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadCustomRulesFromState(t *testing.T) {
	statePath := filepath.Join("testdata", "custom_rule_state.json")

	rules, err := loadCustomRulesFromState(statePath)
	if err != nil {
		t.Fatalf("loadCustomRulesFromState error: %v", err)
	}
	if len(rules) != 1 {
		t.Fatalf("expected 1 custom rule, got %d", len(rules))
	}

	rule := rules[0]
	if rule.ID != "custom-rule-1" || rule.Name != "Custom Rule One" || rule.ResourceName != "rule" {
		t.Fatalf("unexpected custom rule: %+v", rule)
	}
	if rule.Severity != "HIGH" || rule.CloudProvider != "aws" {
		t.Fatalf("unexpected basic fields: %+v", rule)
	}
	if len(rule.Attributes) != 1 || rule.Attributes[0].Name != "bucketEncryption" {
		t.Fatalf("unexpected attributes: %+v", rule.Attributes)
	}
	if len(rule.Rules) != 1 || len(rule.Rules[0].Conditions) != 2 {
		t.Fatalf("unexpected event rules: %+v", rule.Rules)
	}
}

func TestAppendCustomRuleHCL(t *testing.T) {
	item := customRuleItem{
		ID:               "custom-rule-1",
		Name:             "Custom Rule One",
		Description:      "Check encryption",
		CloudProvider:    "aws",
		RemediationNotes: "Turn on encryption\nStep two",
		Service:          "S3",
		ResourceType:     "s3-bucket",
		Severity:         "HIGH",
		Enabled:          true,
		Categories:       []string{"security"},
		Attributes: []customRuleAttributeItem{
			{
				Name:     "bucketEncryption",
				Path:     "data.BucketEncryption",
				Required: true,
			},
		},
		Rules: []customRuleEventRuleItem{
			{
				Description: "S3",
				Operator:    "all",
				Conditions: []customRuleConditionItem{
					{
						Operator: "equal",
						Fact:     "bucketEncryption",
						Path:     "$.Status",
						Value:    "Enabled",
					},
					{
						Operator: "equal",
						Fact:     "bucketEncryption",
						Path:     "$.Kms",
						Value:    "true",
					},
				},
			},
		},
	}

	var lines []string
	appendCustomRuleHCL(&lines, item, "custom_rule_one")

	expected := strings.Join([]string{
		"resource \"visionone_crm_custom_rule\" \"custom_rule_one\" {",
		"  name = \"Custom Rule One\"",
		"  description = \"Check encryption\"",
		"  risk_level = \"HIGH\"",
		"  cloud_provider = \"aws\"",
		"  service = \"S3\"",
		"  resource_type = \"s3-bucket\"",
		"  enabled = true",
		"  categories = [\"security\"]",
		"  remediation_note = \"Turn on encryption\\nStep two\"",
		"  attribute {",
		"    name = \"bucketEncryption\"",
		"    path = \"data.BucketEncryption\"",
		"    required = true",
		"  }",
		"  event_rule {",
		"    description = \"S3\"",
		"    conditions {",
		"      operator = \"all\"",
		"      condition {",
		"        operator = \"equal\"",
		"        fact = \"bucketEncryption\"",
		"        path = \"$.Status\"",
		"        value = \"Enabled\"",
		"      }",
		"      condition {",
		"        operator = \"equal\"",
		"        fact = \"bucketEncryption\"",
		"        path = \"$.Kms\"",
		"        value = true",
		"      }",
		"    }",
		"  }",
		"}",
		"",
	}, "\n") + "\n"

	if got := strings.Join(lines, "\n") + "\n"; got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}

func TestFormatCustomRuleValue(t *testing.T) {
	cases := map[string]string{
		"true":         "true",
		"FALSE":        "false",
		"null":         "null",
		"5":            "5",
		"5.5":          "5.5",
		"{\"days\":7}": "jsonencode({\"days\":7})",
		"Enabled":      "\"Enabled\"",
	}

	for input, expected := range cases {
		if got := formatCustomRuleValue(input); got != expected {
			t.Fatalf("unexpected value for %s: %s", input, got)
		}
	}
}
