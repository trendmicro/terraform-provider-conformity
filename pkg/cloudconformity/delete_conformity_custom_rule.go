package cloudconformity

import "fmt"

// DeleteCustomRule allows a user to delete a custom rule
func (c *Client) DeleteCustomRule(id string) (*deleteResponse, error) {

	response := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/custom-rules/%s", id), nil, "", &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
