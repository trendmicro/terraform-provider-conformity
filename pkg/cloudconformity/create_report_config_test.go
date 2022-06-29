package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReportConfigSuccess(t *testing.T) {
	// prepare the test
	expectedReportID := "C1LBzx2:report-config:H19NxMi5-"
	response := testCreateupdateReportGroupResponseSuccess
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	reportID, err := client.CreateReportConfig(ReportConfigDetails{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, reportID, expectedReportID)
}

func TestCreateReportConfigFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()
	// run the code
	reportID, err := client.CreateReportConfig(ReportConfigDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Equal(t, reportID, "")
}
func TestCreateReportConfigFailUnprocessableEntity(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnprocessableEntity, errResponseUnprocessableEntity)
	defer ts.Close()
	// run the code
	reportID, err := client.CreateReportConfig(ReportConfigDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnprocessableEntity)
	assert.Equal(t, reportID, "")
}

var testCreateupdateReportGroupResponseSuccess = `
{
	"data": {
		"type": "accounts",
		"id": "C1LBzx2:report-config:H19NxMi5-",
		"attributes": {
			"type": "report-config",
			"enabled": true,
			"configuration": {
				"title": "Daily report of IAM",
				"sendEmail": false,
				"emails": [
					"string"
				],
				"filter": {
					"services": [
						"EC2"
					],
					"resourceTypes": [
						"kms-key"
					],
					"regions": [
						"us-west-1"
					],
					"ruleIds": [
						"EC2-001"
					],
					"tags": [
						"string"
					],
					"filterTags": [
						"string"
					],
					"text": "S3",
					"createdLessThanDays": 5,
					"createdMoreThanDays": 0,
					"newerThanDays": 5,
					"olderThanDays": 5,
					"categories": [
						"security"
					],
					"riskLevels": ["HIGH"],
					"complianceStandards": [
						"NIST4",
						"AWAF"
					],
					"reportComplianceStandardId": "NIST4",
					"statuses": [
						"SUCCESS"
					],
					"suppressedFilterMode": "v1",
					"suppressed": true,
					"providers": [
						"aws"
					],
					"resource": "string",
					"resourceSearchMode": "text",
					"message": true,
					"withChecks": false,
					"withoutChecks": false
				},
				"scheduled": false,
				"frequency": "* * *",
				"tz": "Australia/Sydney",
				"generateReportType": "COMPLIANCE-STANDARD",
				"includeChecks": false,
				"shouldEmailIncludePdf": true,
				"shouldEmailIncludeCsv": true
			},
			"created-by": "f5dBnv_",
			"created-date": 0,
			"is-account-level": false,
			"is-group-level": false,
			"is-organisation-level": false
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
			},
			"group": {
				"data": {
					"type": "groups",
					"id": "DaZbc2jd2"
				}
			}
		}
	}
}
`
