package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSsoLegacyUserSuccess(t *testing.T) {
	// prepare the test
	expectedUserID := "uUmE2v0ns"
	response := testInviteUpdateLegacyUserSuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	userID, err := client.CreateSsoLegacyUser(UserDetails{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, userID, expectedUserID)
}

func TestCreateSsoLegacyUserFailUnAuthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()

	// run the code
	userID, err := client.CreateSsoLegacyUser(UserDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Equal(t, userID, "")
}

func TestCreateSsoLegacyUserFailUnprocessableEntity(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, errResponseUnprocessableEntity)
	defer ts.Close()

	// run the code
	userID, err := client.CreateSsoLegacyUser(UserDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnprocessableEntity)
	assert.Equal(t, userID, "")
}
