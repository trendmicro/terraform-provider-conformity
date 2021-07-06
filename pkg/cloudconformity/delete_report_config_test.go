package cloudconformity

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteReportConfigSuccess(t *testing.T) {
	// prepare the test
	response := `{ "meta": { "status": "deleted" } }`
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	// run the code
	deleteReportResponse, err := client.DeleteReportConfig("delete-report-id")
	// check the results
	assert.Nil(t, err)
	assert.Equal(t, "deleted", deleteReportResponse.Meta.Status)
}
func TestDeleteReportConfigFailUnauthorized(t *testing.T) {
	// prepare the test
	client, ts := createHttpTestClient(t, http.StatusForbidden, errResponseUnauthorized)
	defer ts.Close()
	// run the code
	deleteReportResponse, err := client.DeleteReportConfig("delete-report-id")
	// check the results
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponseUnauthorized)
	assert.Nil(t, deleteReportResponse)
}
