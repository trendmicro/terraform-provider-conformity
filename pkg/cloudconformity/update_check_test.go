package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCheckSuccess(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusOK, testUpdateCheckSuccessResponse)
	defer ts.Close()
	// run the code
	check, err := client.UpdateCheck("ccc:8d99dfce-dca2-4f13-8699-20631a5c77c9:Resources-001:Resources:global:%2Fsubscriptions%2Fae9124e0-d61c-4d7d-833d-c58e6f9941f8%2Fproviders%2FMicrosoft.Authorization%2FroleDefinitions%2F00482a5a-887f-4fb3-b363-3b7fe8e74483", CheckDetails{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, check.Data.Id, "ccc:8d99dfce-dca2-4f13-8699-20631a5c77c9:Resources-001:Resources:global:/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483")
	assert.Equal(t, check.Data.Type, "checks")
	assert.Equal(t, check.Data.Attributes.Suppressed, false)
	assert.Equal(t, check.Data.Attributes.Region, "global")
	assert.Equal(t, check.Data.Attributes.Resource, "/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483")
	assert.Equal(t, check.Data.Relationships.Rule.Data.Id, "Resources-001")
	assert.Equal(t, check.Data.Relationships.Account.Data.Id, "8d99dfce-dca2-4f13-8699-20631a5c77c9")
}

func TestUpdateCheckFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	_, err := client.UpdateCheck("ccc:8d99dfce-dca2-4f13-8699-20631a5c77c9:Resources-001:Resources:global:%2Fsubscriptions%2Fae9124e0-d61c-4d7d-833d-c58e6f9941f8%2Fproviders%2FMicrosoft.Authorization%2FroleDefinitions%2F00482a5a-887f-4fb3-b363-3b7fe8e74483", CheckDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
}

func TestUpdateCheckInvalidCheckId(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()
	// run the code
	checkDetails, err := client.UpdateCheck("invalid-check-id", CheckDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid check-id: 'invalid-check-id'")
	assert.Nil(t, checkDetails)
}

var testUpdateCheckSuccessResponse = `
{
  "data": {
    "type": "checks",
    "id": "ccc:8d99dfce-dca2-4f13-8699-20631a5c77c9:Resources-001:Resources:global:/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483",
    "attributes": {
      "region": "global",
      "status": "FAILURE",
      "risk-level": "LOW",
      "pretty-risk-level": "Low",
      "message": "00482a5a-887f-4fb3-b363-3b7fe8e74483 has [Owner, CostID, Application, DemandID, Stage, ITServiceID, ITServiceName] tags missing",
      "resource": "/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483",
      "descriptorType": "access-control-roles",
      "link-title": "/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483",
      "resourceName": "00482a5a-887f-4fb3-b363-3b7fe8e74483",
      "last-modified-date": 1656482101900,
      "created-date": 1653906818310,
      "categories": [
        "security",
        "reliability",
        "performance-efficiency",
        "cost-optimisation",
        "operational-excellence"
      ],
      "compliances": [
        "NIST5",
        "NIST-CSF",
        "AGISM",
        "HITRUST",
        "MAS",
        "CSA"
      ],
      "suppressed": false,
      "failure-discovery-date": 1656482101900,
      "ccrn": "ccrn:azure:8d99dfce-dca2-4f13-8699-20631a5c77c9:AccessControl:global:/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483",
      "extradata": [
        {
          "name": "DETAILED_STATUS",
          "label": "Resource tags status for 00482a5a-887f-4fb3-b363-3b7fe8e74483",
          "value": "{\"service\":\"AccessControl\",\"resourceName\":\"00482a5a-887f-4fb3-b363-3b7fe8e74483\",\"tags\":[{\"key\":\"Owner\",\"hasValue\":false},{\"key\":\"CostID\",\"hasValue\":false},{\"key\":\"Application\",\"hasValue\":false},{\"key\":\"DemandID\",\"hasValue\":false},{\"key\":\"Stage\",\"hasValue\":false},{\"key\":\"ITServiceID\",\"hasValue\":false},{\"key\":\"ITServiceName\",\"hasValue\":false}]}",
          "type": "META",
          "internal": true
        }
      ],
      "tags": [],
      "suppressed-until": null,
      "not-scored": false,
      "excluded": false,
      "rule-title": "Tags",
      "provider": "azure",
      "resolution-page-url": "https://www.cloudconformity.com/knowledge-base/azure/Resources/use-tags-organize-resources.html#503824401549",
      "service": "Resources"
    },
    "relationships": {
      "rule": {
        "data": {
          "type": "rules",
          "id": "Resources-001"
        }
      },
      "account": {
        "data": {
          "type": "accounts",
          "id": "8d99dfce-dca2-4f13-8699-20631a5c77c9"
        }
      }
    }
  }
}`
