package conformity

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceConformityCommSetting() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformityCommSettingCreate,
		ReadContext:   resourceConformityCommSettingRead,
		UpdateContext: resourceConformityCommSettingUpdate,
		DeleteContext: resourceConformityCommSettingDelete,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"email": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"users": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"sms": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"users": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"ms_teams": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"channel_name": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringLenBetween(0, 20),
						},
						"display_extra_data": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"display_resource": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"display_tags": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
						"display_introduced_by": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"slack": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringMatch(
								regexp.MustCompile("^[#|@]"), `channel must start with "#" or "@"`),
						},
						"channel_name": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringLenBetween(0, 20),
						},
						"display_extra_data": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"display_resource": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"display_tags": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
						"display_introduced_by": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"sns": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arn": {
							Type:     schema.TypeString,
							Required: true,
						},
						"channel_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"pager_duty": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"webhook": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_token": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"categories": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"security", "cost-optimisation", "reliability", "performance-efficiency",
									"operational-excellence"}, true),
							},
						},
						"compliances": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"AWAF", "CISAWSF", "CISAZUREF", "CISAWSTTW", "PCI", "HIPAA", "GDPR", "APRA",
									"NIST4", "SOC2", "NIST-CSF", "ISO27001", "AGISM", "ASAE-3150", "MAS", "FEDRAMP"}, true),
							},
						},
						"statuses": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"SUCCESS", "FAILURE"}, true),
							},
						},
						"filter_tags": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"regions": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									"global", "us-east-1", "us-east-2", "us-west-1", "us-west-2", "ap-south-1", "ap-northeast-2", "ap-southeast-1", "ap-southeast-2",
									"ap-northeast-1", "eu-central-1", "eu-west-1", "eu-west-2", "eu-west-3", "eu-north-1", "sa-east-1", "ca-central-1", "me-south-1", "ap-east-1",
									"eastasia", "southeastasia", "centralus", "eastus", "eastus2", "westus", "northcentralus", "southcentralus", "southcentralus",
									"northeurope", "westeurope", "japanwest", "japaneast", "brazilsouth", "australiaeast", "australiasoutheast", "southindia", "centralindia",
									"westindia", "canadacentral", "canadaeast", "uksouth", "ukwest", "westcentralus", "westus2", "koreacentral", "koreasouth", "francecentral",
									"francesouth", "australiacentral", "australiacentral2", "southafricanorth", "southafricawest",
								}, true),
							},
						},
						"risk_levels": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"LOW", "MEDIUM", "HIGH", "VERY_HIGH", "EXTREME"}, false),
							},
						},
						"rule_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"services": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									"EC2", "ELB", "EBS", "VPC", "S3", "CloudTrail", "Route53", "RDS", "IAM", "RTM", "KMS", "SNS", "SQS", "CloudFormation", "Config",
									"AutoScaling", "Redshift", "CloudWatch", "CloudWatchEvents", "CloudWatchLogs", "ResourceGroup", "SES", "DynamoDB", "ElastiCache",
									"Elasticsearch", "WorkSpaces", "ACM", "Budgets", "Inspector", "TrustedAdvisor", "Shield", "EMR", "WAF", "Lambda", "Support", "Kinesis",
									"Organizations", "EFS", "ElasticBeanstalk", "Macie", "ELBv2", "CloudConformity", "APIGateway", "GuardDuty", "Health", "ConfigService",
									"MQ", "Firehose", "SSM", "Route53Domains", "SageMaker", "DAX", "Neptune", "ECR", "Glue", "XRay", "SecretsManager", "DocumentDB", "DMS",
									"Miscellaneous", "EKS", "Backup", "StorageGateway", "ECS", "SecurityHub", "Comprehend", "WellArchitected", "AccessAnalyzer", "StorageAccounts",
									"SecurityCenter", "ActiveDirectory", "MySQL", "PostgreSQL", "Sql", "Monitor", "AppService", "Network", "ActivityLog", "VirtualMachines",
									"AKS", "KeyVault", "Locks", "AccessControl", "Advisor", "RecoveryServices", "Resources", "Subscriptions", "CosmosDB", "RedisCache", "Search",
									"AppInsights",
								}, false),
							},
						},
						"tags": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"relationships": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"organisation": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
		CustomizeDiff: customdiff.All(
			CustomizeDiffValidateConfiguration,
		),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceConformityCommSettingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	payload := cloudconformity.CommunicationSettings{}

	channel, err := getChannel(d)
	if err != nil {
		return diag.FromErr(err)
	}

	payload.Data.Attributes.Channel = strings.ReplaceAll(channel, "_", "-")
	payload.Data.Attributes.Type = "communication"
	payload.Data.Attributes.Enabled = d.Get("enabled").(bool)

	if v, ok := d.GetOk("filter"); ok && len(v.([]interface{})) > 0 {
		proccessInputCommSettingFilter(&payload, d)
	}
	if v, ok := d.GetOk(channel); ok && len(v.(*schema.Set).List()) > 0 {
		proccessInputCommSettingConfiguration(&payload, d, channel)
	}

	if v, ok := d.GetOk("relationships"); ok && len(v.([]interface{})) > 0 {
		proccessInputCommSettingRelationships(&payload, d)
	}

	CommunicationSettings, err := client.CreateCommunicationSetting(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(CommunicationSettings.Data[0].ID)

	resourceConformityCommSettingRead(ctx, d, m)
	return diags
}

func resourceConformityCommSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	commSettingId := d.Id()

	communicationSettings, err := client.GetCommunicationSetting(commSettingId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("enabled", communicationSettings.Data.Attributes.Enabled); err != nil {
		return diag.FromErr(err)
	}
	filter := flattenCommSettingFilter(communicationSettings.Data.Attributes.Filter)
	if err := d.Set("filter", filter); err != nil {
		return diag.FromErr(err)
	}
	config := flattenCommSettingConfiguration(communicationSettings.Data.Attributes.Configuration, communicationSettings.Data.Attributes.Channel)
	if err := d.Set(strings.ReplaceAll(communicationSettings.Data.Attributes.Channel, "-", "_"), config); err != nil {
		return diag.FromErr(err)
	}

	relationships := flattenCommSettingRelationships(communicationSettings.Data.Relationships)
	if err := d.Set("relationships", relationships); err != nil {
		return diag.FromErr(err)
	}

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceConformityCommSettingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type

	payload := cloudconformity.CommunicationSettings{}

	// enabled,channel,type are required during update
	channel, err := getChannel(d)
	if err != nil {
		return diag.FromErr(err)
	}
	payload.Data.Attributes.Type = "communication"
	payload.Data.Attributes.Enabled = d.Get("enabled").(bool)
	payload.Data.Attributes.Channel = strings.ReplaceAll(channel, "_", "-")

	if d.HasChange("filter") || d.HasChange("configuration") || d.HasChange("relationships") {

		if v, ok := d.GetOk("filter"); ok && len(v.([]interface{})) > 0 {
			proccessInputCommSettingFilter(&payload, d)
		}
		if v, ok := d.GetOk(channel); ok && len(v.(*schema.Set).List()) > 0 {
			proccessInputCommSettingConfiguration(&payload, d, channel)
		}

		if v, ok := d.GetOk("relationships"); ok && len(v.([]interface{})) > 0 {
			proccessInputCommSettingRelationships(&payload, d)
		}

		commSettingId := d.Id()
		_, err := client.UpdateCommunicationSetting(commSettingId, payload)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceConformityCommSettingRead(ctx, d, m)
}

func resourceConformityCommSettingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	commSettingId := d.Id()

	_, err := client.DeleteCommunicationSetting(commSettingId)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.

	d.SetId("")
	return diags
}

func proccessInputCommSettingFilter(payload *cloudconformity.CommunicationSettings, d *schema.ResourceData) {
	if d.Get("filter").([]interface{})[0] != nil {
		f := d.Get("filter").([]interface{})[0].(map[string]interface{})
		filter := cloudconformity.CommunicationFilter{}

		filter.Categories = expandStringList(f["categories"].(*schema.Set).List())
		filter.Compliances = expandStringList(f["compliances"].(*schema.Set).List())
		filter.Statuses = expandStringList(f["statuses"].(*schema.Set).List())
		filter.FilterTags = expandStringList(f["filter_tags"].(*schema.Set).List())
		filter.Regions = expandStringList(f["regions"].(*schema.Set).List())
		filter.RiskLevels = expandStringList(f["risk_levels"].(*schema.Set).List())
		filter.RuleIds = expandStringList(f["rule_ids"].(*schema.Set).List())
		filter.Services = expandStringList(f["services"].(*schema.Set).List())
		filter.Tags = expandStringList(f["tags"].(*schema.Set).List())

		payload.Data.Attributes.Filter = &filter
	}

}

func proccessInputCommSettingConfiguration(payload *cloudconformity.CommunicationSettings, d *schema.ResourceData, ch string) {

	if d.Get(ch).(*schema.Set).List() != nil {
		c := d.Get(ch).(*schema.Set).List()[0].(map[string]interface{})
		configuration := cloudconformity.CommunicationConfiguration{}

		switch ch {
		case "email", "sms":
			// do email/sms here
			log.Printf("[DEBUG] Conformity Communication setting channel: email")
			configuration.Users = expandStringList(c["users"].(*schema.Set).List())
		case "ms_teams", "slack":
			// do ms-teams/slack here
			log.Printf("[DEBUG] Conformity Communication setting channel: ms-teams and slack")
			configuration.Channel = c["channel"].(string)
			configuration.ChannelName = c["channel_name"].(string)
			configuration.DisplayExtraData = c["display_extra_data"].(bool)
			configuration.DisplayIntroducedBy = c["display_introduced_by"].(bool)
			configuration.DisplayResource = c["display_resource"].(bool)
			configuration.DisplayTags = c["display_tags"].(bool)
			configuration.Url = c["url"].(string)
		case "sns":
			// do sns here
			log.Printf("[DEBUG] Conformity Communication setting channel: sns")
			configuration.Arn = c["arn"].(string)
			configuration.ChannelName = c["channel_name"].(string)
		case "pager_duty":
			// do pager-duty here
			log.Printf("[DEBUG] Conformity Communication setting channel: pager-duty")
			configuration.ChannelName = c["channel_name"].(string)
			configuration.ServiceKey = c["service_key"].(string)
			configuration.ServiceName = c["service_name"].(string)
		case "webhook":
			// do webhook here
			log.Printf("[DEBUG] Conformity Communication setting channel: webhook")
			configuration.SecurityToken = c["security_token"].(string)
			configuration.Url = c["url"].(string)
		}

		payload.Data.Attributes.Configuration = &configuration
	}

}

func proccessInputCommSettingRelationships(payload *cloudconformity.CommunicationSettings, d *schema.ResourceData) {

	if v, ok := d.GetOk("relationships.0.account"); ok && len(v.([]interface{})) > 0 {

		a := d.Get("relationships.0.account").([]interface{})[0].(map[string]interface{})
		payload.Data.Relationships.Account.Data.ID = a["id"].(string)
		payload.Data.Relationships.Account.Data.Type = "accounts"

	}
	if v, ok := d.GetOk("relationships.0.organisation"); ok && len(v.([]interface{})) > 0 {

		o := d.Get("relationships.0.organisation").([]interface{})[0].(map[string]interface{})
		payload.Data.Relationships.Organisation.Data.ID = o["id"].(string)
		payload.Data.Relationships.Organisation.Data.Type = "organisations"

	}

}

func flattenCommSettingFilter(f *cloudconformity.CommunicationFilter) []interface{} {

	c := make(map[string]interface{})
	if f == nil {
		return []interface{}{c}
	}

	c["categories"] = f.Categories
	c["compliances"] = f.Compliances
	c["filter_tags"] = f.FilterTags
	c["regions"] = f.Regions
	c["risk_levels"] = f.RiskLevels
	c["rule_ids"] = f.RuleIds
	c["services"] = f.Services
	c["tags"] = f.Tags

	return []interface{}{c}
}
func flattenCommSettingConfiguration(config *cloudconformity.CommunicationConfiguration, channel string) []interface{} {

	c := make(map[string]interface{})
	if config == nil {
		return []interface{}{c}
	}
	switch channel {
	case "email", "sms":
		c["users"] = config.Users
	case "ms-teams", "slack":
		c["channel"] = config.Channel
		c["channel_name"] = config.ChannelName
		c["display_extra_data"] = config.DisplayExtraData
		c["display_resource"] = config.DisplayResource
		c["display_tags"] = config.DisplayTags
		c["url"] = config.Url
		c["display_introduced_by"] = config.DisplayIntroducedBy
	case "sns":
		c["channel_name"] = config.ChannelName
		c["arn"] = config.Arn
	case "pager-duty":
		c["channel_name"] = config.ChannelName
		c["service_key"] = config.ServiceKey
		c["service_name"] = config.ServiceName
	case "webhook":
		c["security_token"] = config.SecurityToken
		c["url"] = config.Url
	}
	return []interface{}{c}
}
func flattenCommSettingRelationships(r cloudconformity.CommunicaitonRelationships) []interface{} {
	c := make(map[string]interface{})

	c["organisation"] = flattenData(r.Organisation.Data)
	c["account"] = flattenData(r.Account.Data)

	return []interface{}{c}
}

func flattenData(data cloudconformity.CommunicaitonRelationshipsData) []interface{} {
	c := make(map[string]interface{})
	c["id"] = data.ID
	return []interface{}{c}
}
func getChannel(d *schema.ResourceData) (string, error) {

	list := []string{"email", "sms", "ms_teams", "slack", "sns", "pager_duty", "webhook"}

	for _, ch := range list {
		if _, ok := d.GetOk(ch); ok {
			return ch, nil
		}
	}
	return "", fmt.Errorf("no channel configuration set found, please provide one ot this: %v", list)
}
