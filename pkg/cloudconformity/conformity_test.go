package cloudconformity

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockServerResponse struct {
	Data map[string]string `json:"data"`
}

func TestClientRequestSuccess(t *testing.T) {
	client, ts := createMockHttpTestClient(t, http.StatusOK, false)
	defer ts.Close()

	assert.NotNil(t, client)

	//do post request
	response := MockServerResponse{}
	body, err := client.ClientRequest(Post{}, []interface{}{"create_aws_account"}, strings.NewReader("{}"), "", &response)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.Equal(t, response.Data["url"], "/accounts/")
	assert.Equal(t, response.Data["apiToken"], "ApiKey TEST-APIKEY")

	// do get request
	response = MockServerResponse{}
	body, err = client.ClientRequest(Get{}, []interface{}{"get_profile", "pf-12345"}, nil, "", &response)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.Equal(t, response.Data["url"], "/profiles/pf-12345")
	assert.Equal(t, response.Data["apiToken"], "ApiKey TEST-APIKEY")
}

func TestV1ClientRequestSuccess(t *testing.T) {
	client, ts := createMockHttpTestClient(t, http.StatusOK, true)
	defer ts.Close()

	assert.NotNil(t, client)

	// do patch request
	response := MockServerResponse{}
	body, err := client.ClientRequest(Patch{}, []interface{}{"update_account_rule_settings", "account-12345", "rule-settings-123"}, nil, "", &response)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.Equal(t, response.Data["url"], "/accounts/account-12345/settings/rules/rule-settings-123")
	assert.Equal(t, response.Data["apiToken"], "Bearer TEST-APIKEY")
}

func TestV1ClientRequestFail(t *testing.T) {
	client, ts := createMockHttpTestClient(t, http.StatusOK, true)
	defer ts.Close()

	assert.NotNil(t, client)

	// do patch request
	response := MockServerResponse{}
	body, err := client.ClientRequest(Patch{}, []interface{}{"create_gcp_account"}, nil, "", &response)
	assert.EqualError(t, err, "some resources are not supported when using VisionOne. Please check the documentation https://docs.conformity.com/api/")
	assert.Nil(t, body)
}

func createMockHttpTestClient(_ *testing.T, statusCode int, useV1Feature bool) (*Client, *httptest.Server) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(fmt.Sprintf(`{"data":{"url":"%s", "apiToken": "%s"}}`, r.URL.String(), r.Header["Authorization"][0])))
	}))
	client := Client{Region: "TEST-REGION", Apikey: "TEST-APIKEY", UseV1Feature: useV1Feature, Url: ts.URL, HttpClient: ts.Client()}
	return &client, ts
}
