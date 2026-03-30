package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadProfilesFromState(t *testing.T) {
	statePath := filepath.Join("testdata", "profile_state.json")

	profiles, err := loadProfilesFromState(statePath)
	if err != nil {
		t.Fatalf("loadProfilesFromState error: %v", err)
	}
	if len(profiles) != 1 {
		t.Fatalf("expected 1 profile, got %d", len(profiles))
	}

	profile := profiles[0]
	if profile.ID != "profile-1" || profile.Name != "Profile One" || profile.ResourceName != "profile_one" {
		t.Fatalf("unexpected profile: %+v", profile)
	}
	if profile.Description != "Example profile" {
		t.Fatalf("unexpected description: %s", profile.Description)
	}
	if len(profile.ScanRules) != 1 {
		t.Fatalf("expected 1 scan rule, got %d", len(profile.ScanRules))
	}

	rule := profile.ScanRules[0]
	if rule.ID != "R1" || rule.Provider != "aws" || !rule.Enabled || rule.RiskLevel != "HIGH" {
		t.Fatalf("unexpected scan rule: %+v", rule)
	}
	if rule.Exceptions == nil || len(rule.Exceptions.FilterTags) != 1 || len(rule.Exceptions.Tags) != 1 || len(rule.Exceptions.ResourceIds) != 1 {
		t.Fatalf("unexpected exceptions: %+v", rule.Exceptions)
	}
	if rule.Exceptions.FilterTags[0] != "prod" {
		t.Fatalf("unexpected filter tags: %+v", rule.Exceptions.FilterTags)
	}
	if rule.Exceptions.Tags[0] != "team" {
		t.Fatalf("unexpected tags: %+v", rule.Exceptions.Tags)
	}

	if len(rule.ExtraSettings) != 3 {
		t.Fatalf("expected 3 extra settings, got %d", len(rule.ExtraSettings))
	}
	if rule.ExtraSettings[0].Name != "regions" || len(rule.ExtraSettings[0].ValueSet) != 2 {
		t.Fatalf("unexpected regions setting: %+v", rule.ExtraSettings[0])
	}
	if rule.ExtraSettings[1].Name != "threshold" || rule.ExtraSettings[1].Value != "5" {
		t.Fatalf("unexpected threshold setting: %+v", rule.ExtraSettings[1])
	}
	if rule.ExtraSettings[2].Name != "toggle" || len(rule.ExtraSettings[2].Values) != 1 {
		t.Fatalf("unexpected toggle setting: %+v", rule.ExtraSettings[2])
	}
}

func TestAppendScanRuleHCL(t *testing.T) {
	rule := scanRuleItem{
		ID:        "R1",
		Provider:  "aws",
		Enabled:   true,
		RiskLevel: "HIGH",
		Exceptions: &ruleExceptionsItem{
			FilterTags:  []string{"prod", "team"},
			ResourceIds: []string{"res-1"},
		},
		ExtraSettings: []extraSettingItem{
			{
				Name: "regions",
				Type: "regions",
				ValueSet: []string{
					"us-east-1",
					"us-west-2",
				},
			},
			{
				Name:  "threshold",
				Type:  "single-number-value",
				Value: "5",
			},
			{
				Name: "toggle",
				Type: "single-string-value",
				Values: []valueItem{
					{Value: "foo", Enabled: boolPtr(true)},
				},
			},
		},
	}

	var lines []string
	appendScanRuleHCL(&lines, rule)

	expected := strings.Join([]string{
		"",
		"  scan_rule {",
		"    id = \"R1\"",
		"    provider = \"aws\"",
		"    enabled = true",
		"    risk_level = \"HIGH\"",
		"    exceptions {",
		"      filter_tags = [\"prod\", \"team\"]",
		"      resource_ids = [\"res-1\"]",
		"    }",
		"    extra_settings {",
		"      name = \"regions\"",
		"      type = \"regions\"",
		"      value_set = [\"us-east-1\", \"us-west-2\"]",
		"    }",
		"    extra_settings {",
		"      name = \"threshold\"",
		"      type = \"single-number-value\"",
		"      value = 5",
		"    }",
		"    extra_settings {",
		"      name = \"toggle\"",
		"      type = \"single-string-value\"",
		"      values {",
		"        value = \"foo\"",
		"        enabled = true",
		"      }",
		"    }",
		"  }",
	}, "\n")

	if got := strings.Join(lines, "\n"); got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}

func TestLoadProfilesFromStateWithCustomizedTags(t *testing.T) {
	statePath := filepath.Join("testdata", "profile_state_with_tags.json")

	profiles, err := loadProfilesFromState(statePath)
	if err != nil {
		t.Fatalf("loadProfilesFromState error: %v", err)
	}
	if len(profiles) != 1 {
		t.Fatalf("expected 1 profile, got %d", len(profiles))
	}

	rule := profiles[0].ScanRules[0]
	if len(rule.ExtraSettings) != 1 {
		t.Fatalf("expected 1 extra setting, got %d", len(rule.ExtraSettings))
	}
	setting := rule.ExtraSettings[0]
	if setting.Type != "choice-multiple-value-with-tags" {
		t.Fatalf("unexpected type: %s", setting.Type)
	}
	if len(setting.Values) != 1 || len(setting.Values[0].CustomizedTags) != 2 {
		t.Fatalf("unexpected customized tags: %+v", setting.Values)
	}
	if setting.Values[0].CustomizedTags[0] != "technical:application" || setting.Values[0].CustomizedTags[1] != "updated:override" {
		t.Fatalf("unexpected customized tags: %+v", setting.Values[0].CustomizedTags)
	}
}

func TestAppendScanRuleHCLWithCustomizedTags(t *testing.T) {
	rule := scanRuleItem{
		ID:        "RG-001",
		Provider:  "aws",
		Enabled:   true,
		RiskLevel: "LOW",
		ExtraSettings: []extraSettingItem{
			{
				Name: "resourceTypes",
				Type: "choice-multiple-value-with-tags",
				Values: []valueItem{
					{
						Value:          "s3-bucket",
						Enabled:        boolPtr(true),
						CustomizedTags: []string{"technical:application", "updated:override"},
					},
				},
			},
		},
	}

	var lines []string
	appendScanRuleHCL(&lines, rule)

	expected := strings.Join([]string{
		"",
		"  scan_rule {",
		"    id = \"RG-001\"",
		"    provider = \"aws\"",
		"    enabled = true",
		"    risk_level = \"LOW\"",
		"    extra_settings {",
		"      name = \"resourceTypes\"",
		"      type = \"choice-multiple-value-with-tags\"",
		"      values {",
		"        value = \"s3-bucket\"",
		"        enabled = true",
		"        customized_tags = [\"technical:application\", \"updated:override\"]",
		"      }",
		"    }",
		"  }",
	}, "\n")

	if got := strings.Join(lines, "\n"); got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}

func TestLoadProfilesFromStateWithCustomizedRiskLevel(t *testing.T) {
	statePath := filepath.Join("testdata", "profile_state_with_risk_level.json")

	profiles, err := loadProfilesFromState(statePath)
	if err != nil {
		t.Fatalf("loadProfilesFromState error: %v", err)
	}
	if len(profiles) != 1 {
		t.Fatalf("expected 1 profile, got %d", len(profiles))
	}

	rule := profiles[0].ScanRules[0]
	if len(rule.ExtraSettings) != 1 {
		t.Fatalf("expected 1 extra setting, got %d", len(rule.ExtraSettings))
	}
	setting := rule.ExtraSettings[0]
	if setting.Type != "choice-multiple-value-with-risk-level" {
		t.Fatalf("unexpected type: %s", setting.Type)
	}
	if len(setting.Values) != 1 || setting.Values[0].CustomRisk != "HIGH" {
		t.Fatalf("unexpected customized risk level: %+v", setting.Values)
	}
}

func TestAppendScanRuleHCLWithCustomizedRiskLevel(t *testing.T) {
	rule := scanRuleItem{
		ID:        "IAM-054",
		Provider:  "aws",
		Enabled:   true,
		RiskLevel: "MEDIUM",
		ExtraSettings: []extraSettingItem{
			{
				Name: "ConfigurationChanges",
				Type: "choice-multiple-value-with-risk-level",
				Values: []valueItem{
					{
						Value:      "CreateLoginProfile",
						Enabled:    boolPtr(true),
						CustomRisk: "HIGH",
					},
				},
			},
		},
	}

	var lines []string
	appendScanRuleHCL(&lines, rule)

	expected := strings.Join([]string{
		"",
		"  scan_rule {",
		"    id = \"IAM-054\"",
		"    provider = \"aws\"",
		"    enabled = true",
		"    risk_level = \"MEDIUM\"",
		"    extra_settings {",
		"      name = \"ConfigurationChanges\"",
		"      type = \"choice-multiple-value-with-risk-level\"",
		"      values {",
		"        value = \"CreateLoginProfile\"",
		"        enabled = true",
		"        customized_risk_level = \"HIGH\"",
		"      }",
		"    }",
		"  }",
	}, "\n")

	if got := strings.Join(lines, "\n"); got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}
