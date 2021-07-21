---
page_title: "Create Profile Settings Guide - cloudconformity_terraform"
subcategory: "Profile Settings"
description: |-
  Provides instruction on how to create Profile Settings on Cloud Conformity using Terraform.
---

# How To Create Profile Settings on Cloud Conformity on a local machine
Provides instruction on how to create Profile Settings on Cloud Conformity using Terraform.

## Things needed:
1. Cloud Conformity API Key

#### Step 1

##### Terraform Configuration

1. To use Cloud Conformity and its resources, make sure that the values for certain variables are declared, Region, Account ID and API Key from Cloud Conformity Account are properly configured on the `terraform.tfvars`.

2. Create `terraform.tfvars` on `example/profile_settings/PATH-TO-PROFILE-CONFIG` folder.
   
3. Note: Change the `PATH-TO-PROFILE-CONFIG` value according to the configuration you want to create (e.g. existing_profile, multiple_extra_settings, values_string_int, with_rules, without_rules).
Example Path: `example/profile_settings/with_rules`

Add the following values on `terraform.tfvars`:
```hcl
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "AWS-ACCOUNT-REGION"
```
Note: You can always change the values declared according to your choice.

4. Add filter or configuration values according to your requirements `main.tf` under `/path/terraform-provider-conformity/example/profile_settings/PATH-TO-PROFILE-CONFIG` directory.
Note: Change the `PATH-TO-PROFILE-CONFIG` value according to the configuration you want to create (e.g. existing_profile, multiple_extra_settings, values_string_int, with_rules, without_rules).
Example Path: `/path/terraform-provider-conformity/example/profile_settings/with_rules`

Example template to guide your configuration:

```hcl
resource "conformity_profile" "profile"{
  // Optional | type: string
  name = ""

  // Optional | type : string
  description = ""
  
  // Optional | type : string
  profile_id = ""

  // optional
  // can be multiple declaration
  included {

    // optional | type: string
    // id of the rule set
    id   = ""

    // optional | type: string
    type = ""

    // optional | type: bool | default: true
    enabled = bool

    // optional | type: string
    provider = ""
    // optional | type: string
    // value can be: "LOW" "MEDIUM" "HIGH" "VERY_HIGH" "EXTREME"
    risk_level = []

    // optional
    exceptions {

      // optional | type: array of string
      filter_tags = []

      // optional | type: array of string
      resources   = []

      // optional | type: array of string
      tags  = []

    }

    // optional
    // can be multiple declaration
    extra_settings {
      // optional | type: bool
      countries = bool

      // optional | type: bool
      multiple = bool

      // optional | type: string
      name = ""

      // optional | type: bool
      regions = bool 

      // required | type: string
      type = ""

      // optional | type: string
      value =  ""


      // optional | type: list
            values {
                // required | type: bool
                enabled = bool
                // required | type: string
                label   = ""
                // required | type: string
                value   = ""
            }

    }
    
  }

}
```

Note: if you want to test the example folder `existing profile`, add the profile_id on the resources.

#### Step 2

##### Run Terraform

#### 1. Navigate to project directory:
```sh
cd /path/terraform-provider-conformity/
```
#### 2. Install dependencies:
```sh
go mod vendor
```
#### 3. Create the Artifact:
```sh
make install
```
#### 4. Now, you can run terraform code:
Note: Change the `PATH-TO-PROFILE-CONFIG` value according to the configuration you want to create (e.g. existing_profile, multiple_extra_settings, values_string_int, with_rules, without_rules).
```sh
cd example/profile_settings/PATH-TO-PROFILE-CONFIG
terraform init
terraform plan
terraform apply
```

Example:
```sh
cd example/profile_settings/with_rules
terraform init
terraform plan
terraform apply
```

#### 5. Bash script that can run to automate the whole process 1-5:
Note: Change the `PATH-TO-PROFILE-CONFIG` value according to the configuration you want to create (e.g. existing_profile, multiple_extra_settings, values_string_int, with_rules, without_rules).

Under script folder run:
```sh
cd script/profile_settings/PATH-TO-PROFILE-CONFIG
sh terraform-PATH-TO-PROFILE-CONFIG-create.sh
```

Example:
```sh
cd script/profile_settings/with_rules
sh terraform-with_rules-create.sh
```
