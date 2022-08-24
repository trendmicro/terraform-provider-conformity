package conformity

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccConformityGetCurrentUser(t *testing.T) {
	user := "data.conformity_current_user.user"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccConformityPreCheck(t) },
		Providers: testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: `data "conformity_current_user" "user" {}`,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(user, "created_date", "0"),
					resource.TestCheckResourceAttr(user, "has_credentials", "false"),
					resource.TestCheckResourceAttr(user, "id", "517uNyIvG"),
					resource.TestCheckResourceAttr(user, "is_api_key_user", "false"),
					resource.TestCheckResourceAttr(user, "is_cloud_one_user", "false"),
					resource.TestCheckResourceAttr(user, "last_login_date", "0"),
					resource.TestCheckResourceAttr(user, "summary_email_opt_out", "true"),
					resource.TestCheckResourceAttr(user, "mfa", "false"),
					resource.TestCheckResourceAttr(user, "role", "ADMIN"),
					resource.TestCheckResourceAttr(user, "type", "users"),
				),
			},
		},
	})
}
