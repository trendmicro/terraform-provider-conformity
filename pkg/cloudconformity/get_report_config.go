package cloudconformity

// allows you to get the details of the specified report config
func (c *Client) GetReportConfig(reportId string) (*ReportConfigDetails, error) {

	reportConfigDetails := ReportConfigDetails{}

	_, err := c.ClientRequest(Get{}, []interface{}{"get_report_config", reportId}, nil, "", &reportConfigDetails)
	if err != nil {
		return nil, err
	}

	return &reportConfigDetails, nil
}
