package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccResourceAwsAccount(t *testing.T) {

	// accountPayload,ruleSetting1,rulesetting2, ruleSetting3 also uses by the other resource testing, expect not empty struct
	// To make sure  all are empty
	accountPayload = cloudconformity.AccountPayload{}
	ruleSetting1 = nil
	ruleSetting2 = nil
	ruleSetting3 = nil

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
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.bot.0.delay", "0"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.bot.0.disabled_regions.0", "ap-east-1"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.bot.0.disabled_regions.1", "ap-south-1"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.rule_id", "RTM-005"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.enabled", "true"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.exceptions.0.tags.0", "some_tag"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.type", "countries"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.values.0.value", "CA"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.values.0.label", "Canada"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.values.1.value", "US"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.values.1.label", "United States"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.1.rule_id", "RTM-011"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.1.settings.0.extra_settings.0.type", "multiple-object-values"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.1.settings.0.extra_settings.0.multiple_object_values.0.event_name", "^(iam.amazonaws.com)"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.2.rule_id", "VPC-013"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.2.settings.0.extra_settings.0.type", "multiple-vpc-gateway-mappings"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.2.settings.0.extra_settings.0.mappings.0.values.0.values.0.value", "nat-001"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.2.settings.0.extra_settings.0.mappings.0.values.0.values.1.value", "nat-002"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.2.settings.0.extra_settings.0.mappings.0.values.1.value", "vpc-001"),
				), ExpectNonEmptyPlan: true,
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
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.bot.0.delay", "1"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.bot.0.disabled_regions.0", "ap-east-1"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.bot.0.disabled_regions.1", "ap-south-1"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.rule_id", "RTM-005"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.enabled", "true"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.exceptions.0.tags.0", "another_tag"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.type", "countries"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.values.0.value", "SG"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.values.0.label", "Singapore"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.values.1.value", "US"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.0.settings.0.extra_settings.0.values.1.label", "United States"),
					resource.TestCheckResourceAttr("conformity_aws_account.aws", "settings.0.rule.1.settings.0.extra_settings.0.mappings.0.values.0.values.0.value", "nat-001"),
				), ExpectNonEmptyPlan: true,
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
		settings {
			bot {
				delay            = 0
				disabled         = false
				disabled_regions = [ "ap-east-1", "ap-south-1" ]
			}
			// implement multiple values
			rule {
				rule_id = "RTM-005"
				note = "test"
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
				note = "test note"
				settings {
					enabled     = true
					risk_level  = "MEDIUM"
					rule_exists = false
					extra_settings {
						name    = "patterns"
						type    = "multiple-object-values"
						multiple_object_values {
							event_name         = "^(iam.amazonaws.com)"
							event_source       = "^(IAM).*"
							user_identity_type = "^(Delete).*"
						}
					}
				}
			}
			// implement mappings
			rule {
				rule_id = "VPC-013"
				note = "test note"
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
		settings {
			bot {
				delay            = 1
				disabled         = false
				disabled_regions = [ "ap-east-1", "ap-south-1" ]
			}
			rule {
				rule_id = "RTM-005"
				note = "test"
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
				note = "test note"
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

func TestFlattenBotDisabledRegions(t *testing.T) {
	res := flattenBotDisabledRegions(cloudconformity.BotDisabledRegions{})
	assert.Equal(t, 0, len(res))

	res = flattenBotDisabledRegions(cloudconformity.BotDisabledRegions{
		EuSouth1: true,
	})
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "eu-south-1", res[0])
}

func TestProcessBotDisabledRegions(t *testing.T) {
	list := make([]interface{}, 0)
	regions := processBotDisabledRegions(list)
	assert.False(t, regions.AfSouth1)
	assert.False(t, regions.UsWest2)

	list = make([]interface{}, 1)
	list[0] = "nonsense"
	regions = processBotDisabledRegions(list)
	assert.False(t, regions.AfSouth1)
	assert.False(t, regions.UsWest2)

	list = make([]interface{}, 2)
	list[0] = "af-south-1"
	list[1] = "us-west-2"
	regions = processBotDisabledRegions(list)
	assert.True(t, regions.AfSouth1)
	assert.True(t, regions.UsWest2)
}
