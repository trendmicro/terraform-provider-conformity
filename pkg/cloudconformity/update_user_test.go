package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateLegacyUserSuccess(t *testing.T) {
	// prepare the test
	expectedUserID := "uUmE2v0ns"
	response := testInviteUpdateLegacyUserSuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	groupId, err := client.UpdateLegacyUser("some-id", UserAccessDetails{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, groupId, expectedUserID)

}

func TestUpdateLegacyUserFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()
	// run the code
	userId, err := client.UpdateLegacyUser("some-id", UserAccessDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Equal(t, userId, "")

}

func TestUpdateLegacyUserFailUnprocessableEntity(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, errResponseUnprocessableEntity)
	defer ts.Close()
	// run the code
	userId, err := client.UpdateLegacyUser("some-id", UserAccessDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnprocessableEntity)
	assert.Equal(t, userId, "")

}
