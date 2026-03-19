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
	if rule.Exceptions == nil || len(rule.Exceptions.FilterTags) != 2 || len(rule.Exceptions.ResourceIds) != 1 {
		t.Fatalf("unexpected exceptions: %+v", rule.Exceptions)
	}
	if rule.Exceptions.FilterTags[0] != "prod" || rule.Exceptions.FilterTags[1] != "team" {
		t.Fatalf("unexpected filter tags: %+v", rule.Exceptions.FilterTags)
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
