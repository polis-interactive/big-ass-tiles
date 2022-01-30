package graphics

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"log"
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
	log.Println("RenderService Shutdown: shutting down")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics != nil {
		s.graphics.shutdown()
	}
}

func (s *service) HandleInputChange(state *domain.InputState) {
	s.graphics.mu.Lock()
	defer s.graphics.mu.Unlock()
	s.graphics.gs.UpdateUniform(string(state.InputType), float32(state.InputValue))
}

func (s *service) GetPb() (pb *util.PixelBuffer, preLockedMutex *sync.RWMutex) {
	s.graphics.mu.RLock()
	return s.graphics.pb, s.graphics.mu
}
