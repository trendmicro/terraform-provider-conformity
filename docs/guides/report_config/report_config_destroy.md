---
page_title: "Destroy Report Configs Guide"
subcategory: "Report Configs"
description: |-
  Provides instruction on how to destroy Report Configs using Terraform.
---

# How To Destroy Cloud Conformity Resources
Provides instruction on how to destroy Report Configs using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder report_config directory:
```sh
cd /path/terraform-provider-conformity/example/report_config/main
```
#### 2. Run terraform destroy:
```sh
terraform destroy
```
#### 3. Bash script that can run to automate the whole process 1-2:

Under script folder run:
```sh
cd script/report_config
sh terraform-report_config-destroy.sh
```