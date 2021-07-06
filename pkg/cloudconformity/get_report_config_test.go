package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReportConfigSuccess(t *testing.T) {
	// prepare the test
	expectedReportID := "C1LBzx2:report-config:H19NxMi5-"
	response := testCreateupdateReportGroupResponseSuccess
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()

	// run the code
	reportConfigDetails, err := client.GetReportConfig("some-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, reportConfigDetails.Data.ID, expectedReportID)
}

func TestGetReportConfigFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusUnauthorized, errResponseUnauthorized)
	defer ts.Close()
	// run the code
	reportConfigDetails, err := client.GetReportConfig("some-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Nil(t, reportConfigDetails)
}
