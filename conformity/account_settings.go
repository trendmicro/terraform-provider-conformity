package conformity

import (
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenAccountSettings(settings *cloudconformity.AccountSettings, rule []cloudconformity.GetRuleSettings) []interface{} {

	c := make(map[string]interface{})

	if settings == nil || (settings.Bot == nil && rule == nil) {
		return nil
	}

	c["bot"] = flattenBotSettings(settings.Bot)
	c["rule"] = flattenRuleSettings(rule)

	return []interface{}{c}
}
func flattenRuleSettings(rules []cloudconformity.GetRuleSettings) []interface{} {
	if rules == nil {
		return make([]interface{}, 0)
	}
	rs := make([]interface{}, len(rules))
	for i, rule := range rules {
		r := make(map[string]interface{})
		r["rule_id"] = rule.Id
		r["settings"] = flattenSettings(rule)
		rs[i] = r
	}

	return rs
}
func flattenBotSettings(bot *cloudconformity.AccountBot) []interface{} {

	c := make(map[string]interface{})
	if bot == nil || bot.Disabled == nil || bot.Delay == nil {
		return make([]interface{}, 0)
	}

	c["disabled"] = *bot.Disabled
	c["delay"] = *bot.Delay
	if regions := flattenBotDisabledRegions(bot.DisabledRegions); len(regions) > 0 {
		c["disabled_regions"] = regions
	}
	return []interface{}{c}
}

// extract a list of regions from the struct which are disabled
func flattenBotDisabledRegions(regions *cloudconformity.BotDisabledRegions) []string {
	regionList := []string{}

	if regions == nil {
		return regionList
	}
	v := reflect.ValueOf(regions).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		value := v.FieldByName(name).Bool()

		if value {
			jsonTag := field.Tag.Get("json")
			// we need the tag name of the field and remove the omitempty
			// ie || EuNorth1     bool `json:"eu-north-1,omitempty"` || extracts eu-north-1
			regionList = append(regionList, strings.Split(jsonTag, ",")[0])
		}
	}
	return regionList
}

func flattenSettings(rules cloudconformity.GetRuleSettings) []interface{} {
	c := make(map[string]interface{})
	c["risk_level"] = rules.RiskLevel
	c["enabled"] = rules.Enabled
	c["rule_exists"] = rules.RuleExists
	if rules.Exceptions != nil {
		c["exceptions"] = flattenExceptions(rules.Exceptions)
	}
	if rules.ExtraSettings != nil {
		c["extra_settings"] = flattenExtraSettings(rules.ExtraSettings)
	}

	return []interface{}{c}
}

func flattenExceptions(exceptions *cloudconformity.RuleSettingExceptions) []interface{} {
	c := make(map[string]interface{})

	c["filter_tags"] = exceptions.FilterTags

	c["resources"] = exceptions.Resources
	if len(exceptions.Tags) == 0{
		c["tags"] = nil
	}else{
		c["tags"] = exceptions.Tags
	}
	return []interface{}{c}
}
func flattenExtraSettings(extra []*cloudconformity.RuleSettingExtra) []interface{} {
	if extra == nil {
		return make([]interface{}, 0)
	}
	ex := make([]interface{}, len(extra))
	for i, v := range extra {
		e := make(map[string]interface{})
		e["name"] = v.Name
		e["type"] = v.Type
		e["value"] = v.Value

		if v.Values != nil {

			values := v.Values.([]interface{})
			switch v.Type {
			case "regions":

				e["regions"] = expandStringList(values)

			case "multiple-object-values":

				e["multiple_object_values"] = flattenRuleMultipleObject(values[0].(map[string]interface{}))

			default:

				e["values"] = flattenRuleValues(values)

			}
		}
		if v.Mappings != nil {

			mappings := v.Mappings.([]interface{})
			e["mappings"] = flattenRuleMappings(mappings)
		}

		ex[i] = e
	}
	return ex
}

func flattenRuleValues(values []interface{}) []interface{} {
	if values == nil {
		return make([]interface{}, 0)
	}
	vs := make([]interface{}, 0, len(values))

	for _, value := range values {

		v := make(map[string]interface{})
		item := value.(map[string]interface{})

		if l, ok := item["label"].(string); ok && l != "" {
			v["label"] = l
		}
		if val, ok := item["value"].(string); ok && val != "" {
			v["value"] = val
		}

		if enabled, ok := item["enabled"]; ok && enabled != nil {
			v["enabled"] = item["enabled"].(bool)

		}
		vs = append(vs, v)

	}
	return vs
}

func flattenRuleMultipleObject(multiple map[string]interface{}) []interface{} {

	if multiple == nil {
		return make([]interface{}, 0)
	}

	mo := make([]interface{}, 0, len(multiple))

	for _, object := range multiple {

		v := make(map[string]interface{})
		item := object.(map[string]interface{})

		if en, ok := item["eventName"].(string); ok && en != "" {
			v["event_name"] = en
		}
		if es, ok := item["eventSource"].(string); ok && es != "" {
			v["event_source"] = es
		}
		if uit, ok := item["userIdentityType"].(string); ok && uit != "" {
			v["user_identity_type"] = uit
		}
		mo = append(mo, v)
	}
	return mo
}
func flattenRuleMappings(mappings []interface{}) []interface{} {
	if mappings == nil {
		return make([]interface{}, 0)
	}
	mp := make([]interface{}, 0, len(mappings))
	for _, mapping := range mappings {
		v := make(map[string]interface{})
		item := mapping.(map[string]interface{})
		if values := item["values"].([]interface{}); len(values) > 0 {
			v["values"] = flattenRuleMappingsValues(values)
		}

		mp = append(mp, v)
	}
	return mp
}
func flattenRuleMappingsValues(val []interface{}) []interface{} {
	if val == nil {
		return make([]interface{}, 0)
	}
	mp := make([]interface{}, 0, len(val))
	for _, values := range val {
		item := values.(map[string]interface{})
		v := make(map[string]interface{})

		if n, ok := item["name"].(string); ok && n != "" {
			v["name"] = n
		}
		if t, ok := item["type"].(string); ok && t != "" {
			v["type"] = t
		}
		if val, ok := item["value"].(string); ok && val != "" {
			v["value"] = val
		}
		if values, ok := item["values"].([]interface{}); ok && len(values) > 0 {
			v["values"] = flattenMappingValue(values)
		}

		mp = append(mp, v)
	}

	return mp
}
func flattenMappingValue(val []interface{}) []interface{} {
	if val == nil {
		return make([]interface{}, 0)
	}
	values := make([]interface{}, 0, len(val))
	for _, v := range val {
		value := make(map[string]interface{})
		item := v.(map[string]interface{})
		value["value"] = item["value"].(string)
		values = append(values, value)
	}
	return values

}
func updateAccountTags(payload cloudconformity.AccountPayload, accountId string, d *schema.ResourceData, client *cloudconformity.Client) error {

	tags := d.Get("tags").(*schema.Set)
	for _, tag := range tags.List() {
		payload.Data.Attributes.Tags = append(payload.Data.Attributes.Tags, tag.(string))
	}
	_, err := client.UpdateAccount(accountId, payload)
	if err != nil {
		return err
	}
	return nil
}

func updateAccountSettings(provider string, accountId string, d *schema.ResourceData, client *cloudconformity.Client) error {
	if i, ok := d.GetOk("settings"); ok && len(i.(*schema.Set).List()) > 0 {
		settings := i.(*schema.Set).List()[0].(map[string]interface{})

		if v, ok := settings["bot"]; ok && len(v.(*schema.Set).List()) > 0 {
			bot := v.(*schema.Set).List()[0].(map[string]interface{})

			botRequest := cloudconformity.AccountBotSettingsRequest{}
			botRequest.Data.Attributes.Settings.Bot = &cloudconformity.AccountBot{}
			botRequest.Data.Type = "accounts"
			delay := bot["delay"].(int)
			disabled := bot["disabled"].(bool)

			botRequest.Data.Attributes.Settings.Bot.Delay = &delay
			botRequest.Data.Attributes.Settings.Bot.Disabled = &disabled
			botRequest.Data.Attributes.Settings.Bot.DisabledRegions = processBotDisabledRegions(bot["disabled_regions"].(*schema.Set).List())

			_, err := client.UpdateAccountBotSettings(accountId, botRequest)
			if err != nil {
				// TODO: show an warning instead of error
				return err
			}
		}

		if v, ok := settings["rule"]; ok && len(v.([]interface{})) > 0 {

			for _, item := range v.([]interface{}) {
				rule := item.(map[string]interface{})
				ruleRequest := &cloudconformity.AccountRuleSettings{}

				ruleId := rule["rule_id"].(string)
				ruleRequest.Data.Attributes.Note = "Note automatically added by provider"

				ruleRequest.Data.Attributes.RuleSetting = processRuleSetting(provider, rule["settings"].(*schema.Set).List(), ruleId)
				_, err := client.UpdateAccountRuleSettings(accountId, ruleId, ruleRequest)

				if err != nil {
					// TODO: show an warning instead of error
					return err
				}
			}

		}
	}
	return nil
}

func processBotDisabledRegions(list []interface{}) *cloudconformity.BotDisabledRegions {
	regions := &cloudconformity.BotDisabledRegions{}

	for _, v := range list {
		region, ok := v.(string)
		if ok && region != "" {
			setDisabledRegion(regions, region)
		}
	}
	return regions
}

func setDisabledRegion(regions *cloudconformity.BotDisabledRegions, region string) {
	v := reflect.ValueOf(regions).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		f := v.FieldByName(field.Name)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		jsonTagFieldOnly := strings.Split(jsonTag, ",")[0]

		if jsonTagFieldOnly == region {
			f.SetBool(true)
		}
	}
}

func processRuleSetting(provider string, rule []interface{}, ruleId string) cloudconformity.RuleSetting {

	ruleSetting := cloudconformity.RuleSetting{}
	if len(rule) == 0 {
		return ruleSetting
	}

	item := rule[0].(map[string]interface{})

	ruleSetting.Enabled = item["enabled"].(bool)
	ruleSetting.RuleExists = item["rule_exists"].(bool)
	if exceptions := item["exceptions"].(*schema.Set).List(); len(exceptions) > 0 {
		ruleSetting.Exceptions = processRuleExceptions(item["exceptions"].(*schema.Set).List())
	}
	if extraSetting := item["extra_settings"].(*schema.Set).List(); len(extraSetting) > 0 {
		ruleSetting.ExtraSettings = processRuleExtraSettings(item["extra_settings"].(*schema.Set).List())
	}

	ruleSetting.Id = ruleId

	ruleSetting.Provider = provider
	ruleSetting.RiskLevel = item["risk_level"].(string)
	return ruleSetting
}

func processRuleExceptions(exceptions []interface{}) cloudconformity.RuleSettingExceptions {

	rse := cloudconformity.RuleSettingExceptions{}
	if len(exceptions) == 0 {
		return rse
	}
	item := exceptions[0].(map[string]interface{})
	rse.FilterTags = expandStringList(item["filter_tags"].(*schema.Set).List())
	rse.Resources = expandStringList(item["resources"].(*schema.Set).List())
	rse.Tags = expandStringList(item["tags"].(*schema.Set).List())

	return rse
}

func processRuleExtraSettings(es []interface{}) []cloudconformity.RuleSettingExtra {

	extraSetting := make([]cloudconformity.RuleSettingExtra, len(es))
	for i, v := range es {
		item := v.(map[string]interface{})
		extraSetting[i].Name = item["name"].(string)
		extraSetting[i].Type = item["type"].(string)

		// check when to use `value` and `values` base on the type
		// single-number-value, ttl, single-value-regex, single-string-value - uses `value` and the other type uses `values`
		switch extraSetting[i].Type {

		case "single-string-value", "single-number-value", "ttl", "single-value-regex":

			extraSetting[i].Value = item["value"].(string)

		case "regions":

			extraSetting[i].Values = expandStringList(item["regions"].(*schema.Set).List())
			regions := true
			extraSetting[i].Regions = &regions

		case "multiple-object-values":

			extraSetting[i].Values = processRuleMultipleIp(item["multiple_object_values"].(*schema.Set).List())
			valueKeys := []string{"eventName", "eventSource", "userIdentityType"}
			extraSetting[i].ValueKeys = &valueKeys

		case "multiple-vpc-gateway-mappings":

			extraSetting[i].Mappings = processRuleMapping(item["mappings"].(*schema.Set).List())

		default:

			extraSetting[i].Values = processRuleValues(item["values"].(*schema.Set).List())

		}
	}

	return extraSetting
}
func processRuleValues(vs []interface{}) []*cloudconformity.RuleSettingValues {
	values := make([]*cloudconformity.RuleSettingValues, 0, len(vs))
	for _, v := range vs {

		profileValues := &cloudconformity.RuleSettingValues{}
		item := v.(map[string]interface{})

		profileValues.Enabled = item["enabled"].(bool)
		profileValues.Label = item["label"].(string)
		profileValues.Value = item["value"].(string)
		values = append(values, profileValues)
	}
	return values
}

func processRuleMultipleIp(mo []interface{}) []*cloudconformity.RuleSettingMultipleObject {

	multiple := make([]*cloudconformity.RuleSettingMultipleObject, 0, len(mo))

	for _, v := range mo {

		multipleObject := &cloudconformity.RuleSettingMultipleObject{}
		item := v.(map[string]interface{})

		multipleObject.Value.EventName = item["event_name"].(string)
		multipleObject.Value.EventSource = item["event_source"].(string)
		multipleObject.Value.UserIdentityType = item["user_identity_type"].(string)
		multiple = append(multiple, multipleObject)
	}
	return multiple
}

func processRuleMapping(m []interface{}) []*cloudconformity.RuleSettingMapping {

	mappings := make([]*cloudconformity.RuleSettingMapping, 0, len(m))
	for _, v := range m {

		ruleMapping := &cloudconformity.RuleSettingMapping{}
		item := v.(map[string]interface{})
		ruleMapping.Values = processMappingValues(item["values"].(*schema.Set).List())
		mappings = append(mappings, ruleMapping)
	}
	return mappings
}
func processMappingValues(mv []interface{}) []*cloudconformity.MappingValues {
	MappingValues := make([]*cloudconformity.MappingValues, 0, len(mv))
	for _, v := range mv {
		item := v.(map[string]interface{})
		values := &cloudconformity.MappingValues{}
		values.Name = item["name"].(string)
		values.Type = item["type"].(string)

		if values.Type == "single-string-value" {

			values.Value = item["value"].(string)

		} else {

			values.Values = processValues(item["values"].(*schema.Set).List())

		}

		MappingValues = append(MappingValues, values)
	}
	return MappingValues
}
func processValues(val []interface{}) interface{} {

	type values struct {
		Value string `json:"value"`
	}

	mv := make([]*values, 0, len(val))
	for _, v := range val {
		item := v.(map[string]interface{})
		value := &values{}
		value.Value = item["value"].(string)
		mv = append(mv, value)
	}

	return mv
}
