package conformity

import (
	"fmt"
	"github.com/trendmicro/terraform-provider-conformity/pkg/cloudconformity"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
    private_key := "-----BEGIN PRIVATE KEY-----key=-----END PRIVATE KEY-----"
    client_email := "iam.gserviceaccount.com"
    client_id := "811129548"
    auth_uri := "https://accounts.google.com/o/oauth2/auth"
    token_uri := "https://oauth2.googleapis.com/token"
    auth_provider_x509_cert_url := "https://www.googleapis.com/oauth2/v1/certs"
    client_x509_cert_url := "https://www.googleapis.com/robot/v1/metadata/x509/cloud-one-conformity-bot%40conformity-346910.iam.gserviceaccount.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccConformityPreCheck(t) },
		Providers:    testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGCPOrgConfigBasic(service_account_name, acc_type, project_id, private_key_id, private_key, client_email, client_id, auth_uri, token_uri, auth_provider_x509_cert_url, client_x509_cert_url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "service_account_name", "MySubscription"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "type", "service_account"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "project_id", "conformity-346910"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "private_key_id", "c1c3688e7c"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "private_key", "-----BEGIN PRIVATE KEY-----key=-----END PRIVATE KEY-----"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "client_email", "iam.gserviceaccount.com"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "client_id", "811129548"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "auth_uri", "https://accounts.google.com/o/oauth2/auth"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "token_uri", "https://oauth2.googleapis.com/token"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "auth_provider_x509_cert_url", "https://www.googleapis.com/oauth2/v1/certs"),
					resource.TestCheckResourceAttr("conformity_gcp_org.gcp_org", "client_x509_cert_url", "https://www.googleapis.com/robot/v1/metadata/x509/cloud-one-conformity-bot%40conformity-346910.iam.gserviceaccount.com"),

				),
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
    output "conformity_gcp_org_service_account_name" {
		value = conformity_gcp_org.gcp_org.service_account_name
	}

	`, service_account_name, acc_type, project_id, private_key_id, private_key, client_email, client_id, auth_uri, token_uri, auth_provider_x509_cert_url, client_x509_cert_url)
}
