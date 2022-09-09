---
page_title: "conformity_azure_subscription"
subcategory: "Azure"
description: |-
 Allows you to list Azure Subscriptions from an onboarded Azure Active Directory
---


# Data Source `conformity_azure_subscriptions`

Allows you to list Azure Subscriptions from an onboarded Azure Active Directory

## Example Usage

resource "conformity_azure_account" "azure" {

    subscription_id = The id which you want to find Azure Subscription Account
    
}