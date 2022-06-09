package cloudconformity

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountSuccess(t *testing.T) {
	// prepare the test
	expectedExternalID := "XTLFTLAXVS7G"
	ts := createHttpTestClientGetAccount(t)
	defer ts.Close()
	client := Client{Region: "TEST-REGION", Apikey: "TEST-APIKEY", Url: ts.URL, HttpClient: ts.Client()}
	// run the code
	accountAccessAndDetails, err := client.GetAccount("some-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, expectedExternalID, accountAccessAndDetails.AccessSettings.Attributes.Configuration.ExternalId)
}

func TestGetAccountFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	accountAccessAndDetails, err := client.GetAccount("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, accountAccessAndDetails)
}

func TestGetAccountAccessSettingsFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	accountAccessAndDetails, err := client.GetAccountAccessSettings("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, accountAccessAndDetails)
}
func TestGetAccountDetailsFail(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponse)
	defer ts.Close()

	// run the code
	accountAccessAndDetails, err := client.GetAccountDetails("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, accountAccessAndDetails)
}

var testAccessSettingsResponse = `{
	"id": "BJ0Ox16Hb:access",
	"type": "settings",
	"attributes": {
	  "type": "access",
	  "configuration": {
		"externalId": "XTLFTLAXVS7G",
		"roleArn": "arn:aws:iam::222274792222:role/myRole"
	  }
	},
	"relationships": {
	  "organisation": {
		"data": {
		  "type": "organisations",
		  "id": "B1nHYYpwx"
		}
	  },
	  "account": {
		"data": {
		  "type": "accounts",
		  "id": "BJ0Ox16Hb"
		}
	  }
	}
  }`
var testAccessDetailsResponse = `{
	"data": {
	  "type": "accounts",
	  "id": "H19NxMi5-",
	  "attributes": {
		"name": "MyAccount",
		"environment": "Prod",
		"awsaccount-id": 123456789012,
		"status": "ACTIVE",
		"security-package": true,
		"cost-package": false,
		"created-date": 1505595441887,
		"last-notified-date": 1505595441887,
		"last-checked-date": 1505595441887,
		"last-monitoring-event-date": 1505595441887,
		"billing-account-id": "r1gyR4cqg",
		"is-billing-account": true,
		"bot-status": "RUNNING",
		"cloud-type": "aws",
		"managed-group-id": "rhGZeSTwT",
		"settings": {
		  "communication": {
			"channels": [
			  {
				"name": "email",
				"users": [
				  null
				],
				"enabled": true,
				"levels": [
				  null
				]
			  }
			]
		  },
		  "rules": [
			{
			  "enabled": false,
			  "id": "S3-021",
			  "riskLevel": "HIGH"
			}
		  ],
		  "bot": {
			"disabledRegions": {
			  "us-east-1": true,
			  "us-west-2": true
			},
			"lastModifiedFrom": "13.237.98.102",
			"disabled": false,
			"disabledUntil": 1505595441887,
			"delay": 2,
			"lastModifiedBy": "NHohT7Gr7"
		  },
		  "access": {
			"type": "CROSS_ACCOUNT",
			"stackId": "arn:aws:cloudformation:us-east-1:123456789012:stack/CloudConformity/56db5b90-7ebb-11e7-8a78-500c28902e99"
		  }
		},
		"cost": {
		  "version": "1.0.14",
		  "last-updated-date": 1599184207280,
		  "summary": {
			"mostExpensiveService": "Amazon Route 53",
			"waste": 4,
			"unit": "USD",
			"costToDate": 1.0849829999999998,
			"billId": "2020-09",
			"optimisationRate": 0,
			"forecast": 10.577014169151289,
			"optimisationOpportunityCount": 1
		  },
		  "billing-account-map": {
			"payerAccount": {
			  "awsId": 917125559992,
			  "id": "A7rUbWdt0"
			},
			"linkedAccounts": [
			  {
				"awsId": 917125559992,
				"id": "A7rUbWdt0"
			  }
			]
		  },
		  "bills": [
			{
			  "accountCost": 9.577014169151289,
			  "id": "2020-09",
			  "status": "pending",
			  "current": true
			}
		  ]
		}
	  },
	  "relationships": {
		"organisation": {
		  "data": {
			"type": "organisations",
			"id": "B1nHYYpwx"
		  }
		}
	  }
	}
  }`

func createHttpTestClientGetAccount(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		var getAccessSettings = regexp.MustCompile(`^/accounts/(.*?)/access$`)
		var getAccessDetails = regexp.MustCompile(`^/accounts/(.*?)$`)
		switch {
		case getAccessSettings.MatchString(r.URL.Path):
			fmt.Println("createHttpTestClientGetAccount - getAccessSettings")
			w.Write([]byte(testAccessSettingsResponse))
		case getAccessDetails.MatchString(r.URL.Path):
			fmt.Println("createHttpTestClientGetAccount - getAccessDetails")
			w.Write([]byte(testAccessDetailsResponse))
		}

	}))
}
