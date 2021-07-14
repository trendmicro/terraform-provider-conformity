---
page_title: "Update Profile Settings Guide - cloudconformity_terraform"
subcategory: "Profile Settings"
description: |-
  Provides instruction on how to update Profile Settings on Cloud Conformity using Terraform.
---

# How To Update Cloud Conformity Resources
Provides instruction on how to update Profile Settings on Cloud Conformity using Terraform.

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

#### 2. Edit `main.tf` values according to the changes you want.

#### 3. Run terraform apply to apply the changes.
```sh
terraform apply
```