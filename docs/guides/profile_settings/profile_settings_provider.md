---
page_title: "Configure Profile Settings Provider - cloudconformity_terraform"
subcategory: "Profile Settings"
description: |-
  Provides instruction on how to configure Providers for Profile Settings resources using Terraform.
---

# How To Configure Cloud Conformity Provider
Provides instruction on how to configure Providers for Profile Settings resources using Terraform.

## Things needed:
1. Cloud Conformity API Key

#### Step 1

##### Terraform Configuration

1. To configure the provider, make sure that the Region and API Key from Cloud Conformity Account are properly configured on the `terraform.tfvars`.

## Example Usage for `terraform.tfvars`
```
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "ACCOUNT-REGION"
```
Note: You can always change the values declared according to your choice.

## Example Usage `provider.tf`
```hcl
terraform {
  required_providers {
    conformity = {
      version = "0.1.0"
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
 - `region` - (Required) The region your organisation resides in. See https://github.com/cloudconformity/documentation-api
   for the available regions.