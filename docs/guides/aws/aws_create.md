---
page_title: "Create Account Guide - cloudconformity_terraform"
subcategory: "Accounts"
description: |-
  Provides instruction on how to create or add AWS account on Cloud Conformity using Terraform.
---

# How To Create or Add AWS Account on Cloud Conformity on a local machine
Provides instruction on how to create or add AWS account on Cloud Conformity using Terraform.

## Things needed:
1. AWS Access Key and Secret Access Key
2. Cloud Conformity API Key

#### Step 1

##### Terraform Configuration

1. To use Cloud Conformity and its resources, add the AWS Access Key, Secret Access Key, and Region where the resources will be created and API Key from Cloud Conformity Account to the `terraform.tfvars`. 

2. Create `terraform.tfvars` on `example/aws` folder and add the following values:

```
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "AWS-ACCOUNT-REGION"
access_key="ACCESS-KEY"
secret_key="SECRET-ACCESS-KEY"
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
cd example/aws
terraform init
terraform plan
terraform apply
```
#### 5. Bash script that can run to automate the whole process 1-5:

Under script folder run:
```sh
cd script/aws
sh terraform-aws-create.sh
```
