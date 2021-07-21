# conformity_aws_account.aws:
resource "conformity_aws_account" "aws" {}

output "aws_account" {
    value = conformity_aws_account.aws
}
