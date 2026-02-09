// example_i2c_sht3x_waveshare-rp2040-zero
// https://github.com/tinygo-org/drivers/blob/v0.26.0/examples/sht3x/main.go

package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/sht3x"
)

func main() {

	// for waveshare rp2040 zero I2C0
	// https://tinygo.org/docs/reference/microcontrollers/waveshare-rp2040-zero/
	machine.I2C0.Configure(machine.I2CConfig{
		SDA: machine.GPIO12,
		SCL: machine.GPIO13,
	})

	sensor := sht3x.New(machine.I2C0)

	time.Sleep(2 * time.Second) // for wait USB CDC connect

	for {
		temp, humidity, _ := sensor.ReadTemperatureHumidity()

		fmt.Printf("Temperature: %.2f, Humidity: %.2f \n", float32(temp)/1000, float32(humidity)/100)

		time.Sleep(983 * time.Millisecond)
	}
}
