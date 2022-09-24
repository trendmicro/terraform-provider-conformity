---
page_title: "Custom Rule Guide"
subcategory: "Azure"
description: |-
  Provides instruction on how to create an Custom Rule
---

# How To create an Azure Active Directory on a local machine
Provides instruction on how to create an  Custom Rule using Terraform.

## Things needed:
1. name,description,remediation_notes,service,resource_type,categories,severity,cloud_provider,enabled,attributes{name,path,required},rules{operation,conditions{fact,operator,value}, event_type}
2. Conformity API Key


#### Step 1

##### Terraform Configuration

1. Create terraform `main.tf` on `example/custom_rule` folder for Custom Rule Module, create the file if not existing and add the following values:
```hcl
resource "conformity_custom_rule" "example"{
    name= "S3 Bucket Custom Rule"
    description      = "This custom rule ensures S3 buckets follow our best practice updated"
	remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	service          = "S3"
	resource_type     = "s3-bucket"
	categories       = ["security"]
	severity         = "HIGH"
	cloud_provider   = "azure"
	enabled          = true
	attributes {
		name     = "bucketName"
		path     = "data.Name"
		required = true
	  }
	  rules {
	    operation = "all"
		conditions {
		  fact     = "bucketName"
		  operator = "pattern"
		  value    = "^([a-zA-Z0-9_-]){1,32}$"
		}
		event_type = "Bucket name is longer than 32 characters"
	  }
}
```
2. To use Conformity and its resources, add the  Region where the resources will be created and API Key from Conformity Organisation to the `variable.tf`. 

3. Create `variable.tf` on `example/custom_rule` folder and add the following values:

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
cd example/custom_rule
terraform init
terraform plan
terraform apply
```
