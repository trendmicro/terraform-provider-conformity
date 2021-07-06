package conformity

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccConformityGetExternalId(t *testing.T) {
	name := "data.conformity_external_id.external_id"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccConformityPreCheck(t) },
		Providers: testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: `data "conformity_external_id" "external_id" {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "external_id"),
					resource.TestCheckResourceAttr(name, "external_id", "3ff84b20-0f4c-11eb-a7b7-7d9b3c0e866e"),
				),
			},
		},
	})
}
