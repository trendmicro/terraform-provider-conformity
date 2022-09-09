---
page_title: "conformity_current_user"
subcategory: "Users"
description: |-
 Allows you to get the Current User
---


# Data Source `conformity_current_user`

Allows you to get the Current User

## Example Usage

```
data "conformity_current_user" "user" {}
output "user_details"{
    value =data.conformity_current_user.user
}

```
