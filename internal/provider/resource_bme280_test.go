package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceBme280(t *testing.T) {
	t.Skip("skipping acceptance test; requires running GPIO backend")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    []resource.TestStep{{Config: testAccResourceBme280}},
	})
}

const testAccResourceBme280 = `
resource "gpio_bme280" "test_bme280" {
	i2cbus = "1"
    i2caddr = "0x77"
}
`
