package domain

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"periph.io/x/periph/experimental/devices/ads1x15"
	"sync"
)

const Program = "big-ass-tiles-pi"

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
	guiControl  ControlType = "GUI_CONTROL"
	grpcControl ControlType = "GRPC_CONTROL"
	adcControl  ControlType = "ADC_CONTROL"
)

var ControlTypes = struct {
	GUI  ControlType
	GRPC ControlType
	ADC  ControlType
}{
	GUI:  guiControl,
	GRPC: grpcControl,
	ADC:  adcControl,
}

type ControlService interface {
	Startup()
	Shutdown()
	GetControllerStates() []InputState
}

type InputType string

const (
	brightness InputType = "brightness"
	program              = "program"
	speed                = "speed"
	value                = "value"
)

var InputTypes = struct {
	BRIGHTNESS InputType
	PROGRAM    InputType
	SPEED      InputType
	VALUE      InputType
}{
	BRIGHTNESS: brightness,
	PROGRAM:    program,
	SPEED:      speed,
	VALUE:      value,
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
	HandleInputChange(*InputState)
	GetPb() (pb *util.PixelBuffer, preLockedMutex *sync.RWMutex)
}
