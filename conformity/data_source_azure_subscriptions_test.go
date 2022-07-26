package conformity

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccConformityGetAzureSubscriptions(t *testing.T) {
	name := "data.conformity_azure_subscriptions.this"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccConformityPreCheck(t) },
		Providers: testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckConformityGetAzureSubscriptionsConfig("invalid"),
				ExpectError: regexp.MustCompile("invalid value for active_directory_id \\('active_directory_id' is not in a valid format.\\)"),
			},
			{
				Config: testAccCheckConformityGetAzureSubscriptionsConfig("2a-1b"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "subscriptions.#", "1"),
					resource.TestCheckResourceAttr(name, "subscriptions.0.type", "subscriptions"),
					resource.TestCheckResourceAttr(name, "subscriptions.0.id", "AZURE_SUBSCRIPTION_ID"),
					resource.TestCheckResourceAttr(name, "subscriptions.0.state", "Enabled"),
					resource.TestCheckResourceAttr(name, "subscriptions.0.added_to_conformity", "true"),
					resource.TestCheckResourceAttr(name, "subscriptions.0.display_name", "A Azure Subscription"),
				),
			},
		},
	})
}

func testAccCheckConformityGetAzureSubscriptionsConfig(id string) string {
	return fmt.Sprintf(`data "conformity_azure_subscriptions" "this" {
	  active_directory_id = "%s"
	}`, id)
}
