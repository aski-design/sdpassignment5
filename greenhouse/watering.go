package greenhouse

type Zone struct {
	Area           float64
	Moisture       float64
	TargetMoisture float64
	Stage          string
}

type WateringStrategy interface {
	Liters(z Zone) float64
}

type DeficitWatering struct{}

func (DeficitWatering) Liters(z Zone) float64 {
	deficit := z.TargetMoisture - z.Moisture
	if deficit < 0 {
		return 0
	}
	return deficit * z.Area * 0.05
}

type StageWatering struct{}

func (StageWatering) Liters(z Zone) float64 {
	norm := 2.0
	switch z.Stage {
	case "seedling":
		norm = 0.5
	case "growing":
		norm = 1.2
	}
	return norm * z.Area
}

type WeatherWatering struct {
	Weather WeatherSource
}

func (w WeatherWatering) Liters(z Zone) float64 {
	evaporation := (w.Weather.Temperature() - 5) * 0.2
	if evaporation < 0 {
		evaporation = 0
	}
	if w.Weather.RainExpected() {
		evaporation = evaporation / 2
	}
	return evaporation * z.Area
}
