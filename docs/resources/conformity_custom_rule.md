---
page_title: "conformity_custom_rule Resource"
subcategory: "Custom Rules"
description: |-
  Allows you to create Custom Rules on Conformity. 
---

# Resource `conformity_custom_rule`
Allows you to create Custom Rules on Conformity

## Example Usage
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
output "customrule" {
  value = conformity_custom_rule.example
}
```

## Argument reference

- `name` (String) -  Name of the custom rule.
- `description` (String) -  description of the custom rule.
- `remediation_notes` (String) - remediation_notes of the custom 
- `service` (String) - service  of the custom rule 
- `resource_type` (String) - resource type of the custom rule
- `categories` (Array of String) -  categories of the custom rule. Enum: ["security", "sustainability", "performance-efficiency", "operational-excellence"]
- `severity` (String) - severity of the custom rule. Enum :"LOW","MEDIUM","HIGH","VERY_HIGH","EXTREME"
- `cloud_provider` (String ) -  Name of the cloud provider. Enum: "aws","azure","gcp".
- `enabled` (Bool) - This attributes determines whether this setting enabled or not (true ,false)

- `attributes` List: Can be multiple declaration 
    * `name` (String) BucketName.
    * `path` (String) Path of the Bucket.
    * `required` (String) This  determines whether the attribute is required or not. 

-  `rule` List: Can be multiple declaration
    * `operation` (String) -  operation of the rule. Enum: "any","all"
    * `condition` List: Can be multiple declaration
        * `fact` (String) - BucketName
        * `operation` (String) - pattern
        *  `value` (String) - value  of the operator
    * `event_type` (String) - Message
