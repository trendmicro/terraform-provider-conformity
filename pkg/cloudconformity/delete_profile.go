package cloudconformity

// allows ADMINs to delete a specified profile and all affiliated rule settings
func (c *Client) DeleteProfile(groupId string) (*deleteResponse, error) {

	deleteProfileResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, []interface{}{"delete_profile", groupId}, nil, "", &deleteProfileResponse)
	if err != nil {
		return nil, err
	}

	return &deleteProfileResponse, nil
}
