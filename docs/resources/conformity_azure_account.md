---
page_title: "conformity_azure_account Resource - cloudconformity_terraform"
subcategory: "Azure"
description: |-
  Provides import resources for cloud conformity azure account.
---

# Resource `conformity_azure_account`
Provides import resources for cloud conformity azure account.

## Example Usage
```terraform
resource "conformity_azure_account" "azure" {}
```

## Import
Azure Account - Can import based on the `Account ID` that can be found under the Conformity web console.

To get the Azure Account ID:
Open Conformity Admin Web console > navigate to Account dashboard > choose the Azure account you would like to import.
Notice the URL, you should be able to get the account ID e.g. https://cloudone.trendmicro.com/conformity/account/account-ID

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_azure_account resources into Terraform. Replace the {CLOUDCONFORMITY-ACCOUNT-ID-AZURE} value.
```hcl
terraform import conformity_azure_account.azure {CLOUDCONFORMITY-ACCOUNT-ID-AZURE}
```

Use the `state` subcommand to verify Terraform successfully imported the conformity_azure_account resources.
```hcl
terraform state show conformity_azure_account.azure
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Be sure to remove the id from the resource
```hcl
terraform show -no-color >> main.tf
```
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_conformity"></a> [conformity](#requirement\_conformity) | 0.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_conformity"></a> [conformity](#provider\_conformity) | 0.1.0 |

## Resources

| Name | Type |
|------|------|

| conformity_azure_account.azure | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_apikey"></a> [apikey](#input\_apikey) | n/a | `string` | `""` | no |
| <a name="input_azure_active_directory_id"></a> [azure\_active\_directory\_id](#input\_azure\_active\_directory\_id) | n/a | `string` | `""` | no |
| <a name="input_azure_environment"></a> [azure\_environment](#input\_azure\_environment) | variable "azure\_name"{ type    = string default = "trendmicro\_azure" } | `string` | `"staging"` | no |
| <a name="input_region"></a> [region](#input\_region) | n/a | `string` | `""` | no |
