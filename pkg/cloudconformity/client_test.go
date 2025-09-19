package cloudconformity

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const errResponse string = `{ "Message": "User is not authorized to access this resource with an explicit deny" }`
const errResponseUnauthorized = `{"errors": [{"status": 401,"details": "You are not authorized to perform such actions"}]}`
const errResponseUnprocessableEntity = `{"errors": [{"status": 422,"source": {"pointer": "/data/attributes/name"},"details": "Name is required"}]}`

func TestConformityNewClientFail(t *testing.T) {

	client, err := NewClient("TEST-REGION", "TEST-APIKEY", false)
	assert.Contains(t, err.Error(), "no such host")
	assert.Nil(t, client)
}

func TestValidateApiKeySuccess(t *testing.T) {
	response := `{ "data": [ { "type": "api-keys", "id": "BJ0Ox16Hb" } ] }`
	client, ts := createHttpTestClient(t, http.StatusOK, response)
	defer ts.Close()
	//run the key validation
	result, err := client.validateApiKey()
	//check output
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, client.Region, "TEST-REGION")
	assert.Equal(t, client.Apikey, "TEST-APIKEY")
	assert.Equal(t, client.Url, ts.URL)
	// we assume there is only one element in the list
	assert.Equal(t, result.Data[0].ID, "BJ0Ox16Hb")
	assert.Equal(t, result.Data[0].Type, "api-keys")

}

func TestValidateApiKeyFail(t *testing.T) {
	response := errResponse
	client, ts := createHttpTestClient(t, http.StatusForbidden, response)
	defer ts.Close()
	result, err := client.validateApiKey()
	assert.NotNil(t, err)
	assert.EqualError(t, err, errResponse)
	assert.Nil(t, result)
}

func TestCreateNewClientWithV1Feature(t *testing.T) {
	// Create client with useV1Feature = true
	client, err := NewClient("us-1", "TEST-APIKEY", true)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, client.Region, "us-1")
	assert.Equal(t, client.Apikey, "TEST-APIKEY")
	assert.Equal(t, client.UseV1Feature, true)
	assert.Equal(t, client.Url, "https://api.xdr.trendmicro.com/beta/c1/conformity/")
}

func TestGetUrlSuccess(t *testing.T) {
	// Test standalone conformity regions with useV1Feature = false
	url, err := getUrl("us-west-2", false)
	assert.Nil(t, err)
	assert.Equal(t, "https://us-west-2-api.cloudconformity.com/v1/", url)

	url, err = getUrl("ap-southeast-2", false)
	assert.Nil(t, err)
	assert.Equal(t, "https://ap-southeast-2-api.cloudconformity.com/v1/", url)

	url, err = getUrl("eu-west-1", false)
	assert.Nil(t, err)
	assert.Equal(t, "https://eu-west-1-api.cloudconformity.com/v1/", url)

	url, err = getUrl("us-1", false)
	assert.Nil(t, err)
	assert.Equal(t, "https://conformity.us-1.cloudone.trendmicro.com/api/", url)

	url, err = getUrl("jp-1", false)
	assert.Nil(t, err)
	assert.Equal(t, "https://conformity.jp-1.cloudone.trendmicro.com/api/", url)

	// any region will be mapped to a url for standalone conformity
	// but only some are supported, it is described with the schema in the provider.go
	url, err = getUrl("eu-west-4", false)
	assert.Nil(t, err)
	assert.Equal(t, "https://conformity.eu-west-4.cloudone.trendmicro.com/api/", url)

	// Test with useV1Feature = true (should use XDR API format)
	url, err = getUrl("jp-1", true)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.xdr.trendmicro.co.jp/beta/c1/conformity/", url)

	url, err = getUrl("gb-1", true)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.uk.xdr.trendmicro.com/beta/c1/conformity/", url)

	url, err = getUrl("de-1", true)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.eu.xdr.trendmicro.com/beta/c1/conformity/", url)

	url, err = getUrl("au-1", true)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.au.xdr.trendmicro.com/beta/c1/conformity/", url)

	url, err = getUrl("us-1", true)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.xdr.trendmicro.com/beta/c1/conformity/", url)

	url, err = getUrl("sg-1", true)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.sg.xdr.trendmicro.com/beta/c1/conformity/", url)

}

func TestGetUrlFail(t *testing.T) {
	// Test with an unsupported region for v1 API
	_, err := getUrl("ap-south-1", true)
	assert.NotNil(t, err)
	assert.Equal(t, "region ap-south-1 is not supported by v1 API", err.Error())
}

func TestGetUrlWhenAPIEnvIsSet(t *testing.T) {
	t.Setenv("CONFORMITY_API_URL", "https://test-%s.cloudconformity.com/api/")
	url, err := getUrl("us-west-2", false)
	assert.Nil(t, err)
	assert.Equal(t, "https://test-us-west-2.cloudconformity.com/api/", url)

	t.Setenv("CONFORMITY_API_URL", "https://test-ap-southeast-2.cloudconformity.com/api/")
	url, err = getUrl("us-west-2", true)
	assert.Nil(t, err)
	assert.Equal(t, "https://test-ap-southeast-2.cloudconformity.com/api/", url)
}

func createHttpTestClient(_ *testing.T, statusCode int, response string) (*Client, *httptest.Server) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	}))
	client := Client{Region: "TEST-REGION", Apikey: "TEST-APIKEY", UseV1Feature: false, Url: ts.URL, HttpClient: ts.Client()}
	return &client, ts
}
