package provider

import (
	"context"

	"github.com/andybaran/fictional-goggles/terragpio/gpioclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


func resource_input_temperature_output_fan() *schema.Resource {
	return &schema.Resource{
		Description: "Resource to control PWM Pins",

		CreateContext: resourceITemp_OFan_Create,
		ReadContext:   resourceITemp_OFan_Read,
		UpdateContext: resourceITemp_OFan_Update,
		DeleteContext: resourceITemp_OFan_Delete,

		Schema: map[string]*schema.Schema{
			"timeInterval": {
				Description: "How often to read the temperature and adjust the fan duty cycle (speed)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"bme280DevicePin": {
				Description: "BME280 device to read temp from",
				Type:        schema.TypeString,
				Required:    true,
			},
			"temperatureMax": {
				Description: "Max temperature (for calculating curve)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"temperatureMin": {
				Description: "Min temperature (for calculating curve)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"fanDevice": {
				Description: "fanDevice",
				Type:        schema.TypeString,
				Required:    true,
			},
			"dutyCycleMax": {
				Description: "Max fan duty cycle (for calculating curve)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"dutyCycleMin": {
				Description: "Min fan duty cycle (for calculating curve)",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceITemp_OFan_Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	var timeInterval = d.Get("timeInterval").(uint64)
	var bme280DevicePin = d.Get("bme280DevicePin").(string)
	var temperatureMax = d.Get("temperatureMax").(uint64)
	var temperatureMin = d.Get("temperatureMin").(uint64)
	var dutyCycleMax = d.Get("dutyCycleMax").(uint64)
	var dutyCycleMin = d.Get("dutyCycleMin").(uint64)

	client.MyClient.StartFanController(gpioclient.StartFanControllerArgs{
		timeInterval:   timeInterval,
		BME280Device:   bme280DevicePin,
		temperatureMax: temperatureMax,
		temperatureMin: temperatureMin,
		fanDevice:      fanDevice,
		dutyCycleMax:   dutyCycleMax,
		dutyCylceMin:   dutyCycleMin})

	//	client.MyClient.SetPWM(client.MyClient(SetPWMArgs{Pin == pin, DutyCycle == dutycycle, Freq == freq}))

	return diag.Errorf("not implemented")
}

func resourceITemp_OFan_Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceITemp_OFan_Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceITemp_OFan_Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
