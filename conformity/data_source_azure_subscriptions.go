package conformity

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAzureSubscription() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAzureSubscriptionRead,
		Schema: map[string]*schema.Schema{
			"active_directory_id": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)+$"),
					"'active_directory_id' is not in a valid format."),
			},
			"subscriptions": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"added_to_conformity": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAzureSubscriptionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	directoryId := d.Get("active_directory_id").(string)

	getAzureSubscriptionsResponse, err := client.GetAzureSubscriptions(directoryId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("subscriptions", flattenSubscription(getAzureSubscriptionsResponse)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenSubscription(response *cloudconformity.AzureSubscriptionsResponse) []interface{} {
	if response == nil {
		return make([]interface{}, 0)
	}
	rs := make([]interface{}, len(response.Data))
	for i, subscription := range response.Data {
		r := make(map[string]interface{})
		r["type"] = subscription.Type
		r["id"] = subscription.ID
		r["display_name"] = subscription.Attributes.DisplayName
		r["state"] = subscription.Attributes.State
		r["added_to_conformity"] = subscription.Attributes.AddedToConformity
		rs[i] = r
	}

	return rs
}
