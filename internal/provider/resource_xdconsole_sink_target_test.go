package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceXDConsoleSinkTarget(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_xdconsole_sink_target" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						endpoint = "endpoint"
						project = "project"
						logstore = "logstore"
						access_key_id = "access_key_id"
						access_key_secret = "access_key_secret"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink_target.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink_target.foo", "spec.0.endpoint", "endpoint"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink_target.foo", "spec.0.project", "project"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink_target.foo", "spec.0.logstore", "logstore"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink_target.foo", "spec.0.access_key_id", "access_key_id"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink_target.foo", "spec.0.access_key_secret", "access_key_secret"),
				),
			},
		},
	})
}
