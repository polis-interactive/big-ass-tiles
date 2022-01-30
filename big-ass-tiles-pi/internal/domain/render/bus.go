package render

import (
	"fyne.io/fyne/v2"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"image/color"
)

type Bus interface {
	CopyLightsToFyneBuff(buff []fyne.CanvasObject, fyneCreator func(c color.Color) fyne.CanvasObject) error
	CopyLightsToColorBuffer(buff [][]util.Color) error
	CopyLightsToUint32Buffer(buff []uint32) error
}
