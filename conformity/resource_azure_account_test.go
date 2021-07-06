package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceAzureAccount(t *testing.T) {

	// accountPayload also uses by the other resource testing, expect not empty struct
	// To make sure the accountPayload is empty
	accountPayload = cloudconformity.AccountPayload{}

	name := "test-name"
	environment := "test-env"
	subscriptionId := "test-subscrition-id"
	activeDirectoryId := "test-active-directory-id"

	updatedName := "test-name-2"
	updatedEnvironment := "test-env-2"
	updatedIubscriptionId := "test-subscrition-id-2"
	updatedActiveDirectoryId := "test-active-directory-id-2"
	updatedTags := []string{"tag1", "tag2"}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		CheckDestroy: testAccCheckAzureAccountDestroy,
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAzureAccountConfigBasic(name, environment, subscriptionId, activeDirectoryId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "name", "test-name"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "environment", "test-env"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "subscription_id", "test-subscrition-id"),
					resource.TestCheckOutput("conformity_account_name", "test-name"),
				),
			},
			{
				Config: testAccCheckAzureAccountConfigUpdate(updatedName, updatedEnvironment, subscriptionId, activeDirectoryId, updatedTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "name", "test-name-2"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "environment", "test-env-2"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "tags.0", "tag1"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "tags.1", "tag2"),
					resource.TestCheckOutput("conformity_account_name", "test-name-2"),
				),
			},
			{
				Config: testAccCheckAzureAccountConfigUpdate(updatedName, updatedEnvironment, updatedIubscriptionId, updatedActiveDirectoryId, updatedTags),
				// No check function is given because we expect this configuration
				// to fail before any infrastructure is created
				ExpectError: regexp.MustCompile("'subscription_id' cannot be changed"),
			},
		},
	})
}

func testAccCheckAzureAccountConfigBasic(name, environment, subscriptionId, activeDirectoryId string) string {
	return fmt.Sprintf(`
	resource "conformity_azure_account" "azure" {
		name = "%s"
		environment = "%s"
		subscription_id = "%s"
		active_directory_id = "%s"
	}
	output "conformity_account_name" {
		value = conformity_azure_account.azure.name
	}

	`, name, environment, subscriptionId, activeDirectoryId)
}
func testAccCheckAzureAccountConfigUpdate(name, environment, subscriptionId, activeDirectoryId string, tags []string) string {
	return fmt.Sprintf(`
	resource "conformity_azure_account" "azure" {
		name = "%s"
		environment = "%s"
		subscription_id = "%s"
		active_directory_id = "%s"
		tags = ["%s","%s"]
	}
	output "conformity_account_name" {
		value = conformity_azure_account.azure.name
	}

	`, name, environment, subscriptionId, activeDirectoryId, tags[0], tags[1])
}

func testAccCheckAzureAccountDestroy(s *terraform.State) error {
	c := testAccConformityProvider.Meta().(*cloudconformity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "conformity_azure_account" {
			continue
		}
		accountId := rs.Primary.ID

		deleteAccount, err := c.DeleteAccount(accountId)
		if deleteAccount.Meta.Status != "sent" {
			return fmt.Errorf("Conformity Azure Account not destroyed")
		}
		if err != nil {
			return err
		}
	}
	testServer.Close()
	return nil
}
