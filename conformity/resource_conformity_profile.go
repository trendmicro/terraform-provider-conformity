package conformity

import (
	"context"
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceConformityProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformityProfileCreate,
		ReadContext:   resourceConformityProfileRead,
		UpdateContext: resourceConformityProfileUpdate,
		DeleteContext: resourceConformityProfileDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"included": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"exceptions": {
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// need to verify which one it depricated both filter_tag and tags
									"filter_tags": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"resources": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"tags": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"extra_settings": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"countries": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"multiple": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"regions": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice([]string{"regions", "multiple-string-values", "multiple-number-values", "multiple-aws-account-values",
											"choice-multiple-value", "choice-single-value", "single-number-value", "single-string-value", "ttl", "single-value-regex", "tags"}, true),
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
									// changes schema to allow []string
									"values_array": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"values": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// attributes `value` is commonly use for all the types therefore It should be required
												// eg. [{"value": "ELBSecurityPolicy-2016-08"}]
												"label": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"enabled": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"provider": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"aws", "azure", "gcp"}, true),
						},
						"risk_level": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"LOW", "MEDIUM", "HIGH", "VERY_HIGH", "EXTREME"}, false),
						},
					},
				},
			},
		},
		CustomizeDiff: customdiff.All(
			CustomizeDiffValidateProfileValue,
		),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceConformityProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := cloudconformity.ProfileSettings{}
	payload.Data.Attributes.Name = d.Get("name").(string)
	payload.Data.Attributes.Description = d.Get("description").(string)
	payload.Data.Type = "profiles"

	if v, ok := d.GetOk("included"); ok && len(v.([]interface{})) > 0 {
		proccessProfileIncluded(&payload, d)
		proccessRelationships(&payload, d)
	}

	profileId, err := client.CreateProfileSetting(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(profileId)
	resourceConformityProfileRead(ctx, d, m)
	return diags
}

func resourceConformityProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	profileId := d.Id()

	profileSettings, err := client.GetProfile(profileId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", profileSettings.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", profileSettings.Data.Attributes.Description); err != nil {
		return diag.FromErr(err)
	}
	included := flattenProfileIncluded(profileSettings.Included)
	if err := d.Set("included", included); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceConformityProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	payload := cloudconformity.ProfileSettings{}
	profileId := d.Id()
	payload.Data.ID = profileId
	payload.Data.Attributes.Name = d.Get("name").(string)
	payload.Data.Attributes.Description = d.Get("description").(string)
	payload.Data.Type = "profiles"

	if d.HasChange("included") {

		if v, ok := d.GetOk("included"); ok && len(v.([]interface{})) > 0 {
			proccessProfileIncluded(&payload, d)
			proccessRelationships(&payload, d)
		}

	}

	_, err := client.UpdateProfile(profileId, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceConformityProfileRead(ctx, d, m)
}

func resourceConformityProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	profileId := d.Id()

	_, err := client.DeleteProfile(profileId)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.

	d.SetId("")
	return diags
}

func proccessProfileIncluded(payload *cloudconformity.ProfileSettings, d *schema.ResourceData) {
	included := d.Get("included").([]interface{})
	c := make([]cloudconformity.ProfileIncluded, len(included))
	for i, v := range included {

		item := v.(map[string]interface{})
		c[i].ID = item["id"].(string)
		c[i].Type = "rules"
		c[i].Attributes.Provider = item["provider"].(string)
		c[i].Attributes.Enabled = item["enabled"].(bool)
		c[i].Attributes.RiskLevel = item["risk_level"].(string)

		if len(item["exceptions"].(*schema.Set).List()) != 0 {
			proccessExceptions(item["exceptions"].(*schema.Set).List()[0].(map[string]interface{}), &c[i].Attributes)
		}
		if len(item["extra_settings"].(*schema.Set).List()) != 0 {
			proccessExtraSettings(item["extra_settings"].(*schema.Set).List(), &c[i].Attributes)

		}

	}
	payload.Included = c

}

func proccessRelationships(payload *cloudconformity.ProfileSettings, d *schema.ResourceData) {

	included := d.Get("included").([]interface{})
	c := make([]cloudconformity.RuleSettingsData, len(included))
	for i, v := range included {
		item := v.(map[string]interface{})
		c[i].ID = item["id"].(string)
		c[i].Type = "rules"
	}

	payload.Data.Relationships.RuleSettings.Data = c
}

func proccessExceptions(v map[string]interface{}, e *cloudconformity.IncludedAttributes) {
	exceptions := cloudconformity.IncludedExceptions{}
	exceptions.FilterTags = expandStringList(v["filter_tags"].(*schema.Set).List())
	exceptions.Resources = expandStringList(v["resources"].(*schema.Set).List())
	exceptions.Tags = expandStringList(v["tags"].(*schema.Set).List())
	e.Exceptions = &exceptions
}

func proccessExtraSettings(v []interface{}, a *cloudconformity.IncludedAttributes) {

	c := make([]cloudconformity.IncludedExtraSettings, len(v))
	for i, es := range v {
		item := es.(map[string]interface{})
		c[i].Countries = item["countries"].(bool)
		c[i].Multiple = item["multiple"].(bool)
		c[i].Name = item["name"].(string)
		c[i].Regions = item["regions"].(bool)
		c[i].Type = item["type"].(string)
		// check when to use `value` and `values` base on the type
		// single-number-value, ttl, single-value-regex, single-string-value - uses `value` and the other type uses `values`
		if c[i].Type == "single-string-value" || c[i].Type == "single-number-value" || c[i].Type == "ttl" || c[i].Type == "single-value-regex" || c[i].Type == "choice-single-value" {
			c[i].Value = item["value"].(string)
		} else if c[i].Type == "regions" {
			processProfileValuesStringsSlices(item["values_array"].(*schema.Set).List(), &c[i])
		} else {
			processProfileValues(item["values"].(*schema.Set).List(), &c[i])
		}

	}
	a.ExtraSettings = c
}

func processProfileValues(v []interface{}, ies *cloudconformity.IncludedExtraSettings) {
	c := make([]interface{}, 0, len(v))

	for _, values := range v {
		profileValues := cloudconformity.ProfileValues{}

		item := values.(map[string]interface{})

		profileValues.Enabled = item["enabled"].(bool)
		profileValues.Label = item["label"].(string)
		profileValues.Value = item["value"].(string)

		c = append(c, profileValues)
	}

	ies.Values = c
}

func processProfileValuesStringsSlices(v []interface{}, ies *cloudconformity.IncludedExtraSettings) {
	c := make([]interface{}, 0, len(v))

	for _, value := range v {
		c = append(c, value.(string))
	}

	ies.Values = c
}

func flattenProfileIncluded(included []cloudconformity.ProfileIncluded) []interface{} {

	if included == nil {
		return make([]interface{}, 0)
	}

	pis := make([]interface{}, len(included))
	for i, includedItem := range included {

		pi := make(map[string]interface{})
		pi["id"] = includedItem.ID
		pi["enabled"] = includedItem.Attributes.Enabled
		pi["provider"] = includedItem.Attributes.Provider
		pi["risk_level"] = includedItem.Attributes.RiskLevel

		if includedItem.Attributes.Exceptions != nil {
			pi["exceptions"] = flattenProfileExceptions(includedItem.Attributes.Exceptions)
		}

		pi["extra_settings"] = flattenProfileExtraSettings(includedItem.Attributes.ExtraSettings)
		pis[i] = pi
	}
	return pis

}
func flattenProfileExceptions(exceptions *cloudconformity.IncludedExceptions) []interface{} {

	c := make(map[string]interface{})

	c["filter_tags"] = exceptions.FilterTags
	c["resources"] = exceptions.Resources
	c["tags"] = exceptions.Tags

	return []interface{}{c}
}

func flattenProfileExtraSettings(extra []cloudconformity.IncludedExtraSettings) []interface{} {

	if extra == nil {
		return make([]interface{}, 0)
	}
	ess := make([]interface{}, len(extra))

	for i, extraSettings := range extra {
		es := make(map[string]interface{})

		es["countries"] = extraSettings.Countries
		es["multiple"] = extraSettings.Multiple
		es["name"] = extraSettings.Name
		es["regions"] = extraSettings.Regions
		es["type"] = extraSettings.Type
		// check when to use `value` and `values` base on the type
		// single-number-value, ttl, single-value-regex, single-string-value - uses `value` and the other type uses `values`
		if extraSettings.Type == "single-string-value" ||
			extraSettings.Type == "single-number-value" ||
			extraSettings.Type == "ttl" ||
			extraSettings.Type == "single-value-regex" ||
			extraSettings.Type == "choice-single-value" {
			es["value"] = fmt.Sprintf("%s", extraSettings.Value)
		} else if extraSettings.Type == "regions" {
			es["values_array"] = extraSettings.Values
		} else {
			es["values"] = flattenProfileValues(extraSettings.Values)
		}

		ess[i] = es
	}
	return ess
}

func flattenProfileValues(values []interface{}) []interface{} {
	if values == nil {
		return make([]interface{}, 0)
	}

	vs := make([]interface{}, 0, len(values))

	for _, value := range values {
		// handle case when ProfileValues
		castedVal := value.(map[string]interface{})
		v := make(map[string]interface{})
		v["enabled"] = castedVal["enabled"]
		v["label"] = castedVal["label"]
		v["value"] = castedVal["value"]
		vs = append(vs, v)
	}
	return vs
}
