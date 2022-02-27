package graphics

import (
	"fmt"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/domain"
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"github.com/polis-interactive/go-lighting-utils/pkg/graphicsShader"
	"log"
	"math"
	"sync"
	"time"
)

type graphics struct {
	grid              util.GridDefinition
	reloadOnUpdate    bool
	pixelSize         int
	graphicsFrequency time.Duration
	gs                *graphicsShader.GraphicsShader
	pb                *util.PixelBuffer
	mu                *sync.RWMutex
	wg                *sync.WaitGroup
	lastTimeStep      time.Time
	speed             float64
	inputMap          graphicsShader.UniformDict
	shaderFiles       []string
	shutdowns         chan struct{}
}

func newGraphics(cfg Config) (*graphics, error) {
	log.Println("graphics, newGraphics: creating")
	pixelSize := cfg.GetGraphicsPixelSize()
	if !cfg.GetGraphicsDisplayOutput() {
		pixelSize = 1
	}

	inputs := cfg.GetInputTypes()
	inputMap := make(graphicsShader.UniformDict)
	for _, input := range inputs {
		if input == domain.InputTypes.SPEED || input == domain.InputTypes.PROGRAM {
			continue
		}
		inputMap[graphicsShader.UniformKey(input)] = 0.0
	}
	inputMap["time"] = 0

	return &graphics{
		grid:              cfg.GetGridDefinition(),
		reloadOnUpdate:    cfg.GetGraphicsReloadOnUpdate(),
		graphicsFrequency: cfg.GetGraphicsFrequency(),
		shaderFiles:       cfg.GetShaderFiles(),
		pixelSize:         pixelSize,
		gs:                nil,
		pb:                nil,
		mu:                &sync.RWMutex{},
		wg:                &sync.WaitGroup{},
		inputMap:          inputMap,
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

func (g *graphics) stepTime() {
	g.mu.Lock()
	defer g.mu.Unlock()
	nt := time.Now()
	timeMultiplier := 32.0*math.Pow(g.speed, 4.0) -
		(45+1/3)*math.Pow(g.speed, 3.0) +
		20*math.Pow(g.speed, 2.0) -
		(1+2/3)*g.speed
	elapsed := nt.Sub(g.lastTimeStep).Seconds() * timeMultiplier
	g.inputMap["time"] += float32(elapsed)
	g.lastTimeStep = nt
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

	g.lastTimeStep = time.Now()
	g.inputMap["time"] = 0.0

	gs, err := g.setupGraphicsShader(int32(gridWidth), int32(gridHeight))
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
				g.stepTime()
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
				g.mu.RLock()
				err = gs.ReadToPixels(g.pb.GetUnsafePointer())
				g.mu.RUnlock()
				if err != nil {
					return err
				}
			}
		}
	}()

	g.gs = nil
	return err

}

func (g *graphics) setupGraphicsShader(width int32, height int32) (*graphicsShader.GraphicsShader, error) {
	gs, err := graphicsShader.NewGraphicsShader(domain.Program, width, height, g.inputMap, g.mu)
	if err != nil {
		return nil, err
	}
	for i, s := range g.shaderFiles {
		err = gs.AttachShader(graphicsShader.ShaderIdentifier{
			Key:      graphicsShader.ShaderKey(rune(i)),
			Filename: s,
		})
		if err != nil {
			return nil, err
		}
	}

	return gs, nil
}
