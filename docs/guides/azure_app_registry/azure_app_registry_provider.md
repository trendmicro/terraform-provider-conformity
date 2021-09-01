---
page_title: "Configure Azure App Registry Provider"
subcategory: "Accounts"
description: |-
  Provides instruction on how to configure Providers to Update Azure Conformity resources using Terraform.
---

# How To Configure Azure App Resgistry Provider
Provides instruction on how to configure Providers to Update Azure Conformity resources using Terraform.

## Things needed:
1. Azure login Access
2. Conformity API Key

#### Step 1

##### Terraform Configuration

1. To configure the provider, make sure that conformity region, api key and azure access are properly configured on the `terraform.tfvars`.

## Example Usage for `terraform.tfvars`
```hcl
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
      azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 2.62.1"
    }
  }
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
 - `subscription_id` - (Optional) This is the Azure Account subscription ID. 
 - `client_id` - (Optional) This is the Azure Account client ID. 
 - `client_secret` - (Optional) This is the Azure Account client secret ID. 
 - `tenant_id` - (Optional) This is the Azure Account tenant ID. 