---
page_title: "Create Communication Settings Guide"
subcategory: "Communication Settings"
description: |-
  Provides instruction on how to create Communication Settings on Conformity using Terraform.
---

# How To Create Communication Settings on Conformity on a local machine
Provides instruction on how to create Communication Settings on Conformity using Terraform.

## Things needed:
1. Conformity API Key

#### Step 1

##### Terraform Configuration

1. To use Conformity and its resources, make sure that the values for certain variables are declared, Region, Account ID and API Key from Conformity Account are properly configured on the `terraform.tfvars`.

2. Create `terraform.tfvars` on `example/communication_setting/PATH-TO-CHANNEL` folder.
   
3. Note: Change the `PATH-TO-CHANNEL` value according to the channel you want to create (e.g. email, ms-teams, slack multiple, sms, sns).
Example Path: `example/communication_setting/email`

Add the following values on `terraform.tfvars`:
```
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "AWS-ACCOUNT-REGION"
```
Note: You can always change the values declared according to your choice.

4. Add filter or configuration values according to your requirements `main.tf` under `/path/terraform-provider-conformity/example/communication_setting/PATH-TO-CHANNEL` directory.
Note: Change the `PATH-TO-CHANNEL` value according to the channel you want to create (e.g. email, ms-teams, slack multiple, sms, sns).
Example Path: `/path/terraform-provider-conformity/example/communication_setting/email`

Example template to guide your configuration:

```hcl
resource "conformity_communication_setting" "comm_setting" {
  
  email {
    // required | type: array of user
    //Array of users with at least readOnly access to the account.
    users = []
  }
  
  ms_teams {
    // optional | type: string
    // Channel name
    channel = ""

    // optional | type: string
    // Label to display in the application (to distinguish between multiple instances of the same channel type).
    channel_name = ""
 
    // optional | type: bool
    // True for adding associated extra data to message.
    display_extra_data = bool

    // optional | type: bool
    // True for adding user to message.
    display_introduced_by = bool

    // optional | type: bool
    // True for adding resource to message.
    display_resource = bool

    // optional | type: bool
    // True for adding associated tags to message.
    display_tags = bool

    // required | type: string
    // Webhook MS teams url.
    url = ""
  }

  pager_duty {
    // optional | type: string
    // Label to display in the application (to distinguish between multiple instances of the same channel type).
    channel_name = ""
 
    // required | type: string
    // Page-duty | Your service key.
    service_key = ""

    // required | type: string
    // Page-duty | Your service name.
    service_name = ""
  }

  slack {
    // required | type: string
    // Channel name
    channel = ""

    // optional | type: string
    // Label to display in the application (to distinguish between multiple instances of the same channel type).
    channel_name = ""
 
    // optional | type: bool
    // True for adding associated extra data to message.
    display_extra_data = bool

    // optional | type: bool
    // True for adding user to message.
    display_introduced_by = bool

    // optional | type: bool
    // True for adding resource to message.
    display_resource = bool

    // optional | type: bool
    // True for adding associated tags to message.
    display_tags = bool

    // required | type: string
    // Webhook slack url.
    url = ""
  }

  sms {
    // required | type: array of user
    //Array of users with at least readOnly access to the account.
    users = []
  }

  sns {
    // optional | type: string
    // Label to display in the application (to distinguish between multiple instances of the same channel type).
    channel_name = ""

    // required | type: string
    // Amazon Resource Name of SNS
    arn = ""
  }

  webhook {
    // optional | type: string
    // webhook security token
    security_token = ""

    // required | type: string
    // Webhook slack url.
    url = ""
  }


  // required | type: string
  // value can be: "email", "sms", "ms-teams", "slack", "sns", "pager-duty", "webhook"
  channel = ""

  // required | type: bool | default: true
  enabled = bool

  // optional | type string
  //default : settings 
  type = ""

  // Optional 
  // max number 1
  configuration {

    // optional | type: string
    // Channel name
    channel = ""

    // optional | type: string
    // Label to display in the application (to distinguish between multiple instances of the same channel type).
    channel_name = ""
    
    // optional | type: array of user
    //Array of users with at least readOnly access to the account.
    users = []

    // optional | type: bool
    // True for adding associated extra data to message.
    display_extra_data = bool

    // optional | type: bool
    // True for adding user to message.
    display_introduced_by = bool

    // optional | type: bool
    // True for adding resource to message.
    display_resource = bool

    // optional | type: bool
    // True for adding associated tags to message.
    display_tags = bool

    // optional | type: string
    // Webhook url.
    url = ""

    // optional | type: string
    // Amazon Resource Name of SNS
    arn = ""

    // optional | type: string
    // Page-duty | Your service key.
    service_key = ""

    // optional | type: string
    // Page-duty | Your service name.
    service_name = ""

    // optional | type: string
    // webhook security token
    security_token = ""

  }

  // Optional 
  // max number 1
  filter {
    //optional | type: array of string
    // value can be: "security" "cost-optimisation" "reliability" "performance-efficiency" "operational-excellence"
    categories = []

    // optional | type: array of string
    // value can be: "AWAF" "CISAWSF" "CISAZUREF" "CISAWSTTW" "PCI" "HIPAA" "GDPR" "APRA" "NIST4" "SOC2" "NIST-CSF" 
    // "ISO27001" "AGISM" "ASAE-3150" "MAS" "FEDRAMP"
    compliances = []

    // optional | type: array of string
    // (only used for SNS and webhook channels) An array of statuses strings from the following: "SUCCESS" | "FAILURE"
    statuses = []

    // optional | type: array of string
    filter_tags = []


  // optional | type: string
 // "global" "us-east-1" "us-east-2" "us-west-1" "us-west-2" "ap-south-1" "ap-northeast-2" "ap-southeast-1"  "ap-southeast-2"
 // "ap-northeast-1" "eu-central-1" "eu-west-1" "eu-west-2" "eu-west-3" "eu-north-1" "sa-east-1" "ca-central-1"
 // "me-south-1" "ap-east-1"
 // for azure:
 // "global" "eastasia" "southeastasia" "centralus" "eastus" "eastus2" "westus" "northcentralus" "southcentralus"
 // "northeurope" "westeurope" "japanwest" "japaneast" "brazilsouth" "australiaeast" "australiasoutheast" "southindia" "centralindia"
 // "westindia" "canadacentral" "canadaeast" "uksouth" "ukwest" "westcentralus" "westus2" "koreacentral" "koreasouth" "francecentral"
 // "francesouth" "australiacentral" "australiacentral2" "southafricanorth" "southafricawest"
    regions = []

    // optional | type: array of string
    // value can be: "LOW" "MEDIUM" "HIGH" "VERY_HIGH" "EXTREME"
    risk_levels = []

    // optional | type: array of string
    // An array of valid AWS rule Ids (e.g. ["S3-016", "EC2-001"]). For more information about rules
    rule_ids = []

    // optional | type: array of string
    // value can be: "EC2" "ELB" "EBS" "VPC" "S3" "CloudTrail" "Route53" "RDS" "IAM" "RTM" "KMS" "SNS" "SQS" "CloudFormation" "Config" "CloudFront" "AutoScaling" 
    # "Redshift" "CloudWatch" "CloudWatchEvents" "CloudWatchLogs" "ResourceGroup" "SES" "DynamoDB" "ElastiCache" "Elasticsearch" "WorkSpaces" "ACM" "Budgets" "Inspector" 
    # "TrustedAdvisor" "Shield" "EMR" "WAF" "Lambda" "Support" "Kinesis" "Organizations" "EFS" "ElasticBeanstalk" "Macie" "ELBv2" "CloudConformity" "APIGateway" 
    # "GuardDuty" "Health" "ConfigService" "MQ" "Firehose" "SSM" "Route53Domains" "SageMaker" "DAX" "Neptune" "ECR" "Glue" "XRay" "SecretsManager" "DocumentDB" "DMS" 
    # "Miscellaneous" "EKS" "Backup" "StorageGateway" "ECS" "SecurityHub" "Comprehend" "WellArchitected" "AccessAnalyzer" "StorageAccounts" "SecurityCenter" "ActiveDirectory" 
    # "MySQL" "PostgreSQL" "Sql" "Monitor" "AppService" "Network" "ActivityLog" "VirtualMachines" "AKS" "KeyVault" "Locks" "AccessControl" "Advisor" "RecoveryServices" 
    # "Resources" "Subscriptions" "CosmosDB" "RedisCache" "Search" "AppInsights"
    services = []
    // optional | type: array of string
    // An array of any assigned metadata tags to your AWS resources.
    tags = []
  }

  relationships {
    // optional
    account {
      // required if account is define | type: string
      id = ""
    }
    organization {
      // required if account is define | type: string
      id = ""
    }
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
Note: Change the `PATH-TO-CHANNEL` value according to the channel you want to create (e.g. email, ms-teams, slack multiple, sms, sns).
```sh
cd example/communication_setting/PATH-TO-CHANNEL
terraform init
terraform plan
terraform apply
```

Example:
```sh
cd example/communication_setting/email
terraform init
terraform plan
terraform apply
```

#### 5. Bash script that can run to automate the whole process 1-5:
Note: Change the `PATH-TO-CHANNEL` value according to the channel you want to create (e.g. email, ms-teams, slack multiple, sms, sns).

Under script folder run:
```sh
cd script/communication_setting/PATH-TO-CHANNEL
sh terraform-PATH-TO-CHANNEL-create.sh
```

Example:
```sh
cd script/communication_setting/email
sh terraform-email-create.sh
```
