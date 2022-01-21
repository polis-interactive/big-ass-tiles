package util

import (
	"image/color"
	"math"
)

type Color struct {
	R uint8
	G uint8
	B uint8
	W uint8
}

func (c *Color) DimColor(b float64) Color {
	return Color{
		R: uint8(float64(c.R) * b),
		G: uint8(float64(c.G) * b),
		B: uint8(float64(c.B) * b),
	}
}

func (c *Color) ToBits() uint32 {
	return uint32(c.W)<<24 | uint32(c.R)<<16 | uint32(c.G)<<8 | uint32(c.B)
}

func (c *Color) ToSysColor() color.RGBA {
	return color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: 255,
	}
}

var black = Color{W: 255}
var white = Color{R: 255, G: 255, B: 255, W: 255}

func (c *Color) BlackOut() {
	c.R = 0
	c.G = 0
	c.B = 0
}

func (c *Color) WhiteOut() {
	c.R = 255
	c.G = 255
	c.B = 255
}

func (c *Color) LerpBetween(c1 *Color, c2 *Color, pct float64) {
	if pct <= 0.0 {
		*c = *c1
	} else if pct >= 1.0 {
		*c = *c2
	} else {
		c.R = uint8(float64(c2.R-c1.R)*pct) + c1.R
		c.G = uint8(float64(c2.G-c1.G)*pct) + c1.G
		c.B = uint8(float64(c2.B-c1.B)*pct) + c1.B
	}
}

func (c *Color) FadeWhite(p float64) {
	if p >= 1.0 {
		*c = white
	} else if p <= 0.0 {
		*c = black
	} else {
		c.LerpBetween(&black, &white, p)
	}
}

func (c *Color) FadeBetween(c1 *Color, pct float64) {
	if pct <= 0.0 {
		return
	} else if pct >= 1.0 {
		*c = *c1
	} else {
		c.R = uint8(math.Min(float64(c1.R-c.R)*pct+float64(c.R), 255))
		c.G = uint8(math.Min(float64(c1.G-c.G)*pct+float64(c.G), 255))
		c.B = uint8(math.Min(float64(c1.B-c.B)*pct+float64(c.B), 255))
	}
}

var red = Color{R: 255, W: 255}
var green = Color{G: 255, W: 255}
var blue = Color{B: 255, W: 255}
var fuchsia = Color{R: 255, B: 255, W: 255}
var aqua = Color{G: 255, B: 255, W: 255}
var yellow = Color{R: 255, G: 255, W: 255}

func NextColor(n uint64) Color {
	np := n % 6
	switch np {
	case 0:
		return red
	case 1:
		return green
	case 2:
		return blue
	case 3:
		return fuchsia
	case 4:
		return aqua
	case 5:
		return yellow
	}
	return black
}
