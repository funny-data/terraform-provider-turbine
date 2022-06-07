package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"endpoint": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("TURBINE_ENDPOINT", nil),
					Description: "Endpoint of Turbine API Server",
				},
				"username": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("TURBINE_USERNAME", nil),
				},
				"password": {
					Type:        schema.TypeString,
					Sensitive:   true,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("TURBINE_PASSWORD", nil),
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"turbine_dummy":           resourceDummy(),
				"turbine_kafka_topic":     resourceKafkaTopic(),
				"turbine_ingest_metadata": resourceIngestMetadata(),
				"turbine_hudi_database":   resourceHudiDatabase(),
				"turbine_hudi_table":      resourceHudiTable(),
				"turbine_xdconsole_sink":  resourceXDConsoleSink(),
				"turbine_sls_sink":        resourceSLSSink(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		userAgent := p.UserAgent("terraform-provider-turbine", version)

		endpoint := d.Get("endpoint").(string)
		var (
			username string
			password string
		)
		if i, ok := d.GetOk("username"); ok {
			username = i.(string)
		}
		if i, ok := d.GetOk("password"); ok {
			password = i.(string)
		}

		return newApiClient(userAgent, endpoint, username, password), nil
	}
}
