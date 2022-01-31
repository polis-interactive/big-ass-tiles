package render

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
)

type Bus interface {
	CopyLightsToColorBuffer(buff [][]util.Color) error
	CopyLightsToUint32Buffer(buff []uint32) error
}
