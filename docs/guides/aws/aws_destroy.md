---
page_title: "Destroy Account Guide"
subcategory: "Accounts"
description: |-
  Provides instruction on how to destroy Conformity Account using Terraform.
---

# How To Destroy AWS and Conformity Resources
Provides instruction on how to destroy Conformity Account using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder aws directory:
```sh
cd /path/terraform-provider-conformity/example/aws
```
#### 2. Run terraform destroy:
```sh
terraform destroy
```
#### 3. Bash script that can run to automate the whole process 1-2:

Under script folder run:
```sh
cd script/aws
sh terraform-aws-destroy.sh
```