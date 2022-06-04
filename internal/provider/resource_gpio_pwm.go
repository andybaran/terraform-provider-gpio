package provider

import (
	"context"

	"github.com/andybaran/fictional-goggles/terragpio/gpioclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TODO: Lean how to use Diagnostics so I can return a Diagnostics of type INFO or equivalent
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
				Required:    true,
			},
			"Dutycycle": {
				// Duty cycle for the PWM pin being configured as "nn%" where nn is 00 - 100
				Description: "Duty cycle",
				Type:        schema.TypeString,
				Required:    true,
			},
			"Frequency": {
				// Frequency of the signal in the format "nM" where "n" is the numerical value and "M" is Megahertz
				Description: "Frequency",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourcePWMCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient)

	//idFromAPI := "my-id"
	//d.SetId(idFromAPI)

	var pin = "GPIO12"
	var dutycycle = "100%"
	var freq = "25000"

	resp, err := client.MyClient.SetPWM(gpioclient.SetPWMArgs{Pin: pin, DutyCycle: dutycycle, Freq: freq})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.PinNumber)

	return diag.Errorf("Not really an error")
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

/*TODO: There has to be a better way to do this than just setting values to zero.
Seems like I should be breaking some kind of connection instead, pulling the pin down maybe?  This way is we haven't really deleted the
resource, the values just happen to be set to 0.
*/
func resourcePWMDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient)

	//idFromAPI := "my-id"
	//d.SetId(idFromAPI)

	var pin = "GPIO12"
	var dutycycle = "0%"
	var freq = "0"

	resp, err := client.MyClient.SetPWM(gpioclient.SetPWMArgs{Pin: pin, DutyCycle: dutycycle, Freq: freq})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.PinNumber)

	return diag.Errorf("Not really an error")
}
