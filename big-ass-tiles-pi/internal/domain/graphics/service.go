package graphics

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"image/color"
	"log"
	"sync"
)

type service struct {
	graphics   *graphics
	brightness float64
	mu         *sync.Mutex
}

var _ domain.GraphicsService = (*service)(nil)

func NewService(cfg Config) *service {
	log.Println("Graphics, NewService: creating")
	return &service{
		graphics: newGraphics(cfg),
		mu:       &sync.Mutex{},
	}
}

func (s *service) Startup() {
	log.Println("GraphicsService Startup: starting")
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.graphics != nil {
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

func (s *service) GetGridColorsNumber() [][]uint32 {
	s.graphics.mu.RLock()
	defer s.graphics.mu.RUnlock()
	grid := s.graphics.grid
	colors := make([][]uint32, grid.Columns)
	for i := 0; i < grid.Columns; i++ {
		colors[i] = make([]uint32, grid.Rows)
		for j := 0; j < grid.Rows; j++ {
			c := s.graphics.cells[i][j].outputColor.DimColor(s.brightness)
			colors[i][j] = c.ToBits()
		}
	}
	return colors
}

func (s *service) GetGridSysColors() [][]color.RGBA {
	s.graphics.mu.RLock()
	defer s.graphics.mu.RUnlock()
	grid := s.graphics.grid
	colors := make([][]color.RGBA, grid.Columns)
	for i := 0; i < grid.Columns; i++ {
		colors[i] = make([]color.RGBA, grid.Rows)
		for j := 0; j < grid.Rows; j++ {
			c := s.graphics.cells[i][j].outputColor.DimColor(s.brightness)
			colors[i][j] = c.ToSysColor()
		}
	}
	return colors
}

func (s *service) GetGridColors() [][]util.Color {
	s.graphics.mu.RLock()
	defer s.graphics.mu.RUnlock()
	grid := s.graphics.grid
	colors := make([][]util.Color, grid.Columns)
	for i := 0; i < grid.Columns; i++ {
		colors[i] = make([]util.Color, grid.Rows)
		for j := 0; j < grid.Rows; j++ {
			colors[i][j] = s.graphics.cells[i][j].outputColor.DimColor(s.brightness)
		}
	}
	return colors
}

func (s *service) HandleInputChange(state *domain.InputState) {
	if state.InputType == domain.InputTypes.BRIGHTNESS {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.brightness = state.InputValue
	} else {
		s.graphics.mu.Lock()
		defer s.graphics.mu.Unlock()
		switch state.InputType {
		case domain.InputTypes.ATTACK:
			s.graphics.attack = state.InputValue
		case domain.InputTypes.SPEED:
			s.graphics.speed = state.InputValue
		case domain.InputTypes.DECAY:
			s.graphics.decay = state.InputValue
		}
	}
}
