package conformity

import (
	"context"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
//     "log"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGCPOrg() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGCPOrgCreate,
		ReadContext:   resourceGCPOrgRead,
		UpdateContext: resourceGCPOrgUpdate,
		DeleteContext: resourceOrgDelete,
		Schema: map[string]*schema.Schema{
			"service_account_name": {
				Type:     schema.TypeString,
				Required: true,
			},
            "type": {
                Type:     schema.TypeString,
                Required: true,
            },
            "project_id": {
                Type:     schema.TypeString,
                Required: true,
            },
            "private_key_id": {
                Type:     schema.TypeString,
                Required: true,
            },
            "private_key": {
                Type:     schema.TypeString,
                Required: true,
                Sensitive:   true,
            },
            "client_email": {
                Type:     schema.TypeString,
                Required: true,
            },
            "client_id": {
                Type:     schema.TypeString,
                Required: true,
            },
            "auth_uri": {
                Type:     schema.TypeString,
                Required: true,
            },
            "token_uri": {
                Type:     schema.TypeString,
                Required: true,
            },
            "auth_provider_x509_cert_url": {
                Type:     schema.TypeString,
                Required: true,
            },
            "client_x509_cert_url": {
                Type:     schema.TypeString,
                Required: true,
            },
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceGCPOrgCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type`
	var diags diag.Diagnostics
    payload := cloudconformity.GCPOrgPayload{}
	payload.Data.ServiceAccountName = d.Get("service_account_name").(string)
	payload.Data.ServiceAccountKeyJson.ProjectId = d.Get("project_id").(string)
	payload.Data.ServiceAccountKeyJson.Type = d.Get("type").(string)
	payload.Data.ServiceAccountKeyJson.PrivateKeyId = d.Get("private_key_id").(string)
	payload.Data.ServiceAccountKeyJson.PrivateKey = d.Get("private_key").(string)
	payload.Data.ServiceAccountKeyJson.ClientEmail = d.Get("client_email").(string)
	payload.Data.ServiceAccountKeyJson.ClientId = d.Get("client_id").(string)
	payload.Data.ServiceAccountKeyJson.AuthUri = d.Get("auth_uri").(string)
	payload.Data.ServiceAccountKeyJson.TokenUri = d.Get("token_uri").(string)
	payload.Data.ServiceAccountKeyJson.AuthProviderX509CertUrl = d.Get("auth_provider_x509_cert_url").(string)
	payload.Data.ServiceAccountKeyJson.ClientX509CertUrl = d.Get("client_x509_cert_url").(string)

	orgResponse, err := client.CreateGCPOrg(payload)

	if err != nil {
		return diag.FromErr(err)
	}

    d.SetId(orgResponse)
 	resourceGCPOrgRead(ctx, d, m)
	return diags
}

func resourceGCPOrgRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags
}

func resourceGCPOrgUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags

}

func resourceOrgDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags
}
