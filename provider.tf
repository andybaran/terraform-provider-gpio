terraform {
  required_providers {
    gpio = {
      source = "andy/gpio"
      version = "0.0.1"
    }
  }
}

provider gpio {
    serveraddr = "10.15.21.124:1234"
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

resource "gpio_input_temperature_output_fan" "my_fan_controller" {
    timeinterval = 5
    bme280devicepin = gpio_bme280.my_bme280.id 
    temperaturemax  = 100
    temperaturemin = 15
    fandevice = gpio_pwm.my_fan.id
    dutycyclemax = 100
    dutycyclemin = 10
}