package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

//  register a new AWS account with Conformity
func (c *Client) CreateAwsAccount(AccountPayload AccountPayload) (string, error) {

	accountResponse := AccountResponse{}

	rb, err := json.Marshal(AccountPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity CreateAwsAccount request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "/accounts/", strings.NewReader(string(rb)), "", &accountResponse)
	if err != nil {
		return "", err
	}

	return accountResponse.Data.ID, nil
}
