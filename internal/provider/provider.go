package provider

import (
	"context"

	"github.com/andybaran/terragpio/gpioclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type gpioapiclient struct {
	c *gpioclient.Client
}

var _ provider.Provider = (*gpioProvider)(nil)

func New() provider.Provider {
    return &gpioProvider{}
}

func (p *gpioProvider) Resources(_ context.Context) []func() resource.Resource {
    return []func() resource.Resource{
        func() resource.Resource {
            return &gpio_pwmResource{}
        },
		func() resource.Resource {
            return &gpio_bme280Resource{}
        },
		func() resource.Resource {
            return &gpio_input_temperature_output_fanResource{}
        },
        /* ... */
    }
}

func (p *gpioProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return []func() datasource.DataSource{
        func() datasource.DataSource {
            return &exampleDataSource{},
        },
        /* ... */
    }
}

func (p *gpioProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "serveraddr": schema.StringAttribute{
                Required: true,
            },
        },
    }
}

func (p *gpioProvider) Configure(ctx context.Context, req provider.ConfigureRequest, res *provider.ConfigureResponse) {
    /* ... */
}

/*func New(version string) func() *schema.Provider {
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
				//"gpio_scaffolding_data_source": dataSourceScaffolding(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"gpio_pwm":                          resource_pwm(),
				"gpio_bme280":                       resource_bme280(),
				"gpio_input_temperature_output_fan": resource_input_temperature_output_fan(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}



func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics
		c, err := gpioclient.NewClient(p.Schema["serveraddr"].GoString())
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to connect to the server",
				Detail:   "Unable to connect to the server address as specified",
			})
		}
		return gpioapiclient{c: c}, diags
	}
}*/
