package cloudconformity

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

// UpdateCustomRule allows you to update a specific custom rule
func (c *Client) UpdateCustomRule(id string, payload CustomRuleRequest) (*CustomRuleResponse, error) {

	response := CustomRuleCreateResponse{}

	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Conformity updateCustomRule request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Put{}, []interface{}{"update_custom_rules", id}, strings.NewReader(string(rb)), "", &response)
	if err != nil {
		return nil, err
	}

	if response.Data.ID == "" {
		return nil, errors.New("conformity server has responded with an internal server error")
	}

	return &response.Data, nil
}
