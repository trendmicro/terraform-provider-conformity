package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProfileSuccess(t *testing.T) {
	// prepare the test
	expectedProfileID := "d9yHTrzP0"
	response := testCreateUpdateProfileSuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	profileID, err := client.CreateProfileSetting(ProfileSettings{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, profileID, expectedProfileID)
}

func TestCreateProfileFailUnprocessableEntity(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, errResponseUnprocessableEntity)
	defer ts.Close()

	// run the code
	profileID, err := client.CreateProfileSetting(ProfileSettings{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnprocessableEntity)
	assert.Equal(t, profileID, "")
}
func TestCreateProfileFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()

	// run the code
	profileID, err := client.CreateProfileSetting(ProfileSettings{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Equal(t, profileID, "")
}

var testCreateUpdateProfileSuccessResponse = `
{
	"data": {
	  "type": "profiles",
	  "id": "d9yHTrzP0",
	  "attributes": {
		"name": "hemen test 1",
		"description": "hemen test 1"
	  },
	  "relationships": {
		"ruleSettings": {
		  "data": [
			{
			  "type": "rules",
			  "id": "EC2-055"
			},
			{
			  "type": "rules",
			  "id": "EC2-071"
			},
			{
			  "type": "rules",
			  "id": "RTM-007"
			},
			{
			  "type": "rules",
			  "id": "S3-006"
			},
			{
			  "type": "rules",
			  "id": "SNS-002"
			}
		  ]
		}
	  }
	},
	"included": [
	  {
		"type": "rules",
		"id": "EC2-055",
		"attributes": {
		  "enabled": false,
		  "extraSettings": [
			{
			  "type": "single-number-value",
			  "name": "cpuUtilizationThreshold",
			  "value": 2
			}
		  ],
		  "riskLevel": "HIGH",
		  "provider": "aws"
		}
	  },
	  {
		"type": "rules",
		"id": "EC2-071",
		"attributes": {
		  "enabled": false,
		  "riskLevel": "HIGH",
		  "provider": "aws"
		}
	  },
	  {
		"type": "rules",
		"id": "RTM-007",
		"attributes": {
		  "enabled": false,
		  "extraSettings": [
			{
			  "type": "multiple-ip-values",
			  "name": "authorisedIps",
			  "values": [
				{
				  "value": null,
				  "default": null
				}
			  ]
			},
			{
			  "name": "ttl",
			  "type": "ttl",
			  "value": 24,
			  "ttl": true
			}
		  ],
		  "riskLevel": "HIGH",
		  "provider": "aws"
		}
	  },
	  {
		"type": "rules",
		"id": "S3-006",
		"attributes": {
		  "enabled": true,
		  "exceptions": {
			"resources": [
			  "fadfad"
			],
			"tags": [
			  "adsfs"
			]
		  },
		  "extraSettings": null,
		  "riskLevel": "VERY_HIGH",
		  "provider": "aws"
		}
	  },
	  {
		"type": "rules",
		"id": "SNS-002",
		"attributes": {
		  "enabled": false,
		  "riskLevel": "HIGH",
		  "provider": "aws"
		}
	  }
	]
  }
`
