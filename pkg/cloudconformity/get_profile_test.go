package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProfileSuccess(t *testing.T) {
	// prepare the test
	expectedProfileID := "d9yHTrzP0"
	response := testCreateUpdateProfileSuccessResponse
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	profileSettings, err := client.GetProfile("some-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, profileSettings.Data.ID, expectedProfileID)
}
func TestGetProfileGenericValuesSuccess(t *testing.T) {
	// prepare the test
	expectedProfileID := "d9yHTrzP0"
	response := testGetProfileSuccessResponseGenericProfileValues
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	profileSettings, err := client.GetProfile("some-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, profileSettings.Data.ID, expectedProfileID)
	for _, ele := range profileSettings.Included {
		assert.NotNil(t, ele.Attributes.ExtraSettings)
		assert.Greater(t, len(ele.Attributes.ExtraSettings), 0)
		for _, extraEle := range ele.Attributes.ExtraSettings {
			assert.Greater(t, len(extraEle.Values), 0)
			for _, genericProfileElement := range extraEle.Values {
				values, isStruct := genericProfileElement.(map[string]interface{})
				if isStruct {
					assert.NotNil(t, genericProfileElement)
					assert.NotEmpty(t, values["value"])
					assert.True(t, values["enabled"].(bool))
					assert.Equal(t, values["label"], "_label")
				} else {
					assert.NotEmpty(t, genericProfileElement)
				}
			}
		}
	}
}
func TestGetProfileFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	profileSettings, err := client.GetProfile("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, profileSettings)
}

/*
Allows testing two cases:
  1- when the value is int/float
  2- when the value is []string
Refer to bug reported in:
https://github.com/trendmicro/terraform-provider-conformity/issues/32
*/
var testGetProfileSuccessResponseGenericProfileValues = `
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
            "id": "RTM-008"
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
        "enabled": true,
        "extraSettings": [
          {
            "name": "commonlyUsedPorts",
            "type": "multiple-number-values",
            "values": [
              { 
				"label": "_label",
				"enabled": true,
                "value": 80
              },
              {
				"label": "_label",
				"enabled": true,
                "value": 443
              },
              {
				"label": "_label",
				"enabled": true,
                "value": 12.34
              }
            ]
          }
        ],
        "riskLevel": "HIGH",
        "provider": "aws"
      }
    },
    {
      "type": "rules",
      "id": "RTM-008",
      "attributes": {
        "enabled": true,
        "extraSettings": [
          {
            "name": "authorisedRegions",
            "regions": true,
            "type": "regions",
            "values": [
              "af-south-1",
              "ap-east-1",
              "ap-northeast-1",
              "ap-northeast-2"
            ]
          }
        ]
      }
    }
  ]
}
`
