package cloudconformity

import (
	"fmt"
	"net/url"
	"strings"
)

// GetCheck get check details
func (c *Client) GetCheck(checkId string) (*CheckDetails, error) {

	checkDetails := CheckDetails{}

	// resource_id in check must be url-encoded
	split := strings.Split(checkId, ":")

	if len(split) < 6 {
		return nil, fmt.Errorf("invalid check-id: '%s'", checkId)
	}

	formattedCheckId := strings.Replace(checkId, split[5], url.QueryEscape(split[5]), 1)

	_, err := c.ClientRequest(Get{}, fmt.Sprintf("/checks/%s", formattedCheckId), nil, "", &checkDetails)

	return &checkDetails, err

}
