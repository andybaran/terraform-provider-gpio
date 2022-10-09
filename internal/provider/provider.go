package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/grpc"

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
				"serverAddr": &schema.Schema{
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("rpiaddr", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"scaffolding_data_source": dataSourceScaffolding(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"gpio_pwm":                     resource_gpio_pwm(),
				"bme280":                       resource_bme280(),
				"input_temperature_output_fan": resource_input_temperature_output_fan(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
	MyClient   gpioclient.Client
	ServerAddr string
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {

		myClient, err := gpioclient.NewClient(p.Schema["serverAddr"].GoString())
		if err != nil {
			log.Fatalf("can't connect to: %v", err)
		}

		//var diags diag.Diagnostics

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure()) //not ready to worry about security just yet
		opts = append(opts, grpc.WithBlock())    //we do this b/c we just want to fail immediately if we can't connect to the server https://pkg.go.dev/google.golang.org/grpc@v1.32.0?utm_source=gopls#WithBlock

		if err != nil {
			log.Fatalf("failed to dial: %v", err)
		}

		return &myClient, nil
	}
}
