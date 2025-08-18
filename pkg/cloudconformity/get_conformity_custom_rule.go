package cloudconformity

import (
	"errors"
)

// GetCustomRule allows you to get the custom rule
func (c *Client) GetCustomRule(id string) (*CustomRuleResponse, error) {

	response := CustomRuleGetResponse{}

	_, err := c.ClientRequest(Get{}, []interface{}{"get_custom_rules", id}, nil, "", &response)
	if err != nil {
		return nil, err
	}

	if len(response.Data) == 0 || response.Data[0].ID == "" {
		return nil, errors.New("conformity server has responded with an internal server error")
	}

	return &response.Data[0], nil
}
