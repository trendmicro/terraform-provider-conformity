package cloudconformity

import (
	"net/url"
)

// allows you to get the details of the specified profile
func (c *Client) GetProfile(profileId string) (*ProfileSettings, error) {

	profileSettings := ProfileSettings{}

	q := url.Values{}
	q.Add("includes", "ruleSettings")

	_, err := c.ClientRequest(Get{}, []interface{}{"get_profile", profileId}, nil, q.Encode(), &profileSettings)
	if err != nil {
		return nil, err
	}

	return &profileSettings, nil
}
