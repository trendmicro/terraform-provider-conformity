---
page_title: "Create an Azure Active Directory"
subcategory: "Azure"
description: |-
  Provides instruction on how to create an Azure Active Directory.
---

# How To Create  an Azure Active Directory
Provides instruction on how to create an create an Azure Active Directory..

## Things needed:
1. Conformity API Key
2. name,directory_id,application_id,application_key


#### Step 1

##### Terraform Configuration

1. Create terraform `main.tf` on `example/azure/azure_active_directory` folder for Azure Active Directory Module, create the file if not existing and add the following values:
```hcl
resource "conformity_azure_active_directory" "azure" {
    name = "Azure Active Directory"
    directory_id    = "761d49c9-8898-5d35-c4db-ed8582f20dbf"
    application_id     = var.application_id
    application_key = var.application_key
}
```
2. To use Conformity and its resources, add the  Region where the resources will be created and API Key from Conformity Organisation to the `variable.tf`. 

3. Create `variable.tf` on `example/azure/azure_active_directory` folder and add the following values:

```hcl
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "GCP-REGION"

```
Note: You can always change the values declared according to your choice.

#### Step 2

##### Run Terraform

#### 1. Navigate to project directory:
```sh
cd /path/terraform-provider-conformity
```
#### 2. Install dependencies:
```sh
go mod vendor
```
#### 3. Create the Artifact:
```sh
make install
```
#### 4. Now, you can run terraform code:
```sh
cd example/azure/azure_active_directory
terraform init
terraform plan
terraform apply
```

