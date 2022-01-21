package application

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain/control"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain/render"
)

type applicationBus interface {
	Startup()
	Shutdown()
	BindRenderService(r domain.RenderService)
	BindControlService(b domain.ControlService)
	BindGraphicsService(g domain.GraphicsService)
	render.Bus
	control.Bus
}
