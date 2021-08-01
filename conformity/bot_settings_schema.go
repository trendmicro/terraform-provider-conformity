package conformity

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func BotSettingsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"disabled": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"delay": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  1,
				},
				"disabled_regions": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{
							"af-south-1",
							"ap-east-1",
							"ap-south-1",
							"ap-southeast-1",
							"ap-southeast-2",
							"ap-northeast-1",
							"ap-northeast-2",
							"ap-northeast-3",
							"ca-central-1",
							"eu-central-1",
							"eu-north-1",
							"eu-south-1",
							"eu-west-1",
							"eu-west-2",
							"eu-west-3",
							"me-south-1",
							"sa-east-1",
							"us-east-1",
							"us-east-2",
							"us-west-1",
							"us-west-2",
						}, false),
					},
				},
			},
		},
	}
}
