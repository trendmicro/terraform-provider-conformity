package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteProfileSuccess(t *testing.T) {
	// prepare the test
	response := `{ "meta": { "status": "deleted" } }`
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	deleteProfileResponse, err := client.DeleteProfile("delete-profile-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, "deleted", deleteProfileResponse.Meta.Status)
}

func TestDeleteProfileFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	deleteProfileResponse, err := client.DeleteProfile("delete-profile-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, deleteProfileResponse)
}
