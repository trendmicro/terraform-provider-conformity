---
page_title: "Create Azure App Registry Guide"
subcategory: "Azure App Registry"
description: |-
  Provides instruction on how to create Azure App Registry on Azure account using Terraform.
---

# How To Create or Add Azure App Registry on Conformity on a local machine
Provides instruction on how to create Azure App Registry on Azure account using Terraform.

## Things needed:
1. Azure login Access

#### Step 1

##### Login to Azure

##### Option 1
Sign in with azure cli.
1. Run `az login` command to your terminal.
```sh
az login
```
2. If the CLI can open your default browser, it will do so and load an Azure sign-in page.

Otherwise, open a browser page at https://aka.ms/devicelogin and enter the authorization code displayed in your terminal.

If no web browser is available or the web browser fails to open, use device code flow with az login --use-device-code.

3. Sign in with your account credentials in the browser.

##### Option 2
Sign in with keys.
1. Navigate to azure_app_registry directory and look for `provider.tf`:
```sh
cd /path/terraform-provider-conformity/example/azure_app_registry
```
2. Edit the azurerm resource and uncomment the subscription_id, client_id, client_secret and tenant_id.
   
Example usage: 

provider "azurerm" {
  features {}

  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
}

3. Go to `terraform.tfvars` on `/example/azure_app_registry` folder, create the file if not existing and add the following values:

```hcl
  subscription_id = "SUBSCRIPTION-ID"
  client_id       = "CLIENT_ID"
  client_secret   = "CLIENT_SECRET"
  tenant_id       = "TENANT_ID"
```
Note: You can always change the values declared according to your choice.


##### Terraform Configuration

#### Step 1

##### Run terraform app registration module.

#### 1. Create terraform `main.tf`:

Create terraform `main.tf` on `example/azure_app_registry` folder, create the file if not existing and add the following values:
  
```hcl
module "azure_app_registry" {
  source      = "../"
}
```

#### 2. Now, you can run terraform code on `example/azure_app_registry` directory:
```sh
terraform init
terraform plan
terraform apply
```
#### 3. View outputs.
a. You can view outputs found on the terminal.
b. To view `azuread_application_password`, run
```sh
terraform output app_registration_key
```
c. Save the outputs especially the `active_directory_tenant_id`, `app_registration_application_id` and the `app_registration_key` output from the previous number (Step 2, Number 3, Letter B).

#### Step 2

##### Add Azure Account on Conformity Console
1. Go to Conformity Console.
2. To allow Conformity access to your Azure Subscriptions, you will use the `active_directory_tenant_id`, `app_registration_application_id` and the `app_registration_key` created in the previous setup. Configure it according to the value needed. This will allow the Conformity rule engine to run Rule checks against Subscriptions within your Azure Active Directory.
3. Add the Subscription.

Note: You can also create the pre-requisite for Azure Conformity Account creation manually without using this terraform module. Also, you can check what resources this module is creating. Please refer to this link https://www.cloudconformity.com/help/add-cloud-account/add-an-azure-account.html#add-an-azure-active-directory .