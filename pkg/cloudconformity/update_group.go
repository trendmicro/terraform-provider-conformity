package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

// allows changes to the account name, environment, and code
func (c *Client) UpdateGroup(groupId string, groupPayload GroupDetails) (string, error) {

	groupDetails := GroupDetails{}

	rb, err := json.Marshal(groupPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity UpdateGroup request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Patch{}, []interface{}{"update_group", groupId}, strings.NewReader(string(rb)), "", &groupDetails)
	if err != nil {
		return "", err
	}

	return groupDetails.Data.ID, nil
}
