package cloudconformity

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

//  endpoint allows you to apply profile and rule settings to a set of accounts under your organisation
func (c *Client) CreateApplyProfile(profileId string, profilePayload ApplyProfileSettings) (*ApplyProfileResponse, error) {

	applyProfileResponse := &ApplyProfileResponse{}

	rb, err := json.Marshal(profilePayload)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Conformity CreateApplyProfile request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, fmt.Sprintf("/profiles/%s/apply", profileId), strings.NewReader(string(rb)), "", applyProfileResponse)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Apply profile Response message: %s", applyProfileResponse.Meta.Message)
	log.Printf("[DEBUG] Apply profile Response status: %s", applyProfileResponse.Meta.Status)
	return applyProfileResponse, nil
}
