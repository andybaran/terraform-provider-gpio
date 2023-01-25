package provider

import (
	"context"

	"github.com/andybaran/terragpio/gpioclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TODO: Lean how to use Diagnostics so I can return a Diagnostics of type INFO or equivalent
func resource_pwm() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource to control PWM Pins",

		CreateContext: resourcePWMCreate,
		ReadContext:   resourcePWMRead,
		UpdateContext: resourcePWMCreate, // Functionally an Update is the same as a Create
		DeleteContext: resourcePWMDelete,

		Schema: map[string]*schema.Schema{
			"pin": {
				// GPIO to be configured for PWM in GPIO standard format (i.e. GPIO6)
				Description: "GPIO Pin",
				Type:        schema.TypeString,
				Required:    true,
			},
			"dutycycle": {
				// Duty cycle for the PWM pin being configured as "nn%" where nn is 00 - 100
				Description: "Duty cycle",
				Type:        schema.TypeString,
				Required:    true,
			},
			"frequency": {
				// Frequency of the signal in the format "nM" where "n" is the numerical value and "M" is Megahertz
				Description: "Frequency",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourcePWMCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(gpioapiclient)

	var pin = d.Get("pin").(string)             //Example: "GPIO12"
	var dutycycle = d.Get("dutycycle").(string) //Example: "100%"
	var freq = d.Get("frequency").(string)      //Example: "25000"

	resp, err := client.c.SetPWM(gpioclient.SetPWMArgs{Pin: pin, DutyCycle: dutycycle, Freq: freq})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.PinNumber)

	return diag.Errorf("Not really an error")
}

func resourcePWMRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return diag.Errorf("not implemented")
}

func resourcePWMDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*gpioclient.Client)

	var pin = d.Get("Pin").(string) //Example: "GPIO12"
	var dutycycle = "0%"
	var freq = "0"

	resp, err := client.SetPWM(gpioclient.SetPWMArgs{Pin: pin, DutyCycle: dutycycle, Freq: freq})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.PinNumber)

	return diag.Errorf("Not really an error")
}
