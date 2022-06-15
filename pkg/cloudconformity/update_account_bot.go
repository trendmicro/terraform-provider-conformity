package cloudconformity

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// allows changes to the account name, environment, and code
func (c *Client) UpdateAccountBotSettings(accountId string, accountBotSettings AccountBotSettingsRequest) (string, error) {

	accountResponse := AccountBotSettingsReponse{}

	rb, err := json.Marshal(accountBotSettings)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("/accounts/%s/settings/bot", accountId)
	log.Printf("[DEBUG] Conformity UpdateAccountBotSettings request url: %s\n", url)
	log.Printf("[DEBUG] Conformity UpdateAccountBotSettings request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Patch{}, url, strings.NewReader(string(rb)), "", &accountResponse)
	if err != nil {
		return "", err
	}

	return accountResponse.Data[0].Id, nil
}
