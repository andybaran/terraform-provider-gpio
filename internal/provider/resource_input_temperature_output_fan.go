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

var _ resource.Resource = (*fanControllerResource)(nil)
var _ resource.ResourceWithConfigure = (*fanControllerResource)(nil)

type fanControllerResource struct{ client *gpioclient.Client }

type fanControllerModel struct {
	ID              types.String `tfsdk:"id"`
	TimeInterval    types.String `tfsdk:"timeinterval"`
	BME280DevicePin types.String `tfsdk:"bme280devicepin"`
	TemperatureMax  types.String `tfsdk:"temperaturemax"`
	TemperatureMin  types.String `tfsdk:"temperaturemin"`
	FanDevice       types.String `tfsdk:"fandevice"`
	DutyCycleMax    types.String `tfsdk:"dutycyclemax"`
	DutyCycleMin    types.String `tfsdk:"dutycyclemin"`
}

func NewFanControllerResource() resource.Resource { return &fanControllerResource{} }

func (r *fanControllerResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gpio_input_temperature_output_fan"
}

func (r *fanControllerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Control a fan PWM output based on BME280 temperature readings.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Internal identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"timeinterval":    schema.StringAttribute{Required: true, Description: "Read interval in seconds"},
			"bme280devicepin": schema.StringAttribute{Required: true, Description: "BME280 device pin id"},
			"temperaturemax":  schema.StringAttribute{Required: true, Description: "Max temperature"},
			"temperaturemin":  schema.StringAttribute{Required: true, Description: "Min temperature"},
			"fandevice":       schema.StringAttribute{Required: true, Description: "Fan PWM device pin id"},
			"dutycyclemax":    schema.StringAttribute{Required: true, Description: "Max duty cycle"},
			"dutycyclemin":    schema.StringAttribute{Required: true, Description: "Min duty cycle"},
		},
	}
}

func (r *fanControllerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	if c, ok := req.ProviderData.(*gpioclient.Client); ok {
		r.client = c
	}
}

func (r *fanControllerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan fanControllerModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Convert numeric strings where necessary using the API expectation
	apiResp, err := r.client.StartFanController(gpioclient.StartFanControllerArgs{
		TimeInterval:    parseUint(plan.TimeInterval.ValueString()),
		BME280DevicePin: plan.BME280DevicePin.ValueString(),
		TemperatureMax:  parseUint(plan.TemperatureMax.ValueString()),
		TemperatureMin:  parseUint(plan.TemperatureMin.ValueString()),
		FanDevice:       plan.FanDevice.ValueString(),
		DutyCycleMax:    parseUint(plan.DutyCycleMax.ValueString()),
		DutyCylceMin:    parseUint(plan.DutyCycleMin.ValueString()),
	})
	if err != nil {
		resp.Diagnostics.AddError("Fan controller setup failed", err.Error())
		return
	}

	plan.ID = types.StringValue(apiResp.PinCombo)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *fanControllerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state fanControllerModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *fanControllerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Treat as re-creating controller settings
	r.Create(ctx, resource.CreateRequest{Plan: req.Plan}, &resource.CreateResponse{State: resp.State, Diagnostics: resp.Diagnostics})
}

func (r *fanControllerResource) Delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse) {
	// No explicit stop method in API; assume external management
}

// parseUint is a helper to safely parse unsigned integers from strings
func parseUint(s string) uint64 {
	var u uint64
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0
		}
		u = u*10 + uint64(c-'0')
	}
	return u
}
