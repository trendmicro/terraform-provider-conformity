package cloudconformity

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// allows you to update a specific communication setting
func (c *Client) UpdateCommunicationSetting(commSettingId string, commPayload CommunicationSettings) (string, error) {

	commResponse := CommunicationSettings{}

	rb, err := json.Marshal(commPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity UpdateCommunicationSetting request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Patch{}, fmt.Sprintf("/settings/communication/%s", commSettingId), strings.NewReader(string(rb)), "", &commResponse)
	if err != nil {
		return "", err
	}

	return commResponse.Data.ID, nil
}
