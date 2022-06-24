package conformity

import (
	"context"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"
    "strings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceConformityLegacyUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformityUserCreate,
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

func resourceConformityUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	//invite a user to the non-cloudone platform
	userId, err := client.InviteLegacyUser(payload)
	if err != nil {
	    if strings.Contains(err.Error(), "Unable to call this endpoint, use Cloud One UI or API to invite users"){
            diags = append(diags, diag.Diagnostic{
                Severity: diag.Error,
                Summary:  "Unable to Invite Conformity Cloud One user",
                Detail:   "This Terraform service is not applicable to users who are part of the Cloud One Platform.\nPlease refer to Cloud One User Management Documentation - Add and manage users to invite new users. https://cloudone.trendmicro.com/docs/conformity/api-reference/tag/Users#paths/~1users/get",
		    })
		    return diags
	    }
	    return diag.FromErr(err)

	}
	d.SetId(userId)
	resourceConformityUserRead(ctx, d, m)
	return diags
}

func resourceConformityUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	userId := d.Id()

	// get user details of non-cloudone user
	userDetails, err := client.GetLegacyUser(userId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("email", userDetails.Data.Attributes.Email); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("first_name", userDetails.Data.Attributes.ResFirstName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_name", userDetails.Data.Attributes.ResLastName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("role", userDetails.Data.Attributes.Role); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mfa", userDetails.Data.Attributes.Mfa); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_login", userDetails.Data.Attributes.LastLogIn); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("access_list", userDetails.Data.Attributes.AccessList); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceConformityUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.HasChange("role") || d.HasChange("access_list") {

		payload := cloudconformity.UserAccessDetails{}
		payload.Data.Role = d.Get("role").(string)
		access_list := d.Get("access_list").(*schema.Set).List()
		userAccessList := proccessUserAccessList(access_list)
		payload.Data.AccessList = userAccessList

		userId := d.Id()
		// update user Role/Access of a non-cloudone platform
		_, err := client.UpdateLegacyUser(userId, payload)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChange("email") || d.HasChange("first_name") || d.HasChange("last_name") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to Update Conformity User",
			Detail:   "'email', 'first_name' and 'last_name' cannot be changed",
		})

		return diags
	}
	return resourceConformityUserRead(ctx, d, m)
}

func resourceConformityUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	userId := d.Id()
	// revoke a user of non-cloudone platform
	_, err := client.RevokeLegacyUser(userId)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.

	d.SetId("")

	return diags
}

func proccessUserAccessList(accessList []interface{}) []cloudconformity.UserAccountAccessList {

	userAccessList := make([]cloudconformity.UserAccountAccessList, len(accessList))

	for i, items := range accessList {

		item := items.(map[string]interface{})

		userAccessList[i].Account = item["account"].(string)
		userAccessList[i].Level = item["level"].(string)

	}
	return userAccessList
}
