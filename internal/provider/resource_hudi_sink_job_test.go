package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceHudiSinkJob(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_hudi_sink_job" "foo" {
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
						delta_streamer_task {
							enabled = true
							operation = "INSERT"
							interval = 300
							batch_size_limit = 209715200
							batch_records_limit = 2000000
							backoff = 900
							max_backoff = 3600
							default_lag_records = 10000
							default_avg_record_size = 1000
							parallelism = 20
							extra_properties = {}
						}
						sync_task {
							enabled = true
							interval = 900
							backoff = 900
							max_backoff = 3600
							extra_properties = {}
						}
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_hudi_sink_job.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_sink_job.foo", "spec.0.table.0.database", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_sink_job.foo", "spec.0.table.0.table", "demo"),
				),
			},
			{
				Config: `
				resource "turbine_hudi_sink_job" "foo" {
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
						delta_streamer_task {
							enabled = true
							operation = "INSERT"
							interval = 300
							batch_size_limit = 209715200
							batch_records_limit = 2000000
							backoff = 900
							max_backoff = 3600
							default_lag_records = 10000
							default_avg_record_size = 1000
							parallelism = 25
							extra_properties = {}
						}
						sync_task {
							enabled = true
							interval = 900
							backoff = 900
							max_backoff = 3600
							extra_properties = {}
						}
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_hudi_sink_job.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_sink_job.foo", "spec.0.table.0.database", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_sink_job.foo", "spec.0.table.0.table", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_sink_job.foo", "spec.0.delta_streamer_task.0.parallelism", "25"),
				),
			},
		},
	})
}
