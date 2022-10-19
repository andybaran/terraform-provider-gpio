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