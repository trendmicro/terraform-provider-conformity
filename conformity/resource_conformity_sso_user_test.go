package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceconformitySsoLegacyUser(t *testing.T) {

	// userAccessDetails and userDetails also uses by the other resource testing, expect not empty struct
	// To make sure the both userAccessDetails and userDetails  is empty
	userAccessDetails = cloudconformity.UserAccessDetails{}
	userDetails = cloudconformity.UserDetails{}

	createUserInfo := testPersonalInfo{
		firstName: "John",
		lastName:  "Smith",
		email:     "john.smith@company.com",
		role:      "ADMIN",
	}

	userAccessList := make([]accessList, 2)
	userAccessList[0].account = "some_id_1"
	userAccessList[0].level = "NONE"
	userAccessList[1].account = "some_id_2"
	userAccessList[1].level = "NONE"
	createUserInfo.accessList = userAccessList

	updateUserInfo := createUserInfo
	updateUserInfo.role = "USER"
	updateUserAccessList := make([]accessList, 2)
	updateUserAccessList[0].account = "some_id_3"
	updateUserAccessList[0].level = "FULL"
	updateUserAccessList[1].account = "some_id_4"
	updateUserAccessList[1].level = "READONLY"
	updateUserInfo.accessList = updateUserAccessList

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		CheckDestroy: testAccCheckConformityLegacyUserDestroy,
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConformitySsoLegacyUserConfigBasic(createUserInfo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "first_name", "John"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "last_name", "Smith"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "email", "john.smith@company.com"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "role", "ADMIN"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "access_list.0.account", "some_id_1"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "access_list.0.level", "NONE"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "access_list.1.account", "some_id_2"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "access_list.1.level", "NONE"),
				), ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckConformitySsoLegacyUserConfigBasic(updateUserInfo),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "first_name", "John"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "last_name", "Smith"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "email", "john.smith@company.com"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "role", "USER"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "access_list.0.account", "some_id_3"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "access_list.0.level", "FULL"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "access_list.1.account", "some_id_4"),
					resource.TestCheckResourceAttr("conformity_sso_user.sso_user", "access_list.1.level", "READONLY"),
				), ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckConformitySsoLegacyUserConfigBasic(info testPersonalInfo) string {
	return fmt.Sprintf(`
	resource "conformity_sso_user" "sso_user" {
		first_name = "%s"
		last_name  = "%s"
		email      = "%s"
		role       = "%s"
	  
		access_list {
		  account = "%s"
		  level   = "%s"
		}
		access_list {
			account = "%s"
			level   = "%s"
		}
	}
	`, info.firstName,
		info.lastName,
		info.email,
		info.role,
		info.accessList[0].account,
		info.accessList[0].level,
		info.accessList[1].account,
		info.accessList[1].level,
	)
}
