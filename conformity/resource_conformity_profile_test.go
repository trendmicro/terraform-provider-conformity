package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceconformityProfile(t *testing.T) {
	ttl := "72"
	ttlUpdate := "71"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		CheckDestroy: testAccConformityProfileDestroy,
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConformityProfileBasic(ttl),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_profile.rtm002", "name", "test-with-rules"),
					resource.TestCheckResourceAttr("conformity_profile.rtm002", "included.0.id", "RTM-002"),
					resource.TestCheckResourceAttr("conformity_profile.rtm002", "included.0.extra_settings.0.value", ttl),
					resource.TestCheckResourceAttr("conformity_profile.rtm002", "included.1.extra_settings.0.values.0.value", "includeConformityOrganization"),
				), ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckConformityProfileBasic(ttlUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_profile.rtm002", "included.0.extra_settings.0.value", ttlUpdate),
				), ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckConformityProfileBasic(ttl string) string {
	return fmt.Sprintf(`
	resource "conformity_profile" "rtm002" {
		description = "conformity development - rules included"
		name        = "test-with-rules"

		included {
			id         = "RTM-002"
			provider   = "aws"

			extra_settings {
				name      = "ttl"
				type      = "ttl"
				value     = "%s"
			}
		}
		included {
			id         = "SNS-002"
			provider   = "aws"
			extra_settings {
				name      = "conformityOrganization"
				type      = "choice-multiple-value"
				values {
					enabled = false
					label   = "All within this Conformity organization"
					value   = "includeConformityOrganization"
				}
				values {
					enabled = true
					label   = "All within this AWS Organization"
					value   = "includeAwsOrganizationAccounts"
				}
			}
		}
	}`, ttl)
}

func testAccConformityProfileDestroy(s *terraform.State) error {
	c := testAccConformityProvider.Meta().(*cloudconformity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "conformity_profile" {
			continue
		}
		profileId := rs.Primary.ID

		deleteResponse, err := c.DeleteProfile(profileId)
		if deleteResponse.Meta.Status != "deleted" {
			return fmt.Errorf("Conformity profile not destroyed")
		}
		if err != nil {
			return err
		}
	}
	testServer.Close()
	return nil
}
