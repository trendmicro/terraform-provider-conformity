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
	    if strings.Contains(string(err), "Unable to call this endpoint, use Cloud One UI or API to invite users"){
	        log_debug(`This Terraform service is not applicable to users who are part of the Cloud One Platform.
	         Please refer to Cloud One User Management Documentation - Add and manage users to invite new users.
	         https://cloudone.trendmicro.com/docs/conformity/api-reference/tag/Users#paths/~1users/get`)

	    }
		return "", err
	}

	return userDetails.Data.ID, nil
}
