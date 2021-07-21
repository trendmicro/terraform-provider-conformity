---
page_title: "cloudconformity_external_id - cloudconformity_terraform"
subcategory: "Data Source"
description: |-
  Provides an external ID for your Cloud Conformity organisation.
---

# Data Source `cloudconformity_external_id`

Provides an external ID for your Cloud Conformity organisation.

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

 - `id` (String) - (Required) The external ID.
