---
page_title: "Create Report Configs Guide - cloudconformity_terraform"
subcategory: "Report Configs"
description: |-
  Provides instruction on how to create Report Configs on Cloud Conformity using Terraform.
---

# How To Create Report Configs on Cloud Conformity on a local machine
Provides instruction on how to create Report Configs on Cloud Conformity using Terraform.

## Things needed:
1. Cloud Conformity API Key

#### Step 1

##### Terraform Configuration

1. To use Cloud Conformity and its resources, make sure that the values for certain variables are declared, Region, Account ID and API Key from Cloud Conformity Account are properly configured on the `terraform.tfvars`.

2. Create `terraform.tfvars` on `example/report_config` folder and add the following values:

```
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "AWS-ACCOUNT-REGION"

// if you want to create account-level report config, uncomment account_id value and provide account ID on the terraform.tfvars
// if you want to create group-level report config, uncomment group_id value and provide account ID on the terraform.tfvars
// if you want to create organisation-level report config, leave account_id and group_id commented.

// account_id = "CLOUD-CONFORMITY-ACCOUNT-ID" 
// group_id = "CLOUD-CONFORMITY-GROUP-ID"
```
Note: You can always change the values declared according to your choice.

3. Add filter or configuration values according to your requirements `main.tf` under `/path/terraform-provider-conformity/example/report_config/main` directory.

Example template to guide your configuration:

```
resource "conformity_report_config" "report" {
  // required | type: string
  account_id = "" 
  // optional | type: string
  group_id   = ""   
  configuration { 
    // optional | type: array of strings
    email = [] 
    // optional | type: string
    frequency = ""
    // optional | type: string | default: GENERIC
    // value can be: "GENERIC", "COMPLIANCE-STANDARD"
    generate_report_type = "" 
    // optional | type: bool | default: false
    include_checks = bool
    // optional | type: bool | default: false
    scheduled = bool
    //optional | type: bool | default: false
    send_email = bool
    // optional | type: bool | default: false
    should_email_include_csv = bool
    // optional | type: bool | default: true
    should_email_include_pdf = bool
    // required | type: string 
    title = ""
    // optional | type: string
    tz = ""
  }
  filter {
    // optional | type: array of strings 
    # value can be : "security", "cost-optimisation", "reliability", "performance-efficiency", "operational-excellence"
    # if none is specified, all categories are set
    categories = []
    // optional | type: array of strings
    # value can be: "AWAF" "CISAWSF", "CISAZUREF", "PCI", "HIPAA", "GDPR", "APRA", "NIST4", "SOC2", "NIST-CSF", "ISO27001", "AGISM", "ASAE-3150", "MAS", "FEDRAMP"
    # if none is specified, all standards are set
    compliance_standards = []
    // optional | type: array of strings 
    filter_tags = []
    // optional | type: bool
    message = bool
    // optional | type: number
    newer_than_days = 0
    // optional | type: number
    older_than_days = 0
    // optional | typel: array of strings
    // value can be: "aws", "azure"
    providers = []
    // optional | typel: array of strings
    // value can be: array of AWS region
    regions = []
    // optional | stype: string
    // value can be: "AWAF", "CISAWSF", "CISAZUREF", "PCI", "HIPAA", "GDPR", "APRA", "NIST4", "SOC2", "NIST-CSF", "ISO27001", "AGISM", "ASAE-3150", "MAS", "FEDRAMP"
    report_compliance_standard_id = ""
    // optional | type: string
    resource = ""
    // optional | type: string 
    # value can be a text or regex
    resource_search_mode = ""
    // optional | type: array of strings
    resource_types = []
    // optional | type: string
    # value can be: 
    # "LOW", "MEDIUM", "HIGH", "VERY_HIGH", "EXTREME"
    risk_levels = ""
    // optional | type: array of strings
    rule_ids = []
    // optional | type: array of strings
    # value can be: 
    # "EC2" "ELB" "EBS" "VPC" "S3" "CloudTrail" "Route53" "RDS" "IAM" "RTM" "KMS" "SNS" "SQS" "CloudFormation"  "Config" "CloudFront" "AutoScaling" 
    # "Redshift" "CloudWatch" "CloudWatchEvents" "CloudWatchLogs" "ResourceGroup" "SES" "DynamoDB" "ElastiCache" "Elasticsearch" "WorkSpaces" "ACM" "Budgets" 
    # "Inspector" "TrustedAdvisor" "Shield" "EMR" "WAF" "Lambda" "Support" "Kinesis" "Organizations" "EFS" "ElasticBeanstalk" "Macie" "ELBv2" "CloudConformity" 
    # "APIGateway" "GuardDuty" "Health" "ConfigService" "MQ" "Firehose" "SSM" "Route53Domains" "SageMaker" "DAX" "Neptune" "ECR" "Glue" "XRay" "SecretsManager" 
    # "DocumentDB" "DMS" "Miscellaneous" "EKS" "Backup" "StorageGateway" "ECS" "SecurityHub" "Comprehend" "WellArchitected" "AccessAnalyzer" "StorageAccounts" 
    # "SecurityCenter" "ActiveDirectory" "MySQL" "PostgreSQL" "Sql" "Monitor" "AppService" "Network" "ActivityLog" "VirtualMachines" "AKS" "KeyVault" 
    # "Locks" "AccessControl" "Advisor" "RecoveryServices" "Resources" "Subscriptions" "CosmosDB" "RedisCache" "Search" "AppInsights"
    services = []
    // optional | type: array of strings
    # value can be: "SUCCESS", "FAILURE"
    statuses = []
    // optional | type: bool
    suppressed = bool
    // optional | type: string | default: v1
    # value can be: "v1", "v2"
    suppressed_filter_mode = ""
    // optional | type: array of strings
    // value can be: metadata tags for aws resources
    tags = []
    //optional | type: string
    text = ""
    // optional | type: bool | default: false
    with_checks = bool
    // optional | type: bool | default: false
    without_checks = bool
  }
}
```

#### Step 2

##### Run Terraform

#### 1. Navigate to project directory:
```sh
cd /path/terraform-provider-conformity/
```
#### 2. Install dependencies:
```sh
go mod vendor
```
#### 3. Create the Artifact:
```sh
make install
```
#### 4. Now, you can run terraform code:
```sh
cd example/report_config/main
terraform init
terraform plan
terraform apply
```
#### 5. Bash script that can run to automate the whole process 1-5:

Under script folder run:
```sh
cd script/report_config
sh terraform-report_config-create.sh
```
