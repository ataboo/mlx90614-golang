package sensor

import (
	"fmt"
	"github.com/op/go-logging"

	"os/exec"
	"bytes"
	"strings"
	"strconv"
)


type IrSensor struct {
	AmbientTemp Temp
	ObjectTemp Temp
	Config *Config
}

func NewIrSensor(cnf *Config) *IrSensor {
	sensor := IrSensor{
		Config: cnf,
	}

	return &sensor
}

func (sensor *IrSensor) ReadTemps() error {
	ambientRaw, err := sensor.readWord(sensor.Config.RegTAmb)
	if err != nil {
		sensor.log().Error("failed to read ambient temp")
		sensor.log().Error(err)
		return err
	}

	objRaw, err := sensor.readWord(sensor.Config.RegTObj1)
	if err != nil {
		sensor.log().Error("failed to read obj 1 temp")
		sensor.log().Error(err)
		return err
	}

	sensor.AmbientTemp = Temp(ambientRaw)
	sensor.ObjectTemp = Temp(objRaw)

	return nil
}

func (sensor *IrSensor) log() *logging.Logger {
	return sensor.Config.Logger
}

func (sensor *IrSensor) readWord(reg string) (int16, error) {
	cmd := exec.Command("i2cget", "-y", "1", sensor.Config.I2CAddr, reg, "w")

	out := bytes.Buffer{}
	stdErr := bytes.Buffer{}

	cmd.Stdout = &out
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stdErr.String())
		return 0, err
	}

	rawTemp := strings.Trim(out.String(), "\n")
	temp64, err := strconv.ParseInt(rawTemp, 0, 16)

	return int16(temp64), nil
}
