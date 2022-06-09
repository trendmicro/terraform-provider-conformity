package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

// register a new Azure Subscription with an already onboarded Active Directory on Conformity
func (c *Client) CreateAzureAccount(AzurePayload AccountPayload) (string, error) {

	accountResponse := AccountResponse{}

	rb, err := json.Marshal(AzurePayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity CreateAzureAccount request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "/accounts/azure/", strings.NewReader(string(rb)), "", &accountResponse)
	if err != nil {
		return "", err
	}

	return accountResponse.Data.ID, nil
}
