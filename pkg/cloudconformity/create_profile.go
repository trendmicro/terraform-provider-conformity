package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

// allows you to create a new profile and subsequently add rule settings to the new profile.
// Saving rule settings via this endpoint will overwrite existing settings with those passed in the request.
func (c *Client) CreateProfileSetting(profilePayload ProfileSettings) (string, error) {

	profileSettings := ProfileSettings{}

	rb, err := json.Marshal(profilePayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity CreateProfileSetting request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "/profiles/", strings.NewReader(string(rb)), "", &profileSettings)
	if err != nil {
		return "", err
	}

	return profileSettings.Data.ID, nil
}
