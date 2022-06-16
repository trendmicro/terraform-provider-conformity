---
page_title: "conformity_gcp_account Resource"
subcategory: "GCP"
description: |-
  Provides a Conformity Account.
---

# Resource `conformity_gcp_account`
Provides a Conformity GCP Account.

## Example Usage With GCP Conformity To Create Account Only
```hcl

resource "conformity_gcp_account" "gcp" {
  name                      = "MyProject"
  project_id                = "conformity-346910"
  project_name              = "conformity"
  service_account_unique_id = "10307221"
  environment               = "dev"
  tags = ["staging"]
  settings {
      bot {
          delay            = 1
          disabled         = false
          disabled_regions = [ "ap-east-1", "ap-south-1" ]
      }
      // implement multiple-object-values
      rule {
          rule_id = "CloudAPI-001"
          settings {
              enabled     = true
              risk_level  = "MEDIUM"
              extra_settings {
                  name    = "rotatingPeriod"
                  type    = "single-number-value"
                  value   = 90
              }
          }
      }
  }
}
```

## Argument reference
 - `name` (String) - (Required) The name of your account.
 - `environment` (String) - (Required) The environment for your account.
 - `projectId` (String) - (Required) The ID of your GCP Project.
 - `projectName` (String) - (Required) The name of your GCP Project.
 - `serviceAccountUniqueId` (String) - (Required) The unique ID of your GCP Service Account.
 - `settings` - (Optional) List: (Can be multiple declaration)
  
  Inside `settings` there can be a `bot` set.

 - `bot` - (Optional) List: (Can be multiple declaration)
     * `disabled` (Bool) - (Optional) True to disable or false to enable the Conformity Bot.
     * `disabled_regions` (Array of Strings) - (Optional) - Possible values are "af-south-1", "ap-east-1", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ca-central-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-west-1", "us-west-2". This field can only be applied to AWS accounts. An attribute object containing a list of AWS regions for which Conformity Bot runs will be disabled.
     * `delay` (Int) - (Optional) Sets the number of hours delay between Conformity Bot runs.

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

 - `id` - The ID of the AWS account in Conformity managed by this resource

Example usage on the template:

```hcl
account {
    id = conformity_aws_account.aws.id
}
```
