package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

// not applicable to users who are part of the Cloud One Platform
// allows you to invite a user to your organisation
func (c *Client) InviteLegacyUser(userPayload UserDetails) (string, error) {

	userDetails := UserDetails{}

	rb, err := json.Marshal(userPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity InviteUser request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "/users/", strings.NewReader(string(rb)), "", &userDetails)
	if err != nil {
		return "", err
	}

	return userDetails.Data.ID, nil
}
