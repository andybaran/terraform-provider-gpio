package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourcePwm(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePwm,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"gpio_bme280.test_bme280", "pin", regexp.MustCompile("GPIO[1-9][0-9]?"),
					),
					resource.TestMatchResourceAttr(
						"gpio_bme280.test_bme280", "dutycycle", regexp.MustCompile("[1-9][0-9]?%"),
					),
					resource.TestMatchResourceAttr(
						"gpio_bme280.test_bme280", "frequency", regexp.MustCompile("[:xdigit:]+"),
					),
				),
			},
		},
	})
}

const testAccResourcePwm = `
resource "gpio_pwm" "test_pwm" {
    pin = "GPIO13"
    dutycycle = "10%"
    frequency = 25000
}
`
