---
page_title: "Update Groups and/or Tags on Cloud Conformity"
subcategory: "Groups"
description: |-
  Provides instruction on how to add more tags or group or update their name on Cloud Conformity using Terraform.
---

# How To Update Groups and/or Tags on Cloud Conformity
Provides instruction on how to add more tags or group or update their name on Cloud Conformity using Terraform.

#### Step 1

##### Run Terraform

#### 1. Navigate to folder group directory:
```sh
cd /path/terraform-provider-conformity/example/group
```
#### 2. Edit terraform.tfvars values according to the changes you want.

#### 3. Run terraform apply to apply the changes.
```sh
terraform apply
```
#### 3. Bash script that can run to automate the whole process 1-3. Go to script/group folder and look for `sh terraform-group-update.sh`. Find this and edit it according to the update that you want.

```sh
cat << EOF >> update.tfvars
name_group1 = "UPDATED_NAME_GROUP1"
tag_group1 = ["UPDATED_TAG1_GROUP1","UPDATED_TAG2_GROUP1"]

name_group2 = "UPDATED_NAME_GROUP2"
tag_group2 = ["UPDATED_TAG1_GROUP2", "UPDATED_TAG2_GROUP2"]
EOF
```

After editing the `sh terraform-group-update.sh` file, Run this command:
```sh
cd script/group
sh terraform-group-update.sh
```

