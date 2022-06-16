---
page_title: "conformity_aws_account Resource"
subcategory: "AWS"
description: |-
  Provides a Conformity Account.
---

# Resource `conformity_aws_account`
Provides a Conformity AWS Account.

## Example Usage With AWS Conformity To Create Account Only
```hcl
data "conformity_external_id" "external"{}

resource "conformity_aws_account" "aws" {
    name        = "aws-conformity"
    environment = "development"
    role_arn    = "arn:aws:iam::223334445555:role/CloudConformity"
    external_id = data.conformity_external_id.external.external_id
    settings {
        // implement value
        rule {
            rule_id = "RDS-018"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                exceptions {
                    tags        = [
                        "mysql-backups",
                    ]
                }
                extra_settings {
                    name    = "threshold"
                    type    = "single-number-value"
                    value   = "90"
                }
            }
        }
        // implement multiple values
        rule {
            rule_id = "SNS-002"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                exceptions {
                    tags        = [
                        "some_tag",
                    ]
                }
                extra_settings {
                    name    = "conformityOrganization"
                    type    = "choice-multiple-value"
                    values {
                        enabled = false
                        label   = "All within this Conformity organization"
                        value   = "includeConformityOrganization"
                    }
                    values {
                        enabled = true
                        label   = "All within this AWS Organization"
                        value   = "includeAwsOrganizationAccounts"
                    }
                }
            }
        }
        // implement regions
        rule {
            rule_id = "RTM-008"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                extra_settings {
                    name    = "authorisedRegions"
                    regions = [
                        "ap-southeast-2",
                        "eu-west-1",
                        "us-east-1",
                        "us-west-2",
                    ]
                    type    = "regions"
                }
            }
        }
        // implement multiple_object_values
        rule {
            rule_id = "RTM-011"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                extra_settings {
                    name    = "patterns"
                    type    = "multiple-object-values"
                    multiple_object_values {
                        event_name         = "^(iam.amazonaws.com)"
                        event_source       = "^(IAM).*"
                        user_identity_type = "^(Delete).*"
                    }
                }
            }
        }
        // implement mappings
        rule {
            rule_id = "VPC-013"
            settings {
                enabled     = true
                risk_level  = "LOW"
                rule_exists = false
                extra_settings {
                    name    = "SpecificVPCToSpecificGatewayMapping"
                    type    = "multiple-vpc-gateway-mappings"
                    // can be multiple mappings
                    mappings {
                        // can be multilple value
                        // if mappings is declared, values is required
                        values {
                            // name is required
                            // type is required
                            name = "gatewayIds"
                            type = "multiple-string-values"
                            // can be one of this value/values
                            values {
                                // value is required
                                // validation value should start with nat-
                                value = "nat-001"
                            }
                            values {
                                value = "nat-002"
                            }
                        }
                        values {
                            name  = "vpcId"
                            type  = "single-string-value"
                             // can be one of this value/values
                             // validation value should start with vpc-
                            value = "vpc-001"
                        }
                    }
                }
            }
        }
    }
}
```

## Argument reference
 - `name` (String) - (Required) The name of your account.
 - `environment` (String) - (Required) The environment for your account.
 - `external_id` (String) - (Required) The external ID for your Conformity organisation. Can be referenced from `conformity_external_id` data source.
 - `role_arn` (String) - (Required) The ARN of the role your account can assume. Can be referenced from `aws_cloudformation_stack`.
 - `tags` (Array of Strings) - (Optional) Tags for account.
 - `settings` - (Optional) List: (Can be multiple declaration)
  
  Inside `settings` there can be a `bot` set.

 - `bot` - (Optional) List: (Can be multiple declaration)
     * `disabled` (Bool) - (Optional) True to disable or false to enable the Conformity Bot.
     * `disabled_regions` (Array of Strings) - (Optional) - Possible values are "af-south-1", "ap-east-1", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ca-central-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-west-1", "us-west-2". This field can only be applied to AWS accounts. An attribute object containing a list of AWS regions for which Conformity Bot runs will be disabled.
     * `delay` (Int) - (Optional) Sets the number of hours delay between Conformity Bot runs.

  Inside `settings` there can be multiple `rule` set.

 - `rule` - (Optional) List: (Can be multiple declaration)
     * `rule_id` (String) - (Optional) - Rule ID, same as the one provided in the endpoint.
  
  Inside `settings` under `rule` set, there can be one `settings` set. 

 - `settings` - (Optional)
     * `enabled` (Bool) - (Optional) True for inclusion in bot detection, false for exclusion.
     * `rule_exists` (Bool) - (Optional)  True if rule exists.
     * `risk_level` (String) - (Optional) - Risk level of the Conformity rule. Enum: "LOW" "MEDIUM" "HIGH" "VERY_HIGH" "EXTREME"

  Inside `settings` under `rule` set, there can be one `exceptions` set. 
 
 - `exceptions` - (Optional) List: 
     * `filter_tags` (Array of Strings)- (Optional) An array of resource tags, resource tag keys or resource tag values that are exempted from the rule when it runs, e.g filterTags ["dev"] will exempt resource with tag "environment::dev from the rule".
     * `resources` (Array of Strings) - (Optional) An array of resource IDs that are exempted from the rule when it runs.
     * `tags` (Array of Strings) - (Optional) An array of resource tags that are exempted from the rule when it runs.

   Inside `settings` under `rule` set, there can be multiple `extra_settings` set. 

 - `extra_settings` - (Optional) List: (Can be multiple declaration)
     * `name` (String) - (Optional) (Keyword) Name of the extra setting.
     * `type` (String) - (Required) Rule specific property. Values can be: "multiple-string-values", "multiple-number-values" "multiple-aws-account-values", "choice-multiple-value" "choice-single-value", "single-number-value", "single-string-value", "ttl", "single-value-regex", "countries", "multiple-ip-values", and "tags".
     * `value` (String) - (Optional) Customisable value for rules that take on single name/value pairs.
     * `regions` (Array of Strings) - (Optional) Rule specific property.
     * `multiple-object-values` (Array of Strings) - (Optional) Rule specific property.

  Inside `extra_settings` under `settings` of `rule` set, there can be multiple declaration of `multiple-object-values` set.
  
 - `multiple-object-values` - (Optional) List: (Can be multiple declaration). 
     *  `event_name` (String) - (Optional) Name of the event.
     *  `event_source` (String) - (Optional) Name of the event source
     *  `user_identity_type` (String) - (Required) Type of the Identity of the user.

  Inside `extra_settings` under `settings` of `rule` set, there can be multiple declaration of `mappings` set. And under `mappings` set, here can be multiple declaration of `values` set.
  
 - `values` - (Required) List: (Can be multiple declaration). An array (sometimes of objects) rules that take on a set of of values
     * `name` (String) - (Optional) (Keyword) Name of the values.
     * `type` (String) - (Required) Rule specific property. Values can be: "multiple-string-values", "multiple-number-values" "multiple-aws-account-values", "choice-multiple-value" "choice-single-value", "single-number-value", "single-string-value", "ttl", "single-value-regex", "countries", "multiple-ip-values", and "tags".
     *  `value` (String) - (Required) Description of the checkbox.
    Note: `values` is required when you use `mappings`.

  Inside `values`, there can be multiple declaration of `values` set.
  
 - `values` - (Required) List: (Can be multiple declaration).
     *  `value` (String) - (Required) Description of the checkbox.
    Note: If inside the `values` under the `mappings` has set `values` declared, you cannot use `value` anymore. Inside mappings, its either `values` with `values` set inside it or `values` with declared `value` inside it.

        Note: There is a condition for `type` attribute. If the specified is attribute is `value`, the possible values are "single-number-value", "single-string-value", "single-value-regex" and "ttl". If the specified is attribute is `values`, the declaration of it is inside the extra settings which can be a list and the possible values are "choice-multiple-value", "choice-single-value", "multiple-string-values", "multiple-number-values", "countries", "multiple-ip-values", "multiple-aws-account-values" and "tags". You cannot declare both `values` and `value` at the same time.See the table below:

| type     | possible value                                                                                                                | Sample declaration                                                                                                                                                                                                                    |
|----------|-------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `value`  | single-number-value, single-string-value, single-value-regex, ttl                                                             | included {     ….     exceptions {       ….     }       extra_settings {		 	….                     type = "ttl"         value = "72"       } }                                                                                           |
| `values` | choice-multiple-value, choice-single-value, multiple-string-values, multiple-number-values, countries, multiple-ip-values, multiple-aws-account-values, tags | included {     ….     exceptions {       ….     }       extra_settings {		 	….                     type = "choice-multiple-value"           values {             ….   

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

 - `id` - The ID of the AWS account in Conformity managed by this resource

Example usage on the template:

```hcl
account {
    id = conformity_aws_account.aws.id
}
```

## Import
AWS Account - Can import based on the `Account ID` that can be found under the Conformity web console.

To get the AWS Account ID:
Open Conformity Admin Web console > navigate to Account dashboard > choose the AWS account you would like to import.
Notice the URL, you should be able to get the account ID e.g. https://cloudone.trendmicro.com/conformity/account/account-ID

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_aws_account resources into Terraform. Replace the {CLOUDCONFORMITY-ACCOUNT-ID-AWS} value.
```hcl
terraform import conformity_aws_account.aws {CLOUDCONFORMITY-ACCOUNT-ID-AWS}
```

Use the `state` subcommand to verify Terraform successfully imported the conformity_aws_account resources.
```hcl
terraform state show conformity_aws_account.aws
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Be sure to remove the id from the resource
```hcl
terraform show -no-color >> main.tf
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 3.44.0 |


## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 3.44.0 |


## Resources

| Name | Type |
|------|------|
| [aws_cloudformation_stack.cloud-conformity](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack) | resource |
| conformity_aws_account.aws | resource |
| conformity_external_id.external | data source |