package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRevokeLegacyUserSuccess(t *testing.T) {
	// prepare the test
	response := `{ "meta": { "status": "revoked" } }`
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	RevokeUserResponse, err := client.RevokeLegacyUser("revoke-user-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, "revoked", RevokeUserResponse.Meta.Status)
}

func TestRevokeLegacyUserFailForbidden(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	RevokeUserResponse, err := client.RevokeLegacyUser("revoke-user-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, RevokeUserResponse)
}
func TestRevokeLegacyUserFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()
	// run the code
	RevokeUserResponse, err := client.RevokeLegacyUser("revoke-user-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Nil(t, RevokeUserResponse)
}

func TestRevokeLegacyUserFailUnprocessableEntity(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, testLegacyUserFailUnprocessableEntityUser)
	defer ts.Close()
	// run the code
	RevokeUserResponse, err := client.RevokeLegacyUser("revoke-user-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testLegacyUserFailUnprocessableEntityUser)
	assert.Nil(t, RevokeUserResponse)
}
