package conformity

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
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
				Config: testAccCheckCommunicationSettingServiceNow(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.channel_name", "service-now-channel"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.type", "incident"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.username", "service-now-user"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.password", "service-now-password"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.impact", "1"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.urgency", "3"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.close_code", "Closed/Resolved by Caller"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.close_notes", "Issue resolved"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.creation_override.urgency", "2"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.creation_override.priority", "1"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.resolution_override.close_code", "Closed by Caller"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.resolution_override.close_notes", "Issue closed"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "service_now.0.url", "https://instance.service-now.com/api/now/table/incident"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "filter.0.statuses.0", "FAILURE"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "filter.0.categories.0", "security"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "relationships.0.account.0.id", "awesome-account-id"),
					resource.TestCheckResourceAttr("conformity_communication_setting.service_now", "relationships.0.organisation.0.id", "awesome-org-id"),
				),
				ExpectNonEmptyPlan: true,
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

func testAccCheckCommunicationSettingServiceNow() string {
	return `
	resource "conformity_communication_setting" "service_now" {
		service_now {
			channel_name = "service-now-channel"
			type = "incident"

			username = "service-now-user"
			password = "service-now-password"

			impact = "1" 
			urgency = "3"

			close_code = "Closed/Resolved by Caller"
			close_notes = "Issue resolved"

			creation_override = {
				urgency = "2"
				priority = "3" 
			}

			resolution_override = {
				close_code = "Closed by Caller"
				close_notes = "Issue closed"
			}
			url = "https://instance.service-now.com/api/now/table/incident"
		}
		filter {
			categories  = [ "security" ]
			statuses = ["FAILURE"]
		}
		relationships {
			account {
				id = "awesome-account-id"
			}
			organisation {
				id = "awesome-org-id"
			}
		}
	}
	`
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

func TestFlattenCommSettingConfiguration(t *testing.T) {
	commSetting := cloudconformity.CommunicationConfiguration{
		ChannelName: "my-test-channel",
		Type:        "problem",
		Url:         "https://mytest001.service-now.com",
		UserName:    "admin",
		Password:    "******",
		Assignee:    "admin",
		Impact:      "3",
		Urgency:     "1",
		CloseCode:   "Closed/Resolved by Caller",
		CloseNotes:  "Issue resolved",
		ResolutionOverride: map[string]interface{}{
			"closeCode":  "Closed by Caller",
			"closeNotes": "Issue closed",
		},
		CreationOverride: map[string]interface{}{
			"urgency":  "2",
			"priority": "3",
		},
	}

	flatConfig := flattenCommSettingConfiguration(&commSetting, "service-now")
	assert.Equal(t, 1, len(flatConfig))
	c0 := flatConfig[0].(map[string]interface{})
	assert.Equal(t, "my-test-channel", c0["channel_name"])
	assert.Equal(t, "problem", c0["type"])
	assert.Equal(t, "https://mytest001.service-now.com", c0["url"])
	assert.Equal(t, "admin", c0["username"])
	assert.Equal(t, "******", c0["password"])
	assert.Equal(t, "admin", c0["assignee"])
	assert.Equal(t, "low", c0["impact"])
	assert.Equal(t, "high", c0["urgency"])
	assert.Equal(t, "Closed/Resolved by Caller", c0["close_code"])
	assert.Equal(t, "Issue resolved", c0["close_notes"])

	resolutionOverride := c0["resolution_override"].(map[string]interface{})
	assert.Equal(t, "Closed by Caller", resolutionOverride["closeCode"])
	assert.Equal(t, "Issue closed", resolutionOverride["closeNotes"])

	creationOverride := c0["creation_override"].(map[string]interface{})
	assert.Equal(t, "2", creationOverride["urgency"])
	assert.Equal(t, "3", creationOverride["priority"])
}
