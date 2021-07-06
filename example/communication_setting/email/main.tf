resource "conformity_communication_setting" "email_setting" {

    email {
        users = [
        "urn:tmds:identity:us-east-ds-1:62740:administrator/1915",
        ]
    }

    filter {
        categories  = [
        "security",
        ]
        compliances = [
        "FEDRAMP",
        ]
        filter_tags = [
        "tagKey",
        ]
        regions     = [
        "ap-southeast-1",
        ]
        risk_levels = [
        "MEDIUM",
        ]
        rule_ids    = [
        "S3-016",
        ]
        services    = [
        "EC2",
        "IAM",
        ]
        tags        = [
        "tagName",
        ]
    }
    relationships {
        account {
            id = "80b880c9-336a-490d-b212-4e847956a62d"
        }
        organisation {
            id = "102642678400"
        }
    }
}
output "email_setting" {
    value = conformity_communication_setting.email_setting
}