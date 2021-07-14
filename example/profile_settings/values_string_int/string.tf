# conformity_profile.profile_settings:
resource "conformity_profile" "string_settings" {
    description = "conformity development - string value"
    name        = "conformity-string-rules"

    # type ttl
    # interger converted to string
    included {
        enabled    = true
        id         = "CW-001"
        provider   = "aws"
        risk_level = "MEDIUM"
        extra_settings {
            countries = false
            multiple  = false
            name      = "alarmName"
            regions   = false
            type      = "single-string-value"
            value     = "BillingAlarm"
        }
    }
}