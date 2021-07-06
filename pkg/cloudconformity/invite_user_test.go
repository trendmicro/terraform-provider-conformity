package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInviteLegacyUserSuccess(t *testing.T) {
	// prepare the test
	expectedUserID := "uUmE2v0ns"
	response := testInviteUpdateLegacyUserSuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	userID, err := client.InviteLegacyUser(UserDetails{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, userID, expectedUserID)
}

func TestInviteLegacyUserFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()

	// run the code
	userID, err := client.InviteLegacyUser(UserDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Equal(t, userID, "")
}

func TestInviteLegacyUserFailUnprocessableEntity(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, errResponseUnprocessableEntity)
	defer ts.Close()

	// run the code
	userID, err := client.InviteLegacyUser(UserDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnprocessableEntity)
	assert.Equal(t, userID, "")
}

var testInviteUpdateLegacyUserSuccessResponse = `{
	"data": {
	  "type": "users",
	  "id": "uUmE2v0ns",
	  "attributes": {
		"first-name": "John",
		"last-name": "Smith",
		"role": "ADMIN",
		"email": "john.smith@company.com",
		"status": "ACTIVE",
		"mfa": true,
		"last-login-date": 1608083636734,
		"created-date": 1588825954131
	  },
	  "relationships": {
		"organisation": {
		  "data": {
			"type": "organisations",
			"id": "B1nHYYpwx"
		  }
		}
	  }
	}
  }`
