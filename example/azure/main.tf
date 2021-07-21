resource "conformity_azure_account" "azure" {
    subscription_id = "YOUR-SUBSCRIPTION-ID-FROM-IMPORT"
    name            = var.azure_name
    environment     = var.azure_environment
    active_directory_id = var.azure_active_directory_id
    settings {
        // implement multiple string values
        rule {
            rule_id = "Resources-001"
            note    = "test note"
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
