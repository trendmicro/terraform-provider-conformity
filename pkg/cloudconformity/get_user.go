package cloudconformity

import (
	"fmt"
)

// not applicable to users who are part of the Cloud One Platform
// allows you to get the details of the specified user
func (c *Client) GetLegacyUser(userId string) (*UserDetails, error) {

	userDetails := UserDetails{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/users/%s", userId), nil, "", &userDetails)
	if err != nil {
		return nil, err
	}

	return &userDetails, nil
}
