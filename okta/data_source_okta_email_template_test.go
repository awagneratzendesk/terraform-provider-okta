package okta

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceOktaEmailTemplate_read(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(emailTemplate)
	config := mgr.GetFixtures("datasource.tf", ri, t)

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorChecks(t),
		ProviderFactories: testAccProvidersFactories,
		Steps: []resource.TestStep{
			{
				Config:  config,
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.okta_email_template.forgot_password", "brand_id"),
					resource.TestCheckResourceAttrSet("data.okta_email_template.forgot_password", "name"),
					resource.TestCheckResourceAttr("data.okta_email_template.forgot_password", "name", "ForgotPassword"),
					resource.TestCheckResourceAttrSet("data.okta_email_template.forgot_password", "links"),
				),
			},
		},
	})
}
