package sensor

import (
	"fmt"
	"github.com/op/go-logging"

	"time"

	"github.com/ataboo/mlx90614-golang/config"
	"os"
	"github.com/d2r2/go-i2c"
)


type IrSensor struct {
	AmbientTemp Temp
	ObjectTemp Temp
	Config *config.Config
	i2cBus    *i2c.I2C
	connected bool
}

func NewIrSensor(cnf *config.Config) *IrSensor {
	if cnf == nil {
		cnf = config.DefaultConfig()
	}

	sensor := IrSensor{
		Config: cnf,
		connected: false,
		i2cBus: nil,
	}

	return &sensor
}

func (sensor *IrSensor) Connect() error {
	if sensor.connected {
		return fmt.Errorf("already connected to sensor")
	}

	i2CBus, err := i2c.NewI2C(sensor.Config.I2CAddr, sensor.Config.I2CBus)
	if err != nil {
		sensor.log().Error("failed to connect to i2c")
		sensor.log().Error(err)
		if os.IsPermission(err) {
			sensor.log().Info("try again with sudo?")
		}
		return err
	}

	sensor.takeFive()
	sensor.i2cBus = i2CBus

	//if _, err := sensor.readWord(sensor.Config.RegTAmb); err != nil {
	//	sensor.log().Errorf("failed to read on init")
	//	sensor.Close()
	//	return err
	//}

	sensor.connected = true

	return nil
}

func (sensor *IrSensor) ReadTemps() error {
	if !sensor.connected {
		err := fmt.Errorf("call `Connect()` before reading tempuratures")
		sensor.log().Error(err)
		return err
	}

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

func (sensor *IrSensor) Close() {
	sensor.i2cBus.Close()
	sensor.i2cBus = nil
	sensor.connected = false
}

func (sensor *IrSensor) log() *logging.Logger {
	return sensor.Config.Logger
}

func (sensor *IrSensor) readWord(reg byte) (int16, error) {
	buf := make([]byte, 2)

	if _, err := sensor.i2cBus.WriteBytes([]byte{0, reg}); err != nil {
		sensor.log().Error(err.Error())
	}

	sensor.takeFive()

	if _, err := sensor.i2cBus.ReadBytes(buf); err != nil {
		sensor.log().Error(err.Error())
	}

	sensor.log().Infof("Buff: %+v", buf)

	//word, err := sensor.i2cBus.ReadRegS16BE(reg)
	//if err != nil {
	//	sensor.log().Error(fmt.Sprintf("failed to read from register %#x.", reg))
	//	sensor.log().Error(err.Error())
	//	return 0, err
	//}

	return 0, nil
}

func (sensor *IrSensor) writeWord(reg byte, word uint16) error {
	//TODO tries, give up.

	//if err := sensor.i2cBus.WriteReg(reg, []byte{0}); err != nil {
	//	sensor.log().Errorf(fmt.Sprintf("failed to clear register %#x in prep for write", reg))
	//	sensor.log().Errorf(err.Error())
	//	return err
	//}
	//
	//sensor.takeFive()
	//
	//bytes := make([]byte, 2)
	//binary.BigEndian.PutUint16(bytes, word)
	//if err := sensor.i2cBus.WriteReg(reg, bytes); err != nil {
	//	sensor.log().Errorf(fmt.Sprintf("Failed to write to register %#x", reg))
	//	sensor.log().Errorf(err.Error())
	//	return err
	//}
	//
	//sensor.log().Debug(fmt.Sprintf("Wrote %v to register %#x.", bytes, reg))
	//return nil

	return nil
}

func (sensor *IrSensor) takeFive() {
	time.Sleep(5 * time.Millisecond)
}
