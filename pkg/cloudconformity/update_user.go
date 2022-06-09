package cloudconformity

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// not applicable to users who are part of the Cloud One Platform
// Updates the role and permissions of the specified user
func (c *Client) UpdateLegacyUser(userId string, userPayload UserAccessDetails) (string, error) {

	userDetails := UserDetails{}

	rb, err := json.Marshal(userPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity UpdateUser request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Patch{}, fmt.Sprintf("/users/%s", userId), strings.NewReader(string(rb)), "", &userDetails)
	if err != nil {
		return "", err
	}

	return userDetails.Data.ID, nil
}
