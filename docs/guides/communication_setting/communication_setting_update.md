---
page_title: "Update Communication Settings Guide - cloudconformity_terraform"
subcategory: "Communication Settings"
description: |-
  Provides instruction on how to update Communication Settings on Cloud Conformity using Terraform.
---

# How To Update AWS and Cloud Conformity Resources
Provides instruction on how to update Communication Settings on Cloud Conformity using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder communication settings directory:
Note: Change the `PATH-TO-CHANNEL` value according to the channel you want to create (e.g. email, ms-teams, slack multiple, sms, sns).
```sh
cd example/communication_setting/PATH-TO-CHANNEL
```

Example:
```sh
cd example/communication_setting/email
```

#### 2. Edit `main.tf` values according to the changes you want.

#### 3. Run terraform apply to apply the changes.
```sh
terraform apply
```