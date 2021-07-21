output "azuread_client_config"{
  value = data.azuread_client_config.current.object_id
}
output "app_registration_application_id" {
  value = data.azuread_application.app_registration.application_id
}
output "active_directory_tenant_id"{
  value = data.azurerm_subscription.current.tenant_id
}
output "current_subscription_display_name" {
  value = data.azurerm_subscription.current.display_name
}
output "azurerm_subscription_id"{
  value = data.azurerm_subscription.current.subscription_id
}
output "azurerm_subscription_scope" {
  value = data.azurerm_subscription.primary.id
}
output "app_registration_key" {
  value = azuread_application_password.client_secret.value
  sensitive = true
}
output "service_principal_object_id" {
  value = data.azuread_service_principal.service_principal.object_id
}

output "service_principal_application_id" {
  value = data.azuread_service_principal.service_principal.application_id
}

