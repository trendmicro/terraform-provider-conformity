package cloudconformity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type method interface {
	genericRequest(Client *Client, url_path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error)
}
type Get struct{}

type Post struct{}

type Patch struct{}
type Delete struct{}

func (Post) genericRequest(Client *Client, url_path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do post request
	return newRequest(Client, "POST", url_path, payload, rawQuery, result)
}

func (Get) genericRequest(Client *Client, url_path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do get request
	return newRequest(Client, "GET", url_path, payload, rawQuery, result)
}

func (Patch) genericRequest(Client *Client, url_path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do patch request
	return newRequest(Client, "PATCH", url_path, payload, rawQuery, result)
}

func (Delete) genericRequest(Client *Client, url_path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do delete request
	return newRequest(Client, "DELETE", url_path, payload, rawQuery, result)
}

func (c *Client) headers(request *http.Request) {

	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("ApiKey %s", c.Apikey)},
		"Content-Type":  {"application/vnd.api+json"},
	}
}

func newRequest(c *Client, methodType string, url_path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {

	apiUrl := c.Url
	resource := url_path
	u, _ := url.ParseRequestURI(apiUrl)
	urlString := u.String() + resource

	client := c.HttpClient
	log_debug("Request URL: " + urlString)
	log_debug("Method: " + methodType)
	payload_string := convert_io_to_string(payload)
	log_encrypted(payload_string)

	result_name := reflect.Indirect(reflect.ValueOf(result)).Type().Name()
	req, err := http.NewRequest(methodType, urlString, strings.NewReader(payload_string))
	if err != nil {
		return nil, err
	}
	c.headers(req)

	if rawQuery != "" {
		req.URL.RawQuery = rawQuery
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log_debug("Response Body of " + result_name)
	log_encrypted(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		log_debug("Conformity request error: " + string(resp.StatusCode))
		log_debug("Conformity response body error" + string(body))

		return body, errors.New(string(body))
	}

	return body, nil
}

func (client *Client) ClientRequest(m method, url_path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	return m.genericRequest(client, url_path, payload, rawQuery, result)
}
