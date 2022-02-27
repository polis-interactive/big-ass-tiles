package graphics

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"time"
)

type Config interface {
	GetGraphicsReloadOnUpdate() bool
	GetGraphicsDisplayOutput() bool
	GetGraphicsPixelSize() int
	GetShaderFiles() []string
	GetGraphicsFrequency() time.Duration
	GetGridDefinition() util.GridDefinition
	GetInputTypes() []domain.InputType
}
