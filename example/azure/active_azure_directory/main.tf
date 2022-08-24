resource "conformity_azure_active_directory" "azure" {
    name = "Azure Active Directory"
    directory_id    = "761d49c9-8898-5d35-c4db-ed8582f20dbf"
    application_id     = var.application_id
    application_key = var.application_key
}
