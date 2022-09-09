terraform {
  required_providers {
    conformity = {

      source  = "trendmicro/conformity"
    }
  }
}

provider "conformity" {
  region =  var.region //us-region
  apikey = var.apikey 
}
