---
page_title: "conformity_user Resource"
subcategory: "User"
description: |-
  Allows you to invite a user to your organisation. This resource does not apply to users who are part of the Cloud One Platform.
---

# Resource `conformity_user`
Allows you to invite a user to your organisation. This resource does not apply to users who are part of the Cloud One Platform.

## Example Usage
```hcl
resource "conformity_user" "user" {
  first_name = var.first_name
  last_name  = var.last_name
  email      = var.email
  role       = var.role

  access_list {
    account = var.account01
    level   = var.level01
    }

  access_list {
    account = var.account02
    level   = var.level02
    }
}

output "user"{
    value = conformity_user.user
}
```

## Argument reference

 - `email` (String) - (Required) The email of the sso_user.
 - `first_name` (String) - (Required) The first name of the sso_user.
 - `last_name` (String) - (Required) The last name of the sso_user.
 - `role` (String) - (Required) The role which the sso_user is assigned to. Can be "ADMIN" "USER".
 - `mfa` (Bool) - Shows 'true' if the user has MFA, otherwise 'false'.
 - `access_list` - (Optional) List:
     * `account` (String) - (Required) The account id within the organisation.
     * `level` (String) - (Required) The level of access the user has to the account. Can be "NONE" "READONLY" "FULL".

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

 - `id` - The ID of the user in Conformity managed by this resource
 - `last_login` - Last login of the user

Example usage on the template:

```hcl
report {
    id = conformity_user.user.id
}
```

## Import

User - Can import based on the user_id that can be found under the terraform state or user login.

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_user resources into Terraform. Replace the {user_id} value.
```hcl
terraform import conformity_user.user {user_id}

```

Use the `state` subcommand to verify Terraform successfully imported the conformity_user resources.
```hcl
terraform state show conformity_user.user
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Make sure you remove the id from the imported resource.
```hcl
terraform show -no-color >> main.tf
```
