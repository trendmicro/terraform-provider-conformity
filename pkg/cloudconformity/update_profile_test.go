package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateProfileSuccess(t *testing.T) {
	// prepare the test
	expectedProfileId := "d9yHTrzP0"
	response := testCreateUpdateProfileSuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	profileId, err := client.UpdateProfile("some-id", ProfileSettings{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, profileId, expectedProfileId)
}

func TestUpdateProfileFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	profileId, err := client.UpdateProfile("some-id", ProfileSettings{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Equal(t, profileId, "")
}

func TestUpdateProfileFailUnprocessableEntity(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponseUnprocessableEntity)
	defer ts.Close()
	// run the code
	profileId, err := client.UpdateProfile("some-id", ProfileSettings{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnprocessableEntity)
	assert.Equal(t, profileId, "")
}

func TestUpdateProfileFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()
	// run the code
	profileId, err := client.UpdateProfile("some-id", ProfileSettings{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Equal(t, profileId, "")
}
