resource "conformity_aws_account" "aws" {
    name        = "aws-conformity"
    environment = "development"
    role_arn    = "${aws_cloudformation_stack.cloud-conformity.outputs["CloudConformityRoleArn"]}"
    external_id = data.conformity_external_id.external.external_id
    tags = ["development"]
    
    settings {
          rule {
              rule_id = "S3-021"

              settings {
                  enabled     = false
                  risk_level  = "HIGH"
                  rule_exists = false
                }
            }
        // implement value
        rule {
            rule_id = "RDS-018"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                exceptions {
                    tags        = [
                        "mysql-backups",
                    ]
                }
                extra_settings {
                    name    = "threshold"
                    type    = "single-number-value"
                    value   = 90.90
                }
            }
        }
        // implement multiple values
        rule {
            rule_id = "SNS-002"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                exceptions {
                    tags        = [
                        "some_tag",
                    ]
                }
                extra_settings {
                    name    = "conformityOrganization"
                    type    = "choice-multiple-value"
                    values {
                        enabled = false
                        label   = "All within this Conformity organization"
                        value   = "includeConformityOrganization"
                    }
                    values {
                        enabled = true
                        label   = "All within this AWS Organization"
                        value   = "includeAwsOrganizationAccounts"
                    }
                }
            }
        }
        // implement regions
        rule {
            rule_id = "RTM-008"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                extra_settings {
                    name    = "authorisedRegions"
                    regions = [
                        "ap-southeast-2",
                        "eu-west-1",
                        "us-east-1",
                        "us-west-2",
                    ]
                    type    = "regions"
                }
            }
        }
        // implement multiple_object_values
        rule {
            rule_id = "RTM-011"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                extra_settings {
                    name    = "patterns"
                    type    = "multiple-object-values"
                    multiple_object_values {
                        event_name         = "^(iam.amazonaws.com)"
                        event_source       = "^(IAM).*"
                        user_identity_type = "^(Delete).*"
                    }
                }
            }
        }
        // implement mappings
        rule {
            rule_id = "VPC-013"
            settings {
                enabled     = true
                risk_level  = "LOW"
                rule_exists = false
                extra_settings {
                    name    = "SpecificVPCToSpecificGatewayMapping"
                    type    = "multiple-vpc-gateway-mappings"
                    // can be multiple mappings
                    mappings {
                        // can be multilple value
                        // if mappings is declared, values is required
                        values {
                            // name is required
                            // type is required
                            name = "gatewayIds"
                            type = "multiple-string-values"
                            // can be one of this value/values
                            values {
                                // value is required
                                // validation value should start with nat-
                                value = "nat-001"
                            }
                            values {
                                value = "nat-002"
                            }
                        }
                        values {
                            name  = "vpcId"
                            type  = "single-string-value"
                             // can be one of this value/values
                             // validation value should start with vpc-
                            value = "vpc-001"
                        }
                    }
                }
            }
        }
    }
}
