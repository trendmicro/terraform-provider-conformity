package cloudconformity

// allows you to get the details of the specified group with accounts that you have access to
func (c *Client) GetGroup(groupId string) (*GroupDataList, error) {

	GroupDataList := GroupDataList{}

	_, err := c.ClientRequest(Get{}, []interface{}{"get_group", groupId}, nil, "", &GroupDataList)
	if err != nil {
		return nil, err
	}

	return &GroupDataList, nil
}
