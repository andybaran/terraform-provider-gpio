package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourcePwm(t *testing.T) {
	t.Skip("skipping acceptance test; requires running GPIO backend")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    []resource.TestStep{{Config: testAccResourcePwm}},
	})
}

const testAccResourcePwm = `
resource "gpio_pwm" "test_pwm" {
	pin = "GPIO13"
	dutycycle = "10%"
	frequency = "25000"
}
`
