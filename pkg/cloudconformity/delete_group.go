package cloudconformity

import (
	"fmt"
)

// allows an ADMIN to delete the specified group
func (c *Client) DeleteGroup(groupId string) (*deleteResponse, error) {

	deleteGroupResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/groups/%s", groupId), nil, "", &deleteGroupResponse)
	if err != nil {
		return nil, err
	}

	return &deleteGroupResponse, nil
}
