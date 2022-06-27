package conformity

import (
	"context"
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/robfig/cron"
)

func resourceConformityReportConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConformityReportConfigCreate,
		ReadContext:   resourceConformityReportConfigRead,
		UpdateContext: resourceConformityReportConfigUpdate,
		DeleteContext: resourceConformityReportConfigDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"configuration": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 0,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"emails": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"frequency": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(string)
								// Standard parser without descriptors, excludes time fields
								specParser := cron.NewParser(cron.Dom | cron.Month | cron.Dow)
								_, err := specParser.Parse(v)
								if err != nil {
									errs = append(errs, fmt.Errorf("%q expected cron expression, got: %s", key, v))
								}
								return
							},
						},
						"generate_report_type": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "GENERIC",
							ValidateFunc: validation.StringInSlice([]string{"GENERIC", "COMPLIANCE-STANDARD"}, false),
						},
						"include_checks": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"scheduled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"send_email": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"should_email_include_csv": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"should_email_include_pdf": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"title": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tz": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(string)
								timezones := getTimezones()
								for _, item := range timezones {
									if item == v {
										return
									}
								}
								errs = append(errs, fmt.Errorf(`given value of "tz" is not supported, got: %s`, v))
								return
							},
						},
					},
				},
			},
			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"categories": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"security", "cost-optimisation", "reliability", "performance-efficiency",
									"operational-excellence"}, false),
							},
						},
						"compliance_standards": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"AWAF", "CISAWSF", "CISAZUREF", "PCI", "HIPAA", "GDPR", "APRA", "NIST4", "SOC2",
									"NIST-CSF", "ISO27001", "AGISM", "ASAE-3150", "MAS", "FEDRAMP"}, false),
							},
						},
						"filter_tags": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"message": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"newer_than_days": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"older_than_days": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"providers": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"aws", "azure"}, true),
							},
						},
						"regions": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"report_compliance_standard_id": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{"AWAF", "CISAWSF", "CISAZUREF", "PCI", "HIPAA", "GDPR", "APRA", "NIST4", "SOC2",
								"NIST-CSF", "ISO27001", "AGISM", "ASAE-3150", "MAS", "FEDRAMP"}, false),
						},
						"resource": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"resource_search_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"resource_types": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"risk_levels": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
								ValidateFunc: validation.StringInSlice([]string{"EC2", "ELB", "EBS", "VPC", "S3", "CloudTrail", "Route53", "RDS", "IAM",
									"RTM", "KMS", "SNS", "SQS", "CloudFormation", "Config", "CloudFront", "AutoScaling", "Redshift", "CloudWatch",
									"CloudWatchEvents", "CloudWatchLogs", "ResourceGroup", "SES", "DynamoDB", "ElastiCache", "Elasticsearch", "WorkSpaces",
									"ACM", "Budgets", "Inspector", "TrustedAdvisor", "Shield", "EMR", "WAF", "Lambda", "Support", "Kinesis", "Organizations",
									"EFS", "ElasticBeanstalk", "Macie", "ELBv2", "CloudConformity", "APIGateway", "GuardDuty", "Health", "ConfigService", "MQ",
									"Firehose", "SSM", "Route53Domains", "SageMaker", "DAX", "Neptune", "ECR", "Glue", "XRay", "SecretsManager", "DocumentDB",
									"DMS", "Miscellaneous", "EKS", "Backup", "StorageGateway", "ECS", "SecurityHub", "Comprehend", "WellArchitected",
									"AccessAnalyzer", "StorageAccounts", "SecurityCenter", "ActiveDirectory", "MySQL", "PostgreSQL", "Sql", "Monitor",
									"AppService", "Network", "ActivityLog", "VirtualMachines", "AKS", "KeyVault", "Locks", "AccessControl", "Advisor",
									"RecoveryServices", "Resources", "Subscriptions", "CosmosDB", "RedisCache", "Search", "AppInsights"}, false),
							},
						},
						"statuses": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"SUCCESS", "FAILURE"}, false),
							},
						},
						"suppressed": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"suppressed_filter_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "v1",
							ValidateFunc: validation.StringInSlice([]string{"v1", "v2"}, false),
						},
						"tags": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"with_checks": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"without_checks": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		}, CustomizeDiff: customdiff.Sequence(
			CustomizeDiffValidateFrequency,
			CustomizeDiffValidateTz,
			CustomizeDiffValidateSendEmail,
			CustomizeDiffValidateScheduledFrequency,
			CustomizeDiffValidateScheduledTz,
		),
	}
}

func resourceConformityReportConfigCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := cloudconformity.ReportConfigDetails{}
	payload.Data.Attributes.AccountId = d.Get("account_id").(string)
	payload.Data.Attributes.GroupId = d.Get("group_id").(string)

	if v, ok := d.GetOk("configuration"); ok && len(v.([]interface{})) > 0 {
		proccessInputConfiguration(&payload, d)
	}

	if v, ok := d.GetOk("filter"); ok && len(v.([]interface{})) > 0 {
		proccessInputFilter(&payload, d)
	}

	reportId, err := client.CreateReportConfig(payload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(reportId)

	resourceConformityReportConfigRead(ctx, d, m)
	return diags
}

func resourceConformityReportConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	reportId := d.Id()

	reportConfigDetails, err := client.GetReportConfig(reportId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("account_id", reportConfigDetails.Data.Relationships.Account.Data.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_id", reportConfigDetails.Data.Relationships.Group.Data.ID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("configuration", flattenConfiguration(reportConfigDetails.Data.Attributes.Configuration)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("filter", flattenFilter(reportConfigDetails.Data.Attributes.Configuration.Filter)); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceConformityReportConfigUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	payload := cloudconformity.ReportConfigDetails{}

	if d.HasChange("configuration") || d.HasChange("filter") {

		if v, ok := d.GetOk("configuration"); ok && len(v.([]interface{})) > 0 {
			proccessInputConfiguration(&payload, d)
		}

		if v, ok := d.GetOk("filter"); ok && len(v.([]interface{})) > 0 {
			proccessInputFilter(&payload, d)
		}

		reportId := d.Id()
		_, err := client.UpdateReportConfig(reportId, payload)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChange("account_id") || d.HasChange("group_id") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to Update Conformity Group Config",
			Detail:   "'account_id', and 'group_id' cannot be changed",
		})

		return diags
	}
	return resourceConformityReportConfigRead(ctx, d, m)
}

func resourceConformityReportConfigDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudconformity.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	reportId := d.Id()

	_, err := client.DeleteReportConfig(reportId)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.

	d.SetId("")
	return diags
}

func proccessInputConfiguration(payload *cloudconformity.ReportConfigDetails, d *schema.ResourceData) {

	configuration := d.Get("configuration").([]interface{})[0].(map[string]interface{})

	payload.Data.Attributes.Configuration.Frequency = configuration["frequency"].(string)
	payload.Data.Attributes.Configuration.Emails = expandStringList(configuration["emails"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.GenerateReportType = configuration["generate_report_type"].(string)
	payload.Data.Attributes.Configuration.IncludeChecks = configuration["include_checks"].(bool)
	payload.Data.Attributes.Configuration.Scheduled = configuration["scheduled"].(bool)
	payload.Data.Attributes.Configuration.SendEmail = configuration["send_email"].(bool)
	payload.Data.Attributes.Configuration.ShouldEmailIncludeCsv = configuration["should_email_include_csv"].(bool)
	payload.Data.Attributes.Configuration.ShouldEmailIncludePdf = configuration["should_email_include_pdf"].(bool)
	payload.Data.Attributes.Configuration.Title = configuration["title"].(string)
	payload.Data.Attributes.Configuration.Tz = configuration["tz"].(string)

}

func proccessInputFilter(payload *cloudconformity.ReportConfigDetails, d *schema.ResourceData) {

	f := d.Get("filter").([]interface{})[0].(map[string]interface{})

	payload.Data.Attributes.Configuration.Filter.Categories = expandStringList(f["categories"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.ComplianceStandards = expandStringList(f["compliance_standards"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.FilterTags = expandStringList(f["filter_tags"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.Message = f["message"].(bool)
	payload.Data.Attributes.Configuration.Filter.NewerThanDays = f["newer_than_days"].(int)
	payload.Data.Attributes.Configuration.Filter.OlderThanDays = f["older_than_days"].(int)
	payload.Data.Attributes.Configuration.Filter.Providers = expandStringList(f["providers"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.Regions = expandStringList(f["regions"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.ReportComplianceStandardId = f["report_compliance_standard_id"].(string)
	payload.Data.Attributes.Configuration.Filter.Resource = f["resource"].(string)
	payload.Data.Attributes.Configuration.Filter.ResourceSearchMode = f["resource_search_mode"].(string)
	payload.Data.Attributes.Configuration.Filter.ResourceTypes = expandStringList(f["resource_types"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.RiskLevels = f["risk_levels"].(string)
	payload.Data.Attributes.Configuration.Filter.RuleIds = expandStringList(f["rule_ids"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.Services = expandStringList(f["services"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.Statuses = expandStringList(f["statuses"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.Suppressed = f["suppressed"].(bool)
	payload.Data.Attributes.Configuration.Filter.SuppressedFilterMode = f["suppressed_filter_mode"].(string)
	payload.Data.Attributes.Configuration.Filter.Tags = expandStringList(f["tags"].(*schema.Set).List())
	payload.Data.Attributes.Configuration.Filter.Text = f["text"].(string)
	payload.Data.Attributes.Configuration.Filter.WithChecks = f["with_checks"].(bool)
	payload.Data.Attributes.Configuration.Filter.WithoutChecks = f["without_checks"].(bool)

}
func flattenFilter(f cloudconformity.ReportConfigFilter) []interface{} {

	c := make(map[string]interface{})

	c["categories"] = f.Categories
	c["compliance_standards"] = f.ComplianceStandards
	c["filter_tags"] = f.FilterTags
	c["message"] = f.Message
	c["newer_than_days"] = f.NewerThanDays
	c["older_than_days"] = f.OlderThanDays
	c["providers"] = f.Providers
	c["regions"] = f.Regions
	c["report_compliance_standard_id"] = f.ReportComplianceStandardId
	c["resource"] = f.Resource
	c["resource_search_mode"] = f.ResourceSearchMode
	c["resource_types"] = f.ResourceTypes
	c["risk_levels"] = f.RiskLevels
	c["rule_ids"] = f.RuleIds
	c["services"] = f.Services
	c["statuses"] = f.Statuses
	c["suppressed"] = f.Suppressed
	c["suppressed_filter_mode"] = f.SuppressedFilterMode
	c["tags"] = f.Tags
	c["text"] = f.Text
	c["with_checks"] = f.WithChecks
	c["without_checks"] = f.WithoutChecks

	return []interface{}{c}
}

func flattenConfiguration(config cloudconformity.ReportConfiguration) []interface{} {

	c := make(map[string]interface{})
	c["emails"] = config.Emails
	c["frequency"] = config.Frequency
	c["generate_report_type"] = config.GenerateReportType
	c["include_checks"] = config.IncludeChecks
	c["scheduled"] = config.Scheduled
	c["send_email"] = config.SendEmail
	c["should_email_include_csv"] = config.ShouldEmailIncludeCsv
	c["should_email_include_pdf"] = config.ShouldEmailIncludePdf
	c["title"] = config.Title
	c["tz"] = config.Tz

	return []interface{}{c}
}

func expandStringList(list []interface{}) []string {
	vs := make([]string, 0, len(list))
	for _, v := range list {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
}
