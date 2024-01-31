package conformity

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
)

type ObjectValue struct {
	Days     int    `json:"days"`
	Operator string `json:"operator"`
}

func resourceConformityCustomRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformityCustomRuleCreate,
		ReadContext:   resourceConformityCustomRuleRead,
		UpdateContext: resourceConformityCustomRuleUpdate,
		DeleteContext: resourceConformityCustomRuleDelete,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_provider": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"aws",
					"azure",
					"gcp",
				}, false),
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"remediation_notes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"LOW",
					"MEDIUM",
					"HIGH",
					"VERY_HIGH",
					"EXTREME",
				}, false),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"categories": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"security",
						"cost-optimisation",
						"reliability",
						"performance-efficiency",
						"operational-excellence",
						"sustainability",
					}, false),
				},
			},
			"attributes": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"required": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"rules": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operation": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"any",
								"all",
							}, false),
						},
						"conditions": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fact": {
										Type:     schema.TypeString,
										Required: true,
									},
									"operator": {
										Type:     schema.TypeString,
										Required: true,
									},
									"path": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"event_type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceConformityCustomRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	payload := cloudconformity.CustomRuleRequest{}

	// Required fields
	payload.Name = d.Get("name").(string)
	payload.Provider = d.Get("cloud_provider").(string)
	payload.Description = d.Get("description").(string)
	payload.Service = d.Get("service").(string)
	payload.ResourceType = d.Get("resource_type").(string)
	payload.Severity = d.Get("severity").(string)
	payload.Enabled = d.Get("enabled").(bool)
	payload.Categories = expandStringList(d.Get("categories").(*schema.Set).List())
	if v, ok := d.GetOk("attributes"); ok && len(v.(*schema.Set).List()) > 0 {
		processInputCustomRuleAttributes(&payload, d)
	}
	if v, ok := d.GetOk("rules"); ok && len(v.(*schema.Set).List()) > 0 {
		processInputCustomRuleRules(&payload, d)
	}
	// Optional fields
	if d.Get("remediation_notes") != "" {
		payload.RemediationNotes = d.Get("remediation_notes").(string)
	}

	customRule, err := client.CreateConformityCustomRule(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(customRule.ID)

	resourceConformityCustomRuleRead(ctx, d, m)
	return diags
}

func resourceConformityCustomRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	rule, err := client.GetCustomRule(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(rule.ID)

	if err := d.Set("type", rule.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", rule.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("cloud_provider", rule.Attributes.Provider); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", rule.Attributes.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("remediation_notes", rule.Attributes.RemediationNotes); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("service", rule.Attributes.Service); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_type", rule.Attributes.ResourceType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("severity", rule.Attributes.Severity); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", rule.Attributes.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("categories", rule.Attributes.Categories); err != nil {
		return diag.FromErr(err)
	}
	attributes := flattenAttributes(rule)
	if err := d.Set("attributes", attributes); err != nil {
		return diag.FromErr(err)
	}
	rules := flattenRules(rule)
	if err := d.Set("rules", rules); err != nil {
		return diag.FromErr(err)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceConformityCustomRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type

	payload := cloudconformity.CustomRuleRequest{}
	// Required fields
	payload.Name = d.Get("name").(string)
	payload.Provider = d.Get("cloud_provider").(string)
	payload.Description = d.Get("description").(string)
	payload.Service = d.Get("service").(string)
	payload.ResourceType = d.Get("resource_type").(string)
	payload.Severity = d.Get("severity").(string)
	payload.Enabled = d.Get("enabled").(bool)
	payload.Categories = expandStringList(d.Get("categories").(*schema.Set).List())
	if v, ok := d.GetOk("attributes"); ok && len(v.(*schema.Set).List()) > 0 {
		processInputCustomRuleAttributes(&payload, d)
	}
	if v, ok := d.GetOk("rules"); ok && len(v.(*schema.Set).List()) > 0 {
		processInputCustomRuleRules(&payload, d)
	}
	// Optional fields
	if d.Get("remediation_notes") != "" {
		payload.RemediationNotes = d.Get("remediation_notes").(string)
	}
	id := d.Id()
	_, err := client.UpdateCustomRule(id, payload)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceConformityCustomRuleRead(ctx, d, m)
}

func resourceConformityCustomRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	id := d.Id()

	_, err := client.DeleteCustomRule(id)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.

	d.SetId("")
	return diags
}

// converts terraform state field `attributes` into api object
func processInputCustomRuleAttributes(payload *cloudconformity.CustomRuleRequest, d *schema.ResourceData) {
	attrsIn := d.Get("attributes").(*schema.Set).List()
	attrsOut := make([]cloudconformity.CustomRuleAttributes, len(attrsIn))
	for i, attributes := range attrsIn {
		m := attributes.(map[string]interface{})
		obj := cloudconformity.CustomRuleAttributes{}
		obj.Path = m["path"].(string)
		obj.Name = m["name"].(string)
		obj.Required = m["required"].(bool)
		attrsOut[i] = obj
	}
	payload.Attributes = attrsOut
}

// converts terraform state field `rules` into api object
func processInputCustomRuleRules(payload *cloudconformity.CustomRuleRequest, d *schema.ResourceData) {
	rulesIn := d.Get("rules").(*schema.Set).List()
	rulesOut := make([]cloudconformity.CustomRuleRules, len(rulesIn))
	for i, rules := range rulesIn {
		m := rules.(map[string]interface{})
		obj := cloudconformity.CustomRuleRules{}
		objEvent := cloudconformity.CustomRuleEvent{}
		objEvent.Type = m["event_type"].(string)
		obj.Event = objEvent
		objConditions := processInputCustomRuleConditions(m["conditions"].(*schema.Set).List())
		if m["operation"].(string) == "any" {
			obj.Conditions.Any = objConditions
		} else if m["operation"].(string) == "all" {
			obj.Conditions.All = objConditions
		}
		rulesOut[i] = obj
	}
	payload.Rules = rulesOut
}

// converts terraform state field `conditions` into api object
func processInputCustomRuleConditions(conditionsIn []interface{}) []cloudconformity.CustomRuleCondition {
	conditionsOut := make([]cloudconformity.CustomRuleCondition, len(conditionsIn))
	for i, conditions := range conditionsIn {
		m := conditions.(map[string]interface{})
		obj := cloudconformity.CustomRuleCondition{}
		obj.Fact = m["fact"].(string)
		obj.Path = m["path"].(string)
		/*
			Custom Rule Conditions has an attribute of `value` that can accept a
			string, boolean, integer, or an object. Anything other than string needs
			to be encoded using the built-in Terraform function `jsonencode()`.
			Below we are assigning objValue with an instance of the ObjectValue struct
			that defines the variables that the Custom Rules API will accept.
		*/
		objValue := ObjectValue{}
		if strings.ToLower(m["value"].(string)) == "true" || strings.ToLower(m["value"].(string)) == "false" {
			obj.Value, _ = strconv.ParseBool(m["value"].(string))
		} else if numValue, err := strconv.Atoi(m["value"].(string)); err == nil {
			obj.Value = numValue
		} else if err := json.Unmarshal([]byte(m["value"].(string)), &objValue); err == nil {
			obj.Value = objValue
		} else {
			obj.Value = m["value"]
		}

		if operator, ok := m["operator"]; ok {
			obj.Operator = operator.(string)
		}
		conditionsOut[i] = obj
	}
	return conditionsOut
}

// converts api object into terraform state field `attributes`
func flattenAttributes(rule *cloudconformity.CustomRuleResponse) []interface{} {
	attrsOut := make([]interface{}, len(rule.Attributes.Attributes))
	for i, attributes := range rule.Attributes.Attributes {
		m := make(map[string]interface{})
		m["name"] = attributes.Name
		m["path"] = attributes.Path
		m["required"] = attributes.Required
		attrsOut[i] = m
	}
	return attrsOut
}

// converts api object into terraform state field `rules`
func flattenRules(rule *cloudconformity.CustomRuleResponse) []interface{} {
	rulesOut := make([]interface{}, len(rule.Attributes.Rules))
	for i, rules := range rule.Attributes.Rules {
		m := make(map[string]interface{})
		m["event_type"] = rules.Event.Type
		if len(rules.Conditions.All) > 0 {
			m["operation"] = "all"
			m["conditions"] = flattenConditions(rules.Conditions.All)
		} else if len(rules.Conditions.Any) > 0 {
			m["operation"] = "any"
			m["conditions"] = flattenConditions(rules.Conditions.Any)
		}
		rulesOut[i] = m
	}
	return rulesOut
}

// converts api object into terraform state field `conditions`
func flattenConditions(conditionsIn []cloudconformity.CustomRuleCondition) []interface{} {
	conditionsOut := make([]interface{}, len(conditionsIn))
	for i, conditions := range conditionsIn {
		m := make(map[string]interface{})
		m["fact"] = conditions.Fact
		m["operator"] = conditions.Operator
		m["path"] = conditions.Path
		conditionsValueByte, _ := json.Marshal(conditions.Value)
		m["value"] = string(conditionsValueByte)
		conditionsOut[i] = m
	}
	return conditionsOut
}
