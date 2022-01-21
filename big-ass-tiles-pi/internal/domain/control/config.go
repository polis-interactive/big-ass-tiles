package control

import "github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"

type Config interface {
	GetControlType() domain.ControlType
	GetInputTypes() []domain.InputType
}
