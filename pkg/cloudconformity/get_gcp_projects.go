package cloudconformity

import "fmt"

// GetGcpProjects allows you to get the projects by organisationId
func (c *Client) GetGcpProjects(organizationId string) (*GcpProjectsResponse, error) {

	response := &GcpProjectsResponse{}

	_, err := c.ClientRequest(
		Get{},
		fmt.Sprintf("/gcp/organisations/%s/projects", organizationId),
		nil,
		"",
		&response,
	)

	return response, err
}
