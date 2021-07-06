package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLegacyUserSuccess(t *testing.T) {
	// prepare the test
	expecteduserID := "CClqMqknVb"
	response := testGetLegacyUsersuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	userDetails, err := client.GetLegacyUser("some-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, userDetails.Data.ID, expecteduserID)
}

func TestGetLegacyUserFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()

	// run the code
	userDetails, err := client.GetLegacyUser("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Nil(t, userDetails)
}

func TestGetLegacyUserFailForbidden(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	userDetails, err := client.GetLegacyUser("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, userDetails)
}

func TestGetLegacyUserFailUnprocessableEntity(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, testLegacyUserFailUnprocessableEntityUser)
	defer ts.Close()

	// run the code
	userDetails, err := client.GetLegacyUser("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testLegacyUserFailUnprocessableEntityUser)
	assert.Nil(t, userDetails)
}

var testGetLegacyUsersuccessResponse = `
{
    "data": {
        "type": "users",
        "id": "CClqMqknVb",
        "attributes": {
            "first-name": "Cool",
            "last-name": "Claude",
            "role": "ADMIN",
            "email": "cc@coolclaude.com",
            "status": "ACTIVE",
            "last-login-date": 1523009079960,
            "created-date": 1499359762438,
            "summary-email-opt-out": true,
            "mobile": "15144008080",
            "mobile-country-code": "CA",
            "mobile-verified": true
        },
        "relationships": {
            "organisation": {
                "data": {
                    "type": "organisations",
                    "id": "A9NDYY12z"
                }
            }
        }
    }
}
`
var testLegacyUserFailUnprocessableEntityUser = `
{
	"errors": [
	  {
		"status": 422,
		"source": {
		  "pointer": "/data/attributes"
		},
		"details": "No such user"
	  }
	]
  }`
