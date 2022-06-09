package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

//  allows an ADMIN user to create a new group
func (c *Client) CreateGroup(groupPayload GroupDetails) (string, error) {

	groupDetails := GroupDetails{}

	rb, err := json.Marshal(groupPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity CreateGroup request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "/groups/", strings.NewReader(string(rb)), "", &groupDetails)
	if err != nil {
		return "", err
	}

	return groupDetails.Data.ID, nil
}
