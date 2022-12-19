package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	icebergTableSchema1 = `{"fields":[{"name":"log_id","type":"string"},{"name":"time","type":"timestamp"}]}`
	icebergTableSchema2 = `{"fields":[{"name":"account_id","type":"string"},{"name":"log_id","type":"string"},{"name":"time","type":"timestamp"}]}`
)

func TestAccResourceIcebergTable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "turbine_iceberg_table" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						database = "demo"
						name = "demo"
						schema = %s
						partition_spec {
							fields = [
								{"expression": "days(time)"}
							]
						}
					}
				}
				`, strconv.Quote(icebergTableSchema1)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.database", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.name", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.schema", icebergTableSchema1),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.partition_spec.0.fields.0.expression", "days(time)"),
				),
			},
			{
				Config: fmt.Sprintf(`
				resource "turbine_iceberg_table" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						database = "demo"
						name = "demo"
						schema = %s
						partition_spec {
							fields = [
								{"expression": "days(time)"}
							]
						}
					}
				}
				`, strconv.Quote(icebergTableSchema2)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.database", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.name", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.schema", icebergTableSchema2),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.partition_spec.0.fields.0.expression", "days(time)"),
				),
			},
			{
				Config: fmt.Sprintf(`
				resource "turbine_iceberg_table" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						database = "demo"
						name = "demo"
						schema = %s
						partition_spec {
							fields = [
								{"expression": "months(time)"}
							]
						}
					}
				}
				`, strconv.Quote(icebergTableSchema2)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.database", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.name", "demo"),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.schema", icebergTableSchema2),
					resource.TestCheckResourceAttr(
						"turbine_iceberg_table.foo", "spec.0.partition_spec.0.fields.0.expression", "months(time)"),
				),
			},
		},
	})
}
