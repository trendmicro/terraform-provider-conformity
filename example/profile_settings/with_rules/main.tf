# conformity_profile.profile_settings:
resource "conformity_profile" "profile_settings" {
    description = "conformity development - rules included"
    name        = "conformity-with-rules-testing"

    # without extra settings 
    # included {
    #     enabled    = false
    #     id         = "EC2-001"
    #     provider   = "aws"
    #     risk_level = "MEDIUM"
    #     exceptions {
    #         tags        = [
    #             "some_tag",
    #             "some_tag2",
    #         ]
    #     }
    # }
    # type ttl
    # integer converted to string
    # included {
    #     enabled    = true
    #     id         = "RTM-002"
    #     provider   = "aws"
    #     risk_level = "MEDIUM"

    #     extra_settings {
    #         countries = false
    #         multiple  = false
    #         name      = "ttl"
    #         regions   = false
    #         type      = "ttl"
    #         value     = "72"
    #     }
    # }
    # type choice-multiple-value
    # map of any type (string, bool)
    included {
        id = "KMS-007" 
        provider = "aws" 
        enabled = true 
        risk_level = "HIGH" 
        extra_settings { 
            name = "ConfigurationChanges" 
            type = "choice-multiple-value" 
            values { 
                value = "TagResource" 
                enabled = false
               
                
                } 
            
            }
        }
}
output "profile" {
  value = conformity_profile.profile_settings
}