package main

import (
	"github.com/ataboo/mlx90614-golang/sensor"
	"github.com/ataboo/mlx90614-golang/config"
	"github.com/op/go-logging"
	"time"
	"fmt"
)

func main() {
	logging.SetLevel(logging.DEBUG, "my-logger")
	log := logging.MustGetLogger("my-logger")
	cnf := config.DefaultConfig()
	cnf.I2CAddr = 0x5b
	cnf.Logger = log

	irSensor := sensor.NewIrSensor(cnf)
	if err := irSensor.Connect(); err != nil {
		log.Fatal("failed to connect")
	}

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
		}

	}
}
