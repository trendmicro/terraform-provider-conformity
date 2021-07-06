---
page_title: "Configure Azure App Registry Provider - cloudconformity_terraform"
subcategory: "Accounts"
description: |-
  Provides instruction on how to configure Providers to Update Azure Cloud Conformity resources using Terraform.
---

# How To Configure Azure App Resgistry Provider
Provides instruction on how to configure Providers to Update Azure Cloud Conformity resources using Terraform.

## Things needed:
1. Azure login Access
2. Cloud Conformity API Key

#### Step 1

##### Terraform Configuration

1. To configure the provider, make sure that conformity region, api key and azure access are properly configured on the `terraform.tfvars`.

## Example Usage for `terraform.tfvars`
```
  apikey = "CLOUD-CONFORMITY-API-KEY"
  region = "ACCOUNT-REGION"

# Uncomment this section if you want to login or run terraform using keys.
#  subscription_id = "SUBSCRIPTION-ID"
#  client_id       = "CLIENT_ID"
#  client_secret   = "CLIENT_SECRET"
#  tenant_id       = "TENANT_ID"
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
      azurerm = {
      source  = "hashicorp/azurerm"
      version = "=2.46.0"
    }
  }
}

provider "conformity" {
  region = var.region
  apikey = var.apikey
}

provider "azurerm" {
  features {}

  # Uncomment this section if you are going to use keys for access
  # subscription_id = var.subscription_id
  # client_id       = var.client_id
  # client_secret   = var.client_secret
  # tenant_id       = var.tenant_id
}
```

## Argument Reference
 - `apikey` - (Required) This is the Cloud Conformity API Key. 
 - `region` - (Required) The region your organisation resides in. See https://github.com/cloudconformity/documentation-api for the available regions.
 - `subscription_id` - (Optional) This is the Azure Account subscription ID. 
 - `client_id` - (Optional) This is the Azure Account client ID. 
 - `client_secret` - (Optional) This is the Azure Account client secret ID. 
 - `tenant_id` - (Optional) This is the Azure Account tenant ID. 