package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

//  register a new GCP org with Conformity
func (c *Client) CreateGCPOrg(OrgPayload OrgPayload) (string, error) {

	orgResponse := GCPOrgResponse{}

	rb, err := json.Marshal(OrgPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity CreateGCPOrg request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "/v1/gcp/organisations", strings.NewReader(string(rb)), "", &orgResponse)
	if err != nil {
		return "", err
	}

	return orgResponse.Data.ID, nil
}