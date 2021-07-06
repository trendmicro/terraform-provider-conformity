package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommunicationSettingSuccess(t *testing.T) {
	// prepare the test
	expectedCommunicationSettingID := "ryqs8LNKW:communication:email-Ske1cKKEvM"
	client, ts := createHttpTestClient(t, http.StatusOK, testGetCommunicationSettingSuccessResponse)
	defer ts.Close()

	// run the code
	communicationSetting, err := client.GetCommunicationSetting(expectedCommunicationSettingID)
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, communicationSetting.Data.ID, expectedCommunicationSettingID)
}

func TestGetCommunicationSettingFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	GroupDataList, err := client.GetCommunicationSetting("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, GroupDataList)
}

var testGetCommunicationSettingSuccessResponse = `
{
    "data":
        {
            "type": "settings",
            "id": "ryqs8LNKW:communication:email-Ske1cKKEvM",
            "attributes": {
                "type": "communication",
                "manual": false,
                "enabled": true,
                "filter": {
                    "categories": [
                        "security"
                    ],
                    "suppressed": false
                },
                "configuration": {
                    "users": [
                        "HyL7K6GrZ"
                    ]
                },
                "channel": "email"
            },
            "relationships": {
                "organisation": {
                    "data": {
                        "type": "organisations",
                        "id": "ryqMcJn4b"
                    }
                },
                "account": {
                    "data": {
                        "type": "accounts",
                        "id": "H19NxM15-"
                    }
                }
            }
        }
}
`
