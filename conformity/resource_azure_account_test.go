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

	// accountPayload,ruleSetting1,rulesetting2, ruleSetting3 also uses by the other resource testing, expect not empty struct
	// To make sure  all are empty
	accountPayload = cloudconformity.AccountPayload{}
	ruleSetting1 = nil
	ruleSetting2 = nil
	ruleSetting3 = nil

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
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.rule_id", "SecurityCenter-020"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.enabled", "true"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.exceptions.0.tags.0", "some_tag"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.type", "choice-multiple-value"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.values.0.value", "Azure-CIS-1.1.0"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.values.0.label", "Azure CIS 1.1.0"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.values.1.value", "PCI-DSS-3.2.1"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.values.1.label", "PCI DSS 3.2.1"),
				), ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckAzureAccountConfigUpdate(updatedName, updatedEnvironment, subscriptionId, activeDirectoryId, updatedTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "name", "test-name-2"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "environment", "test-env-2"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "tags.0", "tag1"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "tags.1", "tag2"),
					resource.TestCheckOutput("conformity_account_name", "test-name-2"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.rule_id", "SecurityCenter-020"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.enabled", "true"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.exceptions.0.tags.0", "another_tag"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.type", "choice-multiple-value"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.values.0.value", "ISO-27001"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.values.0.label", "ISO 27001"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.values.1.value", "SOC-TSP"),
					resource.TestCheckResourceAttr("conformity_azure_account.azure", "settings.0.rule.0.settings.0.extra_settings.0.values.1.label", "SOC TSP"),
				), ExpectNonEmptyPlan: true,
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
		settings {
			bot {
				delay            = 0
				disabled         = false
				disabled_regions = [ "ap-east-1", "ap-south-1" ]
			}
			rule {
				rule_id = "SecurityCenter-020"
				note = "test"
				settings {
					enabled     = true
					risk_level  = "MEDIUM"
					rule_exists = false
					exceptions {
						tags = ["some_tag"]
					}
					extra_settings {
						name  = "complianceStandards"
						type  = "choice-multiple-value"
						values {
							value = "Azure-CIS-1.1.0"
							label = "Azure CIS 1.1.0"
							enabled = true
						}
						values {
							value = "PCI-DSS-3.2.1"
							label = "PCI DSS 3.2.1"
							enabled = true
						}
					}
				}
			}
		}
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
		settings {
			bot {
				delay            = 1
				disabled         = false
				disabled_regions = [ "ap-east-1", "ap-south-1" ]
			}
			rule {
				rule_id = "SecurityCenter-020"
				note = "test"
				settings {
					enabled     = true
					risk_level  = "MEDIUM"
					rule_exists = false
					exceptions {
						tags = ["another_tag"]
					}
					extra_settings {
						name  = "complianceStandards"
						type  = "choice-multiple-value"
						values {
							value = "ISO-27001"
							label = "ISO 27001"
							enabled = true
						}
						values {
							value = "SOC-TSP"
							label = "SOC TSP"
							enabled = true
						}
					}
				}
			}
		}
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
