resource "conformity_gcp_account" "gcp" {
    subscription_id = var.gcp_subscription_id
    name            = var.gcp_name
    environment     = var.gcp_environment
    ServiceAccountUniqueId = var.projectId
    ServiceAccountUniqueId = var.ProjectName
    ServiceAccountUniqueId = var.ServiceAccountUniqueId
    settings {
        // implement multiple string values
        rule {
            rule_id = "Resources-001"
            settings {
                enabled     = true
                risk_level  = "MEDIUM"
                rule_exists = false
                exceptions {
                    filter_tags = []
                    resources   = []
                    tags        = [
                        "some_tag",
                    ]
                }
                extra_settings {
                    name    = "tags"
                    type    = "multiple-string-values"
                    values {
                        value = "Environment"
                    }
                    values {
                        value = "Role"
                    }
                    values {
                        value = "Owner"
                    }
                    values {
                        value = "Name"
                    }
                }
            }
        }
    }
}
