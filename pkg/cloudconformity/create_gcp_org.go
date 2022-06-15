package cloudconformity

import (
	"encoding/json"
	"log"
	"strings"
)

//  register a new GCP org with Conformity
func (c *Client) CreateGCPOrg(GCPOrgPayload GCPOrgPayload) (string, error) {

	orgResponse := GCPOrgResponse{}

	rb, err := json.Marshal(GCPOrgPayload)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] Conformity CreateGCPOrg request payload: %v\n", string(rb))

	_, err = c.ClientRequest(Post{}, "/gcp/organisations/", strings.NewReader(string(rb)), "", &orgResponse)
	if err != nil {
		return "", err
	}

	return orgResponse.Data.ID, nil
}
