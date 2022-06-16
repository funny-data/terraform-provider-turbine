package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceTemplateTable(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_template_table" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						database = "default"
						name = "foo"
						schema = <<-EOF
						{
							"fields": [
								{
									"name": "seq",
									"type": "int"
								},
								{
									"name": "sign",
									"type": "string"
								},
								{
									"name": "name",
									"type": "string"
								}
							]
						}
						EOF
						data = <<-EOF
						[
							[1, "feature-3", "feature-3"],
							[1, "feature-5", "feature-5"],
							[1, "feature-21", "feature-21"]
						]
						EOF
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_template_table.foo", "metadata.0.name", "foo"),
				),
			},
			{
				Config: `
				resource "turbine_template_table" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						database = "default"
						name = "foo"
						schema = <<-EOF
						{
							"fields": [
								{
									"name": "seq",
									"type": "int"
								},
								{
									"name": "sign",
									"type": "string"
								},
								{
									"name": "name",
									"type": "string"
								}
							]
						}
						EOF
						data = <<-EOF
						[
							[1, "feature-3", "feature-3"],
							[1, "feature-5", "feature-5"]
						]
						EOF
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_template_table.foo", "metadata.0.name", "foo"),
				),
			},
		},
	})
}
