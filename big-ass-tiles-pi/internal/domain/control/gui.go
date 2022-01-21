package control

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"log"
	"runtime"
)

type guiController struct {
	*controller
	w fyne.Window
}

var _ controllerImpl = (*guiController)(nil)

func newGuiController(c *controller) *guiController {
	g := &guiController{
		controller: c,
		w:          nil,
	}
	return g
}

func (g *guiController) createNewWindow() {
	var a fyne.App
	if fyne.CurrentApp() != nil {
		a = fyne.CurrentApp()
	} else {
		a = app.New()
	}
	g.w = a.NewWindow("Big Ass Tile Controls")

	inputs := make([]fyne.CanvasObject, len(g.inputStates))
	for i := range g.inputStates {
		inputs[i] = g.newSlider(i)
	}
	grid := container.New(layout.NewGridLayout(2), inputs...)
	content := container.NewVBox(
		layout.NewSpacer(),
		grid,
		layout.NewSpacer(),
	)

	g.w.SetContent(content)
	g.w.Resize(fyne.Size{
		Width:  600,
		Height: 200,
	})
}

func (g *guiController) newSlider(sliderPosition int) fyne.CanvasObject {
	input := g.inputStates[sliderPosition]
	f := 0.0
	data := binding.BindFloat(&f)
	listener := binding.NewDataListener(func() {
		g.controller.mu.Lock()
		defer g.controller.mu.Unlock()
		v := f / 255
		g.inputStates[sliderPosition].InputValue = v
		g.bus.HandleControlInputChange(&domain.InputState{
			InputType:  input.InputType,
			InputValue: v,
		})
	})
	data.AddListener(listener)
	slider := widget.NewSliderWithData(0, 255, data)
	label := widget.NewLabel(string(input.InputType))
	return container.NewVBox(
		label,
		slider,
	)
}

func (g *guiController) runNewWindow() {
	runtime.LockOSThread()
	defer func() {
		runtime.UnlockOSThread()
		g.w = nil
	}()
	g.w.ShowAndRun()
}

func (g *guiController) runMainLoop() {
	log.Println("GuiController, RunMainLoop: running")

	g.createNewWindow()

	go g.runNewWindow()

	select {
	case _, ok := <-g.shutdowns:
		if !ok {
			break
		}
	}

	log.Println("GuiController, RunMainLoop: stopping")

	if fyne.CurrentApp() != nil {
		fyne.CurrentApp().Quit()
	}
	g.wg.Done()

	log.Println("GuiController, RunMainLoop: stopped")
}
