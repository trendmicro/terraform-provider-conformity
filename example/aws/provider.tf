terraform {
  required_providers {
    conformity = {
      version = "0.3.0"
      source  = "trendmicro.com/cloudone/conformity"
    }
      aws = {
      source  = "hashicorp/aws"
      version = ">= 2.7.0"
    }
  }
}

provider "conformity" {
  region = var.region
  apikey = var.apikey
}

provider "aws" {
  region = var.region
  access_key = var.access_key
  secret_key = var.secret_key
}
