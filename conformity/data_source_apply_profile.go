package conformity

import (
	"context"
	"errors"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceApplyProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplyProfileCreateRead,
		Schema: map[string]*schema.Schema{
			"profile_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mode": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"fill-gaps", "overwrite", "replace"}, true),
			},
			"notes": {
				Type:     schema.TypeString,
				Required: true,
			},
			"include": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exceptions": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceApplyProfileCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := cloudconformity.ApplyProfileSettings{}
	profileId := d.Get("profile_id").(string)
	payload.Meta.AccountIds = expandStringList(d.Get("account_ids").(*schema.Set).List())
	payload.Meta.Mode = d.Get("mode").(string)
	payload.Meta.Notes = d.Get("notes").(string)
	// [Only rule is supported in the current version] An Array of setting types to be applied to the accounts
	payload.Meta.Types = []string{"rule"}

	// Include exceptions functionality is only available in `overwrite` mode
	err := validateApplyProfileExceptions(d)
	if err != nil {
		return diag.FromErr(err)
	}
	if v, ok := d.GetOk("include"); ok && len(v.(*schema.Set).List()) > 0 {

		include := cloudconformity.ApplyProfileInclude{}

		inc := d.Get("include").(*schema.Set).List()[0].(map[string]interface{})
		include.Exceptions = inc["exceptions"].(bool)

		payload.Meta.Include = &include
	}

	applyProfileResponse, err := client.CreateApplyProfile(profileId, payload)
	if err != nil {
		return diag.FromErr(err)
	}
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Datasource - Conformity_apply_profile",
		Detail:   applyProfileResponse.Meta.Message,
	})

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func validateApplyProfileExceptions(d *schema.ResourceData) error {
	if v, ok := d.GetOk("account_ids"); !ok && len(v.(*schema.Set).List()) < 1 {
		return errors.New(`account_ids should NOT have fewer than 1 items`)
	}
	if _, ok := d.GetOk("include"); !ok {
		return nil
	}
	if v, ok := d.GetOk("mode"); ok && v.(string) != "overwrite" {
		return errors.New(`include exceptions functionality is only available in "overwrite" mode`)
	}
	return nil
}
