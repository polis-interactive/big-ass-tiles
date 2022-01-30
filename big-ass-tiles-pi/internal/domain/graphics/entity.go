package graphics

import (
	"fmt"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util/shader"
	"log"
	"sync"
	"time"
)

type graphics struct {
	grid              util.GridDefinition
	fileHandle        string
	reloadOnUpdate    bool
	pixelSize         int
	graphicsFrequency time.Duration
	gs                *shader.GraphicsShader
	pb                *util.PixelBuffer
	mu                *sync.RWMutex
	wg                *sync.WaitGroup
	inputTypes        []domain.InputType
	shutdowns         chan struct{}
}

func newGraphics(cfg Config) (*graphics, error) {
	log.Println("graphics, newGraphics: creating")
	shaderName := cfg.GetGraphicsShaderName()
	fileHandle, err := shader.GetShaderQualifiedPath(shaderName, domain.Program)
	if err != nil {
		log.Fatalln(fmt.Sprintf("graphics, newGraphics: couldn't find shader %s; %s", shaderName, err.Error()))
		return nil, err
	}
	pixelSize := cfg.GetGraphicsPixelSize()
	if !cfg.GetGraphicsDisplayOutput() {
		pixelSize = 1
	}
	return &graphics{
		grid:              cfg.GetGridDefinition(),
		fileHandle:        fileHandle,
		reloadOnUpdate:    cfg.GetGraphicsReloadOnUpdate(),
		graphicsFrequency: cfg.GetGraphicsFrequency(),
		pixelSize:         pixelSize,
		gs:                nil,
		pb:                nil,
		mu:                &sync.RWMutex{},
		wg:                &sync.WaitGroup{},
		inputTypes:        cfg.GetInputTypes(),
	}, nil
}

func (g *graphics) startup() {

	log.Println("Graphics, startup; starting")

	if g.shutdowns == nil {
		g.shutdowns = make(chan struct{})
		g.wg.Add(1)
		go g.runMainLoop()
	}

	log.Println("Graphics, startup; started")
}

func (g *graphics) shutdown() {

	log.Println("Graphics, shutdown; shutting down")

	if g.shutdowns != nil {
		close(g.shutdowns)
		g.wg.Wait()
		g.shutdowns = nil
	}
	log.Println("Graphics, shutdown; finished")
}

func (g *graphics) runMainLoop() {
	for {
		err := g.runGraphicsLoop()
		if err != nil {
			log.Println(fmt.Sprintf("Graphics, Main Loop: received error; %s", err.Error()))
		}
		select {
		case _, ok := <-g.shutdowns:
			if !ok {
				goto CloseMainLoop
			}
		case <-time.After(5 * time.Second):
			log.Println("Graphics, Main Loop: retrying window")
		}
	}

CloseMainLoop:
	log.Println("Graphics runMainLoop, Main Loop: closed")
	g.wg.Done()
}

func (g *graphics) runGraphicsLoop() error {

	gridWidth := g.grid.Columns
	gridHeight := g.grid.Rows

	gridWidth = gridWidth * g.pixelSize
	gridHeight = gridHeight * g.pixelSize

	g.mu.Lock()
	g.pb = util.NewPixelBuffer(gridWidth, gridHeight, 0, 0, g.pixelSize)
	g.pb.RawWidth = g.grid.Columns
	g.pb.RawHeight = g.grid.Rows
	g.mu.Unlock()

	inputMap := make(map[string]float32)
	for _, input := range g.inputTypes {
		inputMap[string(input)] = 0.0
	}

	gs, err := shader.NewGraphicsShader(g.fileHandle, gridWidth, gridHeight, inputMap)
	if err != nil {
		return err
	}

	g.gs = gs

	ticker := time.NewTicker(g.graphicsFrequency)

	defer func(g *graphics, t *time.Ticker) {
		t.Stop()
		g.gs.Cleanup()
		g.gs = nil
		g.mu.Lock()
		g.pb.BlackOut()
		g.mu.Unlock()
	}(g, ticker)

	err = func() error {
		for {
			select {
			case _, ok := <-g.shutdowns:
				if !ok {
					return nil
				}
			case <-ticker.C:
				if g.reloadOnUpdate {
					err = g.gs.ReloadShader()
					if err != nil {
						return err
					}
				}
				err = g.gs.RunShader()
				if err != nil {
					return err
				}
				g.mu.Lock()
				err = gs.ReadToPixels(g.pb.GetUnsafePointer())
				g.mu.Unlock()
				if err != nil {
					return err
				}
			}
		}
	}()

	g.gs = nil
	return err

}
