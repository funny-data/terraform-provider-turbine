package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceIcebergSinkJob(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_iceberg_sink_job" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						table {
							database = "demo"
							table = "demo"
						}
						source {
							type = "kafka"
							format = "json"
							topic = "demo"
							starting_offsets = "earliest"
						}
						target_interval = 60
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.table.0.database", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.table.0.table", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.source.0.type", "kafka"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.source.0.format", "json"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.source.0.topic", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.source.0.starting_offsets", "earliest"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.target_interval", "60"),
				),
			},
			{
				Config: `
				resource "turbine_iceberg_sink_job" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						table {
							database = "demo"
							table = "demo"
						}
						source {
							type = "kafka"
							format = "json"
							topic = "demo"
							starting_offsets = "earliest"
						}
						target_interval = 120
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.table.0.database", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.table.0.table", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.source.0.type", "kafka"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.source.0.format", "json"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.source.0.topic", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.source.0.starting_offsets", "earliest"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_sink_job.foo", "spec.0.target_interval", "120"),
				),
			},
		},
	})
}
