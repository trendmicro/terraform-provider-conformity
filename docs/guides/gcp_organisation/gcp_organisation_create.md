---
page_title: "Create GCP Organisation Guide"
subcategory: "Organisation"
description: |-
  Provides instruction on how to create or add GCP Organisation on Conformity using Terraform.
---

# How To Create or Add GCP Organisation on Conformity on a local machine
Provides instruction on how to create or add GCP Organisation on Conformity using Terraform.

## Things needed:
1. service_account_name and service account key JSON
2. Conformity API Key

#### Step 1

##### Terraform Configuration

1. Create terraform `cloud_formation.tf` on `example/gcp_organisation` folder for Cloudformation module, create the file if not existing and add the following values:
```hcl
resource "conformity_gcp_org" "gcp_org" {
    private_key              = var.private_key
    service_account_name     = "MySubscription"
    type                     = "service_account"
    project_id               = "conformity-346910"
    private_key_id           = "c1c3688e7c"
    client_email             = "iam.gserviceaccount.com"
    client_id                = "811129548"
    auth_uri                 = "https://accounts.google.com/o/oauth2/auth"
    token_uri                = "https://oauth2.googleapis.com/token"
    auth_provider_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
    client_x509_cert_url     = "https://www.googleapis.com/robot/v1/metadata/x509/cloud-one-conformity-bot%40conformity-346910.iam.gserviceaccount.com"
}
```
2. To use Conformity and its resources, add the GCP Access Key, Secret Access Key, and Region where the resources will be created and API Key from Conformity Organisation to the `terraform.tfvars`. 

3. Create `terraform.tfvars` on `example/gcp_organisation` folder and add the following values:

```hcl
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "GCP-REGION"
private_key = "-----BEGIN PRIVATE KEY-----\nkey=\n-----END PRIVATE KEY-----\n"
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
cd script/gcp_organisation
sh terraform-gcp-organisation-create.sh
```
