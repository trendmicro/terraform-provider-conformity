terraform {
  required_providers {
    conformity = {
      version = "0.3.3"
      source  = "trendmicro.com/cloudone/conformity"

    }
  }
}

provider "conformity" {
  region = var.region
  apikey = var.apikey
}