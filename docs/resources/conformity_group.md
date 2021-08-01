---
page_title: "conformity_aws_account Resource - cloudconformity_terraform"
subcategory: "Groups"
description: |-
  Provides a Cloud Conformity Group Management.
---

# Resource `conformity_group`
Provides a Cloud Conformity Group Management.

## Example Usage
```hcl
resource "conformity_group" "group" {

    name        = "cloudconformity-group"
    tags        = ["tag1", "tag2"]
}
```

## Argument reference
 - `name` (String) - (Required) The name of your account.
 - `tags` (String) - (Required) The tag name that you want to add.

## Import
Conformity group - Can import based on the Group ID that can be found under the Conformity web console.
To get the Group ID same with the AWS Account ID above, you just need to navigate to the Group you want to import and get the
Group ID in the URL e.g. https://cloudone.trendmicro.com/conformity/group/group-ID

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_group resources into Terraform. Replace the {GROUP-ID} value.
```hcl
terraform import conformity_group.group_1 {GROUP-ID}
```

Use the `state` subcommand to verify Terraform successfully imported the conformity_group resources.
```hcl
terraform state show conformity_group.group_1
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Make sure you remove the id from the imported resource.
```hcl
terraform show -no-color >> main.tf
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_conformity"></a> [conformity](#requirement\_conformity) | 0.3.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_conformity"></a> [conformity](#provider\_conformity) | 0.3.1 |

## Resources

| Name | Type |
|------|------|

| conformity_group.group_1 | resource |

| conformity_group.group_2 | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_apikey"></a> [apikey](#input\_apikey) | n/a | `string` | `""` | yes |
| <a name="input_name_group1"></a> [name\_group1](#input\_name\_group1) | n/a | `string` | `"group1"` | yes |
| <a name="input_name_group2"></a> [name\_group2](#input\_name\_group2) | n/a | `string` | `"group2"` | yes |
| <a name="input_region"></a> [region](#input\_region) | n/a | `string` | `"us-west-2"` | yes |
| <a name="input_tag_group1"></a> [tag\_group1](#input\_tag\_group1) | n/a | `string` | `"group1_tag1"` | yes |
| <a name="input_tag_group2"></a> [tag\_group2](#input\_tag\_group2) | n/a | `string` | `"group2_tag2"` | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_group_name_1"></a> [group\_name\_1](#output\_group\_name\_1) | n/a |
| <a name="output_group_name_2"></a> [group\_name\_2](#output\_group\_name\_2) | n/a |
| <a name="output_group_tag_1"></a> [group\_tag\_1](#output\_group\_tag\_1) | n/a |
| <a name="output_group_tag_2"></a> [group\_tag\_2](#output\_group\_tag\_2) | n/a |