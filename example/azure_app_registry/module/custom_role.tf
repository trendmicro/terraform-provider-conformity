resource "azurerm_role_definition" "custom_role" {
  name        = "Custom Role - Cloud One Conformity"
  scope       = data.azurerm_subscription.primary.id
  description = "Subscription level custom role for Cloud One Conformity access."

  permissions {
    actions     = ["Microsoft.AppConfiguration/configurationStores/ListKeyValue/action",
                    "Microsoft.Network/networkWatchers/queryFlowLogStatus/action",
                    "Microsoft.Web/sites/config/list/Action",
                    "Microsoft.Storage/storageAccounts/queueServices/queues/read"]
    not_actions = []
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id, # /subscriptions/00000000-0000-0000-0000-000000000000
  ]
}


resource "azurerm_role_assignment" "role_assignment_reader" {
  scope                = data.azurerm_subscription.primary.id
  role_definition_name = "Reader"
  principal_id         = data.azuread_service_principal.service_principal.object_id
}


resource "azurerm_role_assignment" "role_assignment_conformity_role" {
  scope              = data.azurerm_subscription.primary.id
  role_definition_id = azurerm_role_definition.custom_role.role_definition_resource_id
  principal_id       = data.azuread_service_principal.service_principal.object_id
}

