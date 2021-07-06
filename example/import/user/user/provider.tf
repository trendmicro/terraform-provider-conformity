terraform {
  required_providers {
    conformity = {
      version = "0.1.0"
      source  = "trendmicro.com/cloudone/conformity"
    }
  }
}

provider "conformity" {
  region = "ap-southeast-2"
  apikey = "3yjFvtsDVFJXkP79kywtX1WYDXQAj85ji2o7fh2aHJ5pjGrQtTt55XwyGERcVuKm"
}