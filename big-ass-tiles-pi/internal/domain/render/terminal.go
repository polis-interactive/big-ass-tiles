package render

import (
	"fmt"
	"log"
)

type terminalRender struct {
	*baseRender
}

var _ render = (*terminalRender)(nil)

func newTerminalRender(base *baseRender) *terminalRender {

	log.Println("terminalRender, newTerminalRender: creating")

	r := &terminalRender{
		baseRender: base,
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

	gridColors := r.bus.GetGridColorsStruct()

	outputString := "START(\n"

	for i, rows := range gridColors {
		for j, color := range rows {
			outputString += fmt.Sprintf("(%d, %d): %#v", i, j, color)
		}
		outputString += "\n"
	}

	outputString += ")END"

	log.Println(outputString)

	return nil
}
