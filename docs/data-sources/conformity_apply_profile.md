---
page_title: "conformity_apply_profile - cloudconformity_terraform"
subcategory: "Profile"
description: |-
 Allows you to apply profile and rule settings to a set of accounts under your organisation.
---

# Data Source `conformity_apply_profile`

Allows you to apply profile and rule settings to a set of accounts under your organisation.

Upon applying this data source using terraform plan or terraform apply, Profile will be applied to the accounts in background.

## Example Usage
```terraform
data "conformity_apply_profile" "profile"{

  profile_id = "THE-ID-OF-THE-PROFILE-SETTINGS-THAT-YOU-WANT-TO-APPLY"
  account_ids = ["LIST-OF-ACCOUNT-IDS-WHERE-YOU-WANT-TO-APPLY-THE-PROFILE-SETTINGS"]
  mode = "MODE-THAT-YOU-WANT-TO APPLY"
  notes = "ADD-YOUR-NOTES-HERE"

    # Note: Include and exceptions are only working on `overwrite` mode.
    include {
        exceptions = false
    }
}
```

## Attributes Reference

 - `profile_id` (String) - (Required) The Cloud Conformity ID of the profile.
 - `account_ids` (Array of Strings) - Account IDs that will be configured by the profile.
 - `mode` (String) - (Required) Mode of how the profile will be applied to the accounts, i.e. "fill-gaps", "overwrite" or "replace":
    * fill-gaps - Merge existing settings with this Profile. If there is a conflict, the account's existing setting will be used.
    * overwrite - Merge existing settings with this Profile. If there is a conflict, the Profile's setting will be used.
    * replace - Clear all existing settings and apply settings from this Profile.
 - `include` (String) - (Optional) [Note: this field can only be used in overwrite mode] An object containing rule setting configurations.
    * `exceptions` (String) - (Optional) A boolean value to allow/prevent the account rule setting's exceptions field from being overwritten when applying a profile.






