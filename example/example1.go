package main

import (
	"github.com/ataboo/mlx90614-golang/sensor"
	"github.com/ataboo/mlx90614-golang/config"
	"time"
	"fmt"
)

func main() {
	irSensor := sensor.NewIrSensor(config.DefaultConfig())
	defer irSensor.Close()
	if err := irSensor.Connect(); err != nil {
		//log.Fatal("failed to connect")
	}

	tick := time.Tick(time.Second * 1)
	for i := 0; i < 10; i++ {
		select {
		case <-tick:
			irSensor.ReadTemps()
			fmt.Printf(
				"\nAmbient: %s,  Object: %s",
				irSensor.AmbientTemp.FahrenheitPretty(),
				irSensor.ObjectTemp.CelsiusPretty(),
			)
		}

	}
}
