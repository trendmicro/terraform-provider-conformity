---
page_title: "Update Account ID and/or Level Under Access List for Invite User - cloudconformity_terraform"
subcategory: "Users"
description: |-
  Provides instruction on how to update account ID and/or level under access list for invite user using Terraform.
---

# How To Update Account ID and/or Level Under Access List for Invite User
Provides instruction on how to update account ID and/or level under access list for invite user using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder user directory:
```sh
cd /path/guardrail-conformity-tf-provider/example/user/user
```
#### 2. Edit terraform.tfvars values according to the changes you want.

#### 3. Run terraform apply to apply the changes.
```sh
terraform apply
```
#### 3. Bash script that can run to automate the whole process 1-3. Go to script/user folder and look for `sh terraform-user-update.sh`. Find this and edit it according to the update that you want.

```sh
cat << EOF >> update.tfvars
role       = "ADMIN"

#level can be "NONE" "READONLY" "FULL"
account01 = "cloud-conformity-account-access"
level01  = "READONLY"

account02 = "cloud-conformity-account-access"
level02  = "READYONLY"
```

After editing the `sh terraform-user-update.sh` file, Run this command:
```sh
cd script/suer
sh terraform-user-update.sh
```

