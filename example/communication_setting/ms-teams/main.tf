resource "conformity_communication_setting" "teams_setting" {
    
    ms_teams {
        channel              = "ms-teams"
        channel_name          = "comformity_teams"
        display_extra_data    = true
        display_introduced_by = true
        display_resource      = true
        display_tags          = true
        url                   = "https://cloud.webhook.office.com/webhookb2/b6467af9-3d96-4f37-8a69-0edf527af2f4@795cfe74-f75f-45c1-ba46-5fb6590cc562/IncomingWebhook/15d0c241a569487ebbc1ae02bbc3cedc/8dd3ce28-2130-4336-bb0a-9a1989f618f0"
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