package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceKafkaTopic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "turbine_kafka_topic" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						name = "foo"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_kafka_topic.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_kafka_topic.foo", "spec.0.name", "foo"),
				),
			},
			{
				Config: `
				resource "turbine_kafka_topic" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						name = "foo"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_kafka_topic.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_kafka_topic.foo", "spec.0.name", "foo"),
				),
			},
			{
				Config: `
				resource "turbine_kafka_topic" "foo" {
					metadata {
						name = "foo"
					}
					spec {
						name = "foo"
						config = {
							"retention.ms": "86400000"
						}
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_kafka_topic.foo", "metadata.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_kafka_topic.foo", "spec.0.name", "foo"),
					resource.TestCheckResourceAttr(
						"turbine_kafka_topic.foo", "spec.0.config.retention.ms", "86400000"),
				),
			},
		},
	})
}
