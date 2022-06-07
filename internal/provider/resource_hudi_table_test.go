package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	hudiTableSchema1 = `{"fields":[{"name":"log_id","type":"string"},{"name":"time","type":"long"}]}`
	hudiTableSchema2 = `{"fields":[{"name":"account_id","type":"string"},{"name":"log_id","type":"string"},{"name":"time","type":"long"}]}`
)

func TestAccResourceHudiTable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "turbine_hudi_table" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						database = "demo"
						name = "demo"
						type = "COPY_ON_WRITE"
						schema = %s
						record_key = "log_id"
						pre_combine_key = "time"
						partition_spec = [
							{"name": "dt", "transform": "hudi_partition_transform_date_from_unixtime(time, 'milliseconds', 'Asia/Shanghai')"}
						]
					}
				}
				`, strconv.Quote(hudiTableSchema1)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.database", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.name", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.type", "COPY_ON_WRITE"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.schema", hudiTableSchema1),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.record_key", "log_id"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.pre_combine_key", "time"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.partition_spec.0.name", "dt"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.partition_spec.0.transform", "hudi_partition_transform_date_from_unixtime(time, 'milliseconds', 'Asia/Shanghai')"),
				),
			},
			{
				Config: fmt.Sprintf(`
				resource "turbine_hudi_table" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						database = "demo"
						name = "demo"
						type = "COPY_ON_WRITE"
						schema = %s
						record_key = "log_id"
						pre_combine_key = "time"
						partition_spec = [
							{"name": "dt", "transform": "hudi_partition_transform_date_from_unixtime(time, 'milliseconds', 'Asia/Shanghai')"}
						]
					}
				}
				`, strconv.Quote(hudiTableSchema2)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.database", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.name", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.type", "COPY_ON_WRITE"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.schema", hudiTableSchema2),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.record_key", "log_id"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.pre_combine_key", "time"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.partition_spec.0.name", "dt"),
					resource.TestCheckResourceAttr(
						"turbine_hudi_table.foo", "spec.0.partition_spec.0.transform", "hudi_partition_transform_date_from_unixtime(time, 'milliseconds', 'Asia/Shanghai')"),
				),
			},
		},
	})
}
