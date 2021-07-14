---
page_title: "Update Account Guide - cloudconformity_terraform"
subcategory: "Accounts"
description: |-
  Provides instruction on how to update AWS Cloud Conformity account name or environment name using Terraform.
---

# How To Update AWS and Cloud Conformity Resources
Provides instruction on how to update AWS Cloud Conformity account name or environment name using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder aws directory:
```sh
cd /path/terraform-provider-conformity/example/aws
```
#### 2. Edit terraform.tfvars values according to the changes you want.

#### 3. Run terraform apply to apply the changes.
```sh
terraform apply
```
#### 3. Bash script that can run to automate the whole process 1-3. Go to script/aws folder and look for `sh terraform-aws-update.sh`. Find this and edit it according to the update that you want.

```sh
cat << EOF >> update.tfvars
name = "UPDATED-NAME"
environment = "UPDATED-ENVIRONMENT"
EOF
```

After editing the `sh terraform-aws-update.sh` file, Run this command:
```sh
cd script/aws
sh terraform-aws-update.sh
```