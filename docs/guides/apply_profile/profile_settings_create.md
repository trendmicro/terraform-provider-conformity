---
page_title: "Apply profile to accounts Guide - cloudconformity_terraform"
subcategory: "Profile Settings"
description: |-
  Provides overview on how apply profile to accounts on Cloud Conformity using Terraform.
---

# Profile Settings on Cloud Conformity on a local machine
  Provides overview on how apply profile to accounts on Cloud Conformity using Terraform.

## Things needed for this Data Source:
1. Cloud Conformity API Key

#### Use Cases
Applying profile settings to account uses a data source.

##### Applying profile settings to account with data source and resource.
This example has a resource for applying profile settings to an account. As you can see, the profile_id on the data source is getting the profile_id of the conformity_profile resource. For the account_ids, you can add multiple accounts_ids where you want to apply the profile settings.

To get the Account ID:
Open Conformity Admin Web console > navigate to Account dashboard > choose the account you would like to apply the profile settings.
Notice the URL, you should be able to get the account ID e.g. https://cloudone.trendmicro.com/conformity/account/account-ID

Example Usage:

```
resource "conformity_profile" "profile_settings"{
  name = "cloudready-without-rules"
  description = "conformity guardrail development - without included"
  }

data "conformity_apply_profile" "profile"{
  profile_id = conformity_profile.profile_settings.profile_id
  account_ids = ["LIST-OF-ACCOUNT-IDS-WHERE-YOU-WANT-TO-APPLY-THE-PROFILE-SETTINGS"]
  mode = "MODE-THAT-YOU-WANT-TO APPLY"
  notes = "ADD-YOU-NOTES-HERE"

    # Note: Include and exceptions are only working on `overwrite` mode.
    include {
        exceptions = false
    }
}
output "data_profile" {
  value = data.conformity_apply_profile.profile
}
```

##### Applying profile settings to account with data source only.
This example has a data source only for applying profile settings to an account. As you can see, the profile_id will be from the the outputs upon creation or on the link on cloud conformity console. For the account_ids, you can add multiple accounts_ids where you want to apply the profile settings.

To get the Account ID:
Open Conformity Admin Web console > navigate to Account dashboard > choose the account you would like to apply the profile settings.
Notice the URL, you should be able to get the account ID e.g. https://cloudone.trendmicro.com/conformity/account/account-ID

To get the profile_id, you just need to navigate to the "Profiles" and look for the profile setting that you want to apply the profile settings on Cloud Conformity console and get the profile_id in the URL e.g. https://cloudone.trendmicro.com/conformity/profiles/profile:{profile_id}.

Example Usage:

```terraform
data "conformity_apply_profile" "profile"{

  profile_id = "THE-ID-OF-THE-PROFILE-SETTINGS-THAT-YOU-WANT-TO-APPLY"
  account_ids = ["LIST-OF-ACCOUNT-IDS-WHERE-YOU-WANT-TO-APPLY-THE-PROFILE-SETTINGS"]
  mode = "MODE-THAT-YOU-WANT-TO APPLY"
  notes = "ADD-YOU-NOTES-HERE"

    # Note: Include and exceptions are only working on `overwrite` mode.
    include {
        exceptions = false
    }
}
```
Note: When running terraform plan with data source only will result to applying the changes without running terraform apply.

##### Modes on applying profile setting to one or more accounts

This endpoint allows you to apply profile and rule settings to a set of accounts under your organisation.

fill-gaps -	Merge existing settings with this Profile. If there is a conflict, the account's existing setting will be used.
overwrite	- Merge existing settings with this Profile. If there is a conflict, the Profile's setting will be used.
replace	- Clear all existing settings and apply settings from this Profile.

When using the `overwrite` mode when applying a profile to account/s, there might be data on an account's rule settings that you want to retain, e.g. You want to replace the enabled, extraSettings and riskLevel but wish to keep the exceptions.