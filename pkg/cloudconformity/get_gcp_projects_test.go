package cloudconformity

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetGcpProjects200(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusOK, testGetGcpProjects200Response)
	defer ts.Close()

	// run the code
	response, err := client.GetGcpProjects("some-id")
	// check the results
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, len(response.Data), 1)
	assert.Equal(t, response.Data[0].Type, "projects")
	assert.Equal(t, response.Data[0].Attributes.Name, "My Project")
	assert.Equal(t, response.Data[0].Attributes.ProjectNumber, "415104041262")
	assert.Equal(t, response.Data[0].Attributes.ProjectID, "project-id-1")
	assert.Equal(t, response.Data[0].Attributes.LifecycleState, "ACTIVE")
	assert.Equal(t, response.Data[0].Attributes.AddedToConformity, true)
	assert.Equal(t, response.Data[0].Attributes.AddedToConformity, true)
	assert.Equal(t, response.Data[0].Attributes.CreateTime.String(), "2021-05-17 11:21:58.012 +0000 UTC")
	assert.Equal(t, response.Data[0].Attributes.Parent.ID, "415104041262")
	assert.Equal(t, response.Data[0].Attributes.Parent.Type, "folder")

}

func TestGetGcpProjects403(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, testGetGcpProjects403Response)
	defer ts.Close()

	// run the code
	response, err := client.GetGcpProjects("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testGetGcpProjects403Response)
	assert.Empty(t, response)
}

func TestGetGcpProjects404(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusNotFound, testGetGcpProjects404Response)
	defer ts.Close()

	// run the code
	response, err := client.GetGcpProjects("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testGetGcpProjects404Response)
	assert.Empty(t, response)
}

func TestGetGcpProjects500(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusInternalServerError, testGetGcpProjects500Response)
	defer ts.Close()

	// run the code
	response, err := client.GetGcpProjects("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testGetGcpProjects500Response)
	assert.Empty(t, response)
}

var testGetGcpProjects200Response = `{
  "data": [
    {
      "type": "projects",
      "attributes": {
        "project-number": "415104041262",
        "project-id": "project-id-1",
        "lifecycle-state": "ACTIVE",
        "added-to-conformity": true,
        "create-time": "2021-05-17T11:21:58.012Z",
        "name": "My Project",
        "parent": {
          "type": "folder",
          "id": "415104041262"
        }
      }
    }
  ]
}`
var testGetGcpProjects403Response = `{
  "Message": "User is not authorized to access this resource with an explicit deny"
}`
var testGetGcpProjects404Response = `{
  "Message": "Group not found."
}`
var testGetGcpProjects500Response = `{
  "errors": [
    {
      "status": 500,
      "detail": "Internal Server Error"
    }
  ]
}`
