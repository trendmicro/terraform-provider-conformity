package cloudconformity

// GetAzureSubscriptions allows you to get the subscriptions by directoryId
func (c *Client) GetAzureSubscriptions(directoryId string) (*AzureSubscriptionsResponse, error) {

	azureSubscriptionsResponse := AzureSubscriptionsResponse{}

	_, err := c.ClientRequest(
		Get{},
		[]interface{}{"get_azure_subscriptions", directoryId},
		nil,
		"",
		&azureSubscriptionsResponse,
	)
	if err != nil {
		return nil, err
	}

	return &azureSubscriptionsResponse, nil
}
