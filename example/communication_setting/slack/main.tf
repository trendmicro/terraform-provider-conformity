resource "conformity_communication_setting" "slack_setting" {
    
    slack {
        channel               = "#git-main"
        channel_name          = "Conformity"
        display_introduced_by = true
        display_resource      = true
        display_tags          = true
        display_extra_data    = true
        url                   = "Your Slack Webhook URL"
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
output "slack_setting" {
    value = conformity_communication_setting.slack_setting
}