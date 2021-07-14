---
page_title: "conformity_aws_account Resource - cloudconformity_terraform"
subcategory: "AWS"
description: |-
  Provides a Cloud Conformity Account.
---

# Resource `conformity_aws_account`
Provides a Cloud Conformity Account.

## Example Usage
```terraform
resource "conformity_aws_account" "aws"{

    name        = "Trendmicro Account"
    environment = "Staging"
    external_id = "..."
    role_arn    = "..."
}
```

## Argument reference
 - `name` (String) - (Required) The name of your account.
 - `environment` (String) - (Required) The environment for your account.
 - `external_id` (String) - (Required) The external ID for your Cloud Conformity organisation.


## Attributes Reference

In addition to all the arguments above, the following attributes are imported from `cloudconformity_external_id` and `aws_cloudformation_stack` resources.

 - `ExternalId` (String) - (Required) The external ID. Imported from `cloudconformity_external_id`.
 - `role_arn` (String) - (Required) The ARN of the role your account can assume. Imported from `aws_cloudformation_stack`.
  
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
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 2.7.0 |
| <a name="requirement_conformity"></a> [conformity](#requirement\_conformity) | 0.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 2.7.0 |
| <a name="provider_conformity"></a> [conformity](#provider\_conformity) | 0.1.0 |

## Resources

| Name | Type |
|------|------|
| [aws_cloudformation_stack.cloud-conformity](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack) | resource |

| conformity_aws_account.aws | resource |

| conformity_external_id.external | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_access_key"></a> [access\_key](#input\_access\_key) | n/a | `string` | `""` | yes |
| <a name="input_apikey"></a> [apikey](#input\_apikey) | n/a | `string` | `""` | yes |
| <a name="input_environment"></a> [environment](#input\_environment) | n/a | `string` | `"Staging"` | yes |
| <a name="input_name"></a> [name](#input\_name) | n/a | `string` | `"Cloudconformity"` | yes |
| <a name="input_region"></a> [region](#input\_region) | n/a | `string` | `"us-west-2"` | yes |
| <a name="input_secret_key"></a> [secret\_key](#input\_secret\_key) | n/a | `string` | `""` | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_aws_account_name"></a> [aws\_account\_name](#output\_aws\_account\_name) | n/a |
| <a name="output_aws_environment"></a> [aws\_environment](#output\_aws\_environment) | n/a |
| <a name="output_aws_role_arn"></a> [aws\_role\_arn](#output\_aws\_role\_arn) | n/a |
| <a name="output_external_id"></a> [external\_id](#output\_external\_id) | n/a |
| <a name="output_role_arn"></a> [role\_arn](#output\_role\_arn) | n/a |