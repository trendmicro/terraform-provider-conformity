package cloudconformity

// allows a user to delete a report config
func (c *Client) DeleteReportConfig(reportId string) (*deleteResponse, error) {

	deleteReportResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, []interface{}{"delete_report_config", reportId}, nil, "", &deleteReportResponse)
	if err != nil {
		return nil, err
	}

	return &deleteReportResponse, nil
}
