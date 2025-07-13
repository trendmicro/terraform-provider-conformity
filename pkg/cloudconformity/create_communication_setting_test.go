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
	assert.Equal(t, expectedChannelName, response.Data[0].Attributes.Configuration.ChannelName)

	assert.Equal(t, "testSNSChannel", response.Data[1].Attributes.Configuration.ChannelName)
	assert.Equal(t, "SUCCESS", response.Data[1].Attributes.Filter.Statuses[0])
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
    }
  ]
}
`
