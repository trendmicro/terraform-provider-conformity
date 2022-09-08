package conformity

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceAzureActiveDirectory(t *testing.T) {

	name := "Azure Active Directory"
	directory_id := "761d49c9-8898-5d35-c4db-ed8582f20dbf"
	application_id := "c5187d37-8de6-5920-99df-4d8eb3f8cc05"
	application_key := "kjx9Q~.CeN4.AxZZVvT8qFRmcx9v9HDVBxgA3mc1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccConformityPreCheck(t) },
		Providers: testAccConformityProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAzureActiveDirectoryConfigBasic(name, directory_id, application_id, application_key),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("conformity_azure_active_directory.azure", "name", name),
					resource.TestCheckResourceAttr("conformity_azure_active_directory.azure", "directory_id", directory_id),
					resource.TestCheckResourceAttr("conformity_azure_active_directory.azure", "application_id", application_id),
					resource.TestCheckResourceAttr("conformity_azure_active_directory.azure", "application_key", application_key),
				),
			},
		},
	})
}
func testAccCheckAzureActiveDirectoryConfigBasic(name, directory_id, application_id, application_key string) string {
	return fmt.Sprintf(`
	resource "conformity_azure_active_directory" "azure" {
	    name = "%s"
	    directory_id    = "%s"
	    application_id     = "%s"
	    application_key = "%s"
	}
`, name, directory_id, application_id, application_key)

}
