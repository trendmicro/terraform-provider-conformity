package cloudconformity

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	V1HostMap = map[string]string{
		"au":    "api.au.xdr.trendmicro.com",    // Australia
		"eu":    "api.eu.xdr.trendmicro.com",    // Europe
		"in":    "api.in.xdr.trendmicro.com",    // India
		"jp":    "api.xdr.trendmicro.co.jp",     // Japan
		"sg":    "api.sg.xdr.trendmicro.com",    // Singapore
		"mea":   "api.mea.xdr.trendmicro.com",   // United Arab Emirates
		"us":    "api.xdr.trendmicro.com",       // US
		"usgov": "api.usgov.xdr.trendmicro.com", // US Gov
	}
)

type Client struct {
	Region       string
	Apikey       string
	Url          string
	HttpClient   *http.Client
	UseV1Feature bool
}

// create a client with region and apiKey
func NewClient(region string, apikey string, useV1Feature bool) (*Client, error) {
	Url, err := getUrl(region, useV1Feature)
	if err != nil {
		return nil, err
	}

	client := Client{Region: region, Apikey: apikey, UseV1Feature: useV1Feature, Url: Url, HttpClient: &http.Client{
		Timeout: time.Second * 30,
	}}

	// only validate API key if not using v1 feature
	if !useV1Feature {
		_, err = client.validateApiKey()
		if err != nil {
			return nil, err
		}
	}

	return &client, nil
}

// Validate ApiKey by sending API request using the API key provided
func (c *Client) validateApiKey() (*apiKeyList, error) {

	apiKeyListResult := apiKeyList{}
	_, err := c.ClientRequest(Get{}, "/api-keys/", nil, "", &apiKeyListResult)
	if err != nil {
		return nil, err
	}
	return &apiKeyListResult, nil
}
func stringInSlice(str string, list []string) bool {
	for _, b := range list {
		if b == str {
			return true
		}
	}
	return false
}

// generate Valid conformity URI
func getUrl(region string, useV1Feature bool) (string, error) {
	// if useV1Feature is true, use the v1 API URL format
	if useV1Feature {
		regionMap := map[string]string{
			"eu-west-1":      "eu",
			"us-west-2":      "us",
			"ap-southeast-2": "au",
			"us-1":           "us",
			"in-1":           "in",
			"gb-1":           "eu",
			"jp-1":           "jp",
			"de-1":           "eu",
			"au-1":           "au",
			"ca-1":           "us",
			"sg-1":           "sg",
		}

		v1Region, found := regionMap[region]

		if !found {
			return "", fmt.Errorf("region %s is not supported by v1 API", region)
		}

		host := V1HostMap[v1Region]

		return fmt.Sprintf("https://%s/beta/c1/conformity/", host), nil
	}

	// cloud one conformity API URL format
	urlFormat := "https://conformity.%s.cloudone.trendmicro.com/api/"
	if stringInSlice(region, []string{"eu-west-1", "us-west-2", "ap-southeast-2"}) {
		// standalone conformity API URL format
		urlFormat = "https://%s-api.cloudconformity.com/v1/"
	}

	// check if CONFORMITY_API_URL is set in environment variables
	// if set, use it instead of the default format
	apiURL, ok := os.LookupEnv("CONFORMITY_API_URL")
	if ok {
		urlFormat = apiURL
	}

	return fmt.Sprintf(urlFormat, region), nil
}
