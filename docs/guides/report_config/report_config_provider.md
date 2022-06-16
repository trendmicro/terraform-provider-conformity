---
page_title: "Configure Report Configs Provider"
subcategory: "Report Configs"
description: |-
  Provides instruction on how to configure Providers for Report Configs resources using Terraform.
---

# How To Configure Conformity Provider
Provides instruction on how to configure Providers for Report Configs resources using Terraform.

## Things needed:
1. Conformity API Key

#### Step 1

##### Terraform Configuration

1. To configure the provider, make sure that the Region and API Key from Conformity Account are properly configured on the `terraform.tfvars`.

## Example Usage for `terraform.tfvars`
```hcl
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "ACCOUNT-REGION"
```
Note: You can always change the values declared according to your choice.

## Example Usage `provider.tf`
```hcl
terraform {
  required_providers {
    conformity = {
      source  = "trendmicro/conformity"
    }
  }
}

provider "conformity" {
  region = var.region
  apikey = var.apikey
}
```

## Argument Reference
 - `apikey` - (Required) This is the Conformity API Key. 
 - `region` - (Required) The region your organisation resides in. See https://github.com/cloudconformity/documentation-api
   for the available regions.
