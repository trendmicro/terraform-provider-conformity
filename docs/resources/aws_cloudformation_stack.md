---
page_title: "aws_cloudformation_stack Resource"
subcategory: "AWS"
description: |-
  Provides a CloudFormation Stack resource from AWS provider.
---

# Resource `aws_cloudformation_stack`
Provides a CloudFormation Stack resource from AWS provider.

## Example Usage
```hcl
data "conformity_external_id" "all"{}

resource "aws_cloudformation_stack" "cloud-conformity" {
  name         = "CloudConformity"
  template_url = "https://s3-us-west-2.amazonaws.com/cloudconformity/CloudConformity.template"
  parameters = {
    AccountId  = "717210094962"
    ExternalId = data.conformity_external_id.all.external_id
  }
  capabilities = ["CAPABILITY_NAMED_IAM"]
}
```

## Argument reference

 - `name` (String) - (Required) The name of your CloudFormation Stack (Do not change the value).
 - `template_url` (String) - (Required) Default CloudFormation template (Do not change the value).
 - `AccountId` (String) - (Required) Default Conformity AWS Account ID (Do not change the value).

## Attributes Reference

In addition to all the arguments above, the following attributes are imported from `cloudconformity_external_id` resource.

 - `ExternalId` (String) - (Required) The external ID. Imported from `cloudconformity_external_id`.

## Import
Cloudformation Stacks can be imported using the `name`. e.g.

```hcl
terraform import aws_cloudformation_stack.stack CloudConformity
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