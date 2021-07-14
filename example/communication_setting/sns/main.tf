resource "conformity_communication_setting" "sns_setting" {
    
    sns {
        arn          = "arn:aws:sns:us-west-2:301854035555:CloudConformity"
        channel_name = "cloud_sns"
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
output "sns_setting" {
    value = conformity_communication_setting.sns_setting
}