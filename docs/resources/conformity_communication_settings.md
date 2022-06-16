---
page_title: "conformity_communication_setting Resource"
subcategory: "Communication Setting"
description: |-
  Allows you to create Communication Setting on Conformity. 
---

# Resource `conformity_communication_setting`
Allows you to create Communication Setting on Conformity. 

## Example Usage
You can create multiple channel.

```hcl
resource "conformity_communication_setting" "email_setting" {

    email {
        users = [
        "urn:tmds:identity:us-east-ds-1:62740:administrator/1915",
        ]
    }

    filter {
        categories  = [
        "security",
        ]
        comp 7liances = [
        "FEDRAMP",
        ]
        filter_tags = [
        "tagKey",
        ]
        regions     = [
        "ap-southeast-1",
        ]
        risk_levels = [
        "MEDIUM",
        ]
        rule_ids    = [
        "S3-016",
        ]
        services    = [
        "EC2",
        "IAM",
        ]
        tags        = [
        "tagName",
        ]
    }
    relationships {
        account {
            id = "80b880c9-336a-490d-b212-4e847956a62d"
        }
        organisation {
            id = "102642678400"
        }
    }
}
output "email_setting" {
    value = conformity_communication_setting.email_setting
}

```

## Argument reference
- `email` - (Optional) Max number 1. List of possible values:
     * `users` (Array of String) - (Required) Array of users with at least readOnly access to the account.
- `ms_teams` - (Optional) Max number 1. List of possible values:
     * `channel` (String) - (Optional) Channel name of slack or ms-teams (For slack, ms-teams communication setting only).
     * `channel_name` (String) - (Optional) Label to display in the application (to distinguish between multiple instances of the same channel type).
     * `display_extra_data` (Bool) - (Optional) True for adding associated extra data to message.
     * `display_introduced_by` (Bool) - (Optional) True for adding user to message.
     * `display_resource` (Bool) - (Optional) True for adding resource to message.
     * `display_tags` (Bool) - (Optional) True for adding associated tags to message.
     * `url` (String) - (Required) Webhook MS teams url. To set up webhook url of teams, visit https://support.itglue.com/hc/en-us/articles/360004934997-Setting-up-Microsoft-Teams-webhook-notifications#:~:text=Configuring%20Microsoft%20Teams&text=From%20the%20More%20Options%20menu,Then%2C%20click%20Create.
- `pager_duty` - (Optional) Max number 1. List of possible values:
     * `channel_name` (String) - (Optional) Label to display in the application (to distinguish between multiple instances of the same channel type).
     * `service_key` (String) - (Required) Page-duty | Your service key.
     * `service_name` (String) - (Required) Page-duty | Your service name.
- `slack` - (Optional) Max number 1. List of possible values:
     * `channel` (String) - (Required) Channel name of slack or ms-teams (For slack, ms-teams communication setting only).
     * `channel_name` (String) - (Optional) Label to display in the application (to distinguish between multiple instances of the same channel type).
     * `display_extra_data` (Bool) - (Optional) True for adding associated extra data to message.
     * `display_introduced_by` (Bool) - (Optional) True for adding user to message.
     * `display_resource` (Bool) - (Optional) True for adding resource to message.
     * `display_tags` (Bool) - (Optional) True for adding associated tags to message.
     * `url` (String) - (Required) Webhook slack url.
- `sms` - (Optional) Max number 1. List of possible values:
     * `users` (Array of String) - (Required) Array of users with at least readOnly access to the account.
- `sns` - (Optional) Max number 1. List of possible values:
     * `channel_name` (String) - (Optional) Label to display in the application (to distinguish between multiple instances of the same channel type).
     * `arn` (String) - (Required) Amazon Resource Name of SNS.
- `webhook` - (Optional) Max number 1. List of possible values:
     * `security_token` (String) - (Optional) webhook security token.
     * `url` (String) - (Required) Webhook url.

Other arguments:
 - `filter` - (Optional) Max number 1. List:
     * `categories` (Array of String) - (Optional) An array of category (AWS well-architected framework category) strings from the following:
          security | cost-optimisation | reliability | performance-efficiency | operational-excellence
     * `compliances` (Array of String) - (Optional) An array of supported standard or framework ids. Possible values: ["AWAF" | "CISAWSF" | "CISAZUREF" | "CISAWSTTW" | "PCI" | "HIPAA" | "GDPR" | "APRA" | "NIST4" | "SOC2" | "NIST-CSF" | "ISO27001" | "AGISM" | "ASAE-3150" | "MAS" | "FEDRAMP"]
     * `filter_tags` (Array of String) - (Optional) An array of any assigned metadata tags, tag keys or tag values to your AWS resources. e.g filterTags ["dev"] will match resource with tag "environment::dev" in the filter.
     * `regions` (String) - (Optional) 	An array of valid region strings.
     For AWS:
          "global" "us-east-1" "us-east-2" "us-west-1" "us-west-2" "ap-south-1" "ap-northeast-2" "ap-southeast-1" "ap-southeast-2" "ap-northeast-1" "eu-central-1" "eu-west-1" "eu-west-2" "eu-west-3" "eu-north-1" "sa-east-1" "ca-central-1" "me-south-1" "ap-east-1"
     For Azure:
          "global" "eastasia" "southeastasia" "centralus" "eastus" "eastus2" "westus" "northcentralus" "southcentralus" "northeurope" "westeurope" "japanwest" "japaneast" "brazilsouth" "australiaeast" "australiasoutheast" "southindia" "centralindia" "westindia" "canadacentral" "canadaeast" "uksouth" "ukwest" "westcentralus" "westus2" "koreacentral" "koreasouth" "francecentral" "francesouth" "australiacentral" "australiacentral2" "southafricanorth" "southafricawest"
      * `risk_levels` (Array of String) - (Optional) An array of risk-level strings.
      * `rule_ids` (Array of String) - (Optional) An array of valid rule Ids (e.g. ["S3-016", "EC2-001"]). For more information about rules, please refer to https://cloudone.trendmicro.com/docs/conformity/api-reference/tag/Settings#tag/Services/paths/~1services/get.
      * `services` (Array of String) - (Optional) An array of AWS service strings from the following:
          AutoScaling | CloudConformity | CloudFormation | CloudFront | CloudTrail | CloudWatch | CloudWatchEvents | CloudWatchLogs | Config | DynamoDB | EBS | EC2 | ElastiCache | Elasticsearch | ELB | IAM | KMS | RDS | Redshift | ResourceGroup | Route53 | S3 | SES | SNS | SQS | VPC | WAF | ACM | Inspector | TrustedAdvisor | Shield | EMR | Lambda | Support | Organizations | Kinesis | EFS
      * `tags` (Array of String) - (Optional) An array of any assigned metadata tags to your resources.

- `relationships` - (Optional) List: 
    * `account` - (Optional): 
      * `id` (String) - (Optional) required if account is defined.
    * `organisation` - (Optional): 
      Note: If you did not define organisation, it will automatically create the communication settings on organisational level.
      * `id` (String) - (Optional) required if organisation is defined.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

 - `id` - The ID of the Communication Setting in Conformity managed by this resource

Example usage on the template:

```hcl
setting {
    id = conformity_communication_setting.email_setting.id
}
```

## Import
Comunication Settings - Can import based on the resource_id that can be found under the outputs upon creation or inspect element.

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_communication_setting resources into Terraform. Replace the {RESOURCE-NAME} and {RESOURCE-ID} value.
```hcl
terraform import conformity_communication_setting.{RESOURCE-NAME} {RESOURCE-ID}
```

Use the `state` subcommand to verify Terraform successfully imported the conformity_communication_setting resources.
```hcl
terraform state show conformity_communication_setting.{RESOURCE-NAME}
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Make sure you remove the id from the imported resource.
```hcl
terraform show -no-color >> main.tf
```
