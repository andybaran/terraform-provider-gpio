package provider

import (
	"context"

	"github.com/andybaran/terragpio/gpioclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resource_bme280() *schema.Resource {
	return &schema.Resource{
		Description: "Resource to setup BME280 i2c sensor",

		CreateContext: resourceBME280Create,
		ReadContext:   resourceBME280Read,
		//UpdateContext: resourceBME280Create, // Functionally an Update is the same as a Create
		DeleteContext: resourceBME280Delete,

		Schema: map[string]*schema.Schema{
			"i2cbus": {
				// GPIO to be configured for PWM in GPIO standard format (i.e. GPIO6)
				Description: "i2c Bus",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"i2caddr": {
				Description: "bme280 i2c address on bus",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"temperature": {
				Description: "Temperature sensed at time in 'Last_Sensed'",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"humidity": {
				Description: "Humidity sensed at time in 'Last_Sensed'",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"pressure": {
				Description: "Humidity sensed at time in 'Last_Sensed'",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceBME280Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(gpioapiclient)

	var I2CBus = d.Get("i2cbus").(string)   //"1"
	var I2CAddr = d.Get("i2caddr").(string) //"0x77"
	/*var I2CAddrUINT64, err = strconv.ParseUint(I2CAddr, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}*/

	resp, err := client.c.SetBME280(gpioclient.SetBME280Args{I2CBus: I2CBus, I2CAddr: I2CAddr})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.PinNumber)

	return diag.Errorf("Not really an error")

}

func resourceBME280Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("not implemented")
}

func resourceBME280Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("not implemented")
}
