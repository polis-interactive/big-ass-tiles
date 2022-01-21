package graphics

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"time"
)

type Config interface {
	GetGraphicsFrequency() time.Duration
	GetGridDefinition() util.GridDefinition
}
