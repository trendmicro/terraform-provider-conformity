package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

// not applicable to users who are part of the Cloud One Platform
// This is only available for organisations with an external identity provider setup
func (c *Client) CreateSsoLegacyUser(userPayload UserDetails) (string, error) {

	userDetails := UserDetails{}

	rb, err := json.Marshal(userPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity CreateSsoLegacyUser request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "/users/sso/", strings.NewReader(string(rb)), "", &userDetails)
	if err != nil {
		return "", err
	}

	return userDetails.Data.ID, nil
}
