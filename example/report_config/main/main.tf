resource "conformity_report_config" "report" {
  // optional | type: string
  // if you want to create account-level report config, uncomment account_id value and provide account ID on the terraform.tfvars
  // if you want to create group-level report config, uncomment group_id value and provide account ID on the terraform.tfvars
  // if you want to create organisation-level report config, leave account_id and group_id commented.
   account_id = var.account_id
  // group_id = var.group_id

  configuration { 
    // optional | type: bool | default: false
    scheduled = "true"
    // optional | type: array of strings
    emails = ["conformity@cloudone.com","cloudocnformity@cloudone.com"] 
    // optional | type: string
    frequency = "* * *"
    // required | type: string 
    title = "Conformity Report Config"
    // optional | type: string
    tz = "Asia/Manila"
  }
  filter { 
  }

}