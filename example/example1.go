package main

import (
	"github.com/ataboo/mlx90614-golang/sensor"
	"github.com/ataboo/mlx90614-golang/config"
	"time"
	"fmt"
)

func main() {
	irSensor := sensor.NewIrSensor(config.DefaultConfig())

	tick := time.Tick(time.Second * 1)
	for {
		select {
		case <-tick:
			irSensor.ReadTemps()
			fmt.Printf(
				"\nAmbient: %s,  Object: %s",
				irSensor.AmbientTemp.FahrenheitPretty(),
				irSensor.ObjectTemp.CelsiusPretty(),
			)
		default:
			//ctrl-c works
		}
	}
}
