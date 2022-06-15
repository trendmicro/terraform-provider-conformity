package cloudconformity

import (
	"fmt"
)

// allows you to get the details of the specified report config
func (c *Client) GetReportConfig(reportId string) (*ReportConfigDetails, error) {

	reportConfigDetails := ReportConfigDetails{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/report-configs/%s", reportId), nil, "", &reportConfigDetails)
	if err != nil {
		return nil, err
	}

	return &reportConfigDetails, nil
}
