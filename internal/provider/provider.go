package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	//"google.golang.org/grpc"

	"github.com/andybaran/terragpio/gpioclient"
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
				"serveraddr": &schema.Schema{
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("serveraddr", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"gpio_scaffolding_data_source": dataSourceScaffolding(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"gpio_pwm":                          resource_gpio_pwm(),
				"gpio_bme280":                       resource_bme280(),
				"gpio_input_temperature_output_fan": resource_input_temperature_output_fan(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

/*type apiClient struct {
	MyClient   gpioclient.Client
	ServerAddr string //TODO: Why am I passing this around? This is likely not needed

}*/

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {

		var diags diag.Diagnostics

		myClient, err := gpioclient.NewClient(p.Schema["serveraddr"].GoString())
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to connect to the server",
				Detail:   "Unable to connect to the server address as specified",
			})
		}

		return *myClient, diags
	}
}
