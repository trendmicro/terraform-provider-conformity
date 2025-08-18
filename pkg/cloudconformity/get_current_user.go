package cloudconformity

import (
	"encoding/json"
)

func (c *Client) GetCurrentUser() (*UserDetails, error) {

	userDetails := UserDetails{}
	value, err := c.ClientRequest(Get{}, []interface{}{"get_current_user"}, nil, "", &userDetails)
	if err != nil {
		return nil, err
	}

	err1 := json.Unmarshal(value, &userDetails)
	if err1 != nil {
		return nil, err1
	}
	return &userDetails, nil
}
