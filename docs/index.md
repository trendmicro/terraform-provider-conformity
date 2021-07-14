# Cloud Conformity Provider
Terraform relies on plugins called "providers" to interact with remote systems.

Terraform configurations must declare which providers they require so that Terraform can install and use them. Additionally, some providers require configuration (like endpoint URLs or cloud regions) before they can be used.

These are the providers used to set up your Cloud Conformity Account for AWS. This needs to be properly configured by adding the correct credentials before using.

##### What Providers Do

Each provider adds a set of resource types and/or data sources that Terraform can manage.

Every resource type is implemented by a provider; without providers, Terraform can't manage any kind of infrastructure.

Most providers configure a specific infrastructure platform (either cloud or self-hosted). Providers can also offer local utilities for tasks like generating random numbers for unique resource names.

## List of Providers Used
There are two providers required for this to run, the Cloud Conformity Provider and AWS Provider.

### Cloud Conformity Provider
Cloud Conformity Provider Section

```hcl
terraform {
  required_providers {
    conformity = {
      version = "0.1.0"
      source  = "trendmicro.com/cloudone/conformity"
    }
  }
}

provider "conformity" {
  region = var.region
  apikey = var.apikey
}

```
### AWS Provider

```hcl
terraform {
  required_providers {
      aws = {
      source  = "hashicorp/aws"
      version = ">= 2.7.0"
    }
  }
}

provider "aws" {
  region     = var.region
  access_key = var.access_key
  secret_key = var.secret_key
}
```

## Example Usage
```hcl
terraform {
  required_providers {
    conformity = {
      version = "0.1.0"
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
  region     = var.region
  access_key = var.access_key
  secret_key = var.secret_key
}
```
