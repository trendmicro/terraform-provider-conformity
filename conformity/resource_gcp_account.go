package conformity

import (
	"context"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGCPAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGCPAccountCreate,
		ReadContext:   resourceGCPAccountRead,
		UpdateContext: resourceGCPAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_account_unique_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"settings": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bot": BotSettingsSchema(),
						"rule": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rule_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"settings": {
										Type:     schema.TypeSet,
										Required: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"rule_exists": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"exceptions":     ExceptionsSchema(),
												"extra_settings": ExtraSettingSchema(),
												"risk_level": {
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validation.StringInSlice([]string{"LOW", "MEDIUM", "HIGH", "VERY_HIGH", "EXTREME"}, false),
												},
											},
										},
									},
								},
							},
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

func resourceGCPAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := cloudconformity.AccountPayload{}
	payload.Data.Attributes.Name = d.Get("name").(string)
	payload.Data.Attributes.Environment = d.Get("environment").(string)
	payload.Data.Attributes.Access.ProjectId = d.Get("project_id").(string)
	payload.Data.Attributes.Access.ProjectName = d.Get("project_name").(string)
	payload.Data.Attributes.Access.ServiceAccountUniqueId = d.Get("service_account_unique_id").(string)

	accountId, err := client.CreateGCPAccount(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	if v, ok := d.GetOk("tags"); ok && len(v.(*schema.Set).List()) > 0 {
		err := updateAccountTags(payload, accountId, d, client)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	err = updateAccountSettings("gcp", accountId, d, client)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(accountId)
	resourceGCPAccountRead(ctx, d, m)
	return diags
}

func resourceGCPAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	accountId := d.Id()

	// get both account details and access settings
	accountAccessAndDetails, err := client.GetGCPAccount(accountId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", accountAccessAndDetails.AccountDetails.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment", accountAccessAndDetails.AccountDetails.Data.Attributes.Environment); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tags", accountAccessAndDetails.AccountDetails.Data.Attributes.Tags); err != nil {
		return diag.FromErr(err)
	}
	if accountAccessAndDetails.AccountDetails.Data.Attributes.Settings == nil {
		if err := d.Set("settings", nil); err != nil {
			return diag.FromErr(err)
		}
	}else {
		settings := flattenAccountSettings(accountAccessAndDetails.AccountDetails.Data.Attributes.Settings, accountAccessAndDetails.RuleSettings.Data.Attributes.Settings.Rules)
		if err := d.Set("settings", settings); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGCPAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	accountId := d.Id()
	if d.HasChange("name") || d.HasChange("environment") || d.HasChange("tags") {

		payload := cloudconformity.AccountPayload{}
		payload.Data.Attributes.Name = d.Get("name").(string)
		payload.Data.Attributes.Environment = d.Get("environment").(string)

		tags := d.Get("tags").(*schema.Set)
		for _, tag := range tags.List() {
			payload.Data.Attributes.Tags = append(payload.Data.Attributes.Tags, tag.(string))
		}

		_, err := client.UpdateAccount(accountId, payload)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("settings") {
		err := updateAccountSettings("GCP", accountId, d, client)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChange("role_arn") || d.HasChange("external_id") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to Update Conformity GCP account",
			Detail:   "'role_arn' and 'external_id' cannot be changed",
		})

		return diags
	}

	return resourceGCPAccountRead(ctx, d, m)
}
