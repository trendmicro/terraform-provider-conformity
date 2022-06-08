package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
// 	"github.com/stretchr/testify/assert"
)

func TestAccResourceGCPAccount(t *testing.T) {

	// accountPayload,ruleSetting1,rulesetting2, ruleSetting3 also uses by the other resource testing, expect not empty struct
	// To make sure  all are empty
	accountPayload = cloudconformity.AccountPayload{}
	ruleSetting1 = nil
	ruleSetting2 = nil
	ruleSetting3 = nil

	name := "test-name"
	projectId := "conformity-346910"
	projectName := "conformity"
	environment := "test-env"
	serviceAccountUniqueId := "112840099457455417995"
	updatedName := "test-name-2"
	updatedProjectId := "conformity-346910"
	updatedProjectName := "conformity"
	updatedServiceAccountUniqueId := "112840099457455417995"
	updatedTags := []string{"tag1", "tag2"}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		CheckDestroy: testAccCheckGCPAccountDestroy,
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGCPAccountConfigBasic(name, projectId, projectName, serviceAccountUniqueId, environment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "name", "test-name"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "environment", "test-env"),
					//resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "project_id", "conformity-346910"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "tags.0", "staging"),
					resource.TestCheckOutput("conformity_account_name", "test-name"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.bot.0.delay", "0"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.bot.0.disabled_regions.0", "ap-east-1"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.bot.0.disabled_regions.1", "ap-south-1"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.rule_id", "RTM-005"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.enabled", "true"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.exceptions.0.tags.0", "some_tag"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.type", "countries"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.values.0.value", "CA"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.values.0.label", "Canada"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.values.1.value", "US"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.values.1.label", "United States"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.1.rule_id", "RTM-011"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.1.settings.0.extra_settings.0.type", "multiple-object-values"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.2.rule_id", "VPC-013"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.2.settings.0.extra_settings.0.type", "multiple-vpc-gateway-mappings"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.2.settings.0.extra_settings.0.mappings.0.values.0.values.0.value", "nat-001"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.2.settings.0.extra_settings.0.mappings.0.values.0.values.1.value", "nat-002"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.2.settings.0.extra_settings.0.mappings.0.values.1.value", "vpc-001"),
				), ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckGCPAccountConfigUpdate(updatedName, updatedProjectId, updatedProjectName, updatedServiceAccountUniqueId, updatedTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "name", "test-name-2"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "environment", "test-env-2"),
					//resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "project_id", "conformity-346910"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "tags.0", "tag1"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "tags.1", "tag2"),
					resource.TestCheckOutput("conformity_account_name", "test-name-2"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.bot.0.delay", "1"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.bot.0.disabled_regions.0", "ap-east-1"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.bot.0.disabled_regions.1", "ap-south-1"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.rule_id", "RTM-005"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.enabled", "true"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.exceptions.0.tags.0", "another_tag"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.type", "countries"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.values.0.value", "SG"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.values.0.label", "Singapore"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.values.1.value", "US"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.0.settings.0.extra_settings.0.values.1.label", "United States"),
					resource.TestCheckResourceAttr("conformity_gcp_account.gcp", "settings.0.rule.1.settings.0.extra_settings.0.mappings.0.values.0.values.0.value", "nat-001"),
				), ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckGCPAccountConfigBasic(name, project_id, project_name, service_account_unique_id, environment string) string {
	return fmt.Sprintf(`
	resource "conformity_gcp_account" "gcp" {
		name            = "%s"
        project_id       = "%s"
        project_name     = "%s"
        service_account_unique_id = "%s"
        environment = "%s"
		tags = ["staging"]
		settings {
			bot {
				delay            = 0
				disabled         = false
				disabled_regions = [ "ap-east-1", "ap-south-1" ]
			}
			// implement multiple values
			rule {
				rule_id = "RTM-005"
				settings {
					enabled     = true
					risk_level  = "MEDIUM"
					rule_exists = false
					exceptions {
						tags = ["some_tag"]
					}
					extra_settings {
						name  = "authorisedCountries"
						type  = "countries"
						values {
							value = "CA"
							label = "Canada"
						}
						values {
							value = "US"
							label = "United States"
						}
					}
				}
			}
			// implement multiple-object-values
			rule {
				rule_id = "RTM-011"
				settings {
					enabled     = true
					risk_level  = "MEDIUM"
					rule_exists = false
					extra_settings {
						name    = "patterns"
						type    = "multiple-object-values"
					}
				}
			}
			// implement mappings
			rule {
				rule_id = "VPC-013"
				settings {
					enabled     = true
					risk_level  = "LOW"
					rule_exists = false
					extra_settings {
						name    = "SpecificVPCToSpecificGatewayMapping"
						type    = "multiple-vpc-gateway-mappings"
						mappings {
							values {
								name = "gatewayIds"
								type = "multiple-string-values"
								values {
									value = "nat-001"
								}
								values {
									value = "nat-002"
								}
							}
							values {
								name  = "vpcId"
								type  = "single-string-value"
								value = "vpc-001"
							}

						}
					}
				}
			}
		}
	}
	output "conformity_account_name" {
		value = conformity_gcp_account.gcp.name
	}

	`, name, project_id, project_name, service_account_unique_id, environment)
}

func testAccCheckGCPAccountConfigUpdate(name, project_id, project_name, service_account_unique_id string, tags []string) string {
	return fmt.Sprintf(`
	resource "conformity_gcp_account" "gcp" {
		name = "%s"
		project_id       = "%s"
        project_name     = "%s"
        service_account_unique_id = "%s"
        environment = "test-env-2"
		tags = ["%s","%s"]
		settings {
			bot {
				delay            = 1
				disabled         = false
				disabled_regions = [ "ap-east-1", "ap-south-1" ]
			}
			rule {
				rule_id = "RTM-005"
				settings {
					enabled     = true
					risk_level  = "MEDIUM"
					rule_exists = false
					exceptions {
						tags = ["another_tag"]
					}
					extra_settings {
						name  = "authorisedCountries"
						type  = "countries"
						values {
							value = "SG"
							label = "Singapore"
						}
						values {
							value = "US"
							label = "United States"
						}
					}
				}
			}
			rule {
				rule_id = "VPC-013"
				settings {
					enabled     = true
					risk_level  = "LOW"
					rule_exists = false

					extra_settings {
						name    = "SpecificVPCToSpecificGatewayMapping"
						type    = "multiple-vpc-gateway-mappings"
						mappings {
							values {
								name = "gatewayIds"
								type = "multiple-string-values"
								values {
									value = "nat-001"
								}
								values {
									value = "nat-002"
								}
							}
							values {
								name  = "vpcId"
								type  = "single-string-value"
								value = "vpc-001"
							}

						}
					}
				}
			}
		}
	}
	output "conformity_account_name" {
		value = conformity_gcp_account.gcp.name
	}

	`, name, project_id, project_name, service_account_unique_id, tags[0], tags[1])
}

func testAccCheckGCPAccountDestroy(s *terraform.State) error {
	c := testAccConformityProvider.Meta().(*cloudconformity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "conformity_gcp_account" {
			continue
		}
		accountId := rs.Primary.ID

		deleteAccount, err := c.DeleteAccount(accountId)
		if deleteAccount.Meta.Status != "sent" {
			return fmt.Errorf("Conformity gcp Account not destroyed")
		}
		if err != nil {
			return err
		}
	}
	testServer.Close()
	return nil
}