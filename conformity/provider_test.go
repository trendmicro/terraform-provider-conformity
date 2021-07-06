package conformity

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccConformityProviders map[string]*schema.Provider
var testAccConformityProvider *schema.Provider
var accountPayload cloudconformity.AccountPayload
var groupDetails cloudconformity.GroupDetails
var userDetails cloudconformity.UserDetails
var userAccessDetails cloudconformity.UserAccessDetails
var reportConfigDetails cloudconformity.ReportConfigDetails
var communicationSetting cloudconformity.CommunicationSettings
var testServer *httptest.Server

func init() {
	testAccConformityProvider = Provider()
	testAccConformityProvider.ConfigureContextFunc = testProviderConfigure

	testAccConformityProviders = map[string]*schema.Provider{
		"conformity": testAccConformityProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccConformityPreCheck(t *testing.T) {
	os.Setenv("APIKEY", "test-apikey")
}

func testProviderConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	client, server := createConformityMock()
	testServer = server
	return client, diags

}
func readRequestBody(r *http.Request, payload interface{}) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}
	return nil
}

func createConformityMock() (*cloudconformity.Client, *httptest.Server) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var getOrganizationalExternalId = regexp.MustCompile(`^/v1/organisation/external-id/$`)
		var postAccount = regexp.MustCompile(`^/v1/accounts/$`)
		var postAzureAccount = regexp.MustCompile(`^/v1/accounts/azure/$`)
		var accountDetails = regexp.MustCompile(`^/v1/accounts/(.*)$`)
		var getAccountAccess = regexp.MustCompile(`^/v1/accounts/(.*)/access$`)
		var postGroup = regexp.MustCompile(`^/v1/groups/$`)
		var getGroupDetails = regexp.MustCompile(`^/v1/groups/(.*)$`)
		var postUsers = regexp.MustCompile(`^/v1/users/$`)
		var postSsoUsers = regexp.MustCompile(`^/v1/users/sso/$`)
		var getUserDetails = regexp.MustCompile(`^/v1/users/(.*)$`)
		var postReportconfig = regexp.MustCompile(`^/v1/report-configs/$`)
		var getReportconfig = regexp.MustCompile(`^/v1/report-configs/(.*)$`)
		var postCommunicationConfig = regexp.MustCompile(`^/v1/settings/communication/$`)
		var getCommunicationConfig = regexp.MustCompile(`^/v1/settings/(.*)$`)
		switch {
		case getOrganizationalExternalId.MatchString(r.URL.Path):
			w.Write([]byte(`{ "data": { "type": "external-ids", "id": "3ff84b20-0f4c-11eb-a7b7-7d9b3c0e866e" } }`))
		case postAccount.MatchString(r.URL.Path):
			_ = readRequestBody(r, &accountPayload)
			w.Write([]byte(`{ "data": { "type": "accounts", "id": "H19NxMi5-" } }`))
		case postAzureAccount.MatchString(r.URL.Path):
			_ = readRequestBody(r, &accountPayload)
			w.Write([]byte(`{ "data": { "type": "accounts", "id": "H19NxMi5-" } }`))
		case getAccountAccess.MatchString(r.URL.Path):
			w.Write([]byte(`{
				"id": "BJ0Ox16Hb:access",
				"type": "settings",
				"attributes": {
				"type": "access",
				"configuration": {
				"externalId": "test-external-id",
				"roleArn": "test-arn" } } }`))
		case accountDetails.MatchString(r.URL.Path) && r.Method == "GET":
			tags, _ := json.Marshal(accountPayload.Data.Attributes.Tags)
			w.Write([]byte(`{
				"data": {	
				"type": "accounts",
				"id": "H19NxMi5-",
				"attributes": {
				"name": "` + accountPayload.Data.Attributes.Name + `",
				"environment": "` + accountPayload.Data.Attributes.Environment + `", 
				"tags": ` + string(tags) + `,
				"cloud-type": "azure",
				"cloud-data": {
				"azure": {
				"subscriptionId": "test-subscrition-id"
				} } } } }`))

		case accountDetails.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`{
				"meta": {
				"status": "sent" } }`))
		case accountDetails.MatchString(r.URL.Path) && r.Method == "PATCH":
			_ = readRequestBody(r, &accountPayload)

			w.Write([]byte(`{
				"data": {
				"type": "accounts",
				"id": "H19NxMi5-" } }`))
		case postGroup.MatchString(r.URL.Path):
			response := groupResponse(r, &groupDetails)
			w.Write([]byte(response))

		case getGroupDetails.MatchString(r.URL.Path) && r.Method == "GET":

			tags, _ := json.Marshal(groupDetails.Data.Attributes.Tags)
			stringTags := string(tags)

			if groupDetails.Data.Attributes.Name == "no-tag" {
				stringTags = "null"
			}

			w.Write([]byte(`{
				"data": [{
				"type": "groups",
				"id": "uUmE2v0ns",
				"attributes": {
				"name": "` + groupDetails.Data.Attributes.Name + `",
				"tags": ` + stringTags + ` } }]}`))
		case getGroupDetails.MatchString(r.URL.Path) && r.Method == "PATCH":
			response := groupResponse(r, &groupDetails)
			w.Write([]byte(response))
		case getGroupDetails.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`{
				"meta": {
				"status": "deleted" } }`))
		case postUsers.MatchString(r.URL.Path) || postSsoUsers.MatchString(r.URL.Path):
			_ = readRequestBody(r, &userDetails)

			w.Write([]byte(`{
				"data": {
				"type": "users",
				"id": "OhnzPVXY" } }`))

		case getUserDetails.MatchString(r.URL.Path) && r.Method == "PATCH":
			_ = readRequestBody(r, &userAccessDetails)
			w.Write([]byte(`{
				"data": {
				"type": "users",
				"id": "OhnzPVXY" } }`))
		case getUserDetails.MatchString(r.URL.Path) && r.Method == "GET":
			response := getUserResponse()
			w.Write([]byte(response))
		case getUserDetails.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`{
				"meta": {
				"status": "revoked" } }`))

		case postReportconfig.MatchString(r.URL.Path) || (getReportconfig.MatchString(r.URL.Path) && r.Method == "PATCH"):
			_ = readRequestBody(r, &reportConfigDetails)
			w.Write([]byte(`{ "data": {
				"type": "report-config",
				"id": "vO4SPFxrcC" } }`))

		case getReportconfig.MatchString(r.URL.Path) && r.Method == "GET":
			response := getReportconfigResponse()
			w.Write([]byte(response))

		case getReportconfig.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`{
				"meta": {
				"status": "deleted" } }`))

		case postCommunicationConfig.MatchString(r.URL.Path):
			_ = readRequestBody(r, &communicationSetting)
			w.Write([]byte(`{
				"data": [
				  {
					"id": "ryqs8LNKW:communication:email-Ske1cKKEvM"
				  }
				]
			  }`))

		case getCommunicationConfig.MatchString(r.URL.Path) && r.Method == "PATCH":
			_ = readRequestBody(r, &communicationSetting)
			w.Write([]byte(`{
					"data":
					  {
						"id": "ryqs8LNKW:communication:email-Ske1cKKEvM"
					  }
				  }`))

		case getCommunicationConfig.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(`{
				"data":
					{
						"type": "settings",
						"id": "ryqs8LNKW:communication:email-Ske1cKKEvM",
						"attributes": {
							"type": "communication",
							"enabled": true,
							"filter": {
								"categories": [
									"security"
								]
							},
							"configuration": {
								"users": [
									"urn:tmds:identity:us-east-ds-1:62740:administrator/1915"
								],
								"url": "` + communicationSetting.Data.Attributes.Configuration.Url + `",
								"channel": "` + communicationSetting.Data.Attributes.Configuration.Channel + `",
								"channelName": "` + communicationSetting.Data.Attributes.Configuration.ChannelName + `",
								"arn": "` + communicationSetting.Data.Attributes.Configuration.Arn + `"
							},
							"channel": "` + communicationSetting.Data.Attributes.Channel + `"
						},
						"relationships": {
							"account": {
								"data": {
									"type": "accounts",
									"id": "` + communicationSetting.Data.Relationships.Account.Data.ID + `"
								}
							},
							"organisation": {
								"data": {
									"type": "organisations",
									"id": "ryqMcJn4b"
								}
							}
						}
					}
			}`))

		case getCommunicationConfig.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`{
					"meta": {
					"status": "deleted" } }`))
		}

	}))
	// we do not Close() the server, it will be kept alive until all tests are finished
	client := cloudconformity.Client{Region: "TEST-REGION", Apikey: "TEST-APIKEY", Url: server.URL, HttpClient: server.Client()}
	return &client, server
}

func groupResponse(r *http.Request, groupDetails *cloudconformity.GroupDetails) string {
	_ = readRequestBody(r, groupDetails)
	tags, _ := json.Marshal(groupDetails.Data.Attributes.Tags)
	return `{
		"data": {
		"type": "groups",
		"id": "uUmE2v0ns",
		"attributes": {
		"name": "` + groupDetails.Data.Attributes.Name + `",
		"tags": ` + string(tags) + ` } } }`
}
func getUserResponse() string {
	var (
		role     string
		account1 string
		level1   string
		account2 string
		level2   string
	)
	if userAccessDetails.Data.Role == "" {

		role = userDetails.Data.Attributes.Role
		account1 = userDetails.Data.Attributes.AccessList[0].Account
		level1 = userDetails.Data.Attributes.AccessList[0].Level
		account2 = userDetails.Data.Attributes.AccessList[1].Account
		level2 = userDetails.Data.Attributes.AccessList[1].Level

	} else {

		role = userAccessDetails.Data.Role
		account1 = userAccessDetails.Data.AccessList[0].Account
		level1 = userAccessDetails.Data.AccessList[0].Level
		account2 = userAccessDetails.Data.AccessList[1].Account
		level2 = userAccessDetails.Data.AccessList[1].Level

	}
	response := `{
		"data": {
		"type": "users",
		"id": "OhnzPVXY",
		"attributes": {
		"first-name": "` + userDetails.Data.Attributes.FirstName + `",
		"last-name": "` + userDetails.Data.Attributes.LastName + `",
		"role": "` + role + `",
		"email":"` + userDetails.Data.Attributes.Email + `"
		},
		"accountAccessList": [
		{
		"account": "` + account1 + `",
		"level": "` + level1 + `"
		},
		{
		"account": "` + account2 + `",
		"level": "` + level2 + `" } ] } }`
	return response
}

func getReportconfigResponse() string {

	var (
		scheduled             string
		sendEmail             string
		suppressed            string
		includeChecks         string
		shouldEmailIncludePdf string
		shouldEmailIncludeCsv string
		message               string
		withChecks            string
		withoutChecks         string
	)
	if scheduled = "false"; reportConfigDetails.Data.Attributes.Configuration.Scheduled {
		scheduled = "true"
	}
	if sendEmail = "false"; reportConfigDetails.Data.Attributes.Configuration.SendEmail {
		sendEmail = "true"
	}
	if suppressed = "false"; reportConfigDetails.Data.Attributes.Configuration.Filter.Suppressed {
		suppressed = "true"
	}
	if includeChecks = "false"; reportConfigDetails.Data.Attributes.Configuration.IncludeChecks {
		includeChecks = "true"
	}
	if shouldEmailIncludePdf = "false"; reportConfigDetails.Data.Attributes.Configuration.ShouldEmailIncludePdf {
		shouldEmailIncludePdf = "true"
	}
	if shouldEmailIncludeCsv = "false"; reportConfigDetails.Data.Attributes.Configuration.ShouldEmailIncludeCsv {
		shouldEmailIncludeCsv = "true"
	}
	if message = "false"; reportConfigDetails.Data.Attributes.Configuration.Filter.Message {
		message = "true"
	}
	if withChecks = "false"; reportConfigDetails.Data.Attributes.Configuration.Filter.WithChecks {
		withChecks = "true"
	}
	if withoutChecks = "false"; reportConfigDetails.Data.Attributes.Configuration.Filter.WithoutChecks {
		withoutChecks = "true"
	}

	response := `{
		"data": {
			"type": "report-config",
			"id": "HksLj2_",
			"attributes": {
				"type": "report-config",
				"manual": false,
				"enabled": true,
				"configuration": {
					"title": "` + reportConfigDetails.Data.Attributes.Configuration.Title + `",
					"scheduled": ` + scheduled + `,
					"frequency": "` + reportConfigDetails.Data.Attributes.Configuration.Frequency + `",
					"tz": "` + reportConfigDetails.Data.Attributes.Configuration.Tz + `",
					"generateReportType": "` + reportConfigDetails.Data.Attributes.Configuration.GenerateReportType + `",
					"includeChecks": ` + includeChecks + `,
					"shouldEmailIncludePdf": ` + shouldEmailIncludePdf + `,
					"shouldEmailIncludeCsv": ` + shouldEmailIncludeCsv + `,
					"sendEmail": ` + sendEmail + `,
					"emails": [
						"` + reportConfigDetails.Data.Attributes.Configuration.Emails[0] + `"
					],
					"filter": {
						"categories": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.Categories[0] + `"
						],
						"complianceStandards": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.ComplianceStandards[0] + `"
						],
						"filterTags": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.FilterTags[0] + `"
						],
						"message": ` + message + `,
						"newerThanDays": ` + fmt.Sprintf(`%d`, reportConfigDetails.Data.Attributes.Configuration.Filter.NewerThanDays) + `,
						"olderThanDays": ` + fmt.Sprintf(`%d`, reportConfigDetails.Data.Attributes.Configuration.Filter.OlderThanDays) + `,
						"providers": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.Providers[0] + `"
						],
						"regions": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.Regions[0] + `"
						],
						"reportComplianceStandardId": "` + reportConfigDetails.Data.Attributes.Configuration.Filter.ReportComplianceStandardId + `",
						"resource": "` + reportConfigDetails.Data.Attributes.Configuration.Filter.Resource + `",
						"resourceSearchMode": "` + reportConfigDetails.Data.Attributes.Configuration.Filter.ResourceSearchMode + `",
						"resourceTypes": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.ResourceTypes[0] + `"
						],
						"riskLevels": "` + reportConfigDetails.Data.Attributes.Configuration.Filter.RiskLevels + `",
						"ruleIds": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.RuleIds[0] + `"
						],
						"services": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.Services[0] + `"
						],
						"statuses": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.Statuses[0] + `"
						],
						"suppressed": ` + suppressed + `,
						"suppressedFilterMode": "` + reportConfigDetails.Data.Attributes.Configuration.Filter.SuppressedFilterMode + `",
						"tags": [
							"` + reportConfigDetails.Data.Attributes.Configuration.Filter.Tags[0] + `"
						],
						"text": "` + reportConfigDetails.Data.Attributes.Configuration.Filter.Text + `",
						"withChecks": ` + withChecks + `,
						"withoutChecks": ` + withoutChecks + `

					}
				},
				"is-account-level": true,
				"is-group-level": false,
				"is-organisation-level": false
			},
			"relationships": {
				"organisation": {
					"data": {
						"type": "organisations",
						"id": "B1nHYYpwx"
					}
				},
				"account": {
					"data": {
						"type": "accounts",
						"id": "` + reportConfigDetails.Data.Attributes.AccountId + `"
					}
				},
				"group": {
					"data": {
						"type": "groups",
						"id": "` + reportConfigDetails.Data.Attributes.GroupId + `"
					}
				},
				"profile": {
					"data": null
				}
			}
		}
	}
	`
	return response
}
