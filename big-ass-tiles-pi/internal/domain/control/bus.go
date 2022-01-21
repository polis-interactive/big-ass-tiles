package control

import "github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"

type Bus interface {
	HandleControlInputChange(*domain.InputState)
}
