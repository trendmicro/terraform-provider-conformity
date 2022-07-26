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

func dataSourceGcpProjects() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGcpProjectsRead,
		Schema: map[string]*schema.Schema{
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[a-zA-Z0-9]+(-[a-zA-Z0-9]+)+$"),
					"'organisation_id' is not in a valid format."),
			},
			"projects": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_state": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_id": &schema.Schema{
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

func dataSourceGcpProjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	projectId := d.Get("organisation_id").(string)

	getGoogleProjectsResponse, err := client.GetGcpProjects(projectId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("projects", flattenProjects(getGoogleProjectsResponse)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenProjects(response *cloudconformity.GcpProjectsResponse) []interface{} {
	if response == nil {
		return make([]interface{}, 0)
	}
	rs := make([]interface{}, len(response.Data))
	for i, project := range response.Data {
		r := make(map[string]interface{})
		r["type"] = project.Type
		r["project_number"] = project.Attributes.ProjectNumber
		r["project_id"] = project.Attributes.ProjectID
		r["lifecycle_state"] = project.Attributes.LifecycleState
		r["name"] = project.Attributes.Name
		r["create_time"] = project.Attributes.CreateTime.String()
		r["parent_type"] = project.Attributes.Parent.Type
		r["parent_id"] = project.Attributes.Parent.ID
		r["added_to_conformity"] = project.Attributes.AddedToConformity
		rs[i] = r
	}
	return rs
}
