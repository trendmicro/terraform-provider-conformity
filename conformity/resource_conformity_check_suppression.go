package conformity

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"
	"strings"
)

// Manages a check suppression
// A check suppression is a configuration on an existing check and configures
// the `suppressed` attribute on a check.
// When the Terraform resource is created, the flag is set.
// When the Terraform resource is destroyed, the flag is removed.
// When any configuration parameter is changed, the resource must be
// recreated as it belongs to a different check afterwards.
func resourceConformityCheckSuppression() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformityCheckSuppressionCreate,
		ReadContext:   resourceConformityCheckSuppressionRead,
		DeleteContext: resourceConformityCheckSuppressionDelete,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account ID for which the rule should be suppressed. Example: '5d84e365-af8c-4d9f-8393-f1724bf60d8f'",
				ForceNew:    true,
			},
			"rule_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the rule to be suppressed. Example: 'Resources-001'",
				ForceNew:    true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^(\\w+)-(\\d+)$"),
					"'rule_id' is not in a valid format."),
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The region to which the check applies. Either a specific region or 'global'",
				ForceNew:    true,
			},
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the resource if the check should be suppressed only for a specific resource. Example: '/subscriptions/8dfbsdfe-we13-46we-9963-188868997f40/resourceGroups/myDevResources/providers/Microsoft.KeyVault/vaults/myDevKeyVault-eastus'",
				ForceNew:    true,
			},
			"note": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Explains why the given check has been suppressed",
				ForceNew:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// suppresses a check
// example can be found in the API reference
//{
//	"data": {
//		"type": "checks",
//			"attributes": {
//				"suppressed": true,
//				"suppressed-until": 1526574705655
//			}
//		},
//		"meta": {
//			"note": "suppressed for 1 week, failure not-applicable during project xyz"
//		}
//}
// source: https://cloudone.trendmicro.com/docs/conformity/api-reference/tag/Checks#paths/~1checks~1{checkId}/patch
func resourceConformityCheckSuppressionCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)

	checkId := fmt.Sprintf("ccc:%s:%s:%s:%s:%s",
		d.Get("account_id").(string),
		d.Get("rule_id").(string),
		strings.Split(d.Get("rule_id").(string), "-")[0],
		d.Get("region").(string),
		d.Get("resource_id").(string),
	)

	var diags diag.Diagnostics
	payload := cloudconformity.CheckDetails{}
	payload.Data.Type = "checks"
	payload.Data.Attributes.Suppressed = true
	payload.Meta.Note = d.Get("note").(string)

	checkDetails, err := client.UpdateCheck(checkId, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(checkDetails.Data.Id)
	if err := d.Set("account_id", checkDetails.Data.Relationships.Account.Data.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rule_id", checkDetails.Data.Relationships.Rule.Data.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_id", checkDetails.Data.Attributes.Resource); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("region", checkDetails.Data.Attributes.Region); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceConformityCheckSuppressionRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	checkId := d.Id()

	// get both account details
	checkDetails, err := client.GetCheck(checkId)
	if err != nil {
		return diag.FromErr(err)
	}

	if !checkDetails.Data.Attributes.Suppressed {
		// if the check is no longer suppressed, this resource does not exist anymore
		d.SetId("")
		return diags
	}

	if err := d.Set("account_id", checkDetails.Data.Relationships.Account.Data.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rule_id", checkDetails.Data.Relationships.Rule.Data.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_id", checkDetails.Data.Attributes.Resource); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("region", checkDetails.Data.Attributes.Region); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceConformityCheckSuppressionDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	checkId := d.Id()

	payload := cloudconformity.CheckDetails{}
	payload.Data.Type = "checks"
	payload.Data.Attributes.Suppressed = false
	payload.Meta.Note = "re-enabled as suppression has been deleted in Terraform"

	_, err := client.UpdateCheck(checkId, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
