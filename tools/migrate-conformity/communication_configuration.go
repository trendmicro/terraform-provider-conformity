package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type communicationSettingItem struct {
	ID           string
	Name         string
	ResourceName string
	Enabled      bool
	Manual       bool
	AccountID    string
	ChannelLabel string
	Warnings     []string
	ChecksFilter *communicationChecksFilterItem

	EmailConfiguration      *communicationUserIDsItem
	SmsConfiguration        *communicationUserIDsItem
	MsTeamsConfiguration    *communicationMsTeamsItem
	SlackConfiguration      *communicationSlackItem
	SnsConfiguration        *communicationSnsItem
	PagerDutyConfiguration  *communicationPagerDutyItem
	WebhookConfiguration    *communicationWebhookItem
	ServiceNowConfiguration *communicationServiceNowItem
}

type communicationChecksFilterItem struct {
	Regions               []string
	Services              []string
	RuleIDs               []string
	Categories            []string
	RiskLevels            []string
	Tags                  []string
	IgnoredTags           []string
	ComplianceStandardIDs []string
	Statuses              []string
}

type communicationUserIDsItem struct {
	UserIDs []string
}

type communicationMsTeamsItem struct {
	URL                 string
	IncludeIntroducedBy bool
	IncludeResource     bool
	IncludeTags         bool
	IncludeExtraData    bool
}

type communicationSlackItem struct {
	URL                 string
	Channel             string
	IncludeIntroducedBy bool
	IncludeResource     bool
	IncludeTags         bool
	IncludeExtraData    bool
}

type communicationSnsItem struct {
	Arn string
}

type communicationPagerDutyItem struct {
	ServiceName string
	ServiceKey  string
}

type communicationWebhookItem struct {
	URL           string
	SecurityToken string
}

type communicationServiceNowItem struct {
	Type                string
	URL                 string
	Username            string
	Password            string
	Assignee            string
	Impact              string
	Urgency             string
	DictionaryOverrides []communicationServiceNowDictionaryOverrideItem
}

type communicationServiceNowDictionaryOverrideItem struct {
	Trigger       string
	KeyValuePairs []communicationServiceNowKeyValuePairItem
}

type communicationServiceNowKeyValuePairItem struct {
	Key   string
	Value string
}

func loadCommunicationSettingsFromState(path string) ([]communicationSettingItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read state: %w", err)
	}

	var state terraformState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("parse state: %w", err)
	}

	var settings []communicationSettingItem
	for _, res := range state.Resources {
		if res.Mode != "managed" || res.Type != "conformity_communication_setting" {
			continue
		}
		for _, inst := range res.Instances {
			attrs := inst.Attributes
			if attrs == nil {
				continue
			}
			item := parseCommunicationSettingAttributes(attrs)
			if item.Name == "" {
				item.Name = "communication"
			}
			item.ResourceName = resourceNameFromState(res, inst, "communication")
			settings = append(settings, item)
		}
	}

	sort.Slice(settings, func(i, j int) bool {
		return settings[i].ResourceName < settings[j].ResourceName
	})

	return settings, nil
}

func parseCommunicationSettingAttributes(attrs map[string]interface{}) communicationSettingItem {
	item := communicationSettingItem{}
	item.ID = strings.TrimSpace(toStringValue(attrs["id"]))
	item.Enabled = toBoolValue(attrs["enabled"], true)
	item.Manual = toBoolValue(attrs["manual"], false)

	item.AccountID = parseCommunicationAccountID(attrs["relationships"])

	filter := parseCommunicationFilter(attrs["filter"])
	if filter != nil && hasCommunicationFilterValues(filter) {
		item.ChecksFilter = filter
	}

	channelLabel := ""
	if channel := parseCommunicationUserChannel(attrs["email"]); channel != nil {
		item.EmailConfiguration = channel
	} else if channel := parseCommunicationUserChannel(attrs["sms"]); channel != nil {
		item.SmsConfiguration = channel
	} else if channel := parseCommunicationMsTeamsChannel(attrs["ms_teams"]); channel != nil {
		item.MsTeamsConfiguration = channel
		label, warning := parseMsTeamsChannelLabel(attrs["ms_teams"])
		channelLabel = label
		if warning != "" {
			item.Warnings = append(item.Warnings, warning)
		}
	} else if channel := parseCommunicationSlackChannel(attrs["slack"]); channel != nil {
		item.SlackConfiguration = channel
		channelLabel = parseChannelLabel(attrs["slack"])
	} else if channel := parseCommunicationSnsChannel(attrs["sns"]); channel != nil {
		item.SnsConfiguration = channel
		channelLabel = parseChannelLabel(attrs["sns"])
	} else if channel := parseCommunicationPagerDutyChannel(attrs["pager_duty"]); channel != nil {
		item.PagerDutyConfiguration = channel
		channelLabel = parseChannelLabel(attrs["pager_duty"])
	} else if channel := parseCommunicationWebhookChannel(attrs["webhook"]); channel != nil {
		item.WebhookConfiguration = channel
	} else if channel := parseCommunicationServiceNowChannel(attrs["service_now"]); channel != nil {
		item.ServiceNowConfiguration = channel
		channelLabel = parseChannelLabel(attrs["service_now"])
	}

	item.ChannelLabel = strings.TrimSpace(channelLabel)
	item.Name = deriveCommunicationName(item)

	return item
}

func deriveCommunicationName(item communicationSettingItem) string {
	if item.ChannelLabel != "" {
		return item.ChannelLabel
	}
	switch {
	case item.EmailConfiguration != nil:
		return "email"
	case item.SmsConfiguration != nil:
		return "sms"
	case item.MsTeamsConfiguration != nil:
		return "ms_teams"
	case item.SlackConfiguration != nil:
		return "slack"
	case item.SnsConfiguration != nil:
		return "sns"
	case item.PagerDutyConfiguration != nil:
		return "pager_duty"
	case item.WebhookConfiguration != nil:
		return "webhook"
	case item.ServiceNowConfiguration != nil:
		return "service_now"
	default:
		return "communication"
	}
}

func parseCommunicationFilter(value interface{}) *communicationChecksFilterItem {
	entry := firstListEntry(value)
	if entry == nil {
		return nil
	}

	filterTags := toStringSlice(entry["filter_tags"])
	ignoredTags := toStringSlice(entry["tags"])

	filter := &communicationChecksFilterItem{
		Regions:               toStringSlice(entry["regions"]),
		Services:              toStringSlice(entry["services"]),
		RuleIDs:               toStringSlice(entry["rule_ids"]),
		Categories:            toStringSlice(entry["categories"]),
		RiskLevels:            toStringSlice(entry["risk_levels"]),
		Tags:                  uniqueStrings(filterTags),
		IgnoredTags:           uniqueStrings(ignoredTags),
		ComplianceStandardIDs: toStringSlice(entry["compliances"]),
		Statuses:              toStringSlice(entry["statuses"]),
	}

	return filter
}

func hasCommunicationFilterValues(filter *communicationChecksFilterItem) bool {
	if filter == nil {
		return false
	}
	return len(filter.Regions) > 0 ||
		len(filter.Services) > 0 ||
		len(filter.RuleIDs) > 0 ||
		len(filter.Categories) > 0 ||
		len(filter.RiskLevels) > 0 ||
		len(filter.Tags) > 0 ||
		len(filter.ComplianceStandardIDs) > 0 ||
		len(filter.Statuses) > 0
}

func parseCommunicationAccountID(value interface{}) string {
	entry := firstListEntry(value)
	if entry == nil {
		return ""
	}
	account := firstListEntry(entry["account"])
	if account == nil {
		return ""
	}
	return strings.TrimSpace(toStringValue(account["id"]))
}

func parseCommunicationUserChannel(value interface{}) *communicationUserIDsItem {
	entry := firstSetEntry(value)
	if entry == nil {
		return nil
	}
	users := toStringSlice(entry["users"])
	if len(users) == 0 {
		return nil
	}
	return &communicationUserIDsItem{UserIDs: users}
}

func parseCommunicationMsTeamsChannel(value interface{}) *communicationMsTeamsItem {
	entry := firstSetEntry(value)
	if entry == nil {
		return nil
	}
	url := strings.TrimSpace(toStringValue(entry["url"]))
	if url == "" {
		return nil
	}
	return &communicationMsTeamsItem{
		URL:                 url,
		IncludeIntroducedBy: toBoolValue(entry["display_introduced_by"], false),
		IncludeResource:     toBoolValue(entry["display_resource"], false),
		IncludeTags:         toBoolValue(entry["display_tags"], false),
		IncludeExtraData:    toBoolValue(entry["display_extra_data"], false),
	}
}

func parseCommunicationSlackChannel(value interface{}) *communicationSlackItem {
	entry := firstSetEntry(value)
	if entry == nil {
		return nil
	}
	url := strings.TrimSpace(toStringValue(entry["url"]))
	channel := strings.TrimSpace(toStringValue(entry["channel"]))
	if url == "" || channel == "" {
		return nil
	}
	return &communicationSlackItem{
		URL:                 url,
		Channel:             channel,
		IncludeIntroducedBy: toBoolValue(entry["display_introduced_by"], false),
		IncludeResource:     toBoolValue(entry["display_resource"], false),
		IncludeTags:         toBoolValue(entry["display_tags"], false),
		IncludeExtraData:    toBoolValue(entry["display_extra_data"], false),
	}
}

func parseCommunicationSnsChannel(value interface{}) *communicationSnsItem {
	entry := firstSetEntry(value)
	if entry == nil {
		return nil
	}
	arn := strings.TrimSpace(toStringValue(entry["arn"]))
	if arn == "" {
		return nil
	}
	return &communicationSnsItem{Arn: arn}
}

func parseCommunicationPagerDutyChannel(value interface{}) *communicationPagerDutyItem {
	entry := firstSetEntry(value)
	if entry == nil {
		return nil
	}
	serviceName := strings.TrimSpace(toStringValue(entry["service_name"]))
	serviceKey := strings.TrimSpace(toStringValue(entry["service_key"]))
	if serviceName == "" || serviceKey == "" {
		return nil
	}
	return &communicationPagerDutyItem{
		ServiceName: serviceName,
		ServiceKey:  serviceKey,
	}
}

func parseCommunicationWebhookChannel(value interface{}) *communicationWebhookItem {
	entry := firstSetEntry(value)
	if entry == nil {
		return nil
	}
	url := strings.TrimSpace(toStringValue(entry["url"]))
	if url == "" {
		return nil
	}
	return &communicationWebhookItem{
		URL:           url,
		SecurityToken: strings.TrimSpace(toStringValue(entry["security_token"])),
	}
}

func parseCommunicationServiceNowChannel(value interface{}) *communicationServiceNowItem {
	entry := firstSetEntry(value)
	if entry == nil {
		return nil
	}
	url := strings.TrimSpace(toStringValue(entry["url"]))
	username := strings.TrimSpace(toStringValue(entry["username"]))
	password := strings.TrimSpace(toStringValue(entry["password"]))
	snType := strings.TrimSpace(toStringValue(entry["type"]))
	if url == "" || username == "" || password == "" || snType == "" {
		return nil
	}
	return &communicationServiceNowItem{
		Type:                snType,
		URL:                 url,
		Username:            username,
		Password:            password,
		Assignee:            strings.TrimSpace(toStringValue(entry["assignee"])),
		Impact:              strings.TrimSpace(toStringValue(entry["impact"])),
		Urgency:             strings.TrimSpace(toStringValue(entry["urgency"])),
		DictionaryOverrides: parseServiceNowDictionaryOverrides(entry),
	}
}

func parseServiceNowDictionaryOverrides(entry map[string]interface{}) []communicationServiceNowDictionaryOverrideItem {
	triggers := []struct {
		name string
		key  string
	}{
		{name: "creation", key: "creation_override"},
		{name: "resolution", key: "resolution_override"},
	}

	overrides := make([]communicationServiceNowDictionaryOverrideItem, 0, len(triggers))
	for _, trigger := range triggers {
		rawMap := map[string]interface{}{}
		if existing, ok := entry[trigger.key].(map[string]interface{}); ok {
			for key, value := range existing {
				rawMap[key] = value
			}
		}

		if trigger.name == "resolution" {
			injectIfMissing(rawMap, "close_code", entry["close_code"])
			injectIfMissing(rawMap, "close_notes", entry["close_notes"])
			injectIfMissing(rawMap, "close_code", entry["closeCode"])
			injectIfMissing(rawMap, "close_notes", entry["closeNotes"])
		}

		if len(rawMap) == 0 {
			continue
		}

		keys := make([]string, 0, len(rawMap))
		for key := range rawMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		pairs := make([]communicationServiceNowKeyValuePairItem, 0, len(keys))
		for _, key := range keys {
			value := strings.TrimSpace(toStringValue(rawMap[key]))
			if key == "" || value == "" {
				continue
			}
			pairs = append(pairs, communicationServiceNowKeyValuePairItem{
				Key:   key,
				Value: value,
			})
		}

		if len(pairs) == 0 {
			continue
		}
		overrides = append(overrides, communicationServiceNowDictionaryOverrideItem{
			Trigger:       trigger.name,
			KeyValuePairs: pairs,
		})
	}

	if len(overrides) == 0 {
		return nil
	}
	return overrides
}

func injectIfMissing(target map[string]interface{}, key string, value interface{}) {
	if _, exists := target[key]; exists {
		return
	}
	if strings.TrimSpace(toStringValue(value)) == "" {
		return
	}
	target[key] = value
}

func parseChannelLabel(value interface{}) string {
	entry := firstSetEntry(value)
	if entry == nil {
		return ""
	}
	return strings.TrimSpace(toStringValue(entry["channel_name"]))
}

func parseMsTeamsChannelLabel(value interface{}) (string, string) {
	entry := firstSetEntry(value)
	if entry == nil {
		return "", ""
	}
	label := strings.TrimSpace(toStringValue(entry["channel_name"]))
	channel := strings.TrimSpace(toStringValue(entry["channel"]))
	if label != "" && channel != "" && label != channel {
		warning := fmt.Sprintf("MS Teams has both channel_name (%q) and channel (%q); using channel_name for channel_label", label, channel)
		return label, warning
	}
	if label != "" {
		return label, ""
	}
	return channel, ""
}

func appendCommunicationConfigurationHCL(lines *[]string, item communicationSettingItem, resourceName string) {
	*lines = append(*lines, fmt.Sprintf("resource \"visionone_crm_communication_configuration\" \"%s\" {", resourceName))
	*lines = append(*lines, fmt.Sprintf("  enabled = %t", item.Enabled))
	if item.AccountID != "" {
		*lines = append(*lines, fmt.Sprintf("  account_id = \"%s\"", escapeHCL(item.AccountID)))
	}
	if item.ChannelLabel != "" {
		*lines = append(*lines, fmt.Sprintf("  channel_label = \"%s\"", escapeHCL(item.ChannelLabel)))
	}
	for _, warning := range item.Warnings {
		*lines = append(*lines, fmt.Sprintf("  # Warning: %s", warning))
	}
	if item.Manual {
		*lines = append(*lines, "  manual = true")
	}

	if item.EmailConfiguration != nil {
		*lines = append(*lines, "  email_configuration = {")
		*lines = append(*lines, fmt.Sprintf("    user_ids = [%s]", formatQuotedList(item.EmailConfiguration.UserIDs)))
		*lines = append(*lines, "  }")
		*lines = append(*lines, "  # TODO: user_ids must be formatted as {identifierId}#{companyId}")
	}
	if item.SmsConfiguration != nil {
		*lines = append(*lines, "  sms_configuration = {")
		*lines = append(*lines, fmt.Sprintf("    user_ids = [%s]", formatQuotedList(item.SmsConfiguration.UserIDs)))
		*lines = append(*lines, "  }")
		*lines = append(*lines, "  # TODO: user_ids must be formatted as {identifierId}#{companyId}")
	}
	if item.MsTeamsConfiguration != nil {
		*lines = append(*lines, "  ms_teams_configuration = {")
		*lines = append(*lines, fmt.Sprintf("    url = \"%s\"", escapeHCL(item.MsTeamsConfiguration.URL)))
		*lines = append(*lines, fmt.Sprintf("    include_introduced_by = %t", item.MsTeamsConfiguration.IncludeIntroducedBy))
		*lines = append(*lines, fmt.Sprintf("    include_resource = %t", item.MsTeamsConfiguration.IncludeResource))
		*lines = append(*lines, fmt.Sprintf("    include_tags = %t", item.MsTeamsConfiguration.IncludeTags))
		*lines = append(*lines, fmt.Sprintf("    include_extra_data = %t", item.MsTeamsConfiguration.IncludeExtraData))
		*lines = append(*lines, "  }")
	}
	if item.SlackConfiguration != nil {
		*lines = append(*lines, "  slack_configuration = {")
		*lines = append(*lines, fmt.Sprintf("    url = \"%s\"", escapeHCL(item.SlackConfiguration.URL)))
		*lines = append(*lines, fmt.Sprintf("    channel = \"%s\"", escapeHCL(item.SlackConfiguration.Channel)))
		*lines = append(*lines, fmt.Sprintf("    include_introduced_by = %t", item.SlackConfiguration.IncludeIntroducedBy))
		*lines = append(*lines, fmt.Sprintf("    include_resource = %t", item.SlackConfiguration.IncludeResource))
		*lines = append(*lines, fmt.Sprintf("    include_tags = %t", item.SlackConfiguration.IncludeTags))
		*lines = append(*lines, fmt.Sprintf("    include_extra_data = %t", item.SlackConfiguration.IncludeExtraData))
		*lines = append(*lines, "  }")
	}
	if item.SnsConfiguration != nil {
		*lines = append(*lines, "  sns_configuration = {")
		*lines = append(*lines, fmt.Sprintf("    arn = \"%s\"", escapeHCL(item.SnsConfiguration.Arn)))
		*lines = append(*lines, "  }")
	}
	if item.PagerDutyConfiguration != nil {
		*lines = append(*lines, "  pagerduty_configuration = {")
		*lines = append(*lines, fmt.Sprintf("    service_name = \"%s\"", escapeHCL(item.PagerDutyConfiguration.ServiceName)))
		*lines = append(*lines, fmt.Sprintf("    service_key = \"%s\"", escapeHCL(item.PagerDutyConfiguration.ServiceKey)))
		*lines = append(*lines, "  }")
	}
	if item.WebhookConfiguration != nil {
		*lines = append(*lines, "  webhook_configuration = {")
		*lines = append(*lines, fmt.Sprintf("    url = \"%s\"", escapeHCL(item.WebhookConfiguration.URL)))
		if item.WebhookConfiguration.SecurityToken != "" {
			*lines = append(*lines, fmt.Sprintf("    security_token = \"%s\"", escapeHCL(item.WebhookConfiguration.SecurityToken)))
		}
		*lines = append(*lines, "  }")
		*lines = append(*lines, "  # TODO: headers are not available in Conformity; add if needed")
	}
	if item.ServiceNowConfiguration != nil {
		*lines = append(*lines, "  servicenow_configuration = {")
		*lines = append(*lines, fmt.Sprintf("    type = \"%s\"", escapeHCL(item.ServiceNowConfiguration.Type)))
		*lines = append(*lines, fmt.Sprintf("    url = \"%s\"", escapeHCL(item.ServiceNowConfiguration.URL)))
		*lines = append(*lines, fmt.Sprintf("    username = \"%s\"", escapeHCL(item.ServiceNowConfiguration.Username)))
		*lines = append(*lines, fmt.Sprintf("    password = \"%s\"", escapeHCL(item.ServiceNowConfiguration.Password)))
		if item.ServiceNowConfiguration.Assignee != "" {
			*lines = append(*lines, fmt.Sprintf("    assignee = \"%s\"", escapeHCL(item.ServiceNowConfiguration.Assignee)))
		}
		if item.ServiceNowConfiguration.Impact != "" {
			*lines = append(*lines, fmt.Sprintf("    impact = \"%s\"", escapeHCL(item.ServiceNowConfiguration.Impact)))
		}
		if item.ServiceNowConfiguration.Urgency != "" {
			*lines = append(*lines, fmt.Sprintf("    urgency = \"%s\"", escapeHCL(item.ServiceNowConfiguration.Urgency)))
		}
		if len(item.ServiceNowConfiguration.DictionaryOverrides) > 0 {
			*lines = append(*lines, "    dictionary_overrides = [")
			for _, dictionaryOverride := range item.ServiceNowConfiguration.DictionaryOverrides {
				*lines = append(*lines, "      {")
				*lines = append(*lines, fmt.Sprintf("        trigger = \"%s\"", escapeHCL(dictionaryOverride.Trigger)))
				*lines = append(*lines, "        key_value_pairs = [")
				for _, pair := range dictionaryOverride.KeyValuePairs {
					*lines = append(*lines, "          {")
					*lines = append(*lines, fmt.Sprintf("            key = \"%s\"", escapeHCL(pair.Key)))
					*lines = append(*lines, fmt.Sprintf("            value = \"%s\"", escapeHCL(pair.Value)))
					*lines = append(*lines, "          },")
				}
				*lines = append(*lines, "        ]")
				*lines = append(*lines, "      },")
			}
			*lines = append(*lines, "    ]")
		}
		*lines = append(*lines, "  }")
	}

	if item.ChecksFilter != nil && hasCommunicationFilterValues(item.ChecksFilter) {
		appendCommunicationFilterHCL(lines, item)
	}

	*lines = append(*lines, "}")
	*lines = append(*lines, "")
}

func appendCommunicationFilterHCL(lines *[]string, item communicationSettingItem) {
	filter := item.ChecksFilter
	*lines = append(*lines, "  checks_filter = {")
	if len(filter.IgnoredTags) > 0 {
		*lines = append(*lines, fmt.Sprintf("    # Note: filter.tags ignored: [%s]", formatQuotedList(filter.IgnoredTags)))
	}
	if len(filter.Regions) > 0 {
		*lines = append(*lines, fmt.Sprintf("    regions = [%s]", formatQuotedList(filter.Regions)))
	}
	if len(filter.Services) > 0 {
		*lines = append(*lines, fmt.Sprintf("    services = [%s]", formatQuotedList(filter.Services)))
	}
	if len(filter.RuleIDs) > 0 {
		*lines = append(*lines, fmt.Sprintf("    rule_ids = [%s]", formatQuotedList(filter.RuleIDs)))
	}
	if len(filter.Categories) > 0 {
		*lines = append(*lines, fmt.Sprintf("    categories = [%s]", formatQuotedList(filter.Categories)))
	}
	if len(filter.RiskLevels) > 0 {
		*lines = append(*lines, fmt.Sprintf("    risk_levels = [%s]", formatQuotedList(filter.RiskLevels)))
	}
	if len(filter.Tags) > 0 {
		*lines = append(*lines, fmt.Sprintf("    tags = [%s]", formatQuotedList(filter.Tags)))
	}
	if len(filter.ComplianceStandardIDs) > 0 {
		*lines = append(*lines, fmt.Sprintf("    compliance_standard_ids = [%s]", formatQuotedList(filter.ComplianceStandardIDs)))
	}
	if (item.SnsConfiguration != nil || item.WebhookConfiguration != nil) && len(filter.Statuses) > 0 {
		*lines = append(*lines, fmt.Sprintf("    statuses = [%s]", formatQuotedList(filter.Statuses)))
	}
	*lines = append(*lines, "  }")
}

func appendCommunicationMappingLines(mappingLines *[]string, item communicationSettingItem, targetName string) {
	if mappingLines == nil {
		return
	}

	sourceName := item.ResourceName
	if sourceName == "" {
		sourceName = targetName
	}
	sourceType := "conformity_communication_setting"
	targetType := "visionone_crm_communication_configuration"
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

	mapField("enabled", "enabled")
	if item.Manual {
		mapField("manual", "manual")
	}
	if item.AccountID != "" {
		mapField("relationships.account.id", "account_id")
	}

	if item.EmailConfiguration != nil {
		mapField("email.users", "email_configuration.user_ids")
	}
	if item.SmsConfiguration != nil {
		mapField("sms.users", "sms_configuration.user_ids")
	}
	if item.MsTeamsConfiguration != nil {
		mapField("ms_teams.channel_name", "channel_label")
		mapField("ms_teams.channel", "channel_label")
		mapField("ms_teams.url", "ms_teams_configuration.url")
		mapField("ms_teams.display_introduced_by", "ms_teams_configuration.include_introduced_by")
		mapField("ms_teams.display_resource", "ms_teams_configuration.include_resource")
		mapField("ms_teams.display_tags", "ms_teams_configuration.include_tags")
		mapField("ms_teams.display_extra_data", "ms_teams_configuration.include_extra_data")
	}
	if item.SlackConfiguration != nil {
		mapField("slack.channel_name", "channel_label")
		mapField("slack.url", "slack_configuration.url")
		mapField("slack.channel", "slack_configuration.channel")
		mapField("slack.display_introduced_by", "slack_configuration.include_introduced_by")
		mapField("slack.display_resource", "slack_configuration.include_resource")
		mapField("slack.display_tags", "slack_configuration.include_tags")
		mapField("slack.display_extra_data", "slack_configuration.include_extra_data")
	}
	if item.SnsConfiguration != nil {
		mapField("sns.channel_name", "channel_label")
		mapField("sns.arn", "sns_configuration.arn")
	}
	if item.PagerDutyConfiguration != nil {
		mapField("pager_duty.channel_name", "channel_label")
		mapField("pager_duty.service_name", "pagerduty_configuration.service_name")
		mapField("pager_duty.service_key", "pagerduty_configuration.service_key")
	}
	if item.WebhookConfiguration != nil {
		mapField("webhook.url", "webhook_configuration.url")
		if item.WebhookConfiguration.SecurityToken != "" {
			mapField("webhook.security_token", "webhook_configuration.security_token")
		}
	}
	if item.ServiceNowConfiguration != nil {
		mapField("service_now.channel_name", "channel_label")
		mapField("service_now.type", "servicenow_configuration.type")
		mapField("service_now.url", "servicenow_configuration.url")
		mapField("service_now.username", "servicenow_configuration.username")
		mapField("service_now.password", "servicenow_configuration.password")
		if item.ServiceNowConfiguration.Assignee != "" {
			mapField("service_now.assignee", "servicenow_configuration.assignee")
		}
		if item.ServiceNowConfiguration.Impact != "" {
			mapField("service_now.impact", "servicenow_configuration.impact")
		}
		if item.ServiceNowConfiguration.Urgency != "" {
			mapField("service_now.urgency", "servicenow_configuration.urgency")
		}
		if len(item.ServiceNowConfiguration.DictionaryOverrides) > 0 {
			mapField("service_now.creation_override", "servicenow_configuration.dictionary_overrides[creation]")
			mapField("service_now.resolution_override", "servicenow_configuration.dictionary_overrides[resolution]")
			mapField("service_now.close_code", "servicenow_configuration.dictionary_overrides[resolution].close_code")
			mapField("service_now.close_notes", "servicenow_configuration.dictionary_overrides[resolution].close_notes")
		}
	}

	if item.ChecksFilter != nil {
		if len(item.ChecksFilter.Regions) > 0 {
			mapField("filter.regions", "checks_filter.regions")
		}
		if len(item.ChecksFilter.Services) > 0 {
			mapField("filter.services", "checks_filter.services")
		}
		if len(item.ChecksFilter.RuleIDs) > 0 {
			mapField("filter.rule_ids", "checks_filter.rule_ids")
		}
		if len(item.ChecksFilter.Categories) > 0 {
			mapField("filter.categories", "checks_filter.categories")
		}
		if len(item.ChecksFilter.RiskLevels) > 0 {
			mapField("filter.risk_levels", "checks_filter.risk_levels")
		}
		if len(item.ChecksFilter.Tags) > 0 {
			mapField("filter.filter_tags", "checks_filter.tags")
		}
		if len(item.ChecksFilter.ComplianceStandardIDs) > 0 {
			mapField("filter.compliances", "checks_filter.compliance_standard_ids")
		}
		if len(item.ChecksFilter.Statuses) > 0 && (item.SnsConfiguration != nil || item.WebhookConfiguration != nil) {
			mapField("filter.statuses", "checks_filter.statuses")
		}
	}
}

func firstSetEntry(value interface{}) map[string]interface{} {
	set, ok := value.([]interface{})
	if !ok || len(set) == 0 {
		return nil
	}
	entry, ok := set[0].(map[string]interface{})
	if !ok {
		return nil
	}
	return entry
}
