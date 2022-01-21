package main

import (
	"fmt"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain/control"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type testGuiConfig struct{}

var _ control.Config = (*testGuiConfig)(nil)

func (t testGuiConfig) GetControlType() domain.ControlType {
	return domain.ControlTypes.GUI
}

func (t testGuiConfig) GetInputTypes() []domain.InputType {
	return []domain.InputType{
		domain.InputTypes.BRIGHTNESS,
		domain.InputTypes.ATTACK,
		domain.InputTypes.SPEED,
		domain.InputTypes.DECAY,
	}
}

type testGuiBus struct{}

var _ control.Bus = (*testGuiBus)(nil)

func (t testGuiBus) HandleControlInputChange(state *domain.InputState) {
	log.Println(fmt.Sprintf("%s: %f", state.InputType, state.InputValue))
}

func main() {
	conf := &testGuiConfig{}
	bus := &testGuiBus{}
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
