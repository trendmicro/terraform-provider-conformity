package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateAccountBotSettingsSuccess(t *testing.T) {
	// prepare the test
	expectedAccountId := "AgA12vIwb"
	response := "{\"data\": [ {\"type\": \"accounts\",\"id\": \"" + expectedAccountId + "\"} ] }"
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	accountId, err := client.UpdateAccountBotSettings("some-id", (AccountBotSettingsRequest{}))
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, accountId, expectedAccountId)

}

func TestUpdateAccountBotSettingsFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	awsAccountId, err := client.UpdateAccountBotSettings("some-id", (AccountBotSettingsRequest{}))
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Equal(t, awsAccountId, "")

}
