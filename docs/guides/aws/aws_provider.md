---
page_title: "Configure AWS Provider"
subcategory: "Accounts"
description: |-
  Provides instruction on how to configure Providers to create AWS and Conformity resources using Terraform.
---

# How To Configure AWS Provider
Provides instruction on how to configure Providers to create AWS and Conformity resources using Terraform.

## Things needed:
1. AWS Access Key and Secret Access Key
2. Conformity API Key

#### Step 1

##### Terraform Configuration

1. To configure the provider, make sure that the AWS Access Key, Secret Access Key, and Region where the resources will be created and API Key from Conformity Account are properly configured on the `terraform.tfvars`.

## Example Usage for `terraform.tfvars`
```hcl
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "AWS-ACCOUNT-REGION"
access_key="ACCESS-KEY"
secret_key="SECRET-ACCESS-KEY"
```
Note: You can always change the values declared according to your choice.

## Example Usage `provider.tf`
```hcl
terraform {
  required_providers {
    conformity = {
      source  = "trendmicro/conformity"
    }
      aws = {
      source  = "hashicorp/aws"
      version = ">= 3.44.0"
    }
  }
}

provider "conformity" {
  region = var.region
  apikey = var.apikey
}

provider "aws" {
  region     = var.region
  access_key = var.access_key
  secret_key = var.secret_key
}
```

## Argument Reference
 - `apikey` - (Required) This is the Conformity API Key. 
 - `region` - (Required) The region your organisation resides in. See https://github.com/cloudconformity/documentation-api for the available regions.
 - `access_key` - (Required) This is the AWS Access Key. 
 - `secret_key` - (Required) This is the AWS Secret Access Key. 
