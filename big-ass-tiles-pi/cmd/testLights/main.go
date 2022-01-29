package main

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain/render"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"image/color"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type testRenderConfig struct{}

func (t testRenderConfig) GetGpioPin() util.GpioPinType {
	return util.GpioPinTypes.GPIO18
}

func (t testRenderConfig) GetStripType() util.StripType {
	return util.StripTypes.WS2811RGB
}

func (t testRenderConfig) GetGamma() float32 {
	return 1.2
}

func (t testRenderConfig) GetTileSize() int {
	panic("unused")
}

func (t testRenderConfig) GetRenderType() domain.RenderType {
	return domain.RenderTypes.WS2812
}

func (t testRenderConfig) GetRenderFrequency() time.Duration {
	return time.Millisecond * 33
}

func (t testRenderConfig) GetGridDefinition() util.GridDefinition {
	return util.GridDefinition{
		Rows:         3,
		Columns:      11,
		LedPerCell:   8,
		LedPerScoot:  2,
		RowExtension: 0,
	}
}

var _ render.Config = (*testRenderConfig)(nil)

type testRenderBus struct {
	colors    [][]uint32
	increment int
}

func makeColors() [][]uint32 {
	columns := make([][]uint32, 11)
	for i := 0; i < 11; i++ {
		rows := make([]uint32, 3)
		columns[i] = rows
	}
	return columns
}

func min(i int, j int) int {
	if i <= j {
		return i
	} else {
		return j
	}
}

func max(i int, j int) int {
	if i <= j {
		return j
	} else {
		return i
	}
}

func (t testRenderBus) GetGridColorsNumber() [][]uint32 {
	for i := 0; i < 11; i++ {
		for j := 0; j < 3; j++ {
			position := t.increment + i + 10*j%256
			t.colors[i][j] = util.WheelUint32(position)
		}
	}
	return t.colors
}

func (t testRenderBus) GetGridColorsStruct() [][]util.Color {
	panic("unused")
}

func (t testRenderBus) GetGridSysColors() [][]color.RGBA {
	panic("unused")
}

var _ render.Bus = (*testRenderBus)(nil)

func main() {
	conf := &testRenderConfig{}
	bus := &testRenderBus{
		colors:    makeColors(),
		increment: 0,
	}
	r, err := render.NewService(conf, bus)
	if err != nil {
		log.Println("Unable to create control service")
	}
	r.Startup()

	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	<-s

	r.Shutdown()
}
