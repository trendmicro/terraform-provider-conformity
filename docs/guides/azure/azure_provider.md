---
page_title: "Configure Azure Provider - cloudconformity_terraform"
subcategory: "Accounts"
description: |-
  Provides instruction on how to configure Providers to Update Azure Cloud Conformity resources using Terraform.
---

# How To Configure Azure Provider
Provides instruction on how to configure Providers to Update Azure Cloud Conformity resources using Terraform.

## Things needed:
1. Cloud Conformity API Key
2. Configured App Registration on Azure for Cloud Conformity like described here https://www.cloudconformity.com/help/add-cloud-account/add-an-azure-account.html#add-an-azure-active-directory or from our example at `../azure_app_registry`

#### Step 1

##### Terraform Configuration

1. To configure the provider, make sure that region and api key are properly configured on the `terraform.tfvars`.

## Example Usage for `terraform.tfvars`
```
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "ACCOUNT-REGION"
azure_environment = "AZURE-ENVIRONMENT-NAME"
azure_active_directory_id = "SECRET-ACCESS-KEY"
```
Note: You can always change the values declared according to your choice.

## Example Usage `provider.tf`
```hcl
terraform {
  required_providers {
    conformity = {
      version = "0.1.0"
      source  = "cloudone.com/cloud/conformity"
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
