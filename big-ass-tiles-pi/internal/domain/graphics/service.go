package graphics

import (
	"fmt"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"github.com/polis-interactive/go-lighting-utils/pkg/graphicsShader"
	"log"
	"math"
	"sync"
)

type service struct {
	graphics *graphics
	mu       *sync.Mutex
}

var _ domain.GraphicsService = (*service)(nil)

func NewService(cfg Config) (*service, error) {
	log.Println("Graphics, NewService: creating")

	g, err := newGraphics(cfg)
	if err != nil {
		log.Println("Graphics, NewService: error creating graphics")
		return nil, err
	}
	return &service{
		graphics: g,
		mu:       &sync.Mutex{},
	}, nil
}

func (s *service) Startup() {
	log.Println("RenderService Startup: starting")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics != nil {
		s.graphics.startup()
	}
}

func (s *service) Reset() {
	log.Println("RenderService Startup: resetting")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics != nil {
		s.graphics.shutdown()
		s.graphics.startup()
	}
}

func (s *service) Shutdown() {
	log.Println("GraphicsService Shutdown: shutting down")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics != nil {
		s.graphics.shutdown()
	}
}

func (s *service) HandleInputChange(state *domain.InputState) {
	s.graphics.mu.Lock()
	defer s.graphics.mu.Unlock()
	if state.InputType == domain.InputTypes.SPEED {
		s.graphics.speed = state.InputValue
		return
	} else if state.InputType == domain.InputTypes.PROGRAM {
		programCount := float64(len(s.graphics.shaderFiles))
		if programCount == 1 {
			return
		}
		var programKey graphicsShader.ShaderKey
		if state.InputValue == 1 {
			programKey = graphicsShader.ShaderKey(rune(programCount - 1))
		} else {
			selectProgram := int(math.Floor(programCount * state.InputValue))
			programKey = graphicsShader.ShaderKey(rune(selectProgram))
		}
		err := s.graphics.gs.SetShader(programKey)
		if err != nil {
			log.Println(fmt.Sprintf(
				"GraphicsService, HandleInputChange - Program: couldn't set to %s with error %s",
				programKey, err.Error(),
			))
		}
	} else {
		s.graphics.inputMap[graphicsShader.UniformKey(state.InputType)] = float32(state.InputValue)
	}
}

func (s *service) GetPb() (pb *util.PixelBuffer, preLockedMutex *sync.RWMutex) {
	s.graphics.mu.RLock()
	return s.graphics.pb, s.graphics.mu
}
