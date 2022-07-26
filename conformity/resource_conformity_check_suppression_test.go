package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceConformityCheckSuppression(t *testing.T) {
	ruleId := "Resources-001"
	ruleIdUpdate := "Resources-002"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		CheckDestroy: testAccConformityCheckSuppressionDestroy,
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckConformityCheckSuppressionBasic(ruleId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_check_suppression.test01", "account_id", "8d99dfce-dca2-4f13-8699-20631a5c77c9"),
					resource.TestCheckResourceAttr("conformity_check_suppression.test01", "rule_id", "Resources-001"),
					resource.TestCheckResourceAttr("conformity_check_suppression.test01", "region", "global"),
					resource.TestCheckResourceAttr("conformity_check_suppression.test01", "resource_id", "/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483"),
					resource.TestCheckResourceAttr("conformity_check_suppression.test01", "note", "disable for testing purposes"),
				), ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckConformityCheckSuppressionBasic(ruleIdUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_check_suppression.test01", "rule_id", ruleIdUpdate),
				), ExpectNonEmptyPlan: true,
			},
			{
				Config:      testAccCheckConformityCheckSuppressionBasic("invalidruleformat"),
				ExpectError: regexp.MustCompile(`'rule_id' is not in a valid format.`),
			},
		},
	})
}

func testAccCheckConformityCheckSuppressionBasic(ruleId string) string {
	return fmt.Sprintf(`
		resource conformity_check_suppression test01 {
		  account_id  = "8d99dfce-dca2-4f13-8699-20631a5c77c9"
		  rule_id     = "%s"
		  resource_id = "/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483"
		  region      = "global"
		  note        = "disable for testing purposes"
		}
	`, ruleId)
}

func testAccConformityCheckSuppressionDestroy(s *terraform.State) error {
	c := testAccConformityProvider.Meta().(*cloudconformity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "conformity_check_suppression" {
			continue
		}
		checkId := rs.Primary.ID
		payload := cloudconformity.CheckDetails{}
		payload.Data.Type = "checks"
		payload.Data.Attributes.Suppressed = false
		payload.Meta.Note = "re-enabled as suppression has been deleted in Terraform"
		response, err := c.UpdateCheck(checkId, payload)
		if response.Data.Attributes.Suppressed != false {
			return fmt.Errorf("conformity check suppression not destroyed")
		}
		if err != nil {
			return err
		}
	}
	testServer.Close()
	return nil
}
