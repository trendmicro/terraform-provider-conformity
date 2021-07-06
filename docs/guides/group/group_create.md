---
page_title: "Create or Add Groups and Tags on Cloud Conformity - cloudconformity_terraform"
subcategory: "Groups"
description: |-
  Provides instruction on how to create or add groups and tags on Cloud Conformity using Terraform.
---

# How To Create or Add Groups and Tags on Cloud Conformity
Provides instruction on how to create or add groups and tags on Cloud Conformity using Terraform.

## Things needed:
1. Cloud Conformity API Key

#### Step 1

##### Terraform Configuration

1. To add groups and tags on Cloud Conformity, create `terraform.tfvars` on `example/group` folder and add the following values:
   
## Example Usage for `terraform.tfvars`
```
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "AWS-ACCOUNT-REGION"

name_group1 = "NAME-OF-GROUP1"
tag_group1 = ["NAME-OF-GROUP1-TAG1","NAME-OF-GROUP1-TAG2"]

name_group2 = "NAME-OF-GROUP2"
tag_group2 = ["NAME-OF-GROUP2-TAG1", "NAME-OF-GROUP2-TAG2"]
```
Note: You can always change the values declared according to your choice.

#### Step 2

##### Run Terraform

#### 1. Navigate to project directory:
```sh
cd /path/guardrail-conformity-tf-provider
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
cd example/group
terraform init
terraform plan
terraform apply
```
#### 5. Bash script that can run to automate the whole process 1-5:

Under script folder run:
```sh
cd script/group
sh terraform-group-create.sh
```