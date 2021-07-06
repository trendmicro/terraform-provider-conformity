package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateAwsAccountSuccess(t *testing.T) {
	// prepare the test
	expectedAwsAccountId := "H19NxMi5-"
	response := "{\"data\": {\"type\": \"accounts\",\"id\": \"" + expectedAwsAccountId + "\"}}"
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	awsAccountId, err := client.UpdateAccount("some-id", (AccountPayload{}))
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, awsAccountId, expectedAwsAccountId)

}

func TestUpdateAwsAccountFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	awsAccountId, err := client.UpdateAccount("some-id", (AccountPayload{}))
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Equal(t, awsAccountId, "")

}
