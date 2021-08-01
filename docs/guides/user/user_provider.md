---
page_title: "Configure Invite User Provider - cloudconformity_terraform"
subcategory: "Users"
description: |-
  Provides instruction on how to configure invite user to your organisation using Terraform. This resource is not applicable to users who are part of the Cloud One Platform.
---

# How To Configure Invite User Provider
Provides instruction on how to configure invite user to your organisation using Terraform. This resource is not applicable to users who are part of the Cloud One Platform.

## Things needed:
1. AWS Access Key and Secret Access Key
2. Cloud Conformity API Key

#### Step 1

##### Terraform Configuration

1. To configure the provider, make sure that the region and API Key from Cloud Conformity Account are properly configured on the `terraform.tfvars`.

## Example Usage for `terraform.tfvars`
```hcl
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "ACCOUNT-REGION"

# conformity_user
first_name = "John"
last_name  = "Doe"
email      = "john.doe@cloudconformity.com"
role       = "USER"

# access_list01 (can be multiple)
#level can be "NONE" "READONLY" "FULL"
account01 = "cloud-conformity-account-access"
level01  = "ADD-LEVEL"

account02 = "cloud-conformity-account-access"
level02  = "ADD-LEVEL"
```
Note: You can always change the values declared according to your choice.

## Example Usage `provider.tf`
```hcl
terraform {
  required_providers {
    conformity = {
      version = "0.3.1"
      source  = "trendmicro.com/cloudone/conformity"
    }
  }
}

provider "conformity" {
  region = var.region
  apikey = var.apikey
}
```

## Argument Reference
 - `apikey` - (Required) This is the Cloud Conformity API Key. 
 - `region` - (Required) The region your organisation resides in. See https://github.com/cloudconformity/documentation-api for the available regions.