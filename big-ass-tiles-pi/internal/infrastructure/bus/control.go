package bus

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
)

func (b *bus) HandleControlInputChange(state *domain.InputState) {
	b.graphicsService.HandleInputChange(state)
}
