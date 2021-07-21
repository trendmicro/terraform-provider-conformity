data "azurerm_subscription" "primary" {}

data "azuread_client_config" "current" {}

data "azuread_application" "app_registration" {
  display_name = azuread_application.app_registration.display_name
}
data "azurerm_subscription" "current" {}

data "azuread_service_principal" "service_principal" {
  object_id = azuread_service_principal.service_principal.object_id
}
data "azuread_service_principal" "service_principal_application_id" {
  application_id = azuread_service_principal.service_principal.application_id
}
data "azurerm_client_config" "test" {}