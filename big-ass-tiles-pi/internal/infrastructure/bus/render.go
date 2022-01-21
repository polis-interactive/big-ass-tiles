package bus

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"image/color"
)

func (b *bus) GetGridColorsNumber() [][]uint32 {
	return b.graphicsService.GetGridColorsNumber()
}

func (b *bus) GetGridColorsStruct() [][]util.Color {
	return b.graphicsService.GetGridColors()
}

func (b *bus) GetGridSysColors() [][]color.RGBA {
	return b.graphicsService.GetGridSysColors()
}
