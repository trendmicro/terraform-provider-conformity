package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGCPAccountSuccess(t *testing.T) {
	// prepare the test
	expectedAccountID := "3ff84b20-0f4c-11eb-a7b7-7d9b3c0e866e"
	response := "{ \"data\": { \"type\": \"accounts\", \"id\": \"" + expectedAccountID + "\" } }"
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	accountID, err := client.CreateGCPAccount(AccountPayload{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, accountID, expectedAccountID)
}

func TestCreateGCPAccountFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	accountID, err := client.CreateGCPAccount(AccountPayload{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Equal(t, accountID, "")
}
