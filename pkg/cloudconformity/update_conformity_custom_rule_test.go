package cloudconformity

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUpdateConformityCustomRule200(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusOK, testUpdateConformityCustomRule200Response)
	defer ts.Close()

	// run the code
	rule := initializeCustomRuleObj()
	response, err := client.UpdateCustomRule("CUSTOM-123ABC", rule)

	// check the results
	assert.Nil(t, err)
	assert.NotNil(t, response.ID)
	assert.NotNil(t, response.Type)
	assert.Equal(t, response.Attributes.Name, rule.Name)
	assert.Equal(t, response.Attributes.Description, rule.Description)
	assert.Equal(t, response.Attributes.Provider, rule.Provider)
	assert.Equal(t, response.Attributes.Service, rule.Service)
	assert.Equal(t, response.Attributes.ResourceType, rule.ResourceType)
	assert.Equal(t, response.Attributes.Enabled, rule.Enabled)
	assert.True(t, assert.ObjectsAreEqual(response.Attributes.Categories, rule.Categories))
	assert.True(t, assert.ObjectsAreEqual(response.Attributes.Attributes, rule.Attributes))
	assert.True(t, assert.ObjectsAreEqual(response.Attributes.Rules, rule.Rules))
}

func TestUpdateConformityCustomRule403(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusForbidden, testUpdateConformityCustomRule403Response)
	defer ts.Close()

	// run the code
	rule := initializeCustomRuleObj()
	response, err := client.UpdateCustomRule("CUSTOM-123ABC", rule)

	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testUpdateConformityCustomRule403Response)
	assert.Empty(t, response)
}

func TestUpdateConformityCustomRule422(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, testUpdateConformityCustomRule422Response)
	defer ts.Close()

	// run the code
	rule := initializeCustomRuleObj()
	response, err := client.UpdateCustomRule("CUSTOM-123ABC", rule)

	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testUpdateConformityCustomRule422Response)
	assert.Empty(t, response)
}

func TestUpdateConformityCustomRule404(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusNotFound, testUpdateConformityCustomRule404Response)
	defer ts.Close()

	// run the code
	rule := initializeCustomRuleObj()
	response, err := client.UpdateCustomRule("CUSTOM-123ABC", rule)

	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testUpdateConformityCustomRule404Response)
	assert.Empty(t, response)
}

var testUpdateConformityCustomRule200Response = `{
  "data": {
    "type": "CustomRules",
    "id": "CUSTOM-123ABC",
    "attributes": {
      "name": "S3 Bucket Custom Rule",
      "description": "This custom rule ensures S3 buckets follow our best practice",
      "remediationNotes": "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n",
      "service": "S3",
      "resourceType": "s3-bucket",
      "categories": [
        "security"
      ],
      "severity": "HIGH",
      "provider": "aws",
      "enabled": true,
      "attributes": [
        {
          "name": "bucketName",
          "path": "data.Name",
          "required": true
        }
      ],
      "rules": [
        {
          "conditions": {
            "any": [
              {
                "fact": "bucketName",
                "operator": "pattern",
                "value": "^([a-zA-Z0-9_-]){1,32}$"
              }
            ]
          },
          "event": {
            "type": "Bucket name is longer than 32 characters"
          }
        }
      ]
    }
  }
}`

var testUpdateConformityCustomRule403Response = `{
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

var testUpdateConformityCustomRule404Response = `{
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

var testUpdateConformityCustomRule422Response = `{
  "errors": [
    {
      "status": 422,
      "source": {
        "pointer": "/data/attributes/configuration/name"
      },
      "detail": "Configuration name is required"
    }
  ]
}`
