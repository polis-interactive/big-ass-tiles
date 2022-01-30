package application

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain/control"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain/graphics"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain/render"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/infrastructure/bus"
	"log"
	"sync"
)

type Application struct {
	serviceBus   applicationBus
	shutdown     bool
	shutdownLock sync.Mutex
}

func NewApplication(conf *Config) (*Application, error) {

	log.Println("Application, NewApplication: creating")

	/* create application instance */
	app := &Application{
		shutdown: true,
	}

	/* create bus */
	app.serviceBus = bus.NewBus()

	/* create services */

	controlService, err := control.NewService(conf, app.serviceBus)
	if err != nil {
		log.Println("Application, NewApplication: failed to create controller")
		return nil, err
	}
	app.serviceBus.BindControlService(controlService)

	graphicsService, err := graphics.NewService(conf)
	if err != nil {
		log.Println("Application, NewApplication: failed to create graphics")
		return nil, err
	}
	app.serviceBus.BindGraphicsService(graphicsService)

	renderService, err := render.NewService(conf, app.serviceBus)
	if err != nil {
		log.Println("Application, NewApplication: failed to create render")
		return nil, err
	}
	app.serviceBus.BindRenderService(renderService)

	log.Println("Application, NewApplication: created")

	return app, nil

}

func (app *Application) Startup() error {

	log.Println("Application, Startup: starting")

	app.shutdownLock.Lock()
	defer app.shutdownLock.Unlock()
	if app.shutdown == false {
		return nil
	}

	app.shutdown = false

	app.serviceBus.Startup()

	return nil
}

func (app *Application) Shutdown() error {

	log.Println("Application, Shutdown: shutting down")

	app.shutdownLock.Lock()
	defer app.shutdownLock.Unlock()
	if app.shutdown {
		return nil
	}
	app.shutdown = true

	app.serviceBus.Shutdown()

	return nil
}
