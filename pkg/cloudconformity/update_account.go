package cloudconformity

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// allows changes to the account name, environment, and code
func (c *Client) UpdateAccount(accountId string, AccountPayload AccountPayload) (string, error) {

	accountResponse := AccountResponse{}

	rb, err := json.Marshal(AccountPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity UpdateAccount request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Patch{}, fmt.Sprintf("/accounts/%s", accountId), strings.NewReader(string(rb)), "", &accountResponse)
	if err != nil {
		return "", err
	}

	return accountResponse.Data.ID, nil
}
