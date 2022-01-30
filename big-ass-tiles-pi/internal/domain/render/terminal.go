package render

import (
	"fmt"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"log"
)

type terminalRender struct {
	*baseRender
	output [][]util.Color
}

var _ render = (*terminalRender)(nil)

func newTerminalRender(base *baseRender, cfg Config) *terminalRender {

	log.Println("terminalRender, newTerminalRender: creating")

	r := &terminalRender{
		baseRender: base,
	}

	grid := cfg.GetGridDefinition()
	r.output = make([][]util.Color, grid.Columns)
	for i := 0; i < grid.Columns; i++ {
		r.output[i] = make([]util.Color, grid.Rows)
	}
	
	log.Println("terminalRender, newTerminalRender: created")

	return r
}

func (r *terminalRender) runMainLoop() {
	for {
		err := r.runRenderLoop()
		if err != nil {
			log.Println(fmt.Sprintf("terminal, Main Loop: received error; %s", err.Error()))
		}
		select {
		case _, ok := <-r.shutdowns:
			if !ok {
				goto CloseTerminalLoop
			}
		}
	}

CloseTerminalLoop:
	log.Println("terminalRender runMainLoop, Main Loop: closed")
	r.wg.Done()
}

func (r *terminalRender) runRender() error {

	err := r.bus.CopyLightsToColorBuffer(r.output)
	if err != nil {
		return err
	}

	outputString := "START(\n"

	for i, column := range r.output {
		for j, color := range column {
			outputString += fmt.Sprintf("(%d, %d): %#v", i, j, color)
		}
		outputString += "\n"
	}

	outputString += ")END"

	log.Println(outputString)

	return nil
}
