package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceHudiDatabase(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_hudi_database" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						name = "foo"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_hudi_database.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_database.foo", "spec.0.name", "foo"),
				),
			},
		},
	})
}
