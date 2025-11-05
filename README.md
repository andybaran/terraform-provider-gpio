How will provider be used?

We setup a temperature probe and a fan that reacts to it.

- fan = resource_pwm
- temp = resource_BME280
- reaction = resource_InputTemperature_OutputFan

# gpio_pwm
Terraform resource id (used internally) is equal to the pin attribute.

## Attributes
- pin (string)
- dutycycle (string like "75%")
- frequency (string Hz, e.g. "25000")

## Example Usage

```hcl
resource "gpio_pwm" "my_fan" {
    pin = "GPIO13"
    dutycycle = "10%"
    frequency = "25000"
} 
```

# gpio_bme280
Terraform resource id (used internally) is concantenation of I2C Bus and device address of the connected BME280 device.

## Attributes
- i2cbus (string)
- i2caddr (string like "0x77")

## Example Usage

```hcl
resource "gpio_bme280" "my_bme280" {
    i2cbus = "1"
    i2caddr = "0x77"
} 
```

# gpio_input_temperature_output_fan
Terraform resource id (used internally) is concatentation of id's from resource_bme280 (input) and resource_pwm (output).  An error is thrown if you try to setup more than one of these with the same id since it would be targeting the same devices.

The fans speed is determined by it's duty cycle.  In order to set it's duty cycle in relation to the measured temperature we need to plot it on a graph whose X axis is the range of temperatureMin and temperatureMax and Y axis is dutycycleMin and dutyCycleMax.

## Attributes
- timeinterval (string seconds) how often to check and adjust
- bme280devicepin (string) ID from gpio_bme280 resource
- temperaturemax (string) max Celsius for curve
- temperaturemin (string) min Celsius for curve
- fandevice (string) ID from gpio_pwm resource
- dutycyclemax (string) max duty cycle percent
- dutycyclemin (string) min duty cycle percent

## Example Usage

```hcl
resource "gpio_pwm" "my_fan" {
    pin = "GPIO13"
    dutycycle = "10%"
    frequency = "25000"
} 

resource "gpio_bme280" "my_bme280" {
    i2cbus = "1"
    i2caddr = "0x77"
} 

resource "gpio_input_temperature_output_fan" "my_fan_controller" {
    timeinterval = "5"
    bme280devicepin = gpio_bme280.my_bme280.id 
    temperaturemax  = "100"
    temperaturemin = "15"
    fandevice = gpio_pwm.my_fan.id
    dutycyclemax = "100"
    dutycyclemin = "10"
}
```
