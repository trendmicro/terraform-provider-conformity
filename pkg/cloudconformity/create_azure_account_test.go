package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAzureAccountSuccess(t *testing.T) {
	// prepare the test
	expectedAccountID := "3ff84b20-0f4c-11eb-a7b7-7d9b3c0e866e"
	response := "{ \"data\": { \"type\": \"accounts\", \"id\": \"" + expectedAccountID + "\" } }"
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	accountID, err := client.CreateAzureAccount(AccountPayload{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, accountID, expectedAccountID)
}

func TestCreateAzureAccountFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()

	// run the code
	accountID, err := client.CreateAzureAccount(AccountPayload{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Equal(t, accountID, "")
}

func TestCreateAzureAccountFailUnprocessableEntity(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, errResponseUnprocessableEntity)
	defer ts.Close()

	// run the code
	accountID, err := client.CreateAzureAccount(AccountPayload{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnprocessableEntity)
	assert.Equal(t, accountID, "")
}
