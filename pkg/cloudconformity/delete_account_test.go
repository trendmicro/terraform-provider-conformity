package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteAccountSuccess(t *testing.T) {
	// prepare the test
	response := `{ "meta": { "status": "sent" } }`
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	deleteAccountResponse, err := client.DeleteAccount("delete-aacount-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, "sent", deleteAccountResponse.Meta.Status)
}

func TestDeleteAccountFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	deleteAccountResponse, err := client.DeleteAccount("delete-aacount-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, deleteAccountResponse)
}
