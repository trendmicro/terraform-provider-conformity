resource "conformity_custom_rule" "example" {
   
    name= "Bucket testing for custom rule"
    description      = "This custom rule ensures S3 buckets follow our best practice updated"
	remediation_notes = "If this is broken, please follow these steps:\n1. Step one \n2. Step two\n updated"
	service          = "S3"
	resource_type     = "s3-bucket"
	categories       = ["security"]
	severity         = "HIGH"
	cloud_provider   = "aws"
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
		  value=  "^shunyekaa$"
		}
		event_type = "Bucket name should be shunyeka"
	  }
}

output "customrule"{
    value=conformity_custom_rule.example
}

