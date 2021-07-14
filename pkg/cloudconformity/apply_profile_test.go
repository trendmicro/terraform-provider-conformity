package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyProfileSuccess(t *testing.T) {

	response := `{"meta": {"status": "sent","message": "Profile will be applied to the accounts in background"}}`
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	applyProfileResponse, err := client.CreateApplyProfile("some-id", ApplyProfileSettings{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, applyProfileResponse.Meta.Status, "sent")
}

func TestApplyProfileFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	applyProfileResponse, err := client.CreateApplyProfile("some-id", ApplyProfileSettings{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Empty(t, applyProfileResponse)
}
