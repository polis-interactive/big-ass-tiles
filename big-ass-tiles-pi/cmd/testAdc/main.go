package main

import (
	"fmt"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain/control"
	"log"
	"os"
	"os/signal"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/experimental/devices/ads1x15"
	"syscall"
)

type testAdcConfig struct{}

func (t testAdcConfig) GetInputTolerance() float64 {
	return 0.001
}

var _ control.Config = (*testAdcConfig)(nil)

func (t testAdcConfig) GetControlType() domain.ControlType {
	return domain.ControlTypes.ADC
}

func (t testAdcConfig) GetInputTypes() []domain.InputType {
	return []domain.InputType{
		domain.InputTypes.BRIGHTNESS,
	}
}

func (t testAdcConfig) GetInputPins() []domain.InputPin {
	return []domain.InputPin{
		{
			InputType: domain.InputTypes.BRIGHTNESS,
			Pin:       ads1x15.Channel0,
		},
	}
}

func (t testAdcConfig) GetReadFrequency() physic.Frequency {
	return physic.Hertz * 1
}

func (t testAdcConfig) GetReadVoltage() physic.ElectricPotential {
	return physic.MilliVolt * 3300
}

type testAdcBus struct{}

var _ control.Bus = (*testAdcBus)(nil)

func (t testAdcBus) HandleControlInputChange(state *domain.InputState) {
	log.Println(fmt.Sprintf("%s: %f", state.InputType, state.InputValue))
}

func main() {
	conf := &testAdcConfig{}
	bus := &testAdcBus{}
	c, err := control.NewService(conf, bus)
	if err != nil {
		log.Println("Unable to create control service")
	}
	c.Startup()

	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	<-s

	c.Shutdown()
}
