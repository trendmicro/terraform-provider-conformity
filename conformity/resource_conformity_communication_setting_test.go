package conformity

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceConformityCommSetting(t *testing.T) {

	userId := "urn:tmds:identity:us-east-ds-1:62740:administrator/1915"
	slackChannel := "#slack-channel"
	slackChannelName := "slack-channel-name"
	slackUrl := "slack-url"
	snsArn := "sns-arn"
	snsChannelName := "sns-channel-name"
	webhookToken := "#security-token-01"
	webhookURL := "web-hook-url"
	updatedAccountId := "80b880c9-336a-490d-b212-4e847956a62d"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		CheckDestroy: testAccCheckCommunicationSettingDestroy,
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCommunicationSettingConfigBasic(userId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_communication_setting.email", "email.0.users.0", userId),
				),
			},
			{
				Config: testAccCheckCommunicationSettingConfigUpdate(userId, updatedAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_communication_setting.email", "email.0.users.0", userId),
					resource.TestCheckResourceAttr("conformity_communication_setting.email", "relationships.0.account.0.id", updatedAccountId),
				),
			},
			{
				Config: testAccCheckCommunicationSettingSms(userId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_communication_setting.sms", "sms.0.users.0", userId),
				),
			},

			{
				Config: testAccCheckCommunicationSettingSns(snsArn, snsChannelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_communication_setting.sns", "sns.0.channel_name", snsChannelName),
					resource.TestCheckResourceAttr("conformity_communication_setting.sns", "sns.0.arn", snsArn),
					resource.TestCheckResourceAttr("conformity_communication_setting.sns", "filter.0.statuses.0", "SUCCESS"),
				),
				ExpectNonEmptyPlan: true,
			},

			{
				Config: testAccCheckCommunicationSettingWebhook(webhookToken, webhookURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_communication_setting.webhook", "webhook.0.security_token", webhookToken),
					resource.TestCheckResourceAttr("conformity_communication_setting.webhook", "webhook.0.url", webhookURL),
					resource.TestCheckResourceAttr("conformity_communication_setting.webhook", "filter.0.statuses.0", "FAILURE"),
				),
				ExpectNonEmptyPlan: true,
			},

			{
				Config: testAccCheckCommunicationSettingSlack(slackChannel, slackUrl, slackChannelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_communication_setting.slack", "slack.0.channel", slackChannel),
					resource.TestCheckResourceAttr("conformity_communication_setting.slack", "slack.0.channel_name", slackChannelName),
					resource.TestCheckResourceAttr("conformity_communication_setting.slack", "slack.0.url", slackUrl),
				),
			},
			{
				Config:      testAccCheckCommunicationSettingFail(),
				ExpectError: regexp.MustCompile("found multiple channel configuration set, please provide only one"),
			},
			{
				Config:      testAccCheckCommunicationNoSettingFail(),
				ExpectError: regexp.MustCompile("no channel configuration set found"),
			},
		},
	})
}

func testAccCheckCommunicationSettingConfigBasic(userId string) string {
	return fmt.Sprintf(`
	resource "conformity_communication_setting" "email" {
		email {
			users = [ "%s" ]
		}
		filter {
			categories  = [ "security" ]
		}
		relationships {
			account {
				id = "H19NxM15-"
			}
			organisation {
				id = "ryqMcJn4b"
			}
		}
	}
	`, userId)
}
func testAccCheckCommunicationSettingConfigUpdate(userId, accountId string) string {
	return fmt.Sprintf(`
	resource "conformity_communication_setting" "email" {
		email {
			users = [ "%s" ]
		}
		filter {
			categories  = [ "security" ]
		}
		relationships {
			account {
				id = "%s"
			}
			organisation {
				id = "ryqMcJn4b"
			}
		}
	}
	`, userId, accountId)
}
func testAccCheckCommunicationSettingSms(userId string) string {
	return fmt.Sprintf(`
	resource "conformity_communication_setting" "sms" {
		sms {
			users = [ "%s" ]
		}
		filter {
			categories  = [ "security" ]
		}
		relationships {
			account {
				id = "H19NxM15-"
			}
			organisation {
				id = "ryqMcJn4b"
			}
		}
	}
	`, userId)
}
func testAccCheckCommunicationSettingSlack(slackChannel, slackUrl, slackChannelName string) string {
	return fmt.Sprintf(`
	resource "conformity_communication_setting" "slack" {
		slack {
			channel = "%s"
			url = "%s"
			channel_name = "%s"
		}
		filter {
			categories  = [ "security" ]
		}
		relationships {
			account {
				id = "H19NxM15-"
			}
			organisation {
				id = "ryqMcJn4b"
			}
		}
	}
	`, slackChannel, slackUrl, slackChannelName)
}
func testAccCheckCommunicationSettingSns(arn, channelName string) string {
	return fmt.Sprintf(`
	resource "conformity_communication_setting" "sns" {
		sns {
			arn = "%s"
			channel_name = "%s"
		}
		filter {
			categories  = [ "security" ]
			statuses = ["SUCCESS"]
		}
		relationships {
			account {
				id = "H19NxM15-"
			}
			organisation {
				id = "ryqMcJn4b"
			}
		}
	}
	`, arn, channelName)
}
func testAccCheckCommunicationSettingWebhook(webhookToken, webhookURL string) string {
	return fmt.Sprintf(`
	resource "conformity_communication_setting" "webhook" {
		webhook {
			security_token = "%s"
			url = "%s"
		}
		filter {
			categories  = [ "security" ]
			statuses = ["FAILURE"]
		}
		relationships {
			account {
				id = "H19NxM15-"
			}
			organisation {
				id = "ryqMcJn4b"
			}
		}
	}
	`, webhookToken, webhookURL)
}
func testAccCheckCommunicationSettingFail() string {
	return `
	resource "conformity_communication_setting" "multiplefail" {
		email {
			users = [ "testuser1" ]
		}
		sms {
			users = [ "testuser2" ]
		}
	}
	`
}
func testAccCheckCommunicationNoSettingFail() string {
	return `
	resource "conformity_communication_setting" "nochannelfail" {
		relationships {
			account {
				id = "some-id"
			}
		}
	}
	`
}
func testAccCheckCommunicationSettingDestroy(s *terraform.State) error {
	c := testAccConformityProvider.Meta().(*cloudconformity.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "conformity_communication_setting" {
			continue
		}
		communicationId := rs.Primary.ID

		deleteCommunication, err := c.DeleteCommunicationSetting(communicationId)
		if deleteCommunication.Meta.Status != "deleted" {
			return fmt.Errorf("Conformity Communication Setting not destroyed")
		}
		if err != nil {
			return err
		}
	}

	testServer.Close()

	return nil
}
