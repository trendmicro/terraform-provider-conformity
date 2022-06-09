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
	"strconv"
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

		switch {
		case getOrganizationalExternalId.MatchString(r.URL.Path):
			w.Write([]byte(`{ "data": { "type": "external-ids", "id": "3ff84b20-0f4c-11eb-a7b7-7d9b3c0e866e" } }`))
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
