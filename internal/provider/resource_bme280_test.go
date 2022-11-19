package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceBme280(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceBme280,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"gpio_bme280.test_bme280", "i2cbus", regexp.MustCompile("[1-9][0-9]?"),
					),
					resource.TestMatchResourceAttr(
						"gpio_bme280.test_bme280", "i2caddr", regexp.MustCompile("[:xdigit:]"),
					),
				),
			},
		},
	})
}

const testAccResourceBme280 = `
resource "gpio_bme280" "test_bme280" {
	i2cbus = "1"
    i2caddr = "0x77"
}
`
