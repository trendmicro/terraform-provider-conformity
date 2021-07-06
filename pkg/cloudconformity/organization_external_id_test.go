package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//check if external id return valid response and if contains error
//if the external id is empty it will fail
func TestConformityExternalIdSuccess(t *testing.T) {
	// prepare the test
	expectedExternalID := "3ff84b20-0f4c-11eb-a7b7-7d9b3c0e866e"
	response := "{ \"data\": { \"type\": \"external-ids\", \"id\": \"" + expectedExternalID + "\" } }"
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	externalID, err := client.GetExternalId()
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, externalID, "3ff84b20-0f4c-11eb-a7b7-7d9b3c0e866e")
}

func TestConformityExternalIdFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	externalID, err := client.GetExternalId()
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Equal(t, externalID, "")
}
