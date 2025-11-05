package provider

import (
	"context"

	"github.com/andybaran/terragpio/gpioclient"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the provider satisfies the framework interfaces
var _ fwprovider.Provider = (*gpioProvider)(nil)

type gpioProvider struct {
	version string
}

// New returns a new provider factory function for providerserver.Serve
func New(version string) func() fwprovider.Provider {
	return func() fwprovider.Provider {
		return &gpioProvider{version: version}
	}
}

// providerModel maps provider schema data
type providerModel struct {
	ServerAddr types.String `tfsdk:"serveraddr"`
}

func (p *gpioProvider) Metadata(_ context.Context, _ fwprovider.MetadataRequest, resp *fwprovider.MetadataResponse) {
	resp.TypeName = "gpio"
	resp.Version = p.version
}

func (p *gpioProvider) Schema(_ context.Context, _ fwprovider.SchemaRequest, resp *fwprovider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serveraddr": schema.StringAttribute{
				Required:    true,
				Description: "Address of the GPIO gRPC server",
			},
		},
	}
}

func (p *gpioProvider) Configure(ctx context.Context, req fwprovider.ConfigureRequest, resp *fwprovider.ConfigureResponse) {
	var data providerModel

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ServerAddr.IsUnknown() || data.ServerAddr.IsNull() {
		resp.Diagnostics.AddError(
			"Missing server address",
			"The provider requires the 'serveraddr' attribute to connect to the GPIO backend.",
		)
		return
	}

	client, err := gpioclient.NewClient(data.ServerAddr.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create GPIO client",
			err.Error(),
		)
		return
	}

	// Make the client available to resources and data sources
	resp.ResourceData = client
	resp.DataSourceData = client
}

func (p *gpioProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewPWMResource,
		NewBME280Resource,
		NewFanControllerResource,
	}
}

func (p *gpioProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	// No data sources currently
	return nil
}
