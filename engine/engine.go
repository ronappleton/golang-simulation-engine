package engine

import (
	"fmt"
	"github.com/golang-module/carbon"
	gameLoop "github.com/kutase/go-gameloop"
)

//var date carbon.Carbon

type Engine struct {
	date           carbon.Carbon
	gameYearLength int
	elapsed        int
	onUpdate       func(float64)
	loop           *gameLoop.GameLoop
	debug          bool
}

func New(startDate string, gameYearLength int, onUpdate func(float64), debug bool) *Engine {
	defer func() {
		if r := recover(); r != nil {
			// Handle the panic (e.g., log, return default value, etc.)
			fmt.Println("Error initializing Engine:", r)
		}
	}()

	date := carbon.ParseByFormat(startDate, "Y-m-d", "Europe/London")

	return &Engine{
		date:           date,
		gameYearLength: gameYearLength,
		elapsed:        0,
		onUpdate:       onUpdate,
		debug:          debug,
	}
}

func (e *Engine) Start() {
	gl := gameLoop.New(1, func(delta float64) {
		e.date = e.date.AddMicroseconds(e.gameYearLength)
		e.elapsed++
		e.onUpdate(delta)
		if e.debug {
			e.PrintDebugInfo()
		}
	})

	e.loop = gl

	gl.Start()
}

func (e *Engine) Stop() {
	e.loop.Stop()
}

func (e *Engine) PrintDebugInfo() {
	fmt.Printf("date: %s, elapsed: %d\n", e.date.ToDateString(), e.elapsed)
}

func (e *Engine) Date() carbon.Carbon {
	return e.date
}

func (e *Engine) Elapsed() int {
	return e.elapsed
}
