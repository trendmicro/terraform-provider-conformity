---
page_title: "Update Account ID and/or Level Under Access List for Invite SSO User"
subcategory: "Users"
description: |-
  Provides instruction on how to update account ID and/or level under access list for invite SSO user using Terraform.
---

# How To Update Account ID and/or Level Under Access List for Invite User
Provides instruction on how to update account ID and/or level under access list for invite SSO user using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder sso_user directory:
```sh
cd /path/terraform-provider-conformity/example/sso_user/user
```
#### 2. Edit terraform.tfvars values according to the changes you want.

#### 3. Run terraform apply to apply the changes.
```sh
terraform apply
```
#### 3. Bash script that can run to automate the whole process 1-3. Go to script/sso_user folder and look for `sh terraform-sso_user-update.sh`. Find this and edit it according to the update that you want.

```sh
cat << EOF >> update.tfvars
# conformity_sso_user
role       = "ADMIN"

#level can be "NONE" "READONLY" "FULL"
account01 = "cloud-conformity-account-access"
level01  = "READONLY"

account02 = "cloud-conformity-account-access"
level02  = "READYONLY"
```

After editing the `sh terraform-sso_user-update.sh` file, Run this command:
```sh
cd script/suer
sh terraform-sso_user-update.sh
```

