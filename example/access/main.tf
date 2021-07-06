data "conformity_external_id" "external"{}

output "external_id" {
  value = data.conformity_external_id.external.external_id

}
