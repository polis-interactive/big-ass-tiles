package render

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"image/color"
)

type Bus interface {
	GetGridColorsNumber() [][]uint32
	GetGridColorsStruct() [][]util.Color
	GetGridSysColors() [][]color.RGBA
}
