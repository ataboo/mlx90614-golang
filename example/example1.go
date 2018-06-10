package main

import (
	"github.com/ataboo/mlx90614-golang/sensor"
	"time"
	"fmt"
)

func main() {
	config := sensor.DefaultConfig()
	// Change any config values.

	irSensor := sensor.NewIrSensor(config)

	tick := time.Tick(time.Second * 1)
	for {
		select {
		case <-tick:
			irSensor.ReadTemps()
			fmt.Printf(
				"Ambient: %s,  Object: %s\n",
				irSensor.AmbientTemp.FahrenheitPretty(),
				irSensor.ObjectTemp.CelsiusPretty(),
			)
		default:
			//ctrl-c works
		}
	}
}
