package cloudconformity

// allows a user to delete a communication setting
func (c *Client) DeleteCommunicationSetting(commSettingId string) (*deleteResponse, error) {

	deleteCommResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, []interface{}{"delete_communication_setting", commSettingId}, nil, "", &deleteCommResponse)
	if err != nil {
		return nil, err
	}

	return &deleteCommResponse, nil
}
