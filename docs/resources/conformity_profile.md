---
page_title: "conformity_profile Resource"
subcategory: "Profile Settings"
description: |-
  Allows you to create Profile Settings on Conformity. 
---

# Resource `conformity_profile`
Allows you to create Profile Settings on Conformity. 

## Example Usage
```hcl
# conformity_profile.profile_settings:
resource "conformity_profile" "profile_settings" {
    description = "conformity development - rules included"
    name        = "cloud-with-rules"

    # without extra settings 
    included {
        enabled    = false
        id         = "EC2-001"
        provider   = "aws"
        risk_level = "MEDIUM"
        exceptions {
            filter_tags = []
            resources   = []
            tags        = [
                "some_tag",
                "some_tag2",
            ]
        }
    }
    # type ttl
    # integer converted to string
    included {
        enabled    = true
        id         = "RTM-002"
        provider   = "aws"
        risk_level = "MEDIUM"
        exceptions {
            filter_tags = []
            resources   = []
            tags        = []
        }
        extra_settings {
            countries = false
            multiple  = false
            name      = "ttl"
            regions   = false
            type      = "ttl"
            value     = "72"
        }
    }
    # type choice-multiple-value
    # map of any type (string, bool)
    included {
        enabled    = true
        id         = "SNS-002"
        provider   = "aws"
        risk_level = "MEDIUM"
        exceptions {
            filter_tags = []
            resources   = []
            tags        = []
        }
        extra_settings {
            countries = false
            multiple  = false
            name      = "conformityOrganization"
            regions   = false
            type      = "choice-multiple-value"
            values {
                enabled = false
                label   = "All within this Conformity organization"
                value   = "includeConformityOrganization"
            }
            values {
                enabled = true
                label   = "All within this AWS Organization"
                value   = "includeAwsOrganizationAccounts"
            }
        }
    }
}
output "profile" {
  value = conformity_profile.profile_settings
}
```

## Argument reference

 - `name` (String) - (Optional) Name of the profile.
 - `description` (String) - (Optional) description of the profile.
 - `provider` (String) - (Optional) Name of the cloud provider. Enum: "aws" "azure".

 - `included` - (Optional) List: (Can be multiple declaration)
     * `id` (String) - (Optional) Profile ID.
     * `enabled` (Bool) - (Optional) This attribute determines whether this setting is enabled.
     * `risk_level` (String) - (Optional) Risk level of the Conformity rule. Enum: "LOW" "MEDIUM" "HIGH" "VERY_HIGH" "EXTREME".
     * `profile_id` (String) - (Optional) profile_id for save rule settings to an existing Profile. To add a batch of configured rule settings to an empty profile or overwrite existing rule settings and profile details. Check the import description to know how to get the profile_id.

  Inside `included` there will be `execeptions` set.
 
 - `exceptions` - (Optional) List: 
     * `filter_tags` (Array of Strings)- (Optional) An array of resource tags, resource tag keys or resource tag values that are exempted from the rule when it runs, e.g filterTags ["dev"] will exempt resource with tag "environment::dev from the rule".
     * `resources` (Array of Strings) - (Optional) An array of resource IDs that are exempted from the rule when it runs.
     * `tags` (Array of Strings) - (Optional) An array of resource tags that are exempted from the rule when it runs.

  Inside `included` there can be multiple declaration of `extra_settings` set.

 - `extra_settings` - (Optional) List: (Can be multiple declaration)
     * `countries` (Bool) - (Optional) Rule specific property.
     * `multiple` (Bool) - (Optional) Rule specific property.
     * `name` (String) - (Optional) (Keyword) Name of the extra setting.
     * `regions` (Bool) - (Optional) Rule specific property.
     * `type` (String) - (Required) Rule specific property. Values can be: "multiple-string-values", "multiple-number-values" "multiple-aws-account-values", "choice-multiple-value" "choice-single-value", "single-number-value", "single-string-value", "ttl", "single-value-regex", "countries", "multiple-ip-values" and "tags".
     * `value` (String) - (Optional) Customisable value for rules that take on single name/value pairs.
  
  Inside `extra_settings` there can be multiple declaration of `values` set.
  
 - `values` - (Optional) List: (Can be multiple declaration). An array (sometimes of objects) rules that take on a set of of values
     *  `enabled` (Bool) - (Optional) Defines if the checkbox is enabled or not.
     *  `label` (String) - (Optional) Internal key.
     *  `value` (String) - (Required) Description of the checkbox.

        Note: There is a condition for `type` attribute. If the specified is attribute is `value`, the possible values are "single-number-value", "single-string-value", "single-value-regex" and "ttl". If the specified is attribute is `values`, the declaration of it is inside the extra settings which can be a list and the possible values are "choice-multiple-value", "choice-single-value", "multiple-string-values", "multiple-number-values", "countries", "multiple-ip-values", "multiple-aws-account-values" and "tags". You cannot declare both `values` and `value` at the same time.See the table below:

| type     | possible value                                                                                                                | Sample declaration                                                                                                                                                                                                                    |
|----------|-------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `value`  | single-number-value, single-string-value, single-value-regex, ttl                                                             | included {     ….     exceptions {       ….     }       extra_settings {		 	….                     type = "ttl"         value = "72"       } }                                                                                           |
| `values` | choice-multiple-value, choice-single-value, multiple-string-values, multiple-number-values, countries, multiple-ip-values, multiple-aws-account-values, tags | included {     ….     exceptions {       ….     }       extra_settings {		 	….                     type = "choice-multiple-value"           values {             ….           }           values {             ….            }       } } |

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

 - `id` - The ID of the Profile Setting in Conformity managed by this resource

Example usage on the template:

```hcl
profile {
    id = conformity_profile.profile_settings.id
}
```

## Import
Profile Settings - Can import based on the profile_id that can be found under the outputs upon creation or on the link on Conformity console.
To get the profile_id, you just need to navigate to the "Profiles" and look for the profile setting that you want to import on Conformity console and get the profile_id in the URL e.g. https://cloudone.trendmicro.com/conformity/profiles/profile:{profile_id}

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_profile resources into Terraform. Replace the {profile_id} value.
```hcl
terraform import conformity_profile.profile_settings {profile_id}
```

Use the `state` subcommand to verify Terraform successfully imported the conformity_profile resources.
```hcl
terraform state show conformity_profile.profile_settings
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Make sure you remove the id from the imported resource.
```hcl
terraform show -no-color >> main.tf
```
