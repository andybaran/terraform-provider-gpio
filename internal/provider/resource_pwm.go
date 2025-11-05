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

// Ensure the resource implements required interfaces
var _ resource.Resource = (*pwmResource)(nil)
var _ resource.ResourceWithConfigure = (*pwmResource)(nil)

type pwmResource struct {
	client *gpioclient.Client
}

type pwmModel struct {
	ID        types.String `tfsdk:"id"`
	Pin       types.String `tfsdk:"pin"`
	DutyCycle types.String `tfsdk:"dutycycle"`
	Frequency types.String `tfsdk:"frequency"`
}

func NewPWMResource() resource.Resource { return &pwmResource{} }

func (r *pwmResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gpio_pwm"
}

func (r *pwmResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Resource to control PWM Pins",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Internal identifier (pin number)",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"pin": schema.StringAttribute{
				Required:    true,
				Description: "GPIO Pin (e.g. GPIO12)",
			},
			"dutycycle": schema.StringAttribute{
				Required:    true,
				Description: "Duty cycle (e.g. 75%)",
			},
			"frequency": schema.StringAttribute{
				Required:    true,
				Description: "Frequency in Hz (e.g. 25000)",
			},
		},
	}
}

func (r *pwmResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	if c, ok := req.ProviderData.(*gpioclient.Client); ok {
		r.client = c
	}
}

func (r *pwmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan pwmModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.client.SetPWM(gpioclient.SetPWMArgs{
		Pin:       plan.Pin.ValueString(),
		DutyCycle: plan.DutyCycle.ValueString(),
		Freq:      plan.Frequency.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("PWM create failed", err.Error())
		return
	}

	plan.ID = types.StringValue(apiResp.PinNumber)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *pwmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read RPC is available; keep state
	var state pwmModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Nothing to refresh currently
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *pwmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan pwmModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.SetPWM(gpioclient.SetPWMArgs{
		Pin:       plan.Pin.ValueString(),
		DutyCycle: plan.DutyCycle.ValueString(),
		Freq:      plan.Frequency.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("PWM update failed", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *pwmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state pwmModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set PWM to 0 to effectively disable
	_, err := r.client.SetPWM(gpioclient.SetPWMArgs{
		Pin:       state.Pin.ValueString(),
		DutyCycle: "0%",
		Freq:      "0",
	})
	if err != nil {
		resp.Diagnostics.AddError("PWM delete failed", err.Error())
		return
	}
}
