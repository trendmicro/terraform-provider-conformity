package cloudconformity

// DeleteCustomRule allows a user to delete a custom rule
func (c *Client) DeleteCustomRule(id string) (*deleteResponse, error) {

	response := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, []interface{}{"delete_custom_rules", id}, nil, "", &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
