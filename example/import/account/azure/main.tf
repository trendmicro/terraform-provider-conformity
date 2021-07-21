# conformity_azure_account.azure:
resource "conformity_azure_account" "azure" {}

output "azure_account" {
    value = conformity_azure_account.azure
}
