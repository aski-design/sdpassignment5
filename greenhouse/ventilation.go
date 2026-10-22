package greenhouse

type VentStrategy interface {
	Opening(temp float64, humidity float64) float64
}

func limit(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return 100
	}
	return v
}

type ThermostatVent struct{}

func (ThermostatVent) Opening(temp float64, humidity float64) float64 {
	if temp > 28 {
		return 100
	}
	return 0
}

type ProportionalVent struct{}

func (ProportionalVent) Opening(temp float64, humidity float64) float64 {
	return limit((temp - 22) * 10)
}

type HumidityVent struct{}

func (HumidityVent) Opening(temp float64, humidity float64) float64 {
	byHeat := limit((temp - 24) * 10)
	byHumidity := limit((humidity - 70) * 5)
	if byHumidity > byHeat {
		return byHumidity
	}
	return byHeat
}
