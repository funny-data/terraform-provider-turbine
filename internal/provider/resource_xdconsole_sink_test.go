package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceXDConsoleSink(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_xdconsole_sink" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						topic = "foo"
						event = "foo"
						app = "foo"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink.foo", "spec.0.topic", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink.foo", "spec.0.event", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink.foo", "spec.0.app", "foo"),
				),
			},
			{
				Config: `
				resource "turbine_xdconsole_sink" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						topic = "foo"
						event = "foo"
						app = "bar"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink.foo", "spec.0.topic", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink.foo", "spec.0.event", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_xdconsole_sink.foo", "spec.0.app", "bar"),
				),
			},
		},
	})
}
