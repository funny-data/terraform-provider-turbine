package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceDummy(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_dummy" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						arg = "bar"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_dummy.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_dummy.foo", "spec.0.arg", "bar"),
				),
			},
			{
				Config: `
				resource "turbine_dummy" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						arg = "baz"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_dummy.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_dummy.foo", "spec.0.arg", "baz"),
				),
			},
			{
				Config: `
				resource "turbine_dummy" "foo" {
					metadata {
						name = "foo"
						labels = {
							"x" = "y"
						}
					}
					spec {
						arg = "baz"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_dummy.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_dummy.foo", "metadata.0.labels.x", "y"),
					resource.TestCheckResourceAttr(
						"turbine_dummy.foo", "spec.0.arg", "baz"),
				),
			},
		},
	})
}
