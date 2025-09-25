package cloudconformity

// allows you to get your organisation's external ID
func (c *Client) GetExternalId() (string, error) {
	externalId := externalIdData{}
	_, err := c.ClientRequest(Get{}, []interface{}{"get_organisation_external_id"}, nil, "", &externalId)
	if err != nil {
		return "", err
	}
	return externalId.Data.ID, nil
}
