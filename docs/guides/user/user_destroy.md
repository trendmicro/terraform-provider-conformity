---
page_title: "Create An Invite User To Your Organisation"
subcategory: "Users"
description: |-
  Provides instruction on how to destroy invite user to your organisation using Terraform.
---

# How To Destroy Invite User on Cloud Conformity
Provides instruction on how to destroy invite user to your organiation using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder user directory:
```sh
cd /path/terraform-provider-conformity/example/user/user
```
#### 2. Run terraform destroy:
```sh
terraform destroy
```
#### 3. Bash script that can run to automate the whole process 1-2:

Under script folder run:
```sh
cd script/user
sh terraform-user-destroy.sh
```