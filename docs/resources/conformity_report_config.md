---
page_title: "conformity_report_config Resource"
subcategory: "Report Configs"
description: |-
  Allows you to create Report Configs on Conformity. 
---

# Resource `conformity_report_config`
Allows you to create Report Configs on Conformity. 

## Example Usage
```hcl
resource "conformity_report_config" "report" {
  // optional | type: string
  // if you want to create account-level report config, uncomment account_id value and provide account ID on the terraform.tfvars
  // if you want to create group-level report config, uncomment group_id value and provide account ID on the terraform.tfvars
  // if you want to create organisation-level report config, leave account_id and group_id commented.
  // account_id = var.account_id
  // group_id = var.group_id

  configuration { 
    // optional | type: array of strings
    emails = ["userone@trendmicro.com","usertwo@trendmicro.com"]
    // optional | type: string
    frequency = "* * *"
    // required | type: string 
    title = "my first report config"
    // optional | type: string
    tz = "Asia/Manila"
  }
}
```

## Argument reference

 - `account_id` (String) - (Optional) The Conformity ID of the account. Provide to get only report configs for the specified account. 
 - `group_id` (String) - (Optional) The Conformity ID of the group. Provide to get only report configs for the specified group. Notice: if you provided accountId at the same time, groupId would be ignored. 
 
IMPORTANT: Some guidelines about using this endpoint:
Each report config can be account-level, group-level, or organisation-level.
If creating an account-level report config, you must have a valid accountId.
If creating a group-level report config, you must have a valid groupId. If you provided accountId and groupId at the same time, groupId would be ignored.
If creating an organisation-level report config you don't provide any accountId or groupId.
Only ADMIN/POWER users can create organisation-level and group-level report-configs.
 
 - `configuration` - (Required) List:
     * `title` (String)  (Required) This attribute is report title.
     *  `email` (Array of Strings) - (Optional) Represents email addresses that report would be sent to. It must be a array contains valid email addresses. (Array of Strings)
     *  `frequency` (String) - (Optional) This attribute is optional, but when the attribute scheduled is true, it must be a cron expression string that starts with day of month field. For example, daily cron expression would be * * * 
     *  `generate_report_type` (String) - (Optional) Valid values are COMPLIANCE-STANDARD and GENERIC. By default, GENERIC type reports are produced. If COMPLIANCE-STANDARD is specified, filter.reportComplianceStandardId has to be:
      a.) a valid standard or framework id string (the report is generated with this compliance standard's)
      b.) one of the following supported standards: ["AWAF" | "NIST4" | "CISAWSF" | "CISAZUREF" | "SOC2" | "NIST-CSF" | "ISO27001" | "AGISM" | "HIPAA" | "HITRUST" | " "PCI" | "ASAE-3150" | "APRA" | "MAS" ]
        Enum: "GENERIC" "COMPLIANCE-STANDARD"
     * `include_checks` (Bool) - (Optional) Default: `false` Specifies whether or not to include individual checks in PDF reports when total number of checks is below 10000.
     * `scheduled` (Bool) - (Optional) Default: `false` Means whether the report is scheduled. It must be a boolean value when it was provided.
     * `send_email` (Bool) - (Optional) Default: `false` When report was scheduled, aka the attribute scheduled is true, this is a toggle to send report to specific email addresses. It should be boolean when provided.
     * `should_email_include_csv` (Bool) - (Optional) Default: `false` Specifies whether or not to include CSV attachment in email.
     * `should_email_include_pdf` (Bool) - (Optional) Default: `true` Specifies whether or not to include PDF attachment in email.
     * `tz` (String) - (Optional) It's used as which timezone the report schedule is based on, when the attribute scheduled is true. If this attribute was provided, it must be string that is a valid value of timezone database name such as Australia/Sydney. Available timzezones https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.  

- `filter` - (Optional) List:
     * `categories` (Array of Strings) - (Optional) An array of category (Conformity category) strings from the following: [ security | cost-optimisation | operational-excellence | reliability | performance-efficiency ]. If none is specified, all categories are set
     * `compliance_standards` (Array of Strings) - (Optional) An array of supported standard or framework ids. Possible values: ["AWAF" | "CISAWSF" | "CISAZUREF" | "PCI" | "HIPAA" | "HITRUST" | "GDPR" | "APRA" | "NIST4" | "SOC2" | "NIST-CSF" | "ISO27001" | "AGISM" | "ASAE-3150"] | "MAS" | "FEDRAMP" ]. If none is specified, all standards are set
     * `filter_tags` (Array of Strings) - (Optional) An array of any assigned metadata tags, tag keys or tag values to your AWS resources. e.g filterTags ["dev"] will match resource with tag "environment::dev" in the filter
     * `message` (Bool) - (Optional) Filter by message. Will find messages that contain all words regardless of the order. e.g "new message" will find "message new" and "new message"
     * `newer_than_days` (Number) - (Optional) (The filter.olderThanDays and filter.newerThanDays range refers to days to go back from the report's generation date. It converts the number of days entered to the date when the check was created and assigned a status, or where the status changed from "Success" to "Failure" or from "Failure" to "Success". You can use this filter by entering values for the number of days you wish to view before filter[olderThanDays] and after filter[newerThanDays]. You must pass at least 2 days up to 1 day to see any checks for a specific time duration. To display checks from a particular day up to the report's generation date, pass the number of days in filter.newerThanDays and leave filter.olderThanDays blank. Number. e.g. 5.
     * `older_than_days` (Number) - (Optional) To display all checks for up to a particular day, pass a number of days to go back from the report's generation date in filter.olderThanDays and leave filter.newerThanDays blank. Number. e.g. 5.
     * `providers` (Number) - (Optional) Cloud providers. Possible values: ["aws" | "azure"]. 
     * `regions` (Array of Strings) - (Optional) An array of valid region strings. e.g. for AWS ["us-west-1", "us-west-2"], for Azure ["eastus", "westus"]
     * `report_compliance_standard_id` (String) - (Optional) A single standard or framework id string. Possible values: ["AWAF" | "NIST4" | "CISAWSF" | "CISAZUREF" | "PCI" | "SOC2" | "NIST-CSF" | "ISO27001" | "AGISM" | "HIPAA" | "ASAE-3150" | "APRA" | "MAS" | "FEDRAMP" ]
     * `resource` (String) - (Optional) Filter by resource Id for an exact match, e.g "johnSmith", a wildcard, e.g "joh?Smh" or when used with filter[resourceSearchMode]=regex, a regular expression, e.g "joh.?Sm.h".
     * `resource_search_mode` (String) - (Optional) Set the search mode for the resource filter. Valid values are "text" or "regex". Text supports an exact match or the wildcard characters * and ? Defaults to "text"
     * `resource_types` (Array of Strings) - (Optional) An array of resource types. e.g. for AWS ["kms-key", "ec2-instance"], for Azure ["active-directory-users"]
     * `risk_levels` (String) - (Optional) Risk level. Possible values: ["EXTREME" | "VERY_HIGH" | "HIGH" | "MEDIUM" | "LOW"]
     * `rule_ids` (Array of Strings) - (Optional) An array of rule ids. e.g. ["EC2-001", "S3-001"]
     * `services` (Array of Strings) - (Optional) An array of service strings. e.g.     "EC2" "ELB" "EBS" "VPC" "S3" "CloudTrail" "Route53" "RDS" "IAM" "RTM" "KMS" "SNS" "SQS" "CloudFormation"  "Config" "CloudFront" "AutoScaling" "Redshift" "CloudWatch" "CloudWatchEvents" "CloudWatchLogs" "ResourceGroup" "SES" "DynamoDB" "ElastiCache" "Elasticsearch" "WorkSpaces" "ACM" "Budgets" "Inspector" "TrustedAdvisor" "Shield" "EMR" "WAF" "Lambda" "Support" "Kinesis" "Organizations" "EFS" "ElasticBeanstalk" "Macie" "ELBv2" "CloudConformity" "APIGateway" "GuardDuty" "Health" "ConfigService" "MQ" "Firehose" "SSM" "Route53Domains" "SageMaker" "DAX" "Neptune" "ECR" "Glue" "XRay" "SecretsManager" "DocumentDB" "DMS" "Miscellaneous" "EKS" "Backup" "StorageGateway" "ECS" "SecurityHub" "Comprehend" "WellArchitected" "AccessAnalyzer" "StorageAccounts" "SecurityCenter" "ActiveDirectory" "MySQL" "PostgreSQL" "Sql" "Monitor" "AppService" "Network" "ActivityLog" "VirtualMachines" "AKS" "KeyVault" "Locks" "AccessControl" "Advisor" "RecoveryServices" "Resources" "Subscriptions" "CosmosDB" "RedisCache" "Search" "AppInsights"
     * `statuses` (Array of Strings) - (Optional) The status of the check. Valid values: ["SUCCESS" | "FAILURE"]
     * `suppressed` (Bool) - (Optional) Show Suppressed rules. Will default to true for "v1", and omitted for "v2". Valid values: [true |false]
     * `suppressed_filter_mode` (String) - (Optional) Whether to use the "v1" or "v2" suppressed functionality. "v1": Using suppressed=true will return both suppressed and unsuppressed checks, suppressed=false will just return unsuppressed checks. "v2": Using suppressed=true return will just return suppressed checks, suppressed=false will just return unsuppressed checks, and omitting the filter will return both. Defaults to "v1". Valid values: [ "v1" | "v2" ]
     * `tags` (Array of Strings) - (Optional) An array of any assigned metadata tags to your resources.
     * `text` (String) - (Optional) Filter by resource Id, rule title or message. A string. e.g "john", "s3" or "write"
     * `with_checks` (Bool) - (Optional) Displays only controls from PDF reports with one or more associated checks. If withoutChecks is also set to true, then filter has no effect and all checks will be displayed. The default value is false. Valid values: [true |false]
     * `without_checks` (String) - (Optional) Displays only controls from PDF reports with 0 associated checks. If withChecks is also set to true, then filter has no effect and all checks will be displayed. The default value is false. Valid values: [true |false]
      * `risk_levels` (String) - (Optional) Risk level. Possible values: ["EXTREME" | "VERY_HIGH" | "HIGH" | "MEDIUM" | "LOW"] 
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

 - `id` - The ID of the Report Configuration in Conformity managed by this resource

Example usage on the template:

```hcl
report {
    id = conformity_report_config.report.id
}
```

## Import
Report Configs - Can import based on the report-config-id.

To get the report-config-id:
1. Open Conformity Admin Web console > navigate to Account dashboard > choose the AWS account where your report is > Reports > Click View/Edit of the report you want to import
2. Notice the URL, you should be able to get the report-config-id
e.g. https://cloudone.trendmicro.com/conformity/account/account-ID/reports/report?report-config-id=COPY-THIS-LINE-report-config-id
3. Make sure that the report-config-id is not url encoded.
   To decode, visit https://www.urldecoder.org/ and paste your report-config-id
4. Copy the decoded code and paste it to your `terraform import`.

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_user resources into Terraform. Replace the {report-config-id} with the decoded value.
```hcl
terraform import conformity_report_config.report {report-config-id}
```

Use the `state` subcommand to verify Terraform successfully imported the conformity_report_config resources.
```hcl
terraform state show conformity_report_config.report
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Make sure you remove the id from the imported resource.
```hcl
terraform show -no-color >> main.tf
```
