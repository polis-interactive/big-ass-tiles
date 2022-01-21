package graphics

import (
	"github.com/polis-interactive/big-ass-tiles/big-ass-tiles-pi/internal/util"
	"math/rand"
	"time"
)

const (
	MaxSnakeCount = 5
)

type snakes struct {
	*graphics
	speed        float64
	attack       float64
	snakeCount   uint64
	lastTimeStep time.Time
	snakeRunners []snake
}

func newSnakes(g *graphics) *snakes {
	return &snakes{
		graphics: g,
	}
}

func (s *snakes) startSnakes() {
	s.lastTimeStep = time.Now()
}

func timeMultiplier(m float64) float64 {
	return 6*m*m - m
}

func remove(s []snake, i int) []snake {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s *snakes) updateSnakes() {
	speed, attack := func() (float64, float64) {
		s.mu.RLock()
		defer s.mu.RUnlock()
		return s.speed, s.attack
	}()
	nt := time.Now()
	elapsed := int64(float64(nt.Sub(s.lastTimeStep).Milliseconds()) * timeMultiplier(speed))
	s.lastTimeStep = nt
	i := 0
	for i < len(s.snakeRunners) {
		snakeIsAlive := s.snakeRunners[i].tryUpdateSnake(elapsed, attack)
		if !snakeIsAlive {
			s.snakeRunners = remove(s.snakeRunners, i)
			continue
		}
		i += 1
	}
	if i < MaxSnakeCount {
		s.snakeRunners = append(s.snakeRunners, newSnake(s))
	}
}

func (s *snakes) drawSnakes() {
	for i := range s.snakeRunners {
		s.snakeRunners[i].drawSnake()
	}
}

type position struct {
	point *util.Point
	pct   float64
}

type snakeState int

const (
	entering = 1
	riding   = 2
	leaving  = 3
)

var snakeStates = struct {
	ENTERING snakeState
	RIDING   snakeState
	LEAVING  snakeState
}{
	ENTERING: entering,
	RIDING:   riding,
	LEAVING:  leaving,
}

const (
	SnakeMaxLength = 12
	SnakeMinLength = 5
	SnakeMaxStep   = 750
	SnakeMinStep   = 300
)

type snake struct {
	/* global pointers */
	snakes *snakes
	/* generated */
	length int
	color  util.Color
	step   int64
	ride   int
	/* internal */
	state             snakeState
	path              []*position
	accumulatedMillis int64
}

func newSnake(s *snakes) snake {
	sn := snake{
		snakes:            s,
		state:             snakeStates.ENTERING,
		path:              make([]*position, 1),
		accumulatedMillis: 0,
	}
	sn.path[0] = &position{
		point: util.GetRandomPoint(&s.snakes.grid),
	}
	sn.length = rand.Intn(SnakeMaxLength-SnakeMinLength) + SnakeMinLength
	// maybe think of a different way to grab colors, rand?
	sn.color = util.NextColor(s.snakeCount)
	sn.step = int64(rand.Intn(SnakeMaxStep-SnakeMinStep) + SnakeMinStep)
	// this should be random
	sn.ride = 20
	s.snakeCount += 1
	return sn
}

func (s *snake) tryUpdateSnake(elapsed int64, attack float64) bool {
	s.tryStepSnake(elapsed)
	stepPct := float64(s.accumulatedMillis) / float64(s.step)
	snakeIsAlive := false
	for i := range s.path {
		s.updateSnakePath(i, stepPct, attack)
		snakeIsAlive = true
	}
	return snakeIsAlive
}

func (s *snake) tryStepSnake(elapsed int64) {
	s.accumulatedMillis += elapsed
	if s.accumulatedMillis < s.step {
		return
	}
	s.accumulatedMillis -= s.step
	if s.state == snakeStates.ENTERING {
		var nextPoint *util.Point
		if len(s.path) == 1 {
			nextPoint = s.path[0].point.GetNextRandomPoint(&s.snakes.grid)
		} else {
			nextPoint = s.getUniqueNextPoint()
		}
		if nextPoint == nil {
			s.state = snakeStates.LEAVING
			return
		}
		s.path = append(s.path, &position{point: nextPoint})
		if len(s.path) >= s.length {
			s.state = snakeStates.RIDING
		}
	} else if s.state == snakeStates.RIDING {
		nextPoint := s.getUniqueNextPoint()
		if nextPoint == nil {
			s.state = snakeStates.LEAVING
			return
		}
		s.path = s.path[1:]
		s.path = append(s.path, &position{point: nextPoint})
		s.ride -= 1
		if s.ride <= 0 {
			s.state = snakeStates.LEAVING
		}
	} else {
		if len(s.path) <= 1 {
			s.path = []*position{}
		} else {
			s.path = s.path[1:]
		}
	}
}

func (s *snake) getUniqueNextPoint() *util.Point {
	cntr := 0
	for {
		// the snake probably hit itself
		if cntr == 4 {
			return nil
		}
		nextPoint := s.path[len(s.path)-1].point.GetNextRandomPoint(&s.snakes.grid)
		foundMatch := false
		for i := range s.path {
			if s.path[i].point.PointIsEqual(nextPoint) {
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			return nextPoint
		}
		cntr++
	}
}

func (s *snake) updateSnakePath(pos int, stepPct float64, attack float64) {
	// i'm going to need some offset during entering and leaving
	p := s.path[pos]
	if attack >= 1 {
		p.pct = 1.0
	}
}

func (s *snake) drawSnake() {
	for _, pos := range s.path {
		yPos := pos.point.Y
		if yPos >= 0 && yPos < s.snakes.grid.Rows {
			xPos := pos.point.X
			s.snakes.cells[xPos][yPos].FadeBetween(&s.color, pos.pct)
		}
	}
}
