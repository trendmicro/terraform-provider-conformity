terraform {
  required_providers {
    conformity = {
      source  = "trendmicro/conformity"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
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
