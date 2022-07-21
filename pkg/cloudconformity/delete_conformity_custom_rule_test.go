package cloudconformity

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDeleteConformityCustomRule200(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusOK, testDeleteConformityCustomRule200Response)
	defer ts.Close()

	// run the code
	response, err := client.DeleteCustomRule("CUSTOM-123ABC")
	actual := deleteResponse{}
	actual.Meta.Status = "deleted"

	// check the results
	assert.Nil(t, err)
	assert.True(t, assert.ObjectsAreEqual(response, &actual))
}

func TestDeleteConformityCustomRule403(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusForbidden, testDeleteConformityCustomRule403Response)
	defer ts.Close()

	// run the code
	response, err := client.GetCustomRule("CUSTOM-123ABC")

	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testDeleteConformityCustomRule403Response)
	assert.Empty(t, response)
}

func TestDeleteConformityCustomRule404(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusNotFound, testDeleteConformityCustomRule404Response)
	defer ts.Close()

	// run the code
	response, err := client.DeleteCustomRule("CUSTOM-123ABC")

	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testDeleteConformityCustomRule404Response)
	assert.Empty(t, response)
}

var testDeleteConformityCustomRule200Response = `{
  "meta": {
    "status": "deleted"
  }
}`

var testDeleteConformityCustomRule403Response = `{
  "errors": [
    {
      "status": 403,
      "source": {
        "pointer": "/custom-rules"
      },
      "detail": "Forbidden"
    }
  ]
}`

var testDeleteConformityCustomRule404Response = `{
  "errors": [
    {
      "status": 404,
      "source": {
        "pointer": "/custom-rules/CUSTOM-123ABC"
      },
      "detail": "Custom rule not found"
    }
  ]
}`
