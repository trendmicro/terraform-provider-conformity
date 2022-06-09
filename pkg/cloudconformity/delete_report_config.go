package cloudconformity

import (
	"fmt"
)

//	allows a user to delete a report config
func (c *Client) DeleteReportConfig(reportId string) (*deleteResponse, error) {

	deleteReportResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/report-configs/%s", reportId), nil, "", &deleteReportResponse)
	if err != nil {
		return nil, err
	}

	return &deleteReportResponse, nil
}
