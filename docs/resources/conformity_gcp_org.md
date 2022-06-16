---
page_title: "conformity_gcp_org Resource"
subcategory: "GCP"
description: |-
  Provides a Conformity Organisation.
---

# Resource `conformity_gcp_org`
Provides a Conformity GCP Organisation.

## Example Usage With GCP Conformity To Create Account Only
```hcl

resource "conformity_gcp_org" "gcp_org" {
    private_key              = "privetkey"
    service_account_name     = "MySubscription"
    type                     = "service_account"
    project_id               = "conformity-346910"
    private_key_id           = "c1c3688e7c"
    client_email             = "iam.gserviceaccount.com"
    client_id                = "811129548"
    auth_uri                 = "https://accounts.google.com/o/oauth2/auth"
    token_uri                = "https://oauth2.googleapis.com/token"
    auth_provider_x509_cert_url = "https://www.googleapis.com/oauth2/v1/certs"
    client_x509_cert_url     = "https://www.googleapis.com/robot/v1/metadata/x509/cloud-one-conformity-bot%40conformity-346910.iam.gserviceaccount.com"
}
```

## Argument reference
 - `serviceAccountName` (String) - (Required) The name of your organisation.
 
 Other details you will get it from `serviceAccountKeyJson`
 