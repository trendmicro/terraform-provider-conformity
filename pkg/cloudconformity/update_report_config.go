package cloudconformity

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// allows you to update a specific report config
// accountId or groupId could not be changed after report-config was created
func (c *Client) UpdateReportConfig(reportId string, reportConfigPayload ReportConfigDetails) (string, error) {

	reportConfigDetails := ReportConfigDetails{}

	rb, err := json.Marshal(reportConfigPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity UpdateReportConfig request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Patch{}, fmt.Sprintf("/report-configs/%s", reportId), strings.NewReader(string(rb)), "", &reportConfigDetails)
	if err != nil {
		return "", err
	}

	return reportConfigDetails.Data.ID, nil
}
