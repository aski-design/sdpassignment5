package greenhouse

type WateringSelector struct {
	Weather WeatherSource
}

func (s WateringSelector) Choose(z Zone) WateringStrategy {
	if z.Stage == "seedling" {
		return StageWatering{}
	}
	if s.Weather.Temperature() >= 30 {
		return WeatherWatering{Weather: s.Weather}
	}
	return DeficitWatering{}
}

type VentSelector struct{}

func (VentSelector) Choose(humidity float64, hour int) VentStrategy {
	if humidity >= 85 {
		return HumidityVent{}
	}
	if hour >= 22 || hour < 6 {
		return ThermostatVent{}
	}
	return ProportionalVent{}
}
