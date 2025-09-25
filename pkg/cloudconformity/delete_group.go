package cloudconformity

// allows an ADMIN to delete the specified group
func (c *Client) DeleteGroup(groupId string) (*deleteResponse, error) {

	deleteGroupResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, []interface{}{"delete_group", groupId}, nil, "", &deleteGroupResponse)
	if err != nil {
		return nil, err
	}

	return &deleteGroupResponse, nil
}
