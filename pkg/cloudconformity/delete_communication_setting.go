package cloudconformity

import (
	"fmt"
)

// allows a user to delete a communication setting
func (c *Client) DeleteCommunicationSetting(commSettingId string) (*deleteResponse, error) {

	deleteCommResponse := deleteResponse{}

	_, err := c.ClientRequest(Delete{}, fmt.Sprintf("/settings/%s", commSettingId), nil, "", &deleteCommResponse)
	if err != nil {
		return nil, err
	}

	return &deleteCommResponse, nil
}
