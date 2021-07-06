package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteGroupSuccess(t *testing.T) {
	// prepare the test
	response := `{ "meta": { "status": "deleted" } }`
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	deleteGroupResponse, err := client.DeleteGroup("delete-group-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, "deleted", deleteGroupResponse.Meta.Status)
}

func TestDeleteGroupFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	deleteGroupResponse, err := client.DeleteGroup("delete-group-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, deleteGroupResponse)
}
