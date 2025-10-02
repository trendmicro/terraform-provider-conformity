package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCommunicationSettingSuccess(t *testing.T) {
	// prepare the test
	expectedChannelName := "someChannelName"
	client, ts := createHttpTestClient(t, http.StatusOK, testCreateCommunicationSettingSuccessResponse)
	defer ts.Close()

	// run the code
	response, err := client.CreateCommunicationSetting(CommunicationSettings{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, response.Data[0].Attributes.Configuration.ChannelName, expectedChannelName)

	assert.Equal(t, response.Data[1].Attributes.Configuration.ChannelName, "testSNSChannel")
	assert.Equal(t, response.Data[1].Attributes.Filter.Statuses[0], "SUCCESS")

	serviceNowChannel := response.Data[2]
	assert.Equal(t, serviceNowChannel.Attributes.Channel, "service-now")
	assert.Equal(t, serviceNowChannel.Attributes.Configuration.ChannelName, "my-test-channel")
	assert.Equal(t, serviceNowChannel.Attributes.Configuration.Type, "problem")
	assert.Equal(t, serviceNowChannel.Attributes.Configuration.Url, "https://mytest001.service-now.com")
	assert.Equal(t, serviceNowChannel.Attributes.Configuration.UserName, "admin")
	assert.Equal(t, serviceNowChannel.Attributes.Configuration.Assignee, "admin")
	assert.Equal(t, serviceNowChannel.Attributes.Configuration.Impact, "3")
	assert.Equal(t, serviceNowChannel.Attributes.Configuration.Urgency, "1")
	assert.Equal(t, serviceNowChannel.Attributes.Configuration.CloseCode, "Closed/Resolved by Caller")
	assert.Equal(t, serviceNowChannel.Attributes.Configuration.CloseNotes, "Issue resolved")
	assert.Equal(t, (serviceNowChannel.Attributes.Configuration.ResolutionOverride.(map[string]interface{}))["closeCode"], "Closed by Caller")
	assert.Equal(t, (serviceNowChannel.Attributes.Configuration.ResolutionOverride.(map[string]interface{}))["closeNotes"], "Issue closed")
	assert.Equal(t, (serviceNowChannel.Attributes.Configuration.CreationOverride.(map[string]interface{}))["urgency"], "2")
}

func TestCreateCommunicationSettingFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	response, err := client.CreateCommunicationSetting(CommunicationSettings{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, response)
}

var testCreateCommunicationSettingSuccessResponse = `
{
	"data": [
	{
      "id": "communication:email-3JD1mkXfz",
      "attributes": {
        "type": "communication",
        "enabled": true,
        "channel": "email",
        "filter": {
          "regions": ["us-east-1"],
          "services": [
            "EC2"
          ]
        },
        "configuration": {
          "channelName": "someChannelName",
          "users": [
              "t-UoU9CsK"
          ]
        }
      },
      "type": "settings",
      "relationships": {
        "account": {
          "data": {
            "type": "accounts",
            "id": "H19NxM15-"
          }
        },
        "organisation": {
          "data": {
            "type": "organisations",
            "id": "ryqMcJn4b"
          }
        }
      }
    },
    {
      "id": "communication:sns-3JD1mAub8",
      "attributes": {
        "type": "communication",
        "channel": "sns",
        "enabled": true,
        "filter": {
          "regions": ["us-east-1"],
          "services": [
            "EC2"
          ],
          "statuses": ["SUCCESS"]
        },
        "configuration": {
          "channelName": "testSNSChannel",
          "arn": "sns-t-UoU9CsK"
        }
      },
      "type": "settings",
      "relationships": {
        "account": {
          "data": {
            "type": "accounts",
            "id": "H19NxM15-"
          }
        },
        "organisation": {
          "data": {
            "type": "organisations",
            "id": "ryqMcJn4b"
          }
        }
      }
    },
    {
      "type": "settings",
      "id": "5678abcd-1234-abcd-5678-efgh12345678:communication:service-now-001",
      "attributes": {
        "type": "communication",
        "manual": false,
        "enabled": true,
        "configuration": {
          "channelName": "my-test-channel",
          "type": "problem",
          "url": "https://mytest001.service-now.com",
          "username": "admin",
          "assignee": "admin",
          "impact": "3",
          "urgency": "1",
          "closeCode": "Closed/Resolved by Caller",
          "closeNotes": "Issue resolved",
          "creationOverride": {
            "urgency": "2"
          },
          "resolutionOverride": {
            "closeCode": "Closed by Caller",
            "closeNotes": "Issue closed"
          }
        },
        "created-by": "abcdeb4b-1234-abcd-5678-efgh12345678",
        "created-date": 1759372960182,
        "last-modified-date": 1759382176376,
        "channel": "service-now",
        "is-account-level": true,
        "is-group-level": false,
        "is-organisation-level": false
      },
      "relationships": {
        "organisation": {
          "data": {
            "type": "organisations",
            "id": "abcdefgh-1234-abcd-5678-efgh12345678"
          }
        },
        "account": {
          "data": {
            "type": "accounts",
            "id": "efghabcd-1234-abcd-5678-efgh12345678"
          }
        },
        "group": {
          "data": null
        },
        "profile": {
          "data": null
        }
      }
    }
  ]
}
`
