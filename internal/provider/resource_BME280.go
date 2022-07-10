package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBME280() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource to setup BME280 i2c sensor",

		CreateContext: resourceBME280Create,
		ReadContext:   resourceBME280Read,
		UpdateContext: resourceBME280Create, // Functionally an Update is the same as a Create
		DeleteContext: resourceBME280Delete,

		Schema: map[string]*schema.Schema{
			"I2CBus": {
				// GPIO to be configured for PWM in GPIO standard format (i.e. GPIO6)
				Description: "I2C Bus",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"I2CAddr": {
				Description: "BME280 I2C address on bus",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"Temperature": {
				Description: "Temperature sensed at time in 'Last_Sensed'",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"Humidity": {
				Description: "Humidity sensed at time in 'Last_Sensed'",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"Pressure": {
				Description: "Humidity sensed at time in 'Last_Sensed'",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceBME280Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient)

	var I2CBus = d.Get("I2CBus").(string)   //"1"
	var I2CAddr = d.Get("I2CAddr").(string) //"0x77"

	resp, err = client.MyClient.SetBME280(gpioclient.setBME280Args{I2CBus: I2CBus, I2CAddr: I2CAddr})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.PinNumber)

	return diag.Errorf("Not really an error")

}

//TODO: To really implement this I can create computed fields for temp, barometric pressure, humidity and the time at which they were read
func resourceBME280Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("not implemented")
}

/*TODO: This warrants more investigation into the docs. How do I tear down a connection on the I2CBus
 */

func resourceBME280Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("not implemented")
}
