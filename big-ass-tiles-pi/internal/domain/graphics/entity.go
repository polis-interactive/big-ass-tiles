package graphics

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"log"
	"sync"
	"time"
)

type graphics struct {
	*snakes
	grid            util.GridDefinition
	decay           float64
	cells           [][]cell
	updateFrequency time.Duration
	mu              *sync.RWMutex
	wg              *sync.WaitGroup
	shutdowns       chan struct{}
}

func newGraphics(conf Config) *graphics {
	now := time.Now()
	g := conf.GetGridDefinition()
	cells := make([][]cell, g.Columns)
	for i := 0; i < g.Columns; i++ {
		rows := make([]cell, g.Rows)
		for j := 0; j < g.Rows; j++ {
			rows[j] = newCell(now)
		}
		cells[i] = rows
	}
	gr := &graphics{
		grid:            g,
		cells:           cells,
		updateFrequency: conf.GetGraphicsFrequency(),
		mu:              &sync.RWMutex{},
		wg:              &sync.WaitGroup{},
		shutdowns:       nil,
	}

	gr.snakes = newSnakes(gr)
	return gr
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
		select {
		case <-time.After(g.updateFrequency):
			g.updateGraphics()
		case _, ok := <-g.shutdowns:
			if !ok {
				goto CloseMainLoop
			}
		}
	}
CloseMainLoop:
	log.Println("Graphics runMainLoop, Main Loop: closed")
	g.wg.Done()
}

func (g *graphics) updateGraphics() {
	g.updateSnakes()
	g.mu.Lock()
	defer g.mu.Unlock()
	now := time.Now()
	g.drawSnakes(now)
	g.applyDecay(now)
}

func (g *graphics) applyDecay(now time.Time) {
	if g.decay <= 0 {
		for i := 0; i < g.grid.Columns; i++ {
			for j := 0; j < g.grid.Rows; j++ {
				g.cells[i][j].DoNothing()
			}
		}
		return
	}
	if g.decay == 1 {
		for i := 0; i < g.grid.Columns; i++ {
			for j := 0; j < g.grid.Rows; j++ {
				g.cells[i][j].TryBlackoutCell(now)
			}
		}
		return
	}
	pct := (1 - g.decay) * 0.05
	for i := 0; i < g.grid.Columns; i++ {
		for j := 0; j < g.grid.Rows; j++ {
			g.cells[i][j].FadeOut(pct, now)
		}
	}
}
