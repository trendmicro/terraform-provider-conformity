---
page_title: "conformity_external_id"
subcategory: "Data Source"
description: |-
  Provides an external ID for your Cloud Conformity organisation.
---

# Data Source `conformity_external_id`

Provides an external ID for your Cloud Conformity organisation. When using this data source `apikey` and `region` must be provided as terraform variables

## Example Usage
```hcl
data "conformity_external_id" "all"{}

output "external_id" {
  value = data.conformity_external_id.all.external_id
}

resource "conformity_aws_account" "aws"{

    name        = "Trendmicro Account"
    environment = "Staging"
    external_id = "${data.conformity_external_id.all.external_id}"
    role_arn    = "..."
}
```

## Attributes Reference

 - `external_id` - The external ID.
