package conformity

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"log"
)

func resourceConformityAzureActiveDirectory() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformityAzureActiveDirectoryCreate,
		ReadContext:   resourceConformityAzureActiveDirectoryRead,
		DeleteContext: resourceConformityAzureActiveDirectoryDelete,
		UpdateContext: resourceConformityAzureActiveDirectoryUpdate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"directory_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"application_id": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"application_key": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceConformityAzureActiveDirectoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	var diag diag.Diagnostics
	active_directory_details := cloudconformity.ActiveAzureDirectory{}

	active_directory_details.Data.Attributes.Name = d.Get("name").(string)
	active_directory_details.Data.Attributes.DirectoryId = d.Get("directory_id").(string)
	active_directory_details.Data.Attributes.ApplicationId = d.Get("application_id").(string)
	active_directory_details.Data.Attributes.Applicationkey = d.Get("application_key").(string)

	id, err := client.CreateAzureActiveDirectory(active_directory_details)
	if err != nil {
		return diag
	}
	d.SetId(id)
	resourceConformityAzureActiveDirectoryRead(ctx, d, m)
	return diag
}
func resourceConformityAzureActiveDirectoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diag diag.Diagnostics
	Id := d.Id()
	log.Println("[DEBUG]", "id is  ", Id)
	return diag
}
func resourceConformityAzureActiveDirectoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diag diag.Diagnostics

	return diag
}
func resourceConformityAzureActiveDirectoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diag diag.Diagnostics

	return diag
}
