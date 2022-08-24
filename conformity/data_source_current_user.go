package conformity

import (
	"context"

	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	// "regexp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	// "log"
	// "strconv"
	// "time"
)

func dataSourceCurrentUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCurrentUserRead,
		Schema: map[string]*schema.Schema{
			"is_cloud_one_user": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_api_key_user": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"created_date": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"summary_email_opt_out": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_credentials": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mfa": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_login_date": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}
func dataSourceCurrentUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)

	var diags diag.Diagnostics
	data, err := client.GetCurrentUser()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_cloud_one_user", data.Data.Attributes.IsCloudOneUser); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_api_key_user", data.Data.Meta.IsApiKeyUser); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_date", data.Data.Attributes.CreatedDate); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("summary_email_opt_out", data.Data.Attributes.SummaryEmailOptOut); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("has_credentials", data.Data.Attributes.HasCredentials); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("first_name", data.Data.Attributes.FirstName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_name", data.Data.Attributes.LastName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("role", data.Data.Attributes.Role); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("email", data.Data.Attributes.Email); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mfa", data.Data.Attributes.Mfa); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_login_date", data.Data.Attributes.LastLogIn); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", data.Data.Type); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.Data.ID)
	return diags
}
