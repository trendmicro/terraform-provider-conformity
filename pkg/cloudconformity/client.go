package cloudconformity

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	Region     string
	Apikey     string
	Url        string
	HttpClient *http.Client
}

//create a client with region and apiKey
func NewClient(region string, apikey string) (*Client, error) {
	client := Client{Region: region, Apikey: apikey, Url: getUrl(region), HttpClient: &http.Client{
		Timeout: time.Second * 30,
	}}
	_, err := client.validateApiKey()
	if err != nil {
		return nil, err
	}
	return &client, nil
}

//Validate ApiKey by sending API request using the API key provided
func (c *Client) validateApiKey() (*apiKeyList, error) {

	apiKeyListResult := apiKeyList{}
	_, err := c.ClientRequest(Get{}, "/v1/api-keys/", nil, "", &apiKeyListResult)
	if err != nil {
		return nil, err
	}
	return &apiKeyListResult, nil
}

//generate Valid conformity URI
func getUrl(region string) string {
	return fmt.Sprintf("https://%s-api.cloudconformity.com", region)
}
