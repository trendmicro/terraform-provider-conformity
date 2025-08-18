package cloudconformity

// GetGcpProjects allows you to get the projects by organisationId
func (c *Client) GetGcpProjects(organizationId string) (*GcpProjectsResponse, error) {

	response := &GcpProjectsResponse{}

	_, err := c.ClientRequest(
		Get{},
		[]interface{}{"get_gcp_projects", organizationId},
		nil,
		"",
		&response,
	)

	return response, err
}
