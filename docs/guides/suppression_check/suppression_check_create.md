---
page_title: Check Suppression Guide"
subcategory: "Check"
description: |-
  Provides instruction on how to  Check Suppression
---

# How To  Check Suppression on a local machine
Provides instruction on how to create an  Azure Active Directory using Terraform.

## Things needed:
1. account_id,rule_id,region,resource_id,note
2. Conformity API Key

#### Step 1

##### Terraform Configuration

1. Create terraform `main.tf` on `example/check_suppression` folder for Check Suppression Module, create the file if not existing and add the following values:
```hcl
resource "conformity_check_suppression" "check"{
        account_id="f7ca8f9d-2b4b-7624-289c-d1d2fe23b54d"
        rule_id="EC2-054"
        region="us-1"
        resource_id="sg-016c1348bdc0454a4"
        note="suppression check"
}
```
2. To use Conformity and its resources, add the  Region where the resources will be created and API Key from Conformity Organisation to the `variable.tf`. 

3. Create `variable.tf` on `example/check_suppression` folder and add the following values:

```hcl
apikey = "CLOUD-CONFORMITY-API-KEY"
region = "REGION"

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
cd example/check_suppression
terraform init
terraform plan
terraform apply
```

