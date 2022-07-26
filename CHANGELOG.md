## 0.4.6
* Issue 26: Conformity Profile always triggers a change
* Issue 30: conformity group resource does not handle unmanaged deletions
* Issue 31: Conformity Profile doesn't allow GCP
* Issue 33: Sort Included Settings
* Conformits Customer Rules
* Conformity Suppression Check Resource
* Conformity GCP Projects
* Conformity Azure Subscriptions

## 0.4.5
* Extend conformity profile handling

## 0.4.4
* Terraform Plan/Destroy test PR checks
* Added unit testing as part of the PR checks
* Fixes the conformity_report_config: Cannot set risk_levels filter related issue.

## 0.4.3
* conformity_report_config: Cannot set risk_levels filter 
* User creation failure for normal user
* conformity_azure_account does not read tags from backend

## 0.4.2
* Made GCP private key sensitive through schema
* GCP Account Doc
* conformity_azure_account: environment field should not be required

## 0.4.1
* GCP read issue fixed
* Cloudone URL changes and region support
* Made GCP private key sensitive

## 0.4.0
* GCP integration

## 0.3.8
* security vulnerable issue fixed
* Logging nul payload issue fixed

## 0.3.7

* Logging the request and response payload of API
* Improved Account create and Read request and response struct

## 0.3.6

* fix documentation

## 0.3.5

* fix documentation

## 0.3.4

* fix documentation

## 0.3.3

* fix documentation

## 0.3.2

* fix documentation
* fix tags on account creation

## 0.3.1

* fix create azure account
* improve bot & rule settings on update

## 0.3.0

* **New Data Source:** `conformity_apply_profile`
* add more types for rule settings
* add bot settings

## 0.2.0

FEATURES:

* **New Resource:** `conformity_profile`
* **New Data Source:** `conformity_apply_profile`

## 0.1.0

FEATURES:

* **New Data Source:** `conformity_external_id`
* **New Resource:** `conformity_aws_account`
* **New Resource:** `conformity_azure_account`
* **New Resource:** `conformity_gcp_account`
* **New Resource:** `conformity_gcp_org`
* **New Resource:** `conformity_group`
* **New Resource:** `conformity_user`
* **New Resource:** `conformity_sso_user`
* **New Resource:** `conformity_report_config`
* **New Resource:** `conformity_communication_setting`

