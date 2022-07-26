package conformity

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"testing"
)

func TestAccConformityGetGcpProjects(t *testing.T) {
	name := "data.conformity_gcp_projects.this"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccConformityPreCheck(t) },
		Providers: testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckConformityGetGcpProjectsConfig("invalid"),
				ExpectError: regexp.MustCompile("invalid value for organisation_id \\('organisation_id' is not in a valid format.\\)"),
			},
			{
				Config: testAccCheckConformityGetGcpProjectsConfig("2a-1b"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "projects.#", "1"),
					resource.TestCheckResourceAttr(name, "projects.0.type", "projects"),
					resource.TestCheckResourceAttr(name, "projects.0.project_number", "415104041262"),
					resource.TestCheckResourceAttr(name, "projects.0.project_id", "project-id-1"),
					resource.TestCheckResourceAttr(name, "projects.0.lifecycle_state", "ACTIVE"),
					resource.TestCheckResourceAttr(name, "projects.0.added_to_conformity", "true"),
					resource.TestCheckResourceAttr(name, "projects.0.name", "My Project"),
					resource.TestCheckResourceAttr(name, "projects.0.parent_type", "folder"),
					resource.TestCheckResourceAttr(name, "projects.0.parent_id", "415104041262"),
					resource.TestCheckResourceAttr(name, "projects.0.create_time", "2021-05-17 11:21:58.012 +0000 UTC"),
				),
			},
		},
	})
}

func testAccCheckConformityGetGcpProjectsConfig(id string) string {
	return fmt.Sprintf(`data "conformity_gcp_projects" "this" {
	  organisation_id = "%s"
	}`, id)
}
