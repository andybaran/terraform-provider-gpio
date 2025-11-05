package provider

import (
	"context"

	"github.com/andybaran/terragpio/gpioclient"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*bme280Resource)(nil)
var _ resource.ResourceWithConfigure = (*bme280Resource)(nil)

type bme280Resource struct{ client *gpioclient.Client }

type bme280Model struct {
	ID      types.String `tfsdk:"id"`
	I2CBus  types.String `tfsdk:"i2cbus"`
	I2CAddr types.String `tfsdk:"i2caddr"`
}

func NewBME280Resource() resource.Resource { return &bme280Resource{} }

func (r *bme280Resource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gpio_bme280"
}

func (r *bme280Resource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Resource to setup BME280 i2c sensor",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Internal identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"i2cbus": schema.StringAttribute{
				Required:    true,
				Description: "i2c bus number (e.g. 1)",
			},
			"i2caddr": schema.StringAttribute{
				Required:    true,
				Description: "i2c address (e.g. 0x77)",
			},
		},
	}
}

func (r *bme280Resource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	if c, ok := req.ProviderData.(*gpioclient.Client); ok {
		r.client = c
	}
}

func (r *bme280Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan bme280Model
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.client.SetBME280(gpioclient.SetBME280Args{I2CBus: plan.I2CBus.ValueString(), I2CAddr: plan.I2CAddr.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("BME280 setup failed", err.Error())
		return
	}

	plan.ID = types.StringValue(apiResp.PinNumber)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *bme280Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state bme280Model
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// No read available; keep state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *bme280Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Treat as replace/update via same RPC
	var plan bme280Model
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.SetBME280(gpioclient.SetBME280Args{I2CBus: plan.I2CBus.ValueString(), I2CAddr: plan.I2CAddr.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("BME280 update failed", err.Error())
		return
	}
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *bme280Resource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
	// No-op; server may not support teardown
}
