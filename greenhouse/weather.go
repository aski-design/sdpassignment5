package greenhouse

type WeatherSource interface {
	Temperature() float64
	RainExpected() bool
}

type FakeWeather struct {
	Temp float64
	Rain bool
}

func (f FakeWeather) Temperature() float64 {
	return f.Temp
}

func (f FakeWeather) RainExpected() bool {
	return f.Rain
}
