package domain

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"image/color"
	"periph.io/x/periph/experimental/devices/ads1x15"
)

type RenderType string

const (
	ws2812Render   = "WS2812_RENDER"
	terminalRender = "TERMINAL_RENDER"
	windowRender   = "WINDOW_RENDER"
)

var RenderTypes = struct {
	WS2812   RenderType
	TERMINAL RenderType
	WINDOW   RenderType
}{
	WS2812:   ws2812Render,
	TERMINAL: terminalRender,
	WINDOW:   windowRender,
}

type RenderService interface {
	Startup()
	Shutdown()
}

type ControlType string

const (
	guiControl ControlType = "GUI_CONTROL"
	adcControl ControlType = "ADC_CONTROL"
)

var ControlTypes = struct {
	GUI ControlType
	ADC ControlType
}{
	GUI: guiControl,
	ADC: adcControl,
}

type ControlService interface {
	Startup()
	Shutdown()
	GetControllerStates() []InputState
}

type InputType string

const (
	brightnessInput InputType = "Brightness"
	attackInput               = "Attack"
	speedInput                = "Speed"
	decayInput                = "Decay"
)

var InputTypes = struct {
	BRIGHTNESS InputType
	ATTACK     InputType
	SPEED      InputType
	DECAY      InputType
}{
	BRIGHTNESS: brightnessInput,
	ATTACK:     attackInput,
	SPEED:      speedInput,
	DECAY:      decayInput,
}

type InputPin struct {
	InputType InputType
	Pin       ads1x15.Channel
}

type InputPair struct {
	InputNumber int
	InputValue  float64
}

type InputState struct {
	InputType  InputType
	InputValue float64
}

type GraphicsService interface {
	Startup()
	Shutdown()
	GetGridColorsNumber() [][]uint32
	GetGridColors() [][]util.Color
	GetGridSysColors() [][]color.RGBA
	HandleInputChange(*InputState)
}
