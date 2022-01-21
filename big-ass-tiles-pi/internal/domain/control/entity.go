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
	impl        controllerImpl
	bus         Bus
	inputStates []domain.InputState
	mu          *sync.RWMutex
	wg          *sync.WaitGroup
	shutdowns   chan struct{}
}

func newController(cfg Config, bus Bus) (*controller, error) {

	inputTypes := cfg.GetInputTypes()
	inputStates := make([]domain.InputState, len(inputTypes))
	for i, inputType := range inputTypes {
		inputStates[i].InputType = inputType
	}

	base := &controller{
		bus:         bus,
		inputStates: inputStates,
		mu:          &sync.RWMutex{},
		wg:          &sync.WaitGroup{},
	}

	switch cfg.GetControlType() {
	case domain.ControlTypes.GUI:
		base.impl = newGuiController(base)
	case domain.ControlTypes.PHYSICAL:
	}

	return base, nil
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
