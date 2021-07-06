package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCommunicationSettingSuccess(t *testing.T) {
	// prepare the test
	expectedCommunicationSettingId := "communication:email-3JD1mkXfz"
	client, ts := createHttpTestClient(t, http.StatusOK, testUpdateCommunicationSettingSuccessResponse)
	defer ts.Close()
	// run the code
	id, err := client.UpdateCommunicationSetting(expectedCommunicationSettingId, CommunicationSettings{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, id, expectedCommunicationSettingId)

}

func TestUpdateCommunicationSettingFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	groupId, err := client.UpdateCommunicationSetting("some-id", CommunicationSettings{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Equal(t, groupId, "")
}

var testUpdateCommunicationSettingSuccessResponse = `
{
	"data":
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
    }
}
`
