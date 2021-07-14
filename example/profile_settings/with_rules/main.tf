# conformity_profile.profile_settings:
resource "conformity_profile" "profile_settings" {
    description = "conformity guardrail development - rules included"
    name        = "conformity-with-rules"

    # without extra settings 
    included {
        enabled    = false
        id         = "EC2-001"
        provider   = "aws"
        risk_level = "MEDIUM"
        exceptions {
            tags        = [
                "some_tag",
                "some_tag2",
            ]
        }
    }
    # type ttl
    # integer converted to string
    included {
        enabled    = true
        id         = "RTM-002"
        provider   = "aws"
        risk_level = "MEDIUM"

        extra_settings {
            countries = false
            multiple  = false
            name      = "ttl"
            regions   = false
            type      = "ttl"
            value     = "72"
        }
    }
    # type choice-multiple-value
    # map of any type (string, bool)
    included {
        enabled    = true
        id         = "SNS-002"
        provider   = "aws"
        risk_level = "MEDIUM"

        extra_settings {
            countries = false
            multiple  = false
            name      = "conformityOrganization"
            regions   = false
            type      = "choice-multiple-value"
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
output "profile" {
  value = conformity_profile.profile_settings
}