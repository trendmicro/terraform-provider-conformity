---
page_title: "Destroy Groups and Tags on Cloud Conformity - cloudconformity_terraform"
subcategory: "Groups"
description: |-
  Provides instruction on how to destroy groups and tags on Cloud Conformity Account using Terraform.
---

# How To Destroy Groups and Tags on Cloud Conformity
Provides instruction on how to destroy groups and tags on Cloud Conformity Account using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder group directory:
```sh
cd /path/guardrail-conformity-tf-provider/example/group
```
#### 2. Run terraform destroy:
```sh
terraform destroy
```
#### 3. Bash script that can run to automate the whole process 1-2:

Under script folder run:
```sh
cd script/group
sh terraform-group-destroy.sh
```