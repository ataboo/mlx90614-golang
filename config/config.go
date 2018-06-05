package config

import (
	"time"
	"github.com/op/go-logging"
)

type Config struct {
	RegTAmb byte
	RegTObj1 byte
	RegTObj2 byte
	RegToMax byte
	RegToMin byte
	RegPwmCtrl byte
	RegTaRange byte
	RegKE byte
	RegConfig byte
	RegAddress byte
	RegId0 byte
	RegId1 byte
	RegId2 byte
	RegId3 byte
	RegSleep byte

	ReadTimeout time.Duration
	I2CPath string
    I2CAddr int

	Logger  *logging.Logger
}

func DefaultConfig() *Config {
	return &Config{
		RegTAmb: 0x06,
		RegTObj1: 0x07,
		RegTObj2: 0x08,
		RegToMax: 0x20,
		RegToMin: 0x21,
		RegPwmCtrl: 0x22,
		RegTaRange: 0x23,
		RegKE: 0x24,
		RegConfig: 0x25,
		RegAddress: 0x2E,
		RegId0: 0x3C,
		RegId1: 0x3D,
		RegId2: 0x3E,
		RegId3: 0x3F,
		RegSleep: 0xFF,

		ReadTimeout: 1000,
		I2CPath: "/dev/i2c-1",
		I2CAddr: 0x5A,

		Logger: logging.MustGetLogger("mlx90614"),
	}
}

