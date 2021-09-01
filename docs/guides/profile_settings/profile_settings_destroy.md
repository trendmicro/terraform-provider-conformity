---
page_title: "Destroy Profile Settings Guide"
subcategory: "Profile Settings"
description: |-
  Provides instruction on how to destroy Profile Settings using Terraform.
---

# How To Destroy Conformity Resources
Provides instruction on how to destroy Profile Settings using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder communication settings directory:
Note: Change the `PATH-TO-PROFILE-CONFIG` value according to the configuration you want to create (e.g. existing_profile, multiple_extra_settings, values_string_int, with_rules, without_rules).
```sh
cd example/profile_settings/PATH-TO-PROFILE-CONFIG
```

Example:
```sh
cd example/profile_settings/with_rules
```

#### 2. Run terraform destroy:
```sh
terraform destroy
```
#### 3. Bash script that can run to automate the whole process 1-2:
Note: Change the `PATH-TO-PROFILE-CONFIG` value according to the configuration you want to create (e.g. existing_profile, multiple_extra_settings, values_string_int, with_rules, without_rules).

Under script folder run:
```sh
cd script/profile_settings/PATH-TO-PROFILE-CONFIG
sh terraform-PATH-TO-PROFILE-CONFIG-destroy.sh
```

Example:
```sh
cd script/profile_settings/with_rules
sh terraform-with_rules-destroy.sh
```
