package cloudconformity

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// allows you to customize rule setting for the specified rule Id of the specified account
func (c *Client) UpdateAccountRuleSettings(accountId string, ruleId string, accountRuleSettings *AccountRuleSettings) (string, error) {

	accountResponse := AccountRuleSettings{}

	rb, err := json.Marshal(accountRuleSettings)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity UpdateAccountRuleSettings request payload: %v\n", string(rb))
	log.Printf("[DEBUG] Conformity UpdateAccountRuleSettings accountId: %s", accountId)
	log.Printf("[DEBUG] Conformity UpdateAccountRuleSettings ruleId: %s", ruleId)

	_, err = c.ClientRequest(Patch{}, fmt.Sprintf("/accounts/%s/settings/rules/%s", accountId, ruleId), strings.NewReader(string(rb)), "", &accountResponse)
	if err != nil {
		return "", err
	}

	return accountResponse.Data.Id, nil
}
