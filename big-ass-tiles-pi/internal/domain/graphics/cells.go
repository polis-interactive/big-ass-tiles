package graphics

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"time"
)

type cellState int

const (
	on  = 1
	off = 2
)

var cellStates = struct {
	ON  cellState
	OFF cellState
}{
	ON:  on,
	OFF: off,
}

type cell struct {
	baseColor    util.Color
	decayedColor util.Color
	outputColor  util.Color
	lastChange   time.Time
	state        cellState
}
