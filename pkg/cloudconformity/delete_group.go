package cloudconformity

import (
	"fmt"
)

//   allows you to get the details of the specified group with accounts that you have access to
func (c *Client) DeleteGroup(groupId string) (*deleteResponse, error) {

	deleteGroupResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/v1/groups/%s", groupId), nil, &deleteGroupResponse)
	if err != nil {
		return nil, err
	}

	return &deleteGroupResponse, nil
}
