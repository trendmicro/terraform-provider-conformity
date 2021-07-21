resource "aws_cloudformation_stack" "cloud-conformity" {
  name         = "CloudConformity"
  template_url = "https://s3-us-west-2.amazonaws.com/cloudconformity/CloudConformity.template"
  parameters={
    AccountId  = "717210094962"
    ExternalId = "${data.conformity_external_id.external.external_id}"
  }
  capabilities = ["CAPABILITY_NAMED_IAM"]
}