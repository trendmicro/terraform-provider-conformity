package conformity

import (
	"context"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAzureAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAzureAccountCreate,
		ReadContext:   resourceAzureAccountRead,
		UpdateContext: resourceAzureAccountUpdate,
		DeleteContext: resourceAccountDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"active_directory_id": {
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
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAzureAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type

	var diags diag.Diagnostics
	payload := cloudconformity.AccountPayload{}
	payload.Data.Attributes.Name = d.Get("name").(string)
	payload.Data.Attributes.Environment = d.Get("environment").(string)
	payload.Data.Attributes.Access.SubscriptionId = d.Get("subscription_id").(string)
	payload.Data.Attributes.Access.ActiveDirectoryId = d.Get("active_directory_id").(string)
	accountId, err := client.CreateAzureAccount(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(accountId)
	resourceAzureAccountRead(ctx, d, m)
	return diags
}

func resourceAzureAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	accountId := d.Id()

	// get both account details
	accountDetails, err := client.GetAccountDetails(accountId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", accountDetails.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment", accountDetails.Data.Attributes.Environment); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("subscription_id", accountDetails.Data.Attributes.CloudData.Azure.SubscriptionId); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceAzureAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	if d.HasChange("name") || d.HasChange("environment") || d.HasChange("tags") {
		payload := cloudconformity.AccountPayload{}
		payload.Data.Attributes.Name = d.Get("name").(string)
		payload.Data.Attributes.Environment = d.Get("environment").(string)

		tags := d.Get("tags").(*schema.Set)
		for _, tag := range tags.List() {
			payload.Data.Attributes.Tags = append(payload.Data.Attributes.Tags, tag.(string))
		}

		accountId := d.Id()
		_, err := client.UpdateAccount(accountId, payload)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	// by activating the Active Directory Id need to add active_directory_id to this statement
	// for now create function is not yet working so we need to remove active_directory_id for the if statement
	if d.HasChange("subscription_id") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to Update Conformity Azure account",
			Detail:   "'subscription_id' cannot be changed",
		})

		return diags
	}
	return resourceAzureAccountRead(ctx, d, m)
}
