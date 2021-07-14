package conformity

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccConformityApplyProfile(t *testing.T) {

	profileId := "0mBdeD0CI"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccConformityPreCheck(t) },
		Providers: testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConformityApplyProfileBasic(profileId, "overwrite"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.conformity_apply_profile.profile", "profile_id", profileId),
					resource.TestCheckResourceAttr("data.conformity_apply_profile.profile", "account_ids.0", "80b880c9-336a-490d-b212-4e847956a62d"),
					resource.TestCheckResourceAttr("data.conformity_apply_profile.profile", "mode", "overwrite"),
				),
			},
			{
				Config:      testAccCheckConformityApplyProfileBasic(profileId, "fill-gaps"),
				ExpectError: regexp.MustCompile(`include exceptions functionality is only available in "overwrite" mode`),
			},
			{
				Config:      testAccCheckConformityApplyProfileFail(),
				ExpectError: regexp.MustCompile(`account_ids should NOT have fewer than 1 items`),
			},
		},
	})
}

func testAccCheckConformityApplyProfileBasic(profileId, mode string) string {
	return fmt.Sprintf(`
	data "conformity_apply_profile" "profile"{

		profile_id = "%s"
	  
		account_ids = ["80b880c9-336a-490d-b212-4e847956a62d"]
		mode = "%s"
		notes = "conformity development - apply profile"
		include {
		  exceptions = false
		}
	  }
	  `, profileId, mode)
}

func testAccCheckConformityApplyProfileFail() string {
	return `
	data "conformity_apply_profile" "profile"{

		profile_id = "0mBdeD0CI"
	  
		account_ids = []
		mode = "overwrite"
		notes = "conformity development - apply profile"
		include {
		  exceptions = false
		}
	  }
	  `
}
