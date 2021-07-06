package conformity

import (
	"context"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceConformitySsoLegacyUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformitySsoUserCreate,
		ReadContext:   resourceConformityUserRead,
		UpdateContext: resourceConformityUserUpdate,
		DeleteContext: resourceConformityUserDelete,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"),
					"'email' is not in a valid format."),
			},
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"ADMIN", "USER"}, false),
			},
			"mfa": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_login": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"access_list": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Type:     schema.TypeString,
							Required: true,
						},
						"level": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "NONE",
							ValidateFunc: validation.StringInSlice([]string{"NONE", "READONLY", "FULL"}, false),
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

func resourceConformitySsoUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := cloudconformity.UserDetails{}
	payload.Data.Attributes.Email = d.Get("email").(string)
	payload.Data.Attributes.FirstName = d.Get("first_name").(string)
	payload.Data.Attributes.LastName = d.Get("last_name").(string)
	payload.Data.Attributes.Role = d.Get("role").(string)

	accessList := d.Get("access_list").(*schema.Set).List()
	userAccessList := proccessUserAccessList(accessList)
	payload.Data.Attributes.AccessList = userAccessList

	//add SSO user to the non-cloudone platform
	userId, err := client.CreateSsoLegacyUser(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(userId)
	resourceConformityUserRead(ctx, d, m)
	return diags
}
