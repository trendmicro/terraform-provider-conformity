package cloudconformity

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

// CreateConformityCustomRule used to create a custom rule
func (c *Client) CreateConformityCustomRule(payload CustomRuleRequest) (*CustomRuleResponse, error) {

	response := CustomRuleCreateResponse{}

	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Conformity CreateCustomRule request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "custom-rules", strings.NewReader(string(rb)), "", &response)
	if err != nil {
		return nil, err
	}

	if response.Data.ID == "" {
		return nil, errors.New("conformity server has responded with an internal server error")
	}

	return &response.Data, nil
}
