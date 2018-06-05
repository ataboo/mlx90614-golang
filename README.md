# MLX90614-golang
Golang library for reading from an [MLX90614 thermal sensor](https://www.melexis.com/en/documents/documentation/datasheets/datasheet-mlx90614) over i2c.

# Basic Usage
```Go
// Make a config and change any values needed
cnf := config.DefaultConfig()
cnf.I2CPath = "/dev/i2c-2"

// Make a sensor with the config
irSensor := sensor.NewIrSensor(cnf)

// Defer closing the i2c connection
defer irSensor.Close()

// Connect the sensor
if err := irSensor.Connect(); err != nil {
  log.Fatal("failed to connect")
}

// Update tempurature
irSensor.ReadTemps()

// Get temp values
fmt.Printf("\nAmbient temp: %.1f" irSensor.AmbientTemp.Kelvin())
fmt.Println(irSensor.ObjectTemp.CelsiusPretty())
fmt.Println(irSensor.ObjectTemp.FahrenheitPretty())
```
