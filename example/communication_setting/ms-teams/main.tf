resource "conformity_communication_setting" "teams_setting" {
    
    ms_teams {
        channel              = "ms-teams"
        channel_name          = "comformity_teams"
        display_extra_data    = true
        display_introduced_by = true
        display_resource      = true
        display_tags          = true
        url                   = "Your MS Teams Webhook URL"
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
output "teams_setting" {
    value = conformity_communication_setting.teams_setting
}