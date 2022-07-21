package cloudconformity

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetAzureSubscriptions200(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusOK, testGetAzureSubscriptions200Response)
	defer ts.Close()

	// run the code
	response, err := client.GetAzureSubscriptions("some-id")
	// check the results
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, len(response.Data), 1)
	assert.Equal(t, response.Data[0].ID, "AZURE_SUBSCRIPTION_ID")
	assert.Equal(t, response.Data[0].Type, "subscriptions")
	assert.Equal(t, response.Data[0].Attributes.DisplayName, "A Azure Subscription")
	assert.Equal(t, response.Data[0].Attributes.State, "Enabled")
	assert.Equal(t, response.Data[0].Attributes.AddedToConformity, true)
}

func TestGetAzureSubscriptions403(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, testGetAzureSubscriptions403Response)
	defer ts.Close()

	// run the code
	response, err := client.GetAzureSubscriptions("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testGetAzureSubscriptions403Response)
	assert.Empty(t, response)
}

func TestGetAzureSubscriptions404(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusNotFound, testGetAzureSubscriptions404Response)
	defer ts.Close()

	// run the code
	response, err := client.GetAzureSubscriptions("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testGetAzureSubscriptions404Response)
	assert.Empty(t, response)
}

var testGetAzureSubscriptions200Response = `{
  "data": [
    {
      "type": "subscriptions",
      "id": "AZURE_SUBSCRIPTION_ID",
      "attributes": {
        "display-name": "A Azure Subscription",
        "state": "Enabled",
        "added-to-conformity": true
      }
    }
  ]
}`

var testGetAzureSubscriptions403Response = `{
  "Message": "User is not authorized to access this resource with an explicit deny"
}`

var testGetAzureSubscriptions404Response = `{
  "Message": "Not Found"
}`
