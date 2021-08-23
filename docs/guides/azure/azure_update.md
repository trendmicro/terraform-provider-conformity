---
page_title: "Update Account Guide"
subcategory: "Accounts"
description: |-
  Provides instruction on how to update Azure Cloud Conformity account name or environment name using Terraform.
---

# How To Update Azure Cloud Conformity Resources
Provides instruction on how to update Azure Cloud Conformity account name or environment name using Terraform.

#### Step 1

##### Import resources
1. Navigate to folder Azure directory:
```sh
cd /path/terraform-provider-conformity/example/azure
```
2. Import the resources.
Azure Account - Can import based on the `Account ID` that can be found under the Conformity web console.

To get the Azure Account ID:
Open Conformity Admin Web console > navigate to Account dashboard > choose the Azure account you would like to import.
Notice the URL, you should be able to get the account ID e.g. https://cloudone.trendmicro.com/conformity/account/account-ID

3. Run `terraform init`:
```hcl
terraform init
```

4. Import the conformity_azure_account resources into Terraform. Replace the {CLOUDCONFORMITY-ACCOUNT-ID-AZURE} value.
```hcl
terraform import conformity_azure_account.azure {CLOUDCONFORMITY-ACCOUNT-ID-AZURE}
```

5. Use the `state` subcommand to verify Terraform successfully imported the conformity_azure_account resources.
```hcl
terraform state show conformity_azure_account.azure
```

6. Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Be sure to remove the id from the resource
```hcl
terraform show -no-color >> main.tf
```

NOTE: This will import ID and Subscription ID of the Azure Cloud Conformity Account.

#### Step 2

##### Preapare resources and run terraform

1. Create `terraform.tfvars` on `example/azure` folder and add the following values:

```
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "AWS-ACCOUNT-REGION"
azure_environment = "ACCESS-KEY"
azure_active_directory_id = "SECRET-ACCESS-KEY"
```
Note: You can always change the values declared according to your choice.

2. Go to `variable.tf` and uncomment the `environment`, `tags` (if not yet added) and `azure_active_directory_id`.

Here's an example usage:
```
variable "apikey"{
    type    = string
    default = ""
}
variable "region"{
    type    = string
    default = ""
}
 variable "azure_name"{
     type    = string
     default = "trendmicro_azure"
 }
variable "azure_environment"{
    type    = string
    default = "staging"
}

 variable "azure_subscription_id"{
     type    = string
     default = ""
 }

variable "azure_active_directory_id"{
    type    = string
    default = ""
}
```

3. Add some values on the resources such as `environment`, `tags` (if not yet added), and `azure_active_directory_id`. Comment out the `id` resource value. Here's example usage:
```
resource "conformity_azure_account" "azure" {
    subscription_id = "YOUR-SUBSCRIPTION-ID-FROM-IMPORT"
    name            = var.azure_name
    environment     = var.azure_environment
    active_directory_id = var.azure_active_directory_id
    settings {
        // implement multiple string values
        rule {
            rule_id = "Resources-001"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                exceptions {
                    filter_tags = []
                    resources   = []
                    tags        = [
                        "some_tag",
                    ]
                }
                extra_settings {
                    name    = "tags"
                    type    = "multiple-string-values"
                    values {
                        value = "Environment"
                    }
                    values {
                        value = "Role"
                    }
                    values {
                        value = "Owner"
                    }
                    values {
                        value = "Name"
                    }
                }
            }
        }
    }
}
```
NOTE: After the import you will notice that two values will be imported which are the `id` and `subscription_id` of the Azure Cloud Conformity Account. To use this resources which is for update purposes only, remove the `id`.

4. Run `terraform apply`
```sh
terraform apply
```