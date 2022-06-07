package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceSLSSink(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_sls_sink" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						parallelism = 1
						topic = "foo"
						dead_letter_queue_topic = "foo-dlq"
						project = "foo"
						logstore = "foo"
						time_field = "time"
						flush_size = 1000
						flush_interval_seconds = 5
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.parallelism", "1"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.topic", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.dead_letter_queue_topic", "foo-dlq"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.project", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.logstore", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.time_field", "time"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.flush_size", "1000"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.flush_interval_seconds", "5"),
				),
			},
			{
				Config: `
				resource "turbine_sls_sink" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						parallelism = 2
						topic = "foo"
						dead_letter_queue_topic = "foo-dlq"
						project = "foo"
						logstore = "foo"
						time_field = "time"
						flush_size = 1000
						flush_interval_seconds = 5
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.parallelism", "2"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.topic", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.dead_letter_queue_topic", "foo-dlq"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.project", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.logstore", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.time_field", "time"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.flush_size", "1000"),
					resource.TestCheckResourceAttr(
						"turbine_sls_sink.foo", "spec.0.flush_interval_seconds", "5"),
				),
			},
		},
	})
}
