package cloudconformity

import (
	"fmt"
)

// allows you to get the details of the specified communication setting
func (c *Client) GetCommunicationSetting(commSettingId string) (*CommunicationSettings, error) {

	CommunicationSettings := CommunicationSettings{}

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/settings/%s", commSettingId), nil, "", &CommunicationSettings)
	if err != nil {
		return nil, err
	}

	return &CommunicationSettings, nil
}
