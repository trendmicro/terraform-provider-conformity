---
page_title: "conformity_azure_active_directory"
subcategory: "Azure"
description: |-
 Allows you to create an Azure Active Directory
---


# Resource `conformity_azure_active_directory`

 Allows you to create an Azure Active Directory

## Example Usage

```
resource "conformity_azure_active_directory" "azure" {
    name = "Azure Active Directory"
    directory_id    = " directoryId"
    application_id     = " applicationId"
    application_key = " applicationKey"
}
output "ad_detail"{
    value=conformity_azure_active_directory.azure
}

```
