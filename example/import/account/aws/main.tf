# conformity_aws_account.aws:
resource "conformity_aws_account" "aws" {
    name = var.account_tag
    environment = var.environment_tag
    role_arn = var.role_arn
    external_id = var.external_id
}

output "aws_account" {
    value = conformity_aws_account.aws
}
