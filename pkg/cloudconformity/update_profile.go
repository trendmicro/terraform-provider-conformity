package cloudconformity

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// allows you to update profile details and its associated rule settings.
func (c *Client) UpdateProfile(profileId string, profilePayload ProfileSettings) (string, error) {

	profileSettings := ProfileSettings{}

	rb, err := json.Marshal(profilePayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity UpdateProfile request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Patch{}, fmt.Sprintf("/profiles/%s", profileId), strings.NewReader(string(rb)), "", &profileSettings)
	if err != nil {
		return "", err
	}

	return profileSettings.Data.ID, nil
}
