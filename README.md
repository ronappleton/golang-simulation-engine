# golang-simulation-engine
Simple simulation engine using Gregorian calendar

## Example Usage

```go
package main

import (
	"github.com/ronappleton/golang-simulation-engine/engine"
	"github.com/ronappleton/golang-simulation-engine/realtime"
	"os"
	"os/signal"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	simulation := engine.New("2023-01-01", realtime.OneHour, func(delta float64) {
		// Do stuff with your simulation
    }, true)

	simulation.Start()

	for {
		select {
		case <-signals:
			simulation.Stop()
			return
		}
	}
}
```