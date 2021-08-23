---
page_title: "Destroy Communication Settings Guide"
subcategory: "Communication Settings"
description: |-
  Provides instruction on how to destroy Communication Settings using Terraform.
---

# How To Destroy Communication Settings and Cloud Conformity Resources
Provides instruction on how to destroy Communication Settings using Terraform.

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

#### 2. Run terraform destroy:
```sh
terraform destroy
```
#### 3. Bash script that can run to automate the whole process 1-2:
Note: Change the `PATH-TO-CHANNEL` value according to the channel you want to create (e.g. email, ms-teams, slack multiple, sms, sns).

Under script folder run:
```sh
cd script/communication_setting/PATH-TO-CHANNEL
sh terraform-PATH-TO-CHANNEL-destroy.sh
```

Example:
```sh
cd script/communication_setting/email
sh terraform-email-destroy.sh
```
