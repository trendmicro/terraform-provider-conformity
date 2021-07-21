---
page_title: "conformity_sso_user Resource - cloudconformity_terraform"
subcategory: "User"
description: |-
  Allows you to add a SSO user to your organisation. This resource does not apply to users who are part of the Cloud One Platform.
---

# Resource `conformity_sso_user`
Allows you to add a SSO user to your organisation. This resource does not apply to users who are part of the Cloud One Platform.

## Example Usage
```hcl
resource "conformity_sso_user" "user" {

  first_name = var.first_name
  last_name  = var.last_name
  email      = var.email
  role       = var.role

  access_list {
    account01 = var.account01
    level01   = var.level01
    }

  access_list {
    account02 = var.account02
    level02   = var.level02
    }

  }

  output "user"{
    value = conformity_sso_user.user
  }
```

## Argument reference

 - `email` (String) - (Required) The email of the sso_user.
 - `first_name` (String) - (Required) The first name of the sso_user.
 - `last_name` (String) - (Required) The last name of the sso_user.
 - `role` (String) - (Required) The role which the sso_user is assigned to. Can be "ADMIN" "USER".
 - `access_list` - (Required) List:
      `account` (String) - (Required) The account id within the organisation.
      `level` (String) - (Required) The level of access the user has to the account. Can be "NONE" "READONLY" "FULL".

## Attributes Reference

In addition to all the arguments above, the following attributes are added to the terraform state and can be used for output.

 - `last_login` (Int) - User last login.
 - `last_name` (String) - User last name.
 - `mfa` (Bool) - Shows 'true' if the user has MFA, otherwise 'false'.
  
## Import
SSO User - Can import based on the user_id that can be found under the terraform state or user login.

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_sso_user resources into Terraform. Replace the {GROUP-ID} value.
```hcl
terraform import conformity_sso_user.user {user_id}
```

Use the `state` subcommand to verify Terraform successfully imported the conformity_sso_user resources.
```hcl
terraform state show conformity_sso_user.user
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Make sure you remove the id from the imported resource.
```hcl
terraform show -no-color >> main.tf
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_conformity"></a> [conformity](#requirement\_conformity) | 0.3.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_conformity"></a> [conformity](#provider\_conformity) | 0.3.0 |

## Resources

| Name | Type |
|------|------|

| conformity_sso_user.user | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_account01"></a> [account01](#input\_account01) | n/a | `string` | `""` | yes |
| <a name="input_account02"></a> [account02](#input\_account02) | n/a | `string` | `""` | yes |
| <a name="input_apikey"></a> [apikey](#input\_apikey) | n/a | `string` | `""` | yes |
| <a name="input_email"></a> [email](#input\_email) | n/a | `string` | `""` | yes |
| <a name="input_first_name"></a> [first\_name](#input\_first\_name) | n/a | `string` | `""` | yes |
| <a name="input_last_name"></a> [last\_name](#input\_last\_name) | n/a | `string` | `""` | yes |
| <a name="input_level01"></a> [level01](#input\_level01) | n/a | `string` | `""` | yes |
| <a name="input_level02"></a> [level02](#input\_level02) | n/a | `string` | `""` | yes |
| <a name="input_region"></a> [region](#input\_region) | n/a | `string` | `""` | yes |
| <a name="input_role"></a> [role](#input\_role) | n/a | `string` | `""` | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_user"></a> [user](#output\_user) | n/a |