package cloudconformity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
type Put struct{}
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

func (Put) genericRequest(Client *Client, url_path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do patch request
	return newRequest(Client, "PUT", url_path, payload, rawQuery, result)
}

func (Delete) genericRequest(Client *Client, url_path string, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	//do delete request
	return newRequest(Client, "DELETE", url_path, payload, rawQuery, result)
}

func (c *Client) headers(request *http.Request) {
	AuthorizationType := "ApiKey"
	if c.UseV1Feature {
		AuthorizationType = "Bearer"
	}

	request.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("%s %s", AuthorizationType, c.Apikey)},
		"Content-Type":  {"application/vnd.api+json"},
	}

	request.Header["Request-Source"] = []string{"Terraform"}
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

	body, err := io.ReadAll(resp.Body)
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
		log_debug(fmt.Sprintf("Conformity request error: %d", resp.StatusCode))
		log_debug("Conformity response body error" + string(body))

		return body, errors.New(string(body))
	}

	return body, nil
}

func (client *Client) ClientRequest(m method, url_params []interface{}, payload io.Reader, rawQuery string, result interface{}) ([]byte, error) {
	functionName := url_params[0].(string)
	params := url_params[1:]

	var url_path string
	var found bool
	if client.UseV1Feature {
		url_path, found = V1Functions[functionName]
		if !found {
			return nil, fmt.Errorf("some resources are not supported when using VisionOne. Please contact the support team at %s", apiDocumentationURL)
		}
	} else {
		url_path = LegacyFunctions[functionName]
	}
	url_path = fmt.Sprintf(url_path, params...)
	return m.genericRequest(client, url_path, payload, rawQuery, result)
}

var (
	apiDocumentationURL = "https://success.trendmicro.com/en-US/"
	LegacyFunctions     = map[string]string{
		"get_api_keys":                    "/api-keys/",
		"create_aws_account":              "/accounts/",
		"create_azure_account":            "/accounts/azure/",
		"create_azure_active_directories": "/azure/active-directories",
		"create_gcp_account":              "/accounts/gcp/",
		"create_communication_settings":   "/settings/communication/",
		"create_custom_rules":             "/custom-rules/",
		"create_gcp_organisations":        "/gcp/organisations/",
		"create_group":                    "/groups/",
		"create_profile":                  "/profiles/",
		"create_report_config":            "/report-configs/",
		"create_sso_user":                 "/users/sso/",
		"delete_account":                  "/accounts/%s",
		"delete_communication_setting":    "/settings/%s",
		"delete_custom_rules":             "/custom-rules/%s",
		"delete_group":                    "/groups/%s",
		"delete_profile":                  "/profiles/%s",
		"delete_report_config":            "/report-configs/%s",
		"get_account":                     "/accounts/%s",
		"get_account_access":              "/accounts/%s/access",
		"get_account_settings_rules":      "/accounts/%s/settings/rules",
		"get_azure_subscriptions":         "/azure/active-directories/%s/subscriptions",
		"get_checks":                      "/checks/%s",
		"get_communication_settings":      "/settings/%s",
		"get_custom_rules":                "/custom-rules/%s",
		"get_current_user":                "/users/whoami",
		"get_group":                       "/groups/%s",
		"get_profile":                     "/profiles/%s",
		"get_report_config":               "/report-configs/%s",
		"get_user":                        "/users/%s",
		"invite_user":                     "/users/",
		"revoke_user":                     "/users/%s",
		"update_account_bot":              "/accounts/%s/settings/bot",
		"update_account_rule_settings":    "/accounts/%s/settings/rules/%s",
		"update_account":                  "/accounts/%s",
		"update_check":                    "/checks/%s",
		"update_communication_setting":    "/settings/communication/%s",
		"update_custom_rules":             "/custom-rules/%s",
		"update_group":                    "/groups/%s",
		"update_profile":                  "/profiles/%s",
		"update_report_config":            "/report-configs/%s",
		"update_user":                     "/users/%s",
		"apply_profile":                   "/profiles/%s/apply",
		"get_organisation_external_id":    "/organisation/external-id/",
		"get_gcp_projects":                "/gcp/organisations/%s/projects",
	}

	V1Functions = map[string]string{
		"create_communication_settings": "/settings/communication",
		"create_custom_rules":           "/custom-rules",
		"create_group":                  "/groups",
		"create_profile":                "/profiles",
		"create_report_config":          "/report-configs",
		"delete_communication_setting":  "/settings/%s",
		"delete_custom_rules":           "/custom-rules/%s",
		"delete_group":                  "/groups/%s",
		"delete_profile":                "/profiles/%s",
		"delete_report_config":          "/report-configs/%s",
		"get_checks":                    "/checks/%s",
		"get_communication_settings":    "/settings/%s",
		"get_custom_rules":              "/custom-rules/%s",
		"get_group":                     "/groups/%s",
		"get_profile":                   "/profiles/%s",
		"get_report_config":             "/report-configs/%s",
		"update_account_bot":            "/accounts/%s/settings/bot",
		"update_account_rule_settings":  "/accounts/%s/settings/rules/%s",
		"update_check":                  "/checks/%s",
		"update_communication_setting":  "/settings/communication/%s",
		"update_custom_rules":           "/custom-rules/%s",
		"update_group":                  "/groups/%s",
		"update_profile":                "/profiles/%s",
		"update_report_config":          "/report-configs/%s",
		"apply_profile":                 "/profiles/%s/apply",
	}
)
