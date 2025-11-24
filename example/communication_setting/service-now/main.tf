resource "conformity_communication_setting" "service_now_setting" {

    service_now {
        channel_name = "servicenow-002"
        type         = "incident"
        url          = "https://myservicenowinstance.service-now.com"

        username = "admin"
        password = "123456"

        assignee = "admin"

        impact  = "1" # 1
        urgency = "3" # 3

        close_code  = "Closed/Resolved by Caller"
        close_notes = "Issue resolved by"

        creation_override = {
        urgency  = "2"
        priority = "1"
        }

        resolution_override = {
        close_code  = "Closed by Caller"
        close_notes = "Issue closed"
        }
    }

    filter {
        services = ["EC2", "S3", "Lambda"]
        categories = [
        "security",
        ]
        tags = [
        "service-now-test",
        ]
    }

    manual = true

    relationships {
        account {
            id = "80b880c9-336a-490d-b212-4e847956a62d"
        }
        organisation {
            id = "102642678400"
        }
    }
}
output "service_now_setting" {
    value = conformity_communication_setting.service_now_setting
}
