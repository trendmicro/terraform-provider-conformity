---
page_title: "Create An Invite SSO User To Your Organisation - cloudconformity_terraform"
subcategory: "Users"
description: |-
  Provides instruction on how to create an invite SSO user to your organisation using Terraform. This resource is not applicable to users who are part of the Cloud One Platform.
---

# How To Create An Invite SSO User To Your Organisation
Provides instruction on how to create an invite SSO user to your organiation using Terraform. This resource is not applicable to users who are part of the Cloud One Platform.

## Things needed:
1. Cloud Conformity API Key
2. Account ID

To get the AWS Account ID:
Open Conformity Admin Web console > navigate to Account dashboard > choose the AWS account.
Notice the URL, you should be able to get the account ID e.g. https://cloudone.trendmicro.com/conformity/account/account-ID

#### Step 1

##### Terraform Configuration

1. To create an SSO Invite, create `terraform.tfvars` on `example/user/sso_user` folder and add the following values:

## Example Usage for `terraform.tfvars`
```
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "AWS-ACCOUNT-REGION"

# conformity_sso_user
first_name = "John"
last_name  = "Doe"
email      = "john.doe@cloudconformity.com"
role       = "USER"

# access_list01 (can be multiple)
#level can be "NONE" "READONLY" "FULL"
account01 = "cloud-conformity-account-access"
level01  = "ADD-LEVEL"

account02 = "cloud-conformity-account-access"
level02  = "ADD-LEVEL"
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
cd example/sso_user/user
terraform init
terraform plan
terraform apply
```
#### 5. Bash script that can run to automate the whole process 1-5:

Under script folder run:
```sh
cd script/sso_user/user
sh terraform-sso_user-create.sh
```