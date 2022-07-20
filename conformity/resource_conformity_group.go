package conformity

import (
	"context"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConformityGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformityGroupCreate,
		ReadContext:   resourceConformityGroupRead,
		UpdateContext: resourceConformityGroupUpdate,
		DeleteContext: resourceConformityGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceConformityGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := cloudconformity.GroupDetails{}

	proccessGroupInput(&payload, d)

	groupId, err := client.CreateGroup(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(groupId)
	resourceConformityGroupRead(ctx, d, m)
	return diags
}

func resourceConformityGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groupId := d.Id()

	// get group details: name and tags
	GroupDataList, err := client.GetGroup(groupId)
	if !d.IsNewResource() && (err != nil || GroupDataList.Data == nil) {
		// The resource was deleted manually
		d.SetId("")
		return diags
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if GroupDataList.Data == nil || GroupDataList.Data[0].Attributes.Tags == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Conformity Group Tags - Empty",
			Detail:   "Conformity API return empty tag list",
		})
		return diags
	}
	if err := d.Set("name", GroupDataList.Data[0].Attributes.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tags", GroupDataList.Data[0].Attributes.Tags); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceConformityGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)

	if d.HasChange("name") || d.HasChange("tags") {
		payload := cloudconformity.GroupDetails{}

		proccessGroupInput(&payload, d)

		groupId := d.Id()
		_, err := client.UpdateGroup(groupId, payload)
		if err != nil {
			return diag.FromErr(err)
		}

	}

	return resourceConformityGroupRead(ctx, d, m)
}

func resourceConformityGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	groupId := d.Id()

	_, err := client.DeleteGroup(groupId)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.

	d.SetId("")
	return diags
}

func proccessGroupInput(payload *cloudconformity.GroupDetails, d *schema.ResourceData) {
	payload.Data.Attributes.Name = d.Get("name").(string)
	tags := d.Get("tags").(*schema.Set)
	for _, tag := range tags.List() {
		payload.Data.Attributes.Tags = append(payload.Data.Attributes.Tags, tag.(string))
	}
}
