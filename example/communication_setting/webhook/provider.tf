terraform {
  required_providers {
    conformity = {
      version = "0.1.0"
      source  = "trendmicro.com/cloudone/conformity"
    }
  }
}

provider "conformity" {
  region = "us-west-2"
  apikey = "yzDsSZCRZKhjS25YxuT7RPDgh7YK9r8xEzb5e7utc3VvPRSPvVv62c1os6U6jAey"
}