package greenhouse

type WateringContext struct {
	strategy WateringStrategy
}

func NewWateringContext(s WateringStrategy) *WateringContext {
	return &WateringContext{strategy: s}
}

func (c *WateringContext) SetStrategy(s WateringStrategy) {
	c.strategy = s
}

func (c *WateringContext) Water(z Zone) float64 {
	return c.strategy.Liters(z)
}

type VentContext struct {
	strategy VentStrategy
}

func NewVentContext(s VentStrategy) *VentContext {
	return &VentContext{strategy: s}
}

func (c *VentContext) SetStrategy(s VentStrategy) {
	c.strategy = s
}

func (c *VentContext) Open(temp float64, humidity float64) float64 {
	return c.strategy.Opening(temp, humidity)
}
