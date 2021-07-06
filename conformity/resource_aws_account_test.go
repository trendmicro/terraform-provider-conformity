package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceAwsAccount(t *testing.T) {

	// accountPayload also uses by the other resource testing, expect not empty struct
	// To make sure the accountPayload is empty
	accountPayload = cloudconformity.AccountPayload{}

	name := "test-name"
	environment := "test-env"
	roleARN := "test-arn"
	externalID := "test-external-id"

	updatedName := "test-name-2"
	updatedEnvironment := "test-env-2"
	updatedRoleARN := "test-arn-2"
	updatedExternalID := "test-external-id-2"
	updatedTags := []string{"tag1", "tag2"}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		CheckDestroy: testAccCheckAwsAccountDestroy,
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAwsAccountConfigBasic(name, environment, roleARN, externalID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "name", "test-name"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "environment", "test-env"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "role_arn", "test-arn"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "external_id", "test-external-id"),
					resource.TestCheckOutput("conformity_account_name", "test-name"),
				),
			},
			{
				Config: testAccCheckAwsAccountConfigUpdate(updatedName, updatedEnvironment, roleARN, externalID, updatedTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "name", "test-name-2"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "environment", "test-env-2"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "role_arn", "test-arn"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "external_id", "test-external-id"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "tags.0", "tag1"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "tags.1", "tag2"),
					resource.TestCheckOutput("conformity_account_name", "test-name-2"),
				),
			},
			{
				Config: testAccCheckAwsAccountConfigBasic(updatedName, updatedEnvironment, updatedRoleARN, updatedExternalID),
				// No check function is given because we expect this configuration
				// to fail before any infrastructure is created
				ExpectError: regexp.MustCompile("'role_arn' and 'external_id' cannot be changed"),
			},
		},
	})
}

func testAccCheckAwsAccountConfigBasic(name, environment, roleARN, externalID string) string {
	return fmt.Sprintf(`
	resource "conformity_aws_account" "aws" {
		name = "%s"
		environment = "%s"
		role_arn = "%s"
		external_id = "%s"
	}
	output "conformity_account_name" {
		value = conformity_aws_account.aws.name
	}

	`, name, environment, roleARN, externalID)
}

func testAccCheckAwsAccountConfigUpdate(name, environment, roleARN, externalID string, tags []string) string {
	return fmt.Sprintf(`
	resource "conformity_aws_account" "aws" {
		name = "%s"
		environment = "%s"
		role_arn = "%s"
		external_id = "%s"
		tags = ["%s","%s"]
	}
	output "conformity_account_name" {
		value = conformity_aws_account.aws.name
	}

	`, name, environment, roleARN, externalID, tags[0], tags[1])
}

func testAccCheckAwsAccountDestroy(s *terraform.State) error {
	c := testAccConformityProvider.Meta().(*cloudconformity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "conformity_aws_account" {
			continue
		}
		accountId := rs.Primary.ID

		deleteAccount, err := c.DeleteAccount(accountId)
		if deleteAccount.Meta.Status != "sent" {
			return fmt.Errorf("Conformity AWS Account not destroyed")
		}
		if err != nil {
			return err
		}
	}
	testServer.Close()
	return nil
}
