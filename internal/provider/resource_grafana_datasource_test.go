package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGrafanaDatasource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_grafana_datasource" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						name = "foo"
						org_name = "foo"
						hudi_database = "foo"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_grafana_datasource.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_grafana_datasource.foo", "spec.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_grafana_datasource.foo", "spec.0.org_name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_grafana_datasource.foo", "spec.0.hudi_database", "foo"),
				),
			},
		},
	})
}
