package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadCommunicationSettingsFromState(t *testing.T) {
	statePath := filepath.Join("testdata", "communication_setting_state.json")

	settings, err := loadCommunicationSettingsFromState(statePath)
	if err != nil {
		t.Fatalf("loadCommunicationSettingsFromState error: %v", err)
	}
	if len(settings) != 8 {
		t.Fatalf("expected 8 settings, got %d", len(settings))
	}

	byID := map[string]communicationSettingItem{}
	for _, setting := range settings {
		byID[setting.ID] = setting
	}

	first := byID["comm-1"]
	if first.ID != "comm-1" || first.EmailConfiguration == nil {
		t.Fatalf("unexpected email setting: %+v", first)
	}
	if first.ResourceName != "email" {
		t.Fatalf("unexpected resource name: %s", first.ResourceName)
	}
	if first.ChecksFilter == nil || len(first.ChecksFilter.ComplianceStandardIDs) != 3 {
		t.Fatalf("expected compliance standards")
	}

	sms := byID["comm-2"]
	if sms.SmsConfiguration == nil || sms.AccountID != "account-1" {
		t.Fatalf("unexpected sms setting: %+v", sms)
	}
	if sms.ChecksFilter == nil || len(sms.ChecksFilter.Tags) != 2 {
		t.Fatalf("unexpected sms filter: %+v", sms.ChecksFilter)
	}

	slack := byID["comm-3"]
	if slack.ChannelLabel != "SlackChannel" || slack.SlackConfiguration == nil {
		t.Fatalf("unexpected slack setting: %+v", slack)
	}

	msTeams := byID["comm-4"]
	if msTeams.ChannelLabel != "msTeamsChannel" || msTeams.MsTeamsConfiguration == nil {
		t.Fatalf("unexpected ms teams setting: %+v", msTeams)
	}
	if msTeams.ChecksFilter == nil || len(msTeams.ChecksFilter.Tags) == 0 {
		t.Fatalf("unexpected ms teams filter: %+v", msTeams.ChecksFilter)
	}

	serviceNow := byID["comm-5"]
	if serviceNow.ChannelLabel != "ServiceNowChannel" || serviceNow.ServiceNowConfiguration == nil {
		t.Fatalf("unexpected service now setting: %+v", serviceNow)
	}
	if !serviceNow.Manual || serviceNow.AccountID != "account-1" {
		t.Fatalf("unexpected service now flags: %+v", serviceNow)
	}
	if len(serviceNow.ServiceNowConfiguration.DictionaryOverrides) != 2 {
		t.Fatalf("unexpected service now dictionary overrides: %+v", serviceNow.ServiceNowConfiguration.DictionaryOverrides)
	}

	sns := byID["comm-6"]
	if sns.ChannelLabel != "snsChannel" || sns.SnsConfiguration == nil {
		t.Fatalf("unexpected sns setting: %+v", sns)
	}
	if sns.ChecksFilter == nil || len(sns.ChecksFilter.Statuses) != 2 {
		t.Fatalf("unexpected sns filter: %+v", sns.ChecksFilter)
	}

	pagerDuty := byID["comm-7"]
	if pagerDuty.ChannelLabel != "pagerdutyChannel" || pagerDuty.PagerDutyConfiguration == nil {
		t.Fatalf("unexpected pager duty setting: %+v", pagerDuty)
	}

	webhook := byID["comm-8"]
	if webhook.WebhookConfiguration == nil || webhook.AccountID != "account-1" {
		t.Fatalf("unexpected webhook setting: %+v", webhook)
	}
	if webhook.ResourceName != "webhook" {
		t.Fatalf("unexpected webhook resource name: %s", webhook.ResourceName)
	}
	if webhook.ChecksFilter == nil || len(webhook.ChecksFilter.Statuses) != 1 {
		t.Fatalf("unexpected webhook filter: %+v", webhook.ChecksFilter)
	}
}

func TestAppendCommunicationConfigurationHCL_ServiceNowDictionaryOverrides(t *testing.T) {
	item := communicationSettingItem{
		ID:           "comm-5",
		Enabled:      true,
		Manual:       true,
		AccountID:    "account-1",
		ChannelLabel: "ServiceNowChannel",
		ServiceNowConfiguration: &communicationServiceNowItem{
			Type:     "problem",
			URL:      "https://ven12345.service-now.com",
			Username: "conformity.qa.automation",
			Password: "notapassword",
			Assignee: "admin",
			Impact:   "1",
			Urgency:  "1",
			DictionaryOverrides: []communicationServiceNowDictionaryOverrideItem{
				{
					Trigger: "creation",
					KeyValuePairs: []communicationServiceNowKeyValuePairItem{
						{Key: "priority", Value: "2"},
						{Key: "urgency", Value: "2"},
					},
				},
				{
					Trigger: "resolution",
					KeyValuePairs: []communicationServiceNowKeyValuePairItem{
						{Key: "close_code", Value: "Closed/Resolved by Caller"},
						{Key: "close_notes", Value: "Issue resolved"},
					},
				},
			},
		},
	}

	var lines []string
	appendCommunicationConfigurationHCL(&lines, item, "servicenow")

	expected := strings.Join([]string{
		"resource \"visionone_crm_communication_configuration\" \"servicenow\" {",
		"  enabled = true",
		"  account_id = \"account-1\"",
		"  channel_label = \"ServiceNowChannel\"",
		"  manual = true",
		"  servicenow_configuration = {",
		"    type = \"problem\"",
		"    url = \"https://ven12345.service-now.com\"",
		"    username = \"conformity.qa.automation\"",
		"    password = \"notapassword\"",
		"    assignee = \"admin\"",
		"    impact = \"1\"",
		"    urgency = \"1\"",
		"    dictionary_overrides = [",
		"      {",
		"        trigger = \"creation\"",
		"        key_value_pairs = [",
		"          {",
		"            key = \"priority\"",
		"            value = \"2\"",
		"          },",
		"          {",
		"            key = \"urgency\"",
		"            value = \"2\"",
		"          },",
		"        ]",
		"      },",
		"      {",
		"        trigger = \"resolution\"",
		"        key_value_pairs = [",
		"          {",
		"            key = \"close_code\"",
		"            value = \"Closed/Resolved by Caller\"",
		"          },",
		"          {",
		"            key = \"close_notes\"",
		"            value = \"Issue resolved\"",
		"          },",
		"        ]",
		"      },",
		"    ]",
		"  }",
		"}",
		"",
	}, "\n") + "\n"

	if got := strings.Join(lines, "\n") + "\n"; got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}

func TestParseMsTeamsChannelLabelFallback(t *testing.T) {
	item := parseCommunicationSettingAttributes(map[string]interface{}{
		"id": "comm-ms",
		"ms_teams": []interface{}{
			map[string]interface{}{
				"url":                   "https://example.com/hook",
				"channel":               "fallback-channel",
				"display_introduced_by": true,
				"display_resource":      true,
				"display_tags":          false,
				"display_extra_data":    true,
			},
		},
	})

	if item.MsTeamsConfiguration == nil {
		t.Fatalf("expected ms teams configuration")
	}
	if item.ChannelLabel != "fallback-channel" {
		t.Fatalf("expected fallback channel label, got %q", item.ChannelLabel)
	}
}

func TestParseMsTeamsChannelLabelConflictWarning(t *testing.T) {
	item := parseCommunicationSettingAttributes(map[string]interface{}{
		"id": "comm-ms-conflict",
		"ms_teams": []interface{}{
			map[string]interface{}{
				"url":                   "https://example.com/hook",
				"channel_name":          "team-alerts",
				"channel":               "legacy-channel",
				"display_introduced_by": true,
				"display_resource":      true,
				"display_tags":          false,
				"display_extra_data":    true,
			},
		},
	})

	if item.ChannelLabel != "team-alerts" {
		t.Fatalf("expected channel_name to win, got %q", item.ChannelLabel)
	}
	if len(item.Warnings) != 1 {
		t.Fatalf("expected one warning, got %d", len(item.Warnings))
	}
	if !strings.Contains(item.Warnings[0], "using channel_name for channel_label") {
		t.Fatalf("unexpected warning message: %s", item.Warnings[0])
	}

	var lines []string
	appendCommunicationConfigurationHCL(&lines, item, "ms_teams")
	output := strings.Join(lines, "\n")
	if !strings.Contains(output, "# Warning: MS Teams has both channel_name") {
		t.Fatalf("expected warning comment in HCL output:\n%s", output)
	}
}

func TestParseServiceNowDictionaryOverridesFromLegacyCloseFields(t *testing.T) {
	channel := parseCommunicationServiceNowChannel([]interface{}{
		map[string]interface{}{
			"type":        "problem",
			"url":         "https://example.service-now.com",
			"username":    "user",
			"password":    "pass",
			"close_code":  "solved",
			"close_notes": "resolved via migration",
		},
	})

	if channel == nil {
		t.Fatalf("expected service now channel")
	}
	if len(channel.DictionaryOverrides) != 1 {
		t.Fatalf("expected 1 dictionary override, got %d", len(channel.DictionaryOverrides))
	}
	resolution := channel.DictionaryOverrides[0]
	if resolution.Trigger != "resolution" {
		t.Fatalf("expected resolution trigger, got %s", resolution.Trigger)
	}
	if len(resolution.KeyValuePairs) != 2 {
		t.Fatalf("expected 2 key-value pairs, got %d", len(resolution.KeyValuePairs))
	}
	if resolution.KeyValuePairs[0].Key != "close_code" || resolution.KeyValuePairs[0].Value != "solved" {
		t.Fatalf("unexpected first pair: %+v", resolution.KeyValuePairs[0])
	}
	if resolution.KeyValuePairs[1].Key != "close_notes" || resolution.KeyValuePairs[1].Value != "resolved via migration" {
		t.Fatalf("unexpected second pair: %+v", resolution.KeyValuePairs[1])
	}
}

func TestAppendCommunicationConfigurationHCL(t *testing.T) {
	item := communicationSettingItem{
		ID:           "comm-1",
		Enabled:      true,
		AccountID:    "account-1",
		ChannelLabel: "Alerts",
		EmailConfiguration: &communicationUserIDsItem{
			UserIDs: []string{"user-1", "user-2"},
		},
		ChecksFilter: &communicationChecksFilterItem{
			Regions:               []string{"us-east-1"},
			Services:              []string{"S3"},
			RuleIDs:               []string{"S3-001"},
			Categories:            []string{"security"},
			RiskLevels:            []string{"HIGH"},
			Tags:                  []string{"prod", "team"},
			ComplianceStandardIDs: []string{"NIST4"},
		},
	}

	var lines []string
	appendCommunicationConfigurationHCL(&lines, item, "email")

	expected := strings.Join([]string{
		"resource \"visionone_crm_communication_configuration\" \"email\" {",
		"  enabled = true",
		"  account_id = \"account-1\"",
		"  channel_label = \"Alerts\"",
		"  email_configuration = {",
		"    user_ids = [\"user-1\", \"user-2\"]",
		"  }",
		"  # TODO: user_ids must be formatted as {identifierId}#{companyId}",
		"  checks_filter = {",
		"    regions = [\"us-east-1\"]",
		"    services = [\"S3\"]",
		"    rule_ids = [\"S3-001\"]",
		"    categories = [\"security\"]",
		"    risk_levels = [\"HIGH\"]",
		"    tags = [\"prod\", \"team\"]",
		"    compliance_standard_ids = [\"NIST4\"]",
		"  }",
		"}",
		"",
	}, "\n") + "\n"

	if got := strings.Join(lines, "\n") + "\n"; got != expected {
		t.Fatalf("unexpected HCL output:\n%s", got)
	}
}
