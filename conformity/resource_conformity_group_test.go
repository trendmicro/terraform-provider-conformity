package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceconformityGroup(t *testing.T) {

	name := "test-group-name"
	tags := []string{"dev", "prod"}

	updatedName := "test-group-name-1"
	updatedTags := []string{"dev-1", "prod-1"}

	// send a specific name 'no-tag' will trigger the mock to send a response without tags
	noTagName := "no-tag"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		CheckDestroy: testAccCheckConformityGroupDestroy,
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConformityGroupConfigBasic(name, tags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_group.group_1", "name", "test-group-name"),
					resource.TestCheckResourceAttr("conformity_group.group_1", "tags.0", "dev"),
					resource.TestCheckResourceAttr("conformity_group.group_1", "tags.1", "prod"),
					resource.TestCheckOutput("group_1_name", "test-group-name"),
				),
			},
			{
				Config: testAccCheckConformityGroupConfigBasic(updatedName, updatedTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_group.group_1", "name", "test-group-name-1"),
					resource.TestCheckResourceAttr("conformity_group.group_1", "tags.0", "dev-1"),
					resource.TestCheckResourceAttr("conformity_group.group_1", "tags.1", "prod-1"),
					resource.TestCheckOutput("group_1_name", "test-group-name-1"),
				),
			},
			{
				Config: testAccCheckConformityGroupConfigBasic(noTagName, updatedTags),
				// No check function is given because we expect this configuration
				// to fail before any infrastructure is created
				ExpectError: regexp.MustCompile("Conformity API return empty tag list"),
			},
		},
	})
}

func testAccCheckConformityGroupConfigBasic(name string, tags []string) string {
	return fmt.Sprintf(`
	resource "conformity_group" "group_1" {
		name = "%s"
		tags = ["%s","%s"]
	}
	output "group_1_name" {
		value = conformity_group.group_1.name
	}
	`, name, tags[0], tags[1])
}

func testAccCheckConformityGroupDestroy(s *terraform.State) error {
	c := testAccConformityProvider.Meta().(*cloudconformity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "conformity_group" {
			continue
		}
		GroupId := rs.Primary.ID

		deleteGroup, err := c.DeleteGroup(GroupId)
		if deleteGroup.Meta.Status != "deleted" {
			return fmt.Errorf("Conformity Group not destroyed")
		}
		if err != nil {
			return err
		}
	}
	testServer.Close()
	return nil
}
