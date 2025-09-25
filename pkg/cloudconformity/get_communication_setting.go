package cloudconformity

// allows you to get the details of the specified communication setting
func (c *Client) GetCommunicationSetting(commSettingId string) (*CommunicationSettings, error) {

	CommunicationSettings := CommunicationSettings{}

	_, err := c.ClientRequest(Get{}, []interface{}{"get_communication_settings", commSettingId}, nil, "", &CommunicationSettings)
	if err != nil {
		return nil, err
	}

	return &CommunicationSettings, nil
}
