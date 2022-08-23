package conformity

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccConformityProviders map[string]*schema.Provider
var testAccConformityProvider *schema.Provider
var accountPayload cloudconformity.AccountPayload
var gcpOrgPayload cloudconformity.GCPOrgPayload
var groupDetails cloudconformity.GroupDetails
var userDetails cloudconformity.UserDetails
var userAccessDetails cloudconformity.UserAccessDetails
var reportConfigDetails cloudconformity.ReportConfigDetails
var communicationSetting cloudconformity.CommunicationSettings
var profileSetting cloudconformity.ProfileSettings
var botSetting cloudconformity.AccountBotSettingsRequest
var ruleSetting1, ruleSetting2, ruleSetting3 *cloudconformity.AccountRuleSettings
var testServer *httptest.Server
var checkDetails *cloudconformity.CheckDetails
var customRule *cloudconformity.CustomRule
var customRuleResponse = cloudconformity.CustomRuleResponse{}
var counterGroupReads = 0

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

		var getOrganizationalExternalId = regexp.MustCompile(`^/organisation/external-id/$`)
		var getCurrentUser = regexp.MustCompile(`^/users/whoami$`)
		var postApplyProfile = regexp.MustCompile(`^/profiles/(.*)/apply$`)
		var postAccount = regexp.MustCompile(`^/accounts/$`)
		var patchAccountRuleSetting = regexp.MustCompile(`^/accounts/(.*)/settings/rules/(.*)$`)
		var getAccountRuleSetting = regexp.MustCompile(`^/accounts/(.*)/settings/rules$`)
		var postAzureAccount = regexp.MustCompile(`^/accounts/azure/$`)
		var postGCPAccount = regexp.MustCompile(`^/accounts/gcp/$`)
		var postGCPOrg = regexp.MustCompile(`^/gcp/organisations/$`)
		var accountDetails = regexp.MustCompile(`^/accounts/(.*)$`)
		var getAccountAccess = regexp.MustCompile(`^/accounts/(.*)/access$`)
		var postGroup = regexp.MustCompile(`^/groups/$`)
		var getGroupDetails = regexp.MustCompile(`^/groups/(.*)$`)
		var postUsers = regexp.MustCompile(`^/users/$`)
		var postSsoUsers = regexp.MustCompile(`^/users/sso/$`)
		var getUserDetails = regexp.MustCompile(`^/users/(.*)$`)
		var postReportconfig = regexp.MustCompile(`^/report-configs/$`)
		var getReportconfig = regexp.MustCompile(`^/report-configs/(.*)$`)
		var postCommunicationConfig = regexp.MustCompile(`^/settings/communication/$`)
		var getCommunicationConfig = regexp.MustCompile(`^/settings/(.*)$`)
		var postProfile = regexp.MustCompile(`^/profiles/$`)
		var getProfile = regexp.MustCompile(`^/profiles/(.*)$`)
		var patchBotSettings = regexp.MustCompile(`^/accounts/(.*)/settings/bot$`)
		var getCheck = regexp.MustCompile(`^/checks/(.*)$`)
		var patchCheck = regexp.MustCompile(`^/checks/(.*)$`)
		var getAzureSubscriptions = regexp.MustCompile(`^/azure/active-directories/(.*)/subscriptions/?(.*)$`)
		var postAzureActiveDirectory = regexp.MustCompile(`^/azure/active-directories$`)
		var getGcpProjects = regexp.MustCompile(`^/gcp/organisations/(.*)/projects/?(.*)$`)
		var endPointCustomRule = regexp.MustCompile(`^/custom-rules/(.*)$`)

		switch {
		case getOrganizationalExternalId.MatchString(r.URL.Path):
			w.Write([]byte(`{ "data": { "type": "external-ids", "id": "3ff84b20-0f4c-11eb-a7b7-7d9b3c0e866e" } }`))
		case getCurrentUser.MatchString(r.URL.Path):
			w.Write([]byte(`
			{
  "data": {
    "type": "users",
    "id": "517uNyIvG",
    "attributes": {
    "first-name": "John",
	"last-name": "Smith",
	"email": "john.smith@company.com",
	"status": "ACTIVE",
       "created_date":0,
      "has_credentials":false,
       "is_api_key_user" : false,
		"is_cloud_one_user" : false,
		"last_login_date" : 0,
		"mfa" : false,
		"role" : "ADMIN",
		"summary_email_opt_out" : true
	
    },
    "relationships": {
      "organisation": {
        "data": {
          "type": "organisations",
          "id": "B1nHYYpwx"
        }
      },
      "accountAccessList": [
        {
          "account": "A9_DsY12z",
          "level": "NONE"
        }
      ]
    }
  }
}
			
			`))
		case postApplyProfile.MatchString(r.URL.Path):
			w.Write([]byte(`{
				"meta": {
					"status": "sent",
					"message": "Profile will be applied to the accounts in background"
				}
			}`))
		case postAccount.MatchString(r.URL.Path):
			_ = readRequestBody(r, &accountPayload)
			w.Write([]byte(`{ "data": { "type": "accounts", "id": "H19NxMi5-" } }`))
		case postAzureAccount.MatchString(r.URL.Path):
			_ = readRequestBody(r, &accountPayload)
			w.Write([]byte(`{ "data": { "type": "accounts", "id": "H19NxMi5-" } }`))
		case postGCPAccount.MatchString(r.URL.Path):
			_ = readRequestBody(r, &accountPayload)
			w.Write([]byte(`{ "data": { "type": "accounts", "id": "H19NxMi5-" } }`))
		case postGCPOrg.MatchString(r.URL.Path):
			_ = readRequestBody(r, &accountPayload)
			w.Write([]byte(`{ "data": { "type": "gcp-organisations", "id": "H19NxMi5-" } }`))
		case getAccountAccess.MatchString(r.URL.Path):
			w.Write([]byte(`{
				"id": "BJ0Ox16Hb:access",
				"type": "settings",
				"attributes": {
				"type": "access",
				"configuration": {
				"externalId": "test-external-id",
				"roleArn": "test-arn" } } }`))
		case accountDetails.MatchString(r.URL.Path) && !getAccountRuleSetting.MatchString(r.URL.Path) && r.Method == "GET":
			tags, _ := json.Marshal(accountPayload.Data.Attributes.Tags)
			w.Write([]byte(`{
				"data": {	
				"type": "accounts",
				"id": "H19NxMi5-",
				"attributes": {
				"name": "` + accountPayload.Data.Attributes.Name + `",
				"environment": "` + accountPayload.Data.Attributes.Environment + `", 
				"tags": ` + string(tags) + `,
				"settings": {
					"bot": {
						"disabled": false,
						"disabledUntil": 0,
						"delay": ` + strconv.Itoa(*botSetting.Data.Attributes.Settings.Bot.Delay) + `,
						"disabledRegions": {
							"ap-east-1": true,
							"ap-south-1": true
						}
					}
				},
				"cloud-type": "azure",
				"cloud-data": {
				"azure": {
				"subscriptionId": "test-subscrition-id"
				} } } } }`))

		case accountDetails.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`{
				"meta": {
				"status": "sent" } }`))
		case accountDetails.MatchString(r.URL.Path) && !patchBotSettings.MatchString(r.URL.Path) && !patchAccountRuleSetting.MatchString(r.URL.Path) && r.Method == "PATCH":
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

			if counterGroupReads > 1 && groupDetails.Data.Attributes.Name == "404-group" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte(`{
					"errors": [
						{
							"status": 422,
							"detail": "Group ID entered is invalid"
						}
					]
				}`))
				break
			} else if groupDetails.Data.Attributes.Name == "404-group" {
				counterGroupReads++
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

		case postProfile.MatchString(r.URL.Path) || (getProfile.MatchString(r.URL.Path) && r.Method == "PATCH"):
			_ = readRequestBody(r, &profileSetting)
			w.Write([]byte(`{ "data": {
						"type": "profiles",
						"id": "d9yHTrzP0" } }`))

		case getProfile.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(`{
				"included": [
				  {
					"type": "rules",
					"id": "RTM-002",
					"attributes": {
					  "enabled": true,
					  "exceptions": {
						"tags": ["some_tag"]
					  },
					  "provider": "aws",
					  "extraSettings": [
						{
						  "name": "ttl",
						  "type": "ttl",
						  "value": ` + fmt.Sprintf("%v", profileSetting.Included[0].Attributes.ExtraSettings[0].Value) + `,
						  "ttl": true
						}
					  ]
					}
				  },
				  {
					"type": "rules",
					"id": "SNS-002",
					"attributes": {
					  "enabled": true,
					  "provider": "aws",
					  "extraSettings": [
						{
							"name": "conformityOrganization",
							"type": "choice-multiple-value",
							"values": [{
								"label": "All within this AWS Organization",
								"value": "includeAwsOrganizationAccounts",
								"enabled": true
							}, {
								"label": "All within this Conformity organization",
								"value": "includeConformityOrganization"
							}]
						}
					  ]
					}
				  }
				],
				"data": {
				  "type": "profiles",
				  "id": "d9yHTrzP0",
				  "attributes": {
					"name": "test-with-rules",
					"description": "conformity development - rules included"
				  },
				  "relationships": {
					"ruleSettings": {
					  "data": [
						{
						  "type": "rules",
						  "id": "RTM-002"
						}
					  ]
					}
				  }
				}
			  }`))

		case getProfile.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`{
				"meta": {
				"status": "deleted" } }`))

		case patchBotSettings.MatchString(r.URL.Path) && r.Method == "PATCH":
			_ = readRequestBody(r, &botSetting)
			w.Write([]byte(`{
					"data": [
					  {
						"type": "accounts",
						"id": "AgA12vIwb"
					  }
					]
				  }`))
		case patchAccountRuleSetting.MatchString(r.URL.Path) && r.Method == "PATCH":

			if ruleSetting1 == nil {
				_ = readRequestBody(r, &ruleSetting1)
			} else if ruleSetting2 == nil {
				_ = readRequestBody(r, &ruleSetting2)
			} else {
				_ = readRequestBody(r, &ruleSetting3)
			}

			w.Write([]byte(`{"data": {"type": "accounts","id": "AgA12vIwb"}}`))
		case getAccountRuleSetting.MatchString(r.URL.Path):
			rule1 := `null`
			rule2 := `null`
			rule3 := `null`

			if ruleSetting1 != nil {
				values := ruleSetting1.Data.Attributes.RuleSetting.ExtraSettings[0].Values.([]interface{})
				values1 := values[0].(map[string]interface{})
				values2 := values[1].(map[string]interface{})
				rule1 = `
				{
					"riskLevel": "MEDIUM",
					"id": "` + ruleSetting1.Data.Attributes.RuleSetting.Id + `",
					"exceptions": {
						"tags": [
							"` + ruleSetting1.Data.Attributes.RuleSetting.Exceptions.Tags[0] + `"
						]
					},
					"extraSettings": [{
						"name": "` + ruleSetting1.Data.Attributes.RuleSetting.ExtraSettings[0].Name + `",
						"countries": true,
						"type": "` + ruleSetting1.Data.Attributes.RuleSetting.ExtraSettings[0].Type + `",
						"value": null,
							"values": [{
									"value": "` + values1["value"].(string) + `",
									"label": "` + values1["label"].(string) + `",
									"enabled": true
								},
								{
									"value": "` + values2["value"].(string) + `",
									"label": "` + values2["label"].(string) + `",
									"enabled": true
								}
							]
					}],
					"provider": "aws",
					"enabled": false
				}
				`
			}
			if ruleSetting2 != nil {
				rule2 = `
				{
					"ruleExists": false,
					"riskLevel": "MEDIUM",
					"extraSettings": [{
						"name": "patterns",
						"type": "multiple-object-values",
						"valueKeys": ["eventName", "eventSource", "userIdentityType"],
						"values": [{
							"value": {
								"eventSource": "^(IAM).*",
								"eventName": "^(iam.amazonaws.com)",
								"userIdentityType": "^(Delete).*"
							}
						}]
					}],
					"provider": "aws",
					"id": "RTM-011",
					"enabled": true,
					"exceptions": null
				}
				`
			}
			if ruleSetting3 != nil {

				rule3 = `
				{
					"ruleExists": false,
					"riskLevel": "LOW",
					"extraSettings": [{
						"name": "SpecificVPCToSpecificGatewayMapping",
						"type": "multiple-vpc-gateway-mappings",
						"mappings": [{
							"values": [{
								"type": "multiple-string-values",
								"name": "gatewayIds",
								"values": [{
									"value": "nat-001"
								}, {
									"value": "nat-002"
								}]
							}, {
								"type": "single-string-value",
								"name": "vpcId",
								"value": "vpc-001"
							}]
						}]
					}],
					"provider": "aws",
					"id": "VPC-013",
					"enabled": true,
					"exceptions": null
				}
				`

			}
			w.Write([]byte(`
					{"data": 
						{"type": 
							"accounts","id": "H19NxMi5-",
							"attributes": 
							{"settings": 
								{"rules": [` + rule1 + `,` + rule2 + `,` + rule3 + `]}
							}		
						}
					}`))
		case getCheck.MatchString(r.URL.Path) && r.Method == "GET":
			if checkDetails == nil {
				checkDetails.Data.Id = "ccc:8d99dfce-dca2-4f13-8699-20631a5c77c9:Resources-001:Resources:global:/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483"
				checkDetails.Data.Attributes.Region = "global"
				checkDetails.Data.Attributes.Suppressed = false
				checkDetails.Data.Attributes.Resource = "/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483"
				checkDetails.Data.Relationships.Account.Data.Id = "8d99dfce-dca2-4f13-8699-20631a5c77c9"
				checkDetails.Data.Relationships.Rule.Data.Id = "Resources-001"
			}
			w.Write([]byte(getCheckDetailsResponse()))
		case patchCheck.MatchString(r.URL.Path) && r.Method == "PATCH":
			_ = readRequestBody(r, &checkDetails)
			split := strings.Split(r.URL.RawPath, "/")
			checkDetails.Data.Id, _ = url.QueryUnescape(split[2])
			w.Write([]byte(getCheckDetailsResponse()))

		case getAzureSubscriptions.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(testGetAzureSubscriptions200Response))
		case postAzureActiveDirectory.MatchString(r.URL.Path) && r.Method == "POST":
			w.Write([]byte(testPostAzureActiveDirectory200Response))
		case getGcpProjects.MatchString(r.URL.Path) && r.Method == "GET":
			w.Write([]byte(testGetGcpProjects200Response))

		case endPointCustomRule.MatchString(r.URL.Path) && (r.Method == "POST" || r.Method == "PUT"):
			_ = readRequestBody(r, &customRule)
			customRuleResponse.ID = "some_id"
			customRuleResponse.Type = "CustomRule"
			customRuleResponse.Attributes.Name = customRule.Name
			customRuleResponse.Attributes.Description = customRule.Description
			customRuleResponse.Attributes.Enabled = customRule.Enabled
			customRuleResponse.Attributes.Categories = customRule.Categories
			customRuleResponse.Attributes.Severity = customRule.Severity
			customRuleResponse.Attributes.RemediationNotes = customRule.RemediationNotes
			customRuleResponse.Attributes.Service = customRule.Service
			customRuleResponse.Attributes.Provider = customRule.Provider
			customRuleResponse.Attributes.ResourceType = customRule.ResourceType
			customRuleResponse.Attributes.Attributes = customRule.Attributes
			customRuleResponse.Attributes.Rules = customRule.Rules
			bytes, _ := json.Marshal(cloudconformity.CustomRuleCreateResponse{Data: customRuleResponse})
			w.Write(bytes)
		case endPointCustomRule.MatchString(r.URL.Path) && r.Method == "GET":
			bytes, _ := json.Marshal(cloudconformity.CustomRuleGetResponse{Data: []cloudconformity.CustomRuleResponse{customRuleResponse}})
			w.Write(bytes)
		case endPointCustomRule.MatchString(r.URL.Path) && r.Method == "DELETE":
			w.Write([]byte(`{"meta": {"status": "deleted" } }`))
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
						"riskLevels": ["` + reportConfigDetails.Data.Attributes.Configuration.Filter.RiskLevels[0] + `"],
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

func getCheckDetailsResponse() string {
	response := `{
    "data": {
        "type": "checks",
        "id": "ccc:8d99dfce-dca2-4f13-8699-20631a5c77c9:Resources-001:Resources:global:/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483",
        "attributes": {
            "region": "global",
            "status": "FAILURE",
            "risk-level": "LOW",
            "pretty-risk-level": "Low",
            "message": "00482a5a-887f-4fb3-b363-3b7fe8e74483 has [Owner, CostID, Application, DemandID, Stage, ITServiceID, ITServiceName] tags missing",
            "resource": "/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483",
            "descriptorType": "access-control-roles",
            "link-title": "/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483",
            "resourceName": "00482a5a-887f-4fb3-b363-3b7fe8e74483",
            "last-modified-date": 1656482101900,
            "created-date": 1653906818310,
            "categories": [
                "security",
                "reliability",
                "performance-efficiency",
                "cost-optimisation",
                "operational-excellence"
            ],
            "compliances": [
                "NIST5",
                "NIST-CSF",
                "AGISM",
                "HITRUST",
                "MAS",
                "CSA"
            ],
            "suppressed": false,
            "failure-discovery-date": 1656482101900,
            "ccrn": "ccrn:azure:8d99dfce-dca2-4f13-8699-20631a5c77c9:AccessControl:global:/subscriptions/ae9124e0-d61c-4d7d-833d-c58e6f9941f8/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483",
            "extradata": [
                {
                    "name": "DETAILED_STATUS",
                    "label": "Resource tags status for 00482a5a-887f-4fb3-b363-3b7fe8e74483",
                    "value": "{\"service\":\"AccessControl\",\"resourceName\":\"00482a5a-887f-4fb3-b363-3b7fe8e74483\",\"tags\":[{\"key\":\"Owner\",\"hasValue\":false},{\"key\":\"CostID\",\"hasValue\":false},{\"key\":\"Application\",\"hasValue\":false},{\"key\":\"DemandID\",\"hasValue\":false},{\"key\":\"Stage\",\"hasValue\":false},{\"key\":\"ITServiceID\",\"hasValue\":false},{\"key\":\"ITServiceName\",\"hasValue\":false}]}",
                    "type": "META",
                    "internal": true
                }
            ],
            "tags": [],
            "suppressed-until": null,
            "not-scored": false,
            "excluded": false,
            "rule-title": "Tags",
            "provider": "azure",
            "resolution-page-url": "https://www.cloudconformity.com/knowledge-base/azure/Resources/use-tags-organize-resources.html#503824401549",
            "service": "Resources"
        },
        "relationships": {
            "rule": {
                "data": {
                    "type": "rules",
                    "id": "` + strings.Split(checkDetails.Data.Id, ":")[2] + `"
                }
            },
            "account": {
                "data": {
                    "type": "accounts",
                    "id": "8d99dfce-dca2-4f13-8699-20631a5c77c9"
                }
            }
        }
    }
}
`
	return response
}

var testGetGcpProjects200Response = `{
    "data": [
      {
        "type": "projects",
        "attributes": {
          "project-number": "415104041262",
          "project-id": "project-id-1",
          "lifecycle-state": "ACTIVE",
          "added-to-conformity": true,
          "create-time": "2021-05-17T11:21:58.012Z",
          "name": "My Project",
          "parent": {
            "type": "folder",
            "id": "415104041262"
          }
        }
      }
    ]
  }`

var testGetAzureSubscriptions200Response = `{
    "data": [
      {
        "type": "subscriptions",
        "id": "AZURE_SUBSCRIPTION_ID",
        "attributes": {
          "display-name": "A Azure Subscription",
          "state": "Enabled",
          "added-to-conformity": true
        }
      }
    ]
  }`
var testPostAzureActiveDirectory200Response = `{
  "data": {
    "type": "active-directories",
    "id": "CREATED_ACITVE_DIRECTORY_ID",
    "attributes": {
      "name": "MyAzureActiveDirectory",
      "directory-id": "YOUR_ACTIVE_DIRECTORY_ID",
      "created-date": 1635230845449,
      "last-modified-date": 1635230845449
    }
  }
}`
