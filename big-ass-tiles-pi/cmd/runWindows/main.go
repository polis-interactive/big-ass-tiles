package main

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/application"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conf := &application.Config{
		GridDefinition: util.GridDefinition{
			Rows:        3,
			Columns:     11,
			LedPerCell:  2,
			LedPerScoot: 2,
		},
		ControlConfig: &application.ControlConfig{
			ControlType:    domain.ControlTypes.GRPC,
			InputTolerance: 0.001,
			GrpcPort:       5000,
		},
		WindowConfig: &application.WindowConfig{
			TileSize: 0,
		},
		GraphicsConfig: &application.GraphicsConfig{
			ShaderFiles: []string{
				"checkerboard", "stripe-wheel-spread",
				"bar-hoppin",
			},
			ReloadOnUpdate: true,
			DisplayOutput:  true,
			PixelSize:      100,
			Frequency:      33 * time.Millisecond,
		},
		RenderConfig: &application.RenderConfig{
			RenderType:      domain.RenderTypes.TERMINAL,
			RenderFrequency: 10 * time.Second,
		},
		InputTypes: []domain.InputType{
			domain.InputTypes.BRIGHTNESS,
			domain.InputTypes.SPEED,
			domain.InputTypes.PROGRAM,
			domain.InputTypes.VALUE,
		},
	}

	app, err := application.NewApplication(conf)
	if err != nil {
		panic(err)
	}

	err = app.Startup()
	if err != nil {
		log.Println("Main: failed to startup, shutting down")
		err2 := app.Shutdown()
		if err2 != nil {
			log.Println("Main: issue shutting down; ", err2)
		}
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	err = app.Shutdown()
	if err != nil {
		log.Println("Main: issue shutting down; ", err)
	}

}
