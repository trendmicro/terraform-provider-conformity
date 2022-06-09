package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

//  used to create a new report config
func (c *Client) CreateReportConfig(reportConfigPayload ReportConfigDetails) (string, error) {

	reportConfigDetails := ReportConfigDetails{}

	rb, err := json.Marshal(reportConfigPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity CreateReportConfig request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "/report-configs/", strings.NewReader(string(rb)), "", &reportConfigDetails)
	if err != nil {
		return "", err
	}

	return reportConfigDetails.Data.ID, nil
}
