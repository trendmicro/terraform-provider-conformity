# conformity_communication_setting.sns_setting:
resource "conformity_communication_setting" "sns_setting" {

    sns {
        arn          = "arn:aws:sns:us-west-2:123456789012:CloudConformity"
        channel_name = "cloud_sns"
    }

}

output "sns" {
value = conformity_communication_setting.sns_setting
}

resource "conformity_communication_setting" "email_setting" {

    email {
        users = [
            "urn:tmds:identity:us-east-ds-1:62740:administrator/1915",
        ]
    }
}

output "email" {
value = conformity_communication_setting.email_setting
}
