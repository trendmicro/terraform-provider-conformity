package cloudconformity

import (
	"fmt"
)

// allows ADMINs to delete a specified profile and all affiliated rule settings
func (c *Client) DeleteProfile(groupId string) (*deleteResponse, error) {

	deleteProfileResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/profiles/%s", groupId), nil, "", &deleteProfileResponse)
	if err != nil {
		return nil, err
	}

	return &deleteProfileResponse, nil
}
