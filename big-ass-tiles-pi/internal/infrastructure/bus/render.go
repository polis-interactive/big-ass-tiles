package bus

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"sync"
)

func (b *bus) CopyLightsToColorBuffer(rawPbOut [][]util.Color) error {
	pbIn, preLockedGraphicsMutex := b.graphicsService.GetPb()
	defer func(graphicsMu *sync.RWMutex) {
		graphicsMu.RUnlock()
	}(preLockedGraphicsMutex)
	for i := 0; i < pbIn.RawWidth; i++ {
		for j := 0; j < pbIn.RawHeight; j++ {
			rawPbOut[i][j] = pbIn.GetPixel(&util.Point{
				X: i,
				Y: j,
			})
		}
	}
	return nil
}

func (b *bus) CopyLightsToUint32Buffer(pixelToLedMap [][][]int, rawUint32Buff []uint32) error {
	pbIn, preLockedGraphicsMutex := b.graphicsService.GetPb()
	defer func(graphicsMu *sync.RWMutex) {
		graphicsMu.RUnlock()
	}(preLockedGraphicsMutex)

	for i := 0; i < pbIn.RawWidth; i++ {
		for j := 0; j < pbIn.RawHeight; j++ {
			cOut := pbIn.GetPixelPointer(&util.Point{
				X: i,
				Y: j,
			}).ToBits()
			for _, k := range pixelToLedMap[i][j] {
				rawUint32Buff[k] = cOut
			}
		}
	}
	return nil
}
