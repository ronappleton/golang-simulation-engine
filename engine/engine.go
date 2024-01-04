package engine

import (
	"fmt"
	"github.com/golang-module/carbon"
	gameLoop "github.com/kutase/go-gameloop"
)

//var SimulationDate carbon.Carbon

type Engine struct {
	SimulationDate carbon.Carbon `json:"simulation_date"`
	ElapsedSeconds int           `json:"elapsed_seconds"`
	gameYearLength int
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
		SimulationDate: date,
		gameYearLength: gameYearLength,
		ElapsedSeconds: 0,
		onUpdate:       onUpdate,
		debug:          debug,
	}
}

func (e *Engine) Start() {
	gl := gameLoop.New(1, func(delta float64) {
		e.SimulationDate = e.SimulationDate.AddMicroseconds(e.gameYearLength)
		e.ElapsedSeconds++
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
	fmt.Printf("SimulationDate: %s, ElapsedSeconds: %d\n", e.SimulationDate.ToDateString(), e.ElapsedSeconds)
}

func (e *Engine) Date() carbon.Carbon {
	return e.SimulationDate
}

func (e *Engine) Elapsed() int {
	return e.ElapsedSeconds
}
