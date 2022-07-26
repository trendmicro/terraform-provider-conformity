package cloudconformity

import "fmt"

// GetAzureSubscriptions allows you to get the subscriptions by directoryId
func (c *Client) GetAzureSubscriptions(directoryId string) (*AzureSubscriptionsResponse, error) {

	azureSubscriptionsResponse := AzureSubscriptionsResponse{}

	_, err := c.ClientRequest(
		Get{},
		fmt.Sprintf("/azure/active-directories/%s/subscriptions", directoryId),
		nil,
		"",
		&azureSubscriptionsResponse,
	)
	if err != nil {
		return nil, err
	}

	return &azureSubscriptionsResponse, nil
}
