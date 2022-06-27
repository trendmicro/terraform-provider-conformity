# Conformity Provider
Terraform relies on plugins called "providers" to interact with remote systems.

Terraform configurations must declare which providers they require so that Terraform can install and use them. Additionally, some providers require configuration (like endpoint URLs or cloud regions) before they can be used.

These are the providers used to set up your Conformity Account for AWS. This needs to be properly configured by adding the correct credentials before using.

##### What Providers Do

Each provider adds a set of resource types and/or data sources that Terraform can manage.

Every resource type is implemented by a provider; without providers, Terraform can't manage any kind of infrastructure.

Most providers configure a specific infrastructure platform (either cloud or self-hosted). Providers can also offer local utilities for tasks like generating random numbers for unique resource names.

## List of Providers Used
There are two providers required for this to run, the Conformity Provider and AWS Provider.

### Conformity Provider
Conformity Provider Section

```hcl
terraform {
  required_providers {
    conformity = {
      version = "0.4.1"
      source  = "trendmicro/conformity"
    }
  }
}

provider "conformity" {
  region = var.region
  apikey = var.apikey
}
```
## Region Argument reference

- For `legacy accounts` Terraform will just support for below three regions
   * eu-west-1
   * us-west-2
   * ap-southeast-2 
- For `CloudOne accounts` Terraform will Support below regions
   * us-1
   * in-1
   * gb-1
   * jp-1
   * de-1
   * au-1
   * ca-1
   * sg-1
   
  Also click <a href="https://cloudone.trendmicro.com/docs/identity-and-account-management/c1-regions/"> here </a> to refer cloudone accounts region document
   

### AWS Provider

```hcl
terraform {
  required_providers {
      aws = {
      source  = "hashicorp/aws"
      version = ">= 3.44.0"
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
      version = "0.4.3"
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

### Static Credentials

!> **Warning:** Create a file name `terraform.tfvars` and add all necessary variables here.
Ensure `terraform.tfvars` is included in `.gitignore` so these secrets are not accidentally
pushed to a remote git repository. Terraform also provides a way of reading variables from
the environment: https://www.terraform.io/docs/cli/config/environment-variables.html#tf_var_name

## Support

This solution is provided ‘as is’, ‘with all faults’, ‘as available’ under a Trend Micro end user agreement ( www.trendmicro.com/en_sg/about/legal.html?modal=en-mulitcountry-tm-tools-eula.pdf ). This solution should be seen as community supported and Trend Micro will contribute our expertise as and when possible. We do not provide technical support or help in using or troubleshooting the components of the project through our normal support options. The underlying product used (Conformity API) by the solution are supported, but the support is only for the product functionality and not for help in deploying or using this solution itself.

Report an issue or commit a feature improvement to the provider https://github.com/trendmicro/terraform-provider-conformity. You can provide the output logs while reporting the issue to help our team resolve it faster. The instruction is as follows.

## Enable Debug Logs:

Set `TF_LOG` environment variable to `info`, this will print additional debug logs. The logs will be encrypted in form. Our team will be able to decrypt it.

```sh
TF_LOG=info terraform init
TF_LOG=info terraform apply
```

Contact the Trend Micro team at support@cloudconformity.com to report an issue or make a feature request.
