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

	client, err := NewClient("TEST-REGION", "TEST-APIKEY")
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

func createHttpTestClient(t *testing.T, statusCode int, response string) (*Client, *httptest.Server) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	}))
	client := Client{Region: "TEST-REGION", Apikey: "TEST-APIKEY", Url: ts.URL, HttpClient: ts.Client()}
	return &client, ts
}
