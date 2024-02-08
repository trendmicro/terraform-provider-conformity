---
page_title: "conformity_custom_rule Resource"
subcategory: "Custom Rules"
description: |-
  Allows you to create and run Custom Rules on Conformity. 
---

# Resource: conformity_custom_rule
Allows you to create Custom Rules on Trend Cloud One<sup>TM</sup> - Conformity.

## Example Usage

### Using a string value for a custom rule creation

```hcl
resource "conformity_custom_rule" "s3_example" {
  name              = "S3 Bucket Name Character Limit"
  description       = "Limit number of characters used to name a S3 Bucket"
  remediation_notes = "Reduce the number of characters for S3 Bucket name"
  service           = "S3"
  resource_type     = "s3-bucket"
  categories        = ["sustainability"]
  severity          = "MEDIUM"
  cloud_provider    = "aws"
  enabled           = true
  attributes {
    name     = "bucketName"
    path     = "data.Name"
    required = true
  }
  rules {
    operation  = "any"
    conditions {
      fact     = "bucketName"
      operator = "pattern"
      value    = "^([a-zA-Z0-9_-]){1,32}$"
    }
  }
}
```

### Using a date comparison for a custom rule condition

```hcl
resource "conformity_custom_rule" "kms_example" {
  name              = "KMS Key Creation Date within 90 days"
  description       = "Check KMS Key was created less than 90 days ago"
  remediation_notes = "Recreate the KMS Key"
  service           = "KMS"
  resource_type     = "kms-key"
  categories        = ["security"]
  severity          = "HIGH"
  cloud_provider    = "aws"
  enabled           = true
  attributes {
    name     = "creationDate"
    path     = "data.CreationDate"
    required = true
  }
  rules {
    operation  = "all"
    conditions {
      fact     = "creationDate"
      operator = "dateComparison"
      value    = jsonencode({"days"=90,"operator"="within"})
    }
  }
}

```

## Argument Reference

* `name` (Required) Name of the custom rule.
* `description` (Required) Description of the custom rule.
* `remediation_notes` (Optional) Notes or steps relevant to remediating the custom rule
* `service` (Required) The cloud provider service name (e.g. S3), a complete list of supported services can be found [here](https://us-west-2.cloudconformity.com/v1/services).
* `resource_type` (Required) - The type of resource this custom rule applies to (e.g. "s3-bucket"), a complete list of supported services can be found [here](https://us-west-2.cloudconformity.com/v1/resource-types).
* `categories` (Required) An array of categories for the custom rule. Can be any of the following: ["security","cost-optimisation", "reliability", "sustainability", "performance-efficiency", "operational-excellence"].
* `severity` (Required) Risk/severity of the custom rule, can be one of the following: "LOW","MEDIUM","HIGH","VERY_HIGH","EXTREME".
* `cloud_provider` (Required) -  Name of the cloud provider (e.g. "aws","azure", "gcp"), a complete list is available from (Conformity Providers Endpoint)[https://us-west-2.cloudconformity.com/v1/providers].
* `enabled` (Optional) Boolean that indicates the status of this rule (true, false). Disabled rules (i.e. set to false) will not be run by Conformity Bot or Real-Time Threat Monitoring (RTM).
* `attributes` (Required) One or more blocks describing the attribute(s) from the resources needed for this rule. The structure of this block is described below.
* `rules` (Required) One or more blocks describing what needs to be checked from the attribute(s). The structure of this block is described below.

An `attributes` block supports the following:
* `name` (Required) User defined to the value of the result of the path query. This value is used as the `fact` input to the rule condition.
* `path` (Required) JSONPath syntax to resource value.
* `required` (Required) Boolean that determines if this data value is required for the rule to run.

A `rule` block supports the following
* `operation` (Required) Operation of the rule. Enum: "any","all"
* `condition` (Required) 
* `event_type` (Required) Description of the result of the rule set.

A `condition` block supports the following
* `fact` (Required) The input value from the corresponding attribute name.
* `operator` (Required) A string value of the operator used to evaluate the input value.
* `path` (Optional) Secondary JSONPath query to apply to further evaluate nested data.
* `value` (Required) The expected value from the JSONPath query. This can be a string, number, boolean, or object.

~> **NOTE:** If the `value` is either a number, boolean, or object. It **must** be encoded using the built-in `jsonencode` function. e.g. **Number**: `value = jsonencode(86400)`, **Boolean**: `value = jsonencode(true)`, **Null**: `value = jsonencode(null)`, **Object**: `value = jsonencode({"days"=20,"operator"="within"})`

A `value` block can be defined using the built-in Terraform `jsonencode()` function but must follow the structure below
* `days` (Required) The number of days to compare against.
* `operator` (Required) Date comparison operator, e.g. "within", "olderThan"
