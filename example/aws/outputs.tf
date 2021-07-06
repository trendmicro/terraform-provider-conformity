output "external_id" {
  value = data.conformity_external_id.external.external_id
}
output "role_arn" {
  value = "${aws_cloudformation_stack.cloud-conformity.outputs["CloudConformityRoleArn"]}"
}
output "aws_account_name"{
  value = conformity_aws_account.aws.name
}
output "aws_tags"{
  value = conformity_aws_account.aws.tags
}
output "aws_role_arn"{
  value = conformity_aws_account.aws.role_arn
}
output "aws_environment"{
  value = conformity_aws_account.aws.environment
}
