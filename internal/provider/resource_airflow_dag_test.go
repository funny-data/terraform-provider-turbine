package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceAirflowDAG(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_airflow_dag" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						filename = "foo.py"
						code = <<-EOF
							print("hello world")
						EOF
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_airflow_dag.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_airflow_dag.foo", "spec.0.code", "print(\"hello world\")\n"),
				),
			},
			{
				Config: `
				resource "turbine_airflow_dag" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						filename = "foo.py"
						code = <<-EOF
							print("hello Airflow")
						EOF
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_airflow_dag.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_airflow_dag.foo", "spec.0.code", "print(\"hello Airflow\")\n"),
				),
			},
		},
	})
}
