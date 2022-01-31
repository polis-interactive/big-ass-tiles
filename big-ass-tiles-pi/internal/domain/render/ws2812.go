package render

import (
	"fmt"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	"log"
	"math"
	"time"
)

type ws2812Render struct {
	*baseRender
	brightness uint8
	channel    int
	options    *ws2811.Option
	strip      *ws2811.WS2811
	mapLed     [][][]int
}

var _ render = (*ws2812Render)(nil)

func newWs2812Render(base *baseRender, cfg ws2812RenderConfig) *ws2812Render {

	log.Println("ws2812Render, newWs2812Render: creating")

	pinNumber := cfg.GetGpioPin()

	options := ws2811.DefaultOptions
	options.Channels[0].GpioPin = int(pinNumber)
	options.Channels[0].StripeType = int(cfg.GetStripType())
	options.Channels[0].Brightness = 0
	options.Channels[0].Gamma = nil
	options.Channels[0].LedCount = base.grid.Rows * base.grid.Columns * base.grid.LedPerCell
	channel := 0
	if cfg.GetGamma() != 1 {
		options.Channels[0].Gamma = util.MakeGammaTable(float64(cfg.GetGamma()))
	}
	if pinNumber == util.GpioPinTypes.GPIO19 ||
		pinNumber == util.GpioPinTypes.GPIO13 {
		options.Channels = append([]ws2811.ChannelOption{{}}, options.Channels...)
		options.Channels[0].GpioPin = 18
		options.Channels[0].LedCount = 0
		channel = 1
	}

	r := &ws2812Render{
		baseRender: base,
		brightness: 0,
		options:    &options,
		channel:    channel,
		strip:      nil,
		mapLed:     nil,
	}

	r.generateLeds()

	log.Println("ws2812Render, newWs2812Render: created")
	return r
}

func (r *ws2812Render) generateLeds() {
	ledCount := r.options.Channels[r.channel].LedCount
	columns := r.baseRender.grid.Columns
	rows := r.baseRender.grid.Rows
	ledsPerCell := r.baseRender.grid.LedPerCell

	r.mapLed = make([][][]int, columns)
	for i := 0; i < columns; i++ {
		r.mapLed[i] = make([][]int, rows)
		for j := 0; j < rows; j++ {
			r.mapLed[i][j] = make([]int, 0)
		}
	}

	ledsPerRow := columns * ledsPerCell
	ledsPerColumnScoot := r.baseRender.grid.LedPerScoot
	ledsPerRowScoot := ledsPerColumnScoot * columns
	for i := 0; i < ledCount; i++ {
		row := int(float64(i) / float64(ledsPerRow))
		isOddRowScoot := int(math.Floor(float64(i)/float64(ledsPerRowScoot)))%2 == 1
		nominalColumn := int(math.Floor(float64(i)/float64(ledsPerColumnScoot))) % columns

		var actualColumn int
		if isOddRowScoot {
			actualColumn = columns - nominalColumn - 1
		} else {
			actualColumn = nominalColumn
		}
		r.mapLed[actualColumn][row] = append(r.mapLed[actualColumn][row], i)
	}
}

func (r *ws2812Render) runMainLoop() {

	log.Println("ws2812Render, Main Loop: running")

	for {
		err := func(r *ws2812Render) error {
			dev, err := ws2811.MakeWS2811(r.options)
			if err != nil {
				return err
			}
			err = dev.Init()
			if err != nil {
				return err
			}
			defer dev.Fini()
			r.strip = dev
			err = r.runRenderLoop()
			if err != nil {
				return err
			}
			r.tryBlackoutStrip()
			return nil
		}(r)
		r.strip = nil
		if err != nil {
			log.Println(fmt.Sprintf("ws2812Render, Main Loop: received error; %s", err.Error()))
		}
		select {
		case _, ok := <-r.shutdowns:
			if !ok {
				goto CloseWs2812Loop
			}
		case <-time.After(5 * time.Second):
			log.Println("ws2812Render, Main Loop: retrying connection")
		}
	}

CloseWs2812Loop:
	log.Println("ws2812Render, Main Loop: closed")
	r.wg.Done()
}

func (r *ws2812Render) runRender() error {

	err := r.bus.CopyLightsToUint32Buffer(r.mapLed, r.strip.Leds(r.channel))
	if err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	err = r.strip.Render()
	return err
}

func (r *ws2812Render) tryBlackoutStrip() {
	if r.strip == nil {
		log.Println("ws2812Render, tryBlackoutStrip: couldn't do it, strip is null")
		return
	}
	leds := r.strip.Leds(0)
	for i, _ := range leds {
		leds[i] = 0
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	err := r.strip.Render()
	if err != nil {
		log.Println(fmt.Sprintf("ws2812Render, tryBlackoutStrip: failed for some reason; %s", err.Error()))
	}
}
