package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

// used to create a new one-way communication channel setting
func (c *Client) CreateCommunicationSetting(commPayload CommunicationSettings) (*CommunicationResponse, error) {

	commResponse := CommunicationResponse{}

	rb, err := json.Marshal(commPayload)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Conformity CreateCommunicationSetting request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, []interface{}{"create_communication_settings"}, strings.NewReader(string(rb)), "", &commResponse)
	if err != nil {
		return nil, err
	}

	return &commResponse, nil
}
