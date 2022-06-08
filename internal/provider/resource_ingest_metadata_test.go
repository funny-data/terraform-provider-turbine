package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const ingestMetadata1 = `{"name":"sausage_debug","access_keys":[{"id":"rQJEk4mz66","secret":"CcYyCc=="}],"messages":[{"type":"Event","processors":[{"type":"DebugProcessor","processors":null,"config":{"id":"99"}}]}]}`
const ingestMetadata2 = `{"name":"sausage_debug","access_keys":[{"id":"rQJEk4mz66","secret":"CcYyCc=="}],"messages":[{"type":"Event","processors":[{"type":"DebugProcessor","processors":null,"config":{"id":"100"}}]}]}`

func TestAccResourceIngestMetadata(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "turbine_ingest_metadata" "foo" {
					metadata {
						name = "foo"
					}
					spec = %s
				}
				`, strconv.Quote(ingestMetadata1)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_ingest_metadata.foo", "metadata.0.name", "foo"),
					CheckResourceAttrEqualToJson(
						"turbine_ingest_metadata.foo", "spec", ingestMetadata1),
				),
			},
			{
				Config: fmt.Sprintf(`
				resource "turbine_ingest_metadata" "foo" {
					metadata {
						name = "foo"
					}
					spec = %s
				}
				`, strconv.Quote(ingestMetadata2)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"turbine_ingest_metadata.foo", "metadata.0.name", "foo"),
					CheckResourceAttrEqualToJson(
						"turbine_ingest_metadata.foo", "spec", ingestMetadata2),
				),
			},
		},
	})
}
