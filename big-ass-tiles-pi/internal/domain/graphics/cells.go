package graphics

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"time"
)

type cell struct {
	baseColor   *util.Color
	hasChange   bool
	outputColor *util.Color
	lastChange  time.Time
}

func newCell(now time.Time) cell {
	return cell{
		baseColor:   &util.Color{W: 255},
		outputColor: &util.Color{W: 255},
		lastChange:  now,
	}
}

func (c *cell) FadeBetween(c2 *util.Color, pct float64, now time.Time) {
	if c.hasChange {
		return
	}
	c.hasChange = true
	c.lastChange = now
	if pct >= 1.0 {
		*c.baseColor = *c2
		*c.outputColor = *c2
		return
	}
	c.outputColor.LerpBetween(c.baseColor, c2, pct)
}

func (c *cell) TryBlackoutCell(now time.Time) {
	c.hasChange = false
	if c.lastChange.Equal(now) {
		return
	}
	c.baseColor.BlackOut()
	c.outputColor.BlackOut()
}

func (c *cell) DoNothing() {
	c.hasChange = false
}

func (c *cell) FadeOut(pct float64, now time.Time) {
	c.hasChange = false
}
