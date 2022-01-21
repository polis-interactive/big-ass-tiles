package render

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"log"
	"runtime"
)

type windowRender struct {
	*baseRender
	w        fyne.Window
	tileSize float32
}

var _ render = (*windowRender)(nil)

func newWindowRender(b *baseRender, cfg windowRenderConfig) *windowRender {
	r := &windowRender{
		baseRender: b,
		w:          nil,
		tileSize:   float32(cfg.GetTileSize()),
	}
	return r
}

func (r *windowRender) createNewWindow() {
	var a fyne.App
	if fyne.CurrentApp() != nil {
		a = fyne.CurrentApp()
	} else {
		a = app.New()
	}
	r.w = a.NewWindow("Big Ass Tile Output")

	tileCount := r.grid.Columns * r.grid.Rows

	tiles := make([]fyne.CanvasObject, tileCount)
	for i := 0; i < tileCount; i++ {
		tiles[i] = r.newTileAnyColor(i)
	}
	grid := container.New(layout.NewGridLayout(r.grid.Columns), tiles...)
	content := container.NewVBox(
		layout.NewSpacer(),
		grid,
		layout.NewSpacer(),
	)

	r.w.SetContent(content)
}

func (r *windowRender) newTileAnyColor(i int) fyne.CanvasObject {
	c := color.RGBA{A: 255}
	if i%2 == 0 {
		c.R = 255
	}
	if i%3 == 0 {
		c.G = 255
	}
	if i%5 == 0 {
		c.B = 255
	}

	rect := canvas.NewRectangle(c)
	rect.SetMinSize(fyne.Size{
		Width:  r.tileSize,
		Height: r.tileSize,
	})
	return rect
}

func (r *windowRender) newTile(c color.Color) fyne.CanvasObject {
	rect := canvas.NewRectangle(c)
	rect.SetMinSize(fyne.Size{
		Width:  r.tileSize,
		Height: r.tileSize,
	})
	return rect
}

func (r *windowRender) runNewWindow() {
	runtime.LockOSThread()
	defer func() {
		runtime.UnlockOSThread()
	}()
	r.w.ShowAndRun()
}

func (r *windowRender) runMainLoop() {
	log.Println("WindowRender, RunMainLoop: running")

	r.createNewWindow()

	go r.runNewWindow()

	err := r.runRenderLoop()
	if err != nil {
		log.Println(fmt.Sprintf("terminal, Main Loop: received error; %s", err.Error()))
	}

	log.Println("WindowRender, RunMainLoop: stopping")

	if fyne.CurrentApp() != nil {
		fyne.CurrentApp().Quit()
	}
	r.wg.Done()

	log.Println("WindowRender, RunMainLoop: stopped")
}

func (r *windowRender) runRender() error {

	g := r.bus.GetGridSysColors()

	tileCount := r.grid.Columns * r.grid.Rows

	tiles := make([]fyne.CanvasObject, tileCount)

	for j := 0; j < r.grid.Rows; j++ {
		for i := 0; i < r.grid.Columns; i++ {
			posOut := i + (r.grid.Rows-j-1)*r.grid.Columns
			tiles[posOut] = r.newTile(g[i][j])
		}
	}

	grid := container.New(layout.NewGridLayout(r.grid.Columns), tiles...)
	content := container.NewVBox(
		layout.NewSpacer(),
		grid,
		layout.NewSpacer(),
	)

	r.w.SetContent(content)

	return nil
}
