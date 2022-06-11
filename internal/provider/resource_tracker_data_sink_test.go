package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceTrackerDataSink(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_tracker_data_sink" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						topic = "foo"
						environment = "foo"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_tracker_data_sink.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_tracker_data_sink.foo", "spec.0.topic", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_tracker_data_sink.foo", "spec.0.environment", "foo"),
				),
			},
		},
	})
}
