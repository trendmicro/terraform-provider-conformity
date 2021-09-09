---
page_title: "conformity_azure_account"
subcategory: "Azure"
description: |-
  Provides a Conformity Azure Account.
---

# Resource `conformity_azure_account`
Provides a Conformity Azure Account.

# Prerequisite
1. Setup a app registry in Azure by using the `azure_app_registry` in the examples folder or by following the manual steps here https://www.cloudconformity.com/help/add-cloud-account/add-an-azure-account.html
2. In all cases you need to add the Azure Active Directory manually in the Conformity Application like described here https://www.cloudconformity.com/help/add-cloud-account/add-an-azure-account.html#add-an-azure-active-directory

## Example Usage
```hcl
resource "conformity_azure_account" "azure" {
    name                = "azure-conformity"
    environment         = "development"
    active_directory_id = "your-active-directory-id"
    subscription_id     = "your-subscription-id"
}
```

## Argument reference
 - `name` (String) - (Required) The name of your account.
 - `environment` (String) - (Required) The environment for your account.
 - `subscription_id` (String) - (Required) The Azure Subscription ID
 - `active_directory_id` - (String) - (Required) The Azure Active Directory ID
  
  Inside `settings` there can be a `bot` set.

 - `bot` - (Optional) List: (Can be multiple declaration)
     * `disabled` (Bool) - (Optional) True to disable or false to enable the Conformity Bot.
     * `delay` (Int) - (Optional) Sets the number of hours delay between Conformity Bot runs.

  Inside `settings` there can be multiple `rule` set.

 - `rule` - (Optional) List: (Can be multiple declaration)
     * `rule_id` (String) - (Optional) - Rule ID, same as the one provided in the endpoint.
  
  Inside `settings` under `rule` set, there can be one `settings` set. 

 - `settings` - (Optional)
     * `enabled` (Bool) - (Optional) True for inclusion in bot detection, false for exclusion.
     * `rule_exists` (Bool) - (Optional)  True if rule exists.
     * `risk_level` (String) - (Optional) - Risk level of the Conformity rule. Enum: "LOW" "MEDIUM" "HIGH" "VERY_HIGH" "EXTREME"

  Inside `settings` under `rule` set, there can be one `exceptions` set. 
 
 - `exceptions` - (Optional) List: 
     * `filter_tags` (Array of Strings)- (Optional) An array of resource tags, resource tag keys or resource tag values that are exempted from the rule when it runs, e.g filterTags ["dev"] will exempt resource with tag "environment::dev from the rule".
     * `resources` (Array of Strings) - (Optional) An array of resource IDs that are exempted from the rule when it runs.
     * `tags` (Array of Strings) - (Optional) An array of resource tags that are exempted from the rule when it runs.

   Inside `settings` under `rule` set, there can be multiple `extra_settings` set. 

 - `extra_settings` - (Optional) List: (Can be multiple declaration)
     * `name` (String) - (Optional) (Keyword) Name of the extra setting.
     * `type` (String) - (Required) Rule specific property. Values can be: "multiple-string-values", "multiple-number-values" "multiple-aws-account-values", "choice-multiple-value" "choice-single-value", "single-number-value", "single-string-value", "ttl", "single-value-regex", "countries", "multiple-ip-values", and "tags".
     * `value` (String) - (Optional) Customisable value for rules that take on single name/value pairs.
     * `regions` (Array of Strings) - (Optional) Rule specific property.
     * `multiple-object-values` (Array of Strings) - (Optional) Rule specific property.

  Inside `extra_settings` under `settings` of `rule` set, there can be multiple declaration of `multiple-object-values` set.
  
 - `multiple-object-values` - (Optional) List: (Can be multiple declaration). 
     *  `event_name` (String) - (Optional) Name of the event.
     *  `event_source` (String) - (Optional) Name of the event source
     *  `user_identity_type` (String) - (Required) Type of the Identity of the user.

  Inside `extra_settings` under `settings` of `rule` set, there can be multiple declaration of `mappings` set. And under `mappings` set, here can be multiple declaration of `values` set.
  
 - `values` - (Required) List: (Can be multiple declaration). An array (sometimes of objects) rules that take on a set of of values
     * `name` (String) - (Optional) (Keyword) Name of the values.
     * `type` (String) - (Required) Rule specific property. Values can be: "multiple-string-values", "multiple-number-values" "multiple-aws-account-values", "choice-multiple-value" "choice-single-value", "single-number-value", "single-string-value", "ttl", "single-value-regex", "countries", "multiple-ip-values", and "tags".
     *  `value` (String) - (Required) Description of the checkbox.
    Note: `values` is required when you use `mappings`.

  Inside `values`, there can be multiple declaration of `values` set.
  
 - `values` - (Required) List: (Can be multiple declaration).
     *  `value` (String) - (Required) Description of the checkbox.
    Note: If inside the `values` under the `mappings` has set `values` declared, you cannot use `value` anymore. Inside mappings, its either `values` with `values` set inside it or `values` with declared `value` inside it.

        Note: There is a condition for `type` attribute. If the specified is attribute is `value`, the possible values are "single-number-value", "single-string-value", "single-value-regex" and "ttl". If the specified is attribute is `values`, the declaration of it is inside the extra settings which can be a list and the possible values are "choice-multiple-value", "choice-single-value", "multiple-string-values", "multiple-number-values", "countries", "multiple-ip-values", "multiple-aws-account-values" and "tags". You cannot declare both `values` and `value` at the same time.See the table below:

| type     | possible value                                                                                                                | Sample declaration                                                                                                                                                                                                                    |
|----------|-------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `value`  | single-number-value, single-string-value, single-value-regex, ttl                                                             | included {     ….     exceptions {       ….     }       extra_settings {		 	….                     type = "ttl"         value = "72"       } }                                                                                           |
| `values` | choice-multiple-value, choice-single-value, multiple-string-values, multiple-number-values, countries, multiple-ip-values, multiple-aws-account-values, tags | included {     ….     exceptions {       ….     }       extra_settings {		 	….                     type = "choice-multiple-value"           values {             ….   

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

 - `id` - The ID of the Azure account in Conformity managed by this resource

Example usage on the template:

```hcl
account {
    id = conformity_azure_account.azure.id
}
```

## Import
Azure Account - Can import based on the `Account ID` that can be found under the Conformity web console.

To get the Azure Account ID:
Open Conformity Admin Web console > navigate to Account dashboard > choose the Azure account you would like to import.
Notice the URL, you should be able to get the account ID e.g. https://cloudone.trendmicro.com/conformity/account/account-ID

Run `terraform init`:
```hcl
terraform init
```

Import the conformity_azure_account resources into Terraform. Replace the {CLOUDCONFORMITY-ACCOUNT-ID-AZURE} value.
```hcl
terraform import conformity_azure_account.azure {CLOUDCONFORMITY-ACCOUNT-ID-AZURE}
```

Use the `state` subcommand to verify Terraform successfully imported the conformity_azure_account resources.
```hcl
terraform state show conformity_azure_account.azure
```

Run `terraform show -no-color >> main.tf` to import the resources on the `main.tf` file. Be sure to remove the id from the resource
```hcl
terraform show -no-color >> main.tf
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_conformity"></a> [conformity](#requirement\_conformity) | 0.3.6 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_conformity"></a> [conformity](#provider\_conformity) | 0.3.6 |