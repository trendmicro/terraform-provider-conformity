package conformity

import (
	"context"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{

			"apikey": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APIKEY", nil),
			},
			// if region is not specify us-west-2 will be use
			"region": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "us-west-2",
				ValidateFunc: validation.StringInSlice([]string{"eu-west-1", "us-west-2", "ap-southeast-2", "us-1",
				                                                "in-1", "gb-1", "jp-1", "de-1", "au-1", "ca-1",
				                                                "sg-1"}, true),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"conformity_aws_account":           resourceAwsAccount(),
			"conformity_azure_account":         resourceAzureAccount(),
			"conformity_gcp_account":           resourceGCPAccount(),
			"conformity_gcp_org":               resourceGCPOrg(),
			"conformity_group":                 resourceConformityGroup(),
			"conformity_user":                  resourceConformityLegacyUser(),
			"conformity_sso_user":              resourceConformitySsoLegacyUser(),
			"conformity_report_config":         resourceConformityReportConfig(),
			"conformity_communication_setting": resourceConformityCommSetting(),
			"conformity_profile":               resourceConformityProfile(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"conformity_external_id":   dataSourceExternalId(),
			"conformity_apply_profile": dataSourceApplyProfile(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	region := d.Get("region").(string)
	apiKey := d.Get("apikey").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c, err := cloudconformity.NewClient(region, apiKey)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Conformity client",
			Detail:   err.Error(),
		})

		return nil, diags
	}

	return c, diags

}
