---
page_title: "Configure AWS Provider"
subcategory: "Groups"
description: |-
  Provides instruction on how to configure Providers to create or add groups and tags on Conformity using Terraform.
---

# How To Configure Groups and Tags Provider
Provides instruction on how to configure Providers to create or add groups and tags on Conformity using Terraform.

## Things needed:
1. AWS Access Key and Secret Access Key
2. Conformity API Key

#### Step 1

##### Terraform Configuration

1. To configure the provider, make sure that the region and API Key from Conformity Account are properly configured on the `terraform.tfvars`.

## Example Usage for `terraform.tfvars`
```hcl
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "ACCOUNT-REGION"

name_group1 = "NAME-OF-GROUP1"
tag_group1 = ["NAME-OF-GROUP1-TAG1","NAME-OF-GROUP1-TAG2"]

name_group2 = "NAME-OF-GROUP2"
tag_group2 = ["NAME-OF-GROUP2-TAG1", "NAME-OF-GROUP2-TAG2"]
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
 - `region` - (Required) The region your organisation resides in. See https://github.com/cloudconformity/documentation-api for the available regions.