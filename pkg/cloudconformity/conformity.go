package cloudconformity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

type method interface {
	genericRequest(Client *Client, path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error)
}
type Get struct{}

type Post struct{}

type Patch struct{}
type Delete struct{}

func (Post) genericRequest(Client *Client, path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do post request
	return newRequest(Client, "POST", path, payload, rawQuery, result)
}

func (Get) genericRequest(Client *Client, path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do get request
	return newRequest(Client, "GET", path, payload, rawQuery, result)
}

func (Patch) genericRequest(Client *Client, path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do patch request
	return newRequest(Client, "PATCH", path, payload, rawQuery, result)
}

func (Delete) genericRequest(Client *Client, path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do delete request
	return newRequest(Client, "DELETE", path, payload, rawQuery, result)
}

func (c *Client) headers(request *http.Request) {

	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("ApiKey %s", c.Apikey)},
		"Content-Type":  {"application/vnd.api+json"},
	}
}

func getType(myvar interface{}) {
    valueOf := reflect.ValueOf(myvar)
    if valueOf.Type().Kind() == reflect.Ptr {
        return reflect.Indirect(valueOf).Type().Name()
    } else {
        return valueOf.Type().Name()
    }
}

func newRequest(c *Client, methodType string, path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {

	apiUrl := c.Url
	resource := path

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlString := u.String()
	client := c.HttpClient
    log.Printf("[DEBUG] Request URL: %v\n", urlString)
    log.Printf("[DEBUG] Request Payload: %v\n", string(payload))
    log.Printf("[DEBUG] Defined Struct Name: %v\n", getType(result))
	req, err := http.NewRequest(methodType, urlString, payload)
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
    log.Printf("[DEBUG] Response Body: %v\n", string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Printf("[DEBUG] Conformity request error: %v\n", resp.StatusCode)
		log.Printf("[DEBUG] Conformity reponse body error: %v\n", string(body))

		return body, errors.New(string(body))
	}

	return body, nil
}

func (client *Client) ClientRequest(m method, path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	return m.genericRequest(client, path, payload, rawQuery, result)
}
