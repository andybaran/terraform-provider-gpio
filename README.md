# Current State Notes
This provider currently uses the Terraform SDK and should be migrated to the plug-in framework. 

This provider can be tested using the TF files in examples/provider. This folder also contains a .terraformrc to aid in local development line 4 should be changed to the path where you've compiled the provider.  The .terraformrc needs to be in your home directory. 

---


How will provider be used?

We setup a temperature probe and a fan that reacts to it.

- fan = resource_pwm
- temp = resource_BME280
- reaction = resource_InputTemperature_OutputFan

# gpio_pwm
Terraform resource id (used internally) is equal to the pin attribute.

## Attributes
- pin
- dutycycle
- frequency

## Example Usage

```hcl
resource "gpio_pwm" "my_fan" {
    pin = "GPIO13"
    dutycycle = "10%"
    frequency = 25000
} 
```

# gpio_bme280
Terraform resource id (used internally) is concantenation of I2C Bus and device address of the connected BME280 device.

## Attributes
- i2c_bus
- i2c_addr

## Example Usage

```hcl
resource "gpio_bme280" "my_bme280" {
    i2c_bus = "1"
    i2c_addr = "0x77"
} 
```

# gpio_input_temperature_output_fan
Terraform resource id (used internally) is concatentation of id's from resource_bme280 (input) and resource_pwm (output).  An error is thrown if you try to setup more than one of these with the same id since it would be targeting the same devices.

The fans speed is determined by it's duty cycle.  In order to set it's duty cycle in relation to the measured temperature we need to plot it on a graph whose X axis is the range of temperatureMin and temperatureMax and Y axis is dutycycleMin and dutyCycleMax.

## Attributes
- time_interval : int specifying how often (in seconds) to check the temperature and adjust the fan dutycycle (ie: speed)
- bme280_id : the BME280 device to use (this should be from an attribute of resource_bme280)
- temp_max : the max (in celsius) temperature to use for dutycycle calculation
- temp_min : the min (in celsius) temperature to use for dutycycle calculation
- fan_id : the pwm based fan device to use (this should be from an attribute of resource_pwm)
- duty_max : the max (in percent) dutycycle to use
- duty_min : the min (in percent) dutycycle to use

## Example Usage

```hcl
resource "gpio_pwm" "my_fan" {
    pin = "GPIO13"
    dutycycle = "10%"
    frequency = 25000
} 

resource "gpio_bme280" "my_bme280" {
    i2c_bus = "1"
    i2c_addr = "0x77"
} 

resource "gpio_input_temperature_output_fan" "my_fan_controller" {
    time_interval = 5
    bme280_id = gpio_bme280.my_bme280.id 
    temp_max  = 100
    temp_min = 15
    fan_id = gpio_pwm.my_fan.id
    duty_max = 100
    duty_min = 10
}
```
