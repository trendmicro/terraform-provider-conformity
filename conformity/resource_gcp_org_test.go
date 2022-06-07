package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
//	"regexp"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceGCPOrganisation(t *testing.T) {

	// accountPayload,ruleSetting1,rulesetting2, ruleSetting3 also uses by the other resource testing, expect not empty struct
	// To make sure  all are empty
	// accountPayload,ruleSetting1,rulesetting2, ruleSetting3 also uses by the other resource testing, expect not empty struct
	// To make sure  all are empty
	accountPayload = cloudconformity.AccountPayload{}
	ruleSetting1 = nil
	ruleSetting2 = nil
	ruleSetting3 = nil

	service_account_name := "MySubscription"
    acc_type := "service_account"
    project_id := "conformity-346910"
    private_key_id := "c1c3688e7c"
    private_key := "-----BEGIN PRIVATE KEY-----\nkey=\n-----END PRIVATE KEY-----\n"
    client_email := "iam.gserviceaccount.com"
    client_id := "811129548"
    auth_uri := "https://accounts.google.com/o/oauth2/auth"
    token_uri := "https://oauth2.googleapis.com/token"
    auth_provider_x509_cert_url := "https://www.googleapis.com/oauth2/v1/certs"
    client_x509_cert_url := "https://www.googleapis.com/robot/v1/metadata/x509/cloud-one-conformity-bot%40conformity-346910.iam.gserviceaccount.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGCPOrgConfigBasic(service_account_name, acc_type, project_id, private_key_id, private_key, client_email, client_id, auth_uri, token_uri, auth_provider_x509_cert_url, client_x509_cert_url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "name", "test-name"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "type", "service_account"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "project_id", "conformity-346910"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "private_key_id", "c1c3688e7c"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "private_key", "-----BEGIN PRIVATE KEY-----\nkey=\n-----END PRIVATE KEY-----\n"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "client_email", "iam.gserviceaccount.com"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "client_id", "811129548"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "auth_uri", "https://accounts.google.com/o/oauth2/auth"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "token_uri", "https://oauth2.googleapis.com/token"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "auth_provider_x509_cert_url", "https://www.googleapis.com/oauth2/v1/certs"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp", "client_x509_cert_url", "https://www.googleapis.com/robot/v1/metadata/x509/cloud-one-conformity-bot%40conformity-346910.iam.gserviceaccount.com"),

				), ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckGCPOrgConfigBasic(service_account_name, acc_type, project_id, private_key_id, private_key, client_email, client_id, auth_uri, token_uri, auth_provider_x509_cert_url, client_x509_cert_url string) string {
	return fmt.Sprintf(`
	resource "conformity_gcp_org" "gcp_org" {
    service_account_name     = "%s"
    type                     = "%s"
    project_id               = "%s"
    private_key_id           = "%s"
    private_key              = "%s"
    client_email             = "%s"
    client_id                = "%s"
    auth_uri                 = "%s"
    token_uri                = "%s"
    auth_provider_x509_cert_url = "%s"
    client_x509_cert_url     = "%s"
}

	`, service_account_name, acc_type, project_id, private_key_id, private_key, client_email, client_id, auth_uri, token_uri, auth_provider_x509_cert_url, client_x509_cert_url)
}
