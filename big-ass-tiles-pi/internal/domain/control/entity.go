package control

import (
	"fmt"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"log"
	"reflect"
	"sync"
)

type controllerImpl interface {
	runMainLoop()
}

type controller struct {
	impl           controllerImpl
	bus            Bus
	inputStates    []domain.InputState
	inputTolerance float64
	mu             *sync.RWMutex
	wg             *sync.WaitGroup
	shutdowns      chan struct{}
}

func newController(cfg Config, bus Bus) (*controller, error) {

	base := &controller{
		bus:            bus,
		inputTolerance: cfg.GetInputTolerance(),
		mu:             &sync.RWMutex{},
		wg:             &sync.WaitGroup{},
	}

	var err error = nil
	switch cfg.GetControlType() {
	case domain.ControlTypes.GUI:
		base.impl = newGuiController(base, cfg)
	case domain.ControlTypes.ADC:
		base.impl, err = newAdcController(base, cfg)
	}

	return base, err
}

func (c *controller) startup() {

	log.Println(fmt.Sprintf("%s, startup: starting", reflect.TypeOf(c.impl)))

	if c.shutdowns == nil {
		c.shutdowns = make(chan struct{})
		c.wg.Add(1)
		go c.impl.runMainLoop()
		log.Println(fmt.Sprintf("%s, startup: running", reflect.TypeOf(c.impl)))
	}
}

func (c *controller) shutdown() {
	log.Println(fmt.Sprintf("%s, shutdown: shutting down", reflect.TypeOf(c.impl)))
	if c.shutdowns != nil {
		close(c.shutdowns)
		c.wg.Wait()
		c.shutdowns = nil
	}
	log.Println(fmt.Sprintf("%s, shutdown: done", reflect.TypeOf(c.impl)))
}

func (c *controller) setInputValue(inputNumber int, inputValue float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	input := c.inputStates[inputNumber]
	oldVal := input.InputValue
	positiveDiff := oldVal + c.inputTolerance
	negativeBound := oldVal - c.inputTolerance
	if positiveDiff > inputValue && inputValue < negativeBound {
		return
	}
	input.InputValue = inputValue
	c.bus.HandleControlInputChange(&domain.InputState{
		InputType:  input.InputType,
		InputValue: inputValue,
	})
}
