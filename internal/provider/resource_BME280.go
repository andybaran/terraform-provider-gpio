package provider

import (
	"context"

	"github.com/andybaran/fictional-goggles/terragpio/gpioclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBME280() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource to setup BME280 i2c sensor",

		CreateContext: resourceBME280Create,
		ReadContext:   resourceBME280Read,
		UpdateContext: resourceBME280Update,
		DeleteContext: resourceBME280Delete,

		Schema: map[string]*schema.Schema{
			"I2CBus": {
				// GPIO to be configured for PWM in GPIO standard format (i.e. GPIO6)
				Description: "I2C Bus",
				Type:        schema.TypeString,
				Required:    true,
			},
			"I2CAddr": {
				Description: "BME280 I2C address on bus",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceBME280Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	var I2CBus = "1"
	var I2CAddr = "0x77"

	client.MyClient.SetBME280(gpioclient.SetBME280Args{I2CBus: I2CBus, I2CAddr: I2CAddr})

	//	client.MyClient.SetPWM(client.MyClient(SetPWMArgs{Pin == pin, DutyCycle == dutycycle, Freq == freq}))

	return diag.Errorf("not implemented")
}

func resourceBME280Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceBME280Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceBME280Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
