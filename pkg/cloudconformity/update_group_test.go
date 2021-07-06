package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateGroupSuccess(t *testing.T) {
	// prepare the test
	expectedGroupId := "uUmE2v0ns"
	response := testCreateUpdateGroupSuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	groupId, err := client.UpdateGroup("some-id", GroupDetails{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, groupId, expectedGroupId)

}

func TestUpdateGroupFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	groupId, err := client.UpdateGroup("some-id", GroupDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Equal(t, groupId, "")

}
