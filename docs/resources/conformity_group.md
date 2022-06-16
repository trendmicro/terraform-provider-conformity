---
page_title: "conformity_aws_account Resource"
subcategory: "Groups"
description: |-
  Provides a Conformity Group Management.
---

# Resource `conformity_group`
Provides a Conformity Group Management.

## Example Usage
```hcl
resource "conformity_group" "group" {

    name        = "conformity-group"
    tags        = ["tag1", "tag2"]
}
```

## Argument reference
 - `name` (String) - (Required) The name of your account.
 - `tags` (String) - (Required) The tag name that you want to add.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

 - `id` - The ID of the Group in Conformity managed by this resource

Example usage on the template:

```hcl
group {
    id = conformity_group.group.id
}
```

## Import
Conformity group - Can import based on the Group ID that can be found under the Conformity web console.
To get the Group ID same with the AWS Account ID above, you just need to navigate to the Group you want to import and get the
Group ID in the URL e.g. https://cloudone.trendmicro.com/conformity/group/group-ID

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_group resources into Terraform. Replace the {GROUP-ID} value.
```hcl
terraform import conformity_group.group_1 {GROUP-ID}
```

Use the `state` subcommand to verify Terraform successfully imported the conformity_group resources.
```hcl
terraform state show conformity_group.group_1
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Make sure you remove the id from the imported resource.
```hcl
terraform show -no-color >> main.tf
```
