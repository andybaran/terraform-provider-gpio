package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceInputTemperatureOutputFan(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInputTemperatureOutputFan,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"gpio_input_temperature_output_fan.test_temp_and_fan", "timeinterval", regexp.MustCompile("[:digit:]*"),
					),
					resource.TestMatchResourceAttr(
						"gpio_input_temperature_output_fan.test_temp_and_fan", "bme280devicepin", regexp.MustCompile("[:xdigit:]"),
					),
					resource.TestMatchResourceAttr(
						"gpio_input_temperature_output_fan.test_temp_and_fan", "temperaturemax", regexp.MustCompile("[:digit:]*"),
					),
					resource.TestMatchResourceAttr(
						"gpio_input_temperature_output_fan.test_temp_and_fan", "temperaturemin", regexp.MustCompile("[:digit:]*"),
					),
					resource.TestMatchResourceAttr(
						"gpio_input_temperature_output_fan.test_temp_and_fan", "fandevice", regexp.MustCompile("[1-9][0-9]?"),
					),
					resource.TestMatchResourceAttr(
						"gpio_input_temperature_output_fan.test_temp_and_fan", "dutycyclemax", regexp.MustCompile("[:digit:]*"),
					),
					resource.TestMatchResourceAttr(
						"gpio_input_temperature_output_fan.test_temp_and_fan", "dutycyclemin", regexp.MustCompile("[:digit:]*"),
					),
				),
			},
		},
	})
}

const testAccResourceInputTemperatureOutputFan = `
resource "gpio_input_temperature_output_fan" "test_temp_and_fan" {
    timeinterval = 5
    bme280devicepin = gpio_bme280.my_bme280.id 
    temperaturemax  = 100
    temperaturemin = 15
    fandevice = gpio_pwm.my_fan.id
    dutycyclemax = 100
    dutycyclemin = 10
}

resource "gpio_pwm" "my_fan" {
    pin = "GPIO13"
    dutycycle = "10%"
    frequency = 25000
} 

resource "gpio_bme280" "my_bme280" {
    i2cbus = "1"
    i2caddr = "0x77"
} 
`
