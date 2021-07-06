package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGroupSuccess(t *testing.T) {
	// prepare the test
	expectedGroupID := "uUmE2v0ns"
	response := testGetGroupsuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	groupDataList, err := client.GetGroup("some-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, groupDataList.Data[0].ID, expectedGroupID)
}

func TestGetGroupFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	GroupDataList, err := client.GetGroup("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, GroupDataList)
}

var testGetGroupsuccessResponse = `
{
	"data": [{
		"type": "groups",
		"id": "uUmE2v0ns",
		"attributes": {
			"name": "test-group",
			"tags": [],
			"created-date": 1587441074460,
			"last-modified-date": 1590647034893
		},
		"relationships": {
			"organisation": {
				"data": {
					"type": "organisations",
					"id": "B1nHYYpwx"
				}
			},
			"accounts": {
				"data": [{
					"type": "accounts",
					"id": "16gZQXGZf"
				}]
			}
		}
	}]
}
`
