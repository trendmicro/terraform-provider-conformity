terraform {
  required_providers {
    conformity = {
      version = "0.1.0"
      source  = "trendmicro.com/cloudone/conformity"
    }
      azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 2.62.1"
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