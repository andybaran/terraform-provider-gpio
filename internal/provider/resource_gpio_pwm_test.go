package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGpioPwm(t *testing.T) {
	//	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGpioPwmConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"gpio_pwm.my_fan", "pin", regexp.MustCompile("^GPIO")),
				),
			},
		},
	})
}

const testAccResourceGpioPwmConfig = `
resource "gpio_pwm" "my_fan" {
    pin = "GPIO13"
    dutycycle = "10%"
    frequency = 25000
} 
`
