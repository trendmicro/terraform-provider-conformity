package cloudconformity

import (
	"fmt"
)

// not applicable to users who are part of the Cloud One Platform
// allows you to get the details of the specified group with accounts that you have access to
func (c *Client) RevokeLegacyUser(userId string) (*deleteResponse, error) {

	deleteUserResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/users/%s", userId), nil, "", &deleteUserResponse)
	if err != nil {
		return nil, err
	}

	return &deleteUserResponse, nil
}
