package bus

import (
	"fyne.io/fyne/v2"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"image/color"
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

func (b *bus) CopyLightsToFyneBuff(rawFynebuff []fyne.CanvasObject, fyneCreator func(c color.Color) fyne.CanvasObject) error {
	pbIn, preLockedGraphicsMutex := b.graphicsService.GetPb()
	defer func(graphicsMu *sync.RWMutex) {
		graphicsMu.RUnlock()
	}(preLockedGraphicsMutex)
	height := pbIn.Height
	for i := 0; i < pbIn.RawWidth; i++ {
		for j := 0; j < pbIn.RawHeight; j++ {
			pos := j + height*i
			rawFynebuff[pos] = fyneCreator(
				pbIn.GetPixelPointer(&util.Point{
					X: i,
					Y: j,
				}).ToSysColor(),
			)
		}
	}
	return nil
}

func (b *bus) CopyLightsToUint32Buffer(rawUint32BuffOut []uint32) error {
	pbIn, preLockedGraphicsMutex := b.graphicsService.GetPb()
	defer func(graphicsMu *sync.RWMutex) {
		graphicsMu.RUnlock()
	}(preLockedGraphicsMutex)
	height := pbIn.Height
	for i := 0; i < pbIn.RawWidth; i++ {
		for j := 0; j < pbIn.RawHeight; j++ {
			pos := j + height*i
			rawUint32BuffOut[pos] = pbIn.GetPixelPointer(&util.Point{
				X: i,
				Y: j,
			}).ToBits()
		}
	}
	return nil
}
