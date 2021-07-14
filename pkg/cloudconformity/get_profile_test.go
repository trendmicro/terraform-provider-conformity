package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProfileSuccess(t *testing.T) {
	// prepare the test
	expectedProfileID := "d9yHTrzP0"
	response := testCreateUpdateProfileSuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	profileSettings, err := client.GetProfile("some-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, profileSettings.Data.ID, expectedProfileID)
}
func TestGetProfileFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	profileSettings, err := client.GetProfile("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, profileSettings)
}
