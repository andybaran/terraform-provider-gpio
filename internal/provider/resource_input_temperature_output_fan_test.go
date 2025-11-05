package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceInputTemperatureOutputFan(t *testing.T) {
	t.Skip("skipping acceptance test; requires running GPIO backend")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    []resource.TestStep{{Config: testAccResourceInputTemperatureOutputFan}},
	})
}

const testAccResourceInputTemperatureOutputFan = `
resource "gpio_input_temperature_output_fan" "test_temp_and_fan" {
	timeinterval = "5"
    bme280devicepin = gpio_bme280.my_bme280.id 
	temperaturemax  = "100"
	temperaturemin = "15"
    fandevice = gpio_pwm.my_fan.id
	dutycyclemax = "100"
	dutycyclemin = "10"
}

resource "gpio_pwm" "my_fan" {
    pin = "GPIO13"
    dutycycle = "10%"
	frequency = "25000"
} 

resource "gpio_bme280" "my_bme280" {
    i2cbus = "1"
    i2caddr = "0x77"
} 
`
