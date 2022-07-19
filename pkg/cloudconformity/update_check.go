package cloudconformity

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

// UpdateCheck update a custom check
func (c *Client) UpdateCheck(checkId string, checkPayload CheckDetails) (*CheckDetails, error) {

	checkDetails := CheckDetails{}

	rb, err := json.Marshal(checkPayload)
	if err != nil {
		return nil, err
	}

	// resource_id in check must be url-encoded
	split := strings.Split(checkId, ":")

	if len(split) < 6 {
		return nil, fmt.Errorf("invalid check-id: '%s'", checkId)
	}

	formattedCheckId := strings.Replace(checkId, split[5], url.QueryEscape(split[5]), 1)

	log.Printf("[DEBUG] Conformity UpdateCheck request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Patch{}, fmt.Sprintf("/checks/%s", formattedCheckId), strings.NewReader(string(rb)), "", &checkDetails)

	return &checkDetails, err
}
