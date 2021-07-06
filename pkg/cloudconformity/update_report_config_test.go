package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateReportConfigSuccess(t *testing.T) {
	// prepare the test
	expectedReportID := "C1LBzx2:report-config:H19NxMi5-"
	response := testCreateupdateReportGroupResponseSuccess
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	reportID, err := client.UpdateReportConfig("some-id", ReportConfigDetails{})
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, reportID, expectedReportID)

}

func TestUpdateReportConfigFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()
	// run the code
	reportID, err := client.UpdateReportConfig("some-id", ReportConfigDetails{})
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Equal(t, reportID, "")
}
