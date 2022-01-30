package graphics

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"time"
)

type Config interface {
	GetGraphicsShaderName() string
	GetGraphicsReloadOnUpdate() bool
	GetGraphicsDisplayOutput() bool
	GetGraphicsPixelSize() int
	GetGraphicsFrequency() time.Duration
	GetGridDefinition() util.GridDefinition
	GetInputTypes() []domain.InputType
}
