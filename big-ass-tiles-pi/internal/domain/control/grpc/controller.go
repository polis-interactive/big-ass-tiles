package grpc

import (
	"sync"
)

type controller interface {
	GetShutdowns() chan struct{}
	GetWg() *sync.WaitGroup
	SetInputValue(inputNumber int, inputValue float64)
}
