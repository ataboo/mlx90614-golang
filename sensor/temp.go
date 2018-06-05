package sensor

import "fmt"

type Temp uint16

func (t Temp) Kelvin() float32 {
	return float32(t) * 0.02
}

func (t Temp) KelvinPretty() string {
	return fmt.Sprintf("%.2f°K", t.Kelvin())
}

func (t Temp) Celsius() float32 {
	return t.Kelvin() - 273.15
}

func (t Temp) CelsiusPretty() string {
	return fmt.Sprintf("%.2f°C", t.Celsius())
}

// deprecated? ;P
func (t Temp) Fahrenheit() float32 {
	return t.Kelvin() * 9/5 - 459.67
}

func (t Temp) FahrenheitPretty() string {
	return fmt.Sprintf("%.2f°F", t.Fahrenheit())
}

func (t Temp) String() string {
	return t.KelvinPretty()
}