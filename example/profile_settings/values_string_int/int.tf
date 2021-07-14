# conformity_profile.profile_settings:
resource "conformity_profile" "int_settings" {
    description = "conformity development - int value"
    name        = "conformity-integer-rules"

    # type ttl
    # interger converted to string
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
}