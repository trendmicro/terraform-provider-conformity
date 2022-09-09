resource "conformity_custom_rule" "example" {
   
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

