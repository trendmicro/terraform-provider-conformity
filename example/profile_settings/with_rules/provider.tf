terraform {
  required_providers {
    conformity = {

      source  = "trendmicro.com/cloudone/conformity"
    }
  }
}

provider "conformity" {
  region = var.region
  apikey = var.apikey
}
