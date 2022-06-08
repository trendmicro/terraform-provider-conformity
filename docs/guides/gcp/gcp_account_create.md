---
page_title: "Create GCP Account Guide"
subcategory: "Accounts"
description: |-
  Provides instruction on how to create or add GCP account on Conformity using Terraform.
---

# How To Create or Add GCP Account on Conformity on a local machine
Provides instruction on how to create or add GCP account on Conformity using Terraform.

## Things needed:
1. GCP name, project id, project name and service account unique_id
2. Conformity API Key

#### Step 1

##### Terraform Configuration

1. Create terraform `cloud_formation.tf` on `example/gcp` folder for Cloudformation module, create the file if not existing and add the following values:
```hcl
resource "conformity_gcp_account" "gcp" {
    name            = "MyProject"
    project_id       = "conformity-123456"
    project_name     = "conformity"
    service_account_unique_id = "123456"
    environment = "staging"

}
```
2. To use Conformity and its resources, add the GCP Access Key, Secret Access Key, and Region where the resources will be created and API Key from Conformity Account to the `terraform.tfvars`. 

3. Create `terraform.tfvars` on `example/GCP` folder and add the following values:

```hcl
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "CONFORMITY-REGION"
```
Note: You can always change the values declared according to your choice.

#### Step 2

##### Run Terraform

#### 1. Navigate to project directory:
```sh
cd /path/terraform-provider-conformity/
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
cd example/gcp
terraform init
terraform plan
terraform apply
```
#### 5. Bash script that can run to automate the whole process 1-5:

Under script folder run:
```sh
cd script/gcp
sh terraform-gcp-create.sh
```
