package sensor

import (
	"github.com/op/go-logging"
)

type Config struct {
	RegTAmb string
	RegTObj1 string
	RegTObj2 byte

	// `/dev/i2c-1` => 1
	I2CBus      int
    I2CAddr     string

	Logger  *logging.Logger
}

func DefaultConfig() *Config {
	return &Config{
		RegTAmb:  "0x06",
		RegTObj1: "0x07",
		I2CBus:   1,
		I2CAddr:  "0x5A",
		Logger:   logging.MustGetLogger("mlx90614"),
	}
}

