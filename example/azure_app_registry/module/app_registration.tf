resource "azuread_application" "app_registration" {
  display_name     = "Conformity Azure Access"
  owners           = [data.azuread_client_config.current.object_id]
  sign_in_audience = "AzureADMyOrg"

  required_resource_access {
    resource_app_id = "00000003-0000-0000-c000-000000000000"

    resource_access {
      id   = "e1fe6dd8-ba31-4d61-89e7-88639da4683d"
      type = "Scope"
    }

    resource_access {
      id   = "a154be20-db9c-4678-8ab7-66f6cc099a59"
      type = "Scope"
    }

    resource_access {
      id   = "df021288-bdef-4463-88db-98f22de89214"
      type = "Role"
    }
    resource_access {
      id   = "7ab1d382-f21e-4acd-a863-ba3e13f7da61"
      type = "Role"
    }
  }

    required_resource_access {
    resource_app_id = "00000002-0000-0000-c000-000000000000"

    resource_access {
      id   = "311a71cc-e848-46a1-bdf8-97ff7156d8e6"
      type = "Scope"
    }

    resource_access {
      id   = "c582532d-9d9e-43bd-a97c-2667a28ce295"
      type = "Scope"
    }

    resource_access {
      id   = "5778995a-e1bf-45b8-affa-663a9f3f4d04"
      type = "Scope"
    }
    resource_access {
      id   = "5778995a-e1bf-45b8-affa-663a9f3f4d04"
      type = "Role"
    }
  }

}

resource "azuread_application_password" "client_secret" {
  display_name = "Conformity Azure Access"
  application_object_id = azuread_application.app_registration.object_id
  end_date = coalesce(var.end_date, timeadd(timestamp(), "8760h"))
}

resource "azuread_service_principal" "service_principal" {
  application_id               = azuread_application.app_registration.application_id
}
