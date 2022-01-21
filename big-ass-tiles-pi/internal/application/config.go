package application

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"periph.io/x/periph/conn/physic"
	"time"
)

type Ws2812Config struct {
	GpioPin   util.GpioPinType
	StripType util.StripType
	Gamma     float32
	LedCount  uint32
}

func (w *Ws2812Config) GetGpioPin() util.GpioPinType {
	return w.GpioPin
}

func (w *Ws2812Config) GetStripType() util.StripType {
	return w.StripType
}

func (w *Ws2812Config) GetGamma() float32 {
	return w.Gamma
}

type RenderConfig struct {
	RenderType      domain.RenderType
	RenderFrequency time.Duration
}

func (r *RenderConfig) GetRenderType() domain.RenderType {
	return r.RenderType
}

func (r *RenderConfig) GetRenderFrequency() time.Duration {
	return r.RenderFrequency
}

type WindowConfig struct {
	TileSize   int
	InputTypes []domain.InputType
}

func (w *WindowConfig) GetTileSize() int {
	return w.TileSize
}

func (w *WindowConfig) GetInputTypes() []domain.InputType {
	return w.InputTypes
}

type AdcConfig struct {
	InputPins     []domain.InputPin
	ReadFrequency physic.Frequency
	ReadVoltage   physic.ElectricPotential
}

func (a *AdcConfig) GetInputPins() []domain.InputPin {
	return a.InputPins
}

func (a *AdcConfig) GetReadFrequency() physic.Frequency {
	return a.ReadFrequency
}
func (a *AdcConfig) GetReadVoltage() physic.ElectricPotential {
	return a.ReadVoltage
}

type ControlConfig struct {
	ControlType    domain.ControlType
	InputTolerance float64
}

func (c *ControlConfig) GetControlType() domain.ControlType {
	return c.ControlType
}

func (c *ControlConfig) GetInputTolerance() float64 {
	return c.InputTolerance
}

type GraphicsConfig struct {
	GraphicsFrequency time.Duration
}

func (g *GraphicsConfig) GetGraphicsFrequency() time.Duration {
	return g.GraphicsFrequency
}

type Config struct {
	*RenderConfig
	*Ws2812Config
	*WindowConfig
	*AdcConfig
	*ControlConfig
	*GraphicsConfig
	GridDefinition util.GridDefinition
}

func (c *Config) GetGridDefinition() util.GridDefinition {
	return c.GridDefinition
}
