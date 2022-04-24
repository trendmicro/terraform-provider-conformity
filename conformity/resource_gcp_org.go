package conformity

import (
	"context"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGCPOrg() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGCPOrgCreate,
		ReadContext:   resourceGCPOrgRead,
		UpdateContext: resourceGCPOrgUpdate,
		DeleteContext: resourceOrgDelete,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed-group-id": {
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
			"attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
                        "serviceAccountName": {
                            Type:     schema.TypeString,
                            Required: true,
                                },
                        "serviceAccountUniqueId": {
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

func resourceGCPOrgCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := cloudconformity.OrgPayload{}
	payload.Data.serviceAccountName = d.Get("name").(string)
	payload.Data.serviceAccountKeyJson.type = d.Get("type").(string)
	payload.Data.serviceAccountKeyJson.project_id = d.Get("project_id").(string)
	payload.Data.serviceAccountKeyJson.private_key_id = d.Get("private_key_id").(string)
	payload.Data.serviceAccountKeyJson.private_key = d.Get("private_key").(string)
	payload.Data.serviceAccountKeyJson.client_email = d.Get("client_email").(string)
	payload.Data.serviceAccountKeyJson.client_id = d.Get("client_id").(string)
	payload.Data.serviceAccountKeyJson.auth_uri = d.Get("auth_uri").(string)
	payload.Data.serviceAccountKeyJson.token_uri = d.Get("token_uri").(string)
	payload.Data.serviceAccountKeyJson.auth_provider_x509_cert_url = d.Get("auth_provider_x509_cert_url").(string)
	payload.Data.serviceAccountKeyJson.client_x509_cert_url = d.Get("client_x509_cert_url").(string)


	organisationId, err := client.CreateGCPOrg(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(organisationId)
	resourceGCPOrgRead(ctx, d, m)
	return diags
}

func resourceGCPOrgRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
}

func resourceGCPOrgUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

}

func resourceOrgDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

}
