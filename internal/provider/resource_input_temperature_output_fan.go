package provider

import (
	"context"

	"github.com/andybaran/terragpio/gpioclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resource_input_temperature_output_fan() *schema.Resource {
	return &schema.Resource{
		Description: "Resource to control PWM Pins",

		CreateContext: resource_input_temperature_output_fan_Create,
		ReadContext:   resource_input_temperature_output_fan_Read,
		UpdateContext: resource_input_temperature_output_fan_Create, // Functionally the same as a create
		DeleteContext: resource_input_temperature_output_fan_Delete,

		Schema: map[string]*schema.Schema{
			"timeinterval": {
				Description: "How often to read the temperature and adjust the fan duty cycle (speed)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"bme280devicepin": {
				Description: "BME280 device to read temp from",
				Type:        schema.TypeString,
				Required:    true,
			},
			"temperaturemax": {
				Description: "Max temperature (for calculating curve)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"temperaturemin": {
				Description: "Min temperature (for calculating curve)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"fandevice": {
				Description: "fanDevice",
				Type:        schema.TypeString,
				Required:    true,
			},
			"dutycyclemax": {
				Description: "Max fan duty cycle (for calculating curve)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"dutycyclemin": {
				Description: "Min fan duty cycle (for calculating curve)",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resource_input_temperature_output_fan_Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*apiClient)

	var timeInterval = d.Get("timeinterval").(uint64)
	var bme280DevicePin = d.Get("bme280devicepin").(string)
	var temperatureMax = d.Get("temperaturemax").(uint64)
	var temperatureMin = d.Get("temperaturemin").(uint64)
	var dutyCycleMax = d.Get("dutycyclemax").(uint64)
	var dutyCycleMin = d.Get("dutycyclemin").(uint64)
	var fanDevice = d.Get("fandevice").(string)

	resp, err := client.MyClient.StartFanController(gpioclient.StartFanControllerArgs{
		TimeInterval:    timeInterval,
		BME280DevicePin: bme280DevicePin,
		TemperatureMax:  temperatureMax,
		TemperatureMin:  temperatureMin,
		FanDevice:       fanDevice,
		DutyCycleMax:    dutyCycleMax,
		DutyCylceMin:    dutyCycleMin})

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.PinCombo)
	return diag.Errorf("Not really an error")

}

func resource_input_temperature_output_fan_Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resource_input_temperature_output_fan_Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
