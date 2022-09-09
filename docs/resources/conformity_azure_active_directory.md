---
page_title: "conformity_azure_active_directory"
subcategory: "Azure"
description: |-
  Provides an  instruction on how to create an Azure Active Directory.
---

# Resource `conformity_azure_active_directory`
Provides an  instruction on how to create an Azure Active Directory.

## Example Usage
```hcl
resource "conformity_azure_account" "azure" {
    name                = "azure-conformity"
    environment         = "development"
    active_directory_id = "your-active-directory-id"
    subscription_id     = "your-subscription-id"
}
```
