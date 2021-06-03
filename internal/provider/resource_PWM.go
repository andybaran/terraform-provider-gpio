package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TODO: This really only be allowed to be set to proper PWM pins and give a warning if pin(s) other than 13 are used
func resourcePWM() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource to control PWM Pins",

		CreateContext: resourcePWMCreate,
		ReadContext:   resourcePWMRead,
		UpdateContext: resourcePWMUpdate,
		DeleteContext: resourcePWMDelete,

		Schema: map[string]*schema.Schema{
			"Pin": {
				// GPIO to be configured for PWM in GPIO standard format (i.e. GPIO6)
				Description: "GPIO Pin",
				Type:        schema.TypeString,
				Optional:    false,
			},
			"Dutycycle": {
				// Duty cycle for the PWM pin being configured as "nn%" where nn is 00 - 100
				Description: "Duty cycle",
				Type:        schema.TypeString,
				Optional:    false,
			},
			"Frequency": {
				// Frequency of the signal in the format "nM" where "n" is the numerical value and "M" is Megahertz
				Description: "Frequency",
				Type:        schema.TypeString,
				Optional:    false,
			},
		},
	}
}

func resourcePWMCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}

func resourcePWMRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourcePWMUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourcePWMDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
