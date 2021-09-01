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
      version = "0.3.3"
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
      version = "0.3.3"
      source  = "trendmicro.com/cloudone/conformity"
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

This solution is provided ‘as is’, ‘with all faults’, ‘as available’ under a Trend Micro end user agreement<br/>
(www.trendmicro.com/en_sg/about/legal.html?modal=en-mulitcountry-tm-tools-eulapdf). This<br/>
solution should be seen as community supported and Trend Micro will<br/>
contribute our expertise as and when possible. We do not provide<br/>
technical support or help in using or troubleshooting the components of<br/>
the project through our normal support options. The underlying product<br/>
used (Conformity API) by the solution are supported, but the support is<br/>
only for the product functionality and not for help in deploying or<br/>
using this solution itself.<br/>


Report an issue or commit a feature improvement to the provider https://github.com/trendmicro/terraform-provider-conformity.<br/>
Contact the Trend Micro team at support@cloudconformity.com to report an issue or make a feature request<br/>
Report an Issue https://github.com/trendmicro/terraform-provider-conformity/issues.<br/>