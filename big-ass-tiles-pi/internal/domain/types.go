package domain

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"image/color"
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
	guiControl      ControlType = "GUI_CONTROL"
	physicalControl ControlType = "PHYSICAL_CONTROL"
)

var ControlTypes = struct {
	GUI      ControlType
	PHYSICAL ControlType
}{
	GUI:      guiControl,
	PHYSICAL: physicalControl,
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
