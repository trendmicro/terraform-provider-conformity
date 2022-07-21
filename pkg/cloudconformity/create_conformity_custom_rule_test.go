package cloudconformity

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateConformityCustomRule200(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusOK, testCreateConformityCustomRule200Response)
	defer ts.Close()

	// run the code
	rule := initializeCustomRuleObj()
	response, err := client.CreateConformityCustomRule(rule)

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

func TestCreateConformityCustomRule403(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusForbidden, testCreateConformityCustomRule403Response)
	defer ts.Close()

	// run the code
	rule := initializeCustomRuleObj()
	response, err := client.CreateConformityCustomRule(rule)

	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testCreateConformityCustomRule403Response)
	assert.Empty(t, response)
}

func TestCreateConformityCustomRule422(t *testing.T) {
	// mock server
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, testCreateConformityCustomRule422Response)
	defer ts.Close()

	// run the code
	rule := initializeCustomRuleObj()
	response, err := client.CreateConformityCustomRule(rule)

	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, testCreateConformityCustomRule422Response)
	assert.Empty(t, response)
}

func initializeCustomRuleObj() CustomRuleRequest {
	rule := CustomRuleRequest{
		Name:             "S3 Bucket Custom Rule",
		Description:      "This custom rule ensures S3 buckets follow our best practice",
		RemediationNotes: "If this is broken, please follow these steps:\\n1. Step one \\n2. Step two\\n",
		Service:          "S3",
		ResourceType:     "s3-bucket",
		Categories:       []string{"security"},
		Severity:         "HIGH",
		Provider:         "aws",
		Enabled:          true,
		Attributes: []CustomRuleAttributes{
			{
				Name:     "bucketName",
				Path:     "data.Name",
				Required: true,
			},
		},
	}
	conditions := []CustomRuleCondition{
		{
			Fact:     "bucketName",
			Operator: "pattern",
			Value:    "^([a-zA-Z0-9_-]){1,32}$",
		},
	}
	ruleRules := CustomRuleRules{}
	ruleRules.Event = CustomRuleEvent{Type: "Bucket name is longer than 32 characters"}
	ruleRules.Conditions.Any = conditions
	rule.Rules = append(rule.Rules, ruleRules)
	return rule
}

var testCreateConformityCustomRule200Response = `{
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

var testCreateConformityCustomRule403Response = `{
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

var testCreateConformityCustomRule422Response = `{
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
