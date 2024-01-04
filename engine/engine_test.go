package engine_test

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang-simulation-engine/engine"
)

func TestEngineStartStop(t *testing.T) {
	// Initialize the engine with a start date, game year length, an update function, and debug mode disabled
	startDate := "2023-01-01"
	gameYearLength := 1000000 // 1 second in microseconds
	var updateFuncCalled bool
	updateFunc := func(delta float64) {
		updateFuncCalled = true
	}

	eng := engine.New(startDate, gameYearLength, updateFunc, false)

	// Start the engine
	go eng.Start()

	// Let the engine run for some time (e.g., 3 seconds)
	time.Sleep(3 * time.Second)

	// Stop the engine
	eng.Stop()

	// Ensure that the update function was called at least once during the engine run
	assert.True(t, updateFuncCalled, "Update function should have been called")
}

func TestEngineDebugInfo(t *testing.T) {
	// Initialize the engine with a start date, game year length, an update function, and debug mode enabled
	startDate := "2023-01-01"
	gameYearLength := 1000000 // 1 second in microseconds
	updateFunc := func(delta float64) {}

	// Redirect stdout for capturing printed debug information
	old := os.Stdout // Keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	eng := engine.New(startDate, gameYearLength, updateFunc, true)
	eng.PrintDebugInfo()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old // Reset stdout

	// Verify that debug information contains the expected date format
	assert.Contains(t, string(out), "date:", "Debug info should contain date information")
}

func TestDateIncrementation(t *testing.T) {
	// Initialize the engine with a start date, game year length, an update function, and debug mode disabled
	startDate := "2023-01-01"
	gameYearLength := 1000000 // 1 second in microseconds
	updateFunc := func(delta float64) {}

	eng := engine.New(startDate, gameYearLength, updateFunc, false)

	// Run the engine for 5 seconds
	runDuration := 5 * time.Second
	go func() {
		eng.Start()
		time.Sleep(runDuration)
		eng.Stop()
	}()

	// Wait for the engine to finish
	time.Sleep(runDuration + 500*time.Millisecond)

	// Increment the start date by gameYearLength microseconds to compare
	expectedDate := engine.New(startDate, gameYearLength, updateFunc, false).Date().AddMicroseconds(int(runDuration / time.Microsecond))

	// Check if the date incremented correctly
	assert.Equal(t, expectedDate.ToDateString(), eng.Date().ToDateString(), "Date should be incremented correctly")
}

func TestElapsedTimeCalculation(t *testing.T) {
	// Initialize the engine with a start date, game year length, an update function, and debug mode disabled
	startDate := "2023-01-01"
	gameYearLength := 1000000 // 1 second in microseconds
	updateFunc := func(delta float64) {}

	eng := engine.New(startDate, gameYearLength, updateFunc, false)

	// Run the engine for 5 seconds
	runDuration := 5 * time.Second
	go func() {
		eng.Start()
		time.Sleep(runDuration)
		eng.Stop()
	}()

	// Wait for the engine to finish
	time.Sleep(runDuration + 500*time.Millisecond)

	// Check if the elapsed time matches the expected duration
	assert.Equal(t, int(runDuration/time.Second), eng.Elapsed(), "Elapsed time should match the run duration")
}

func TestNewEnginePanicHandling(t *testing.T) {
	// Ensure the 'New' function doesn't panic due to incorrect input format
	defer func() {
		if r := recover(); r != nil {
			// Handle the panic (e.g., set a flag to indicate a panic occurred)
			assert.Fail(t, "Panic occurred during engine initialization")
		}
	}()

	// Simulate incorrect start date format causing a panic
	startDate := "2023/01/01" // Incorrect format (should be "Y-m-d")

	// Initialize the engine (expecting a panic)
	_ = engine.New(startDate, 1000000, nil, false)

	// If a panic occurs, the deferred function will handle it
	// This assertion ensures that the test doesn't fail due to an unhandled panic
	assert.True(t, true, "Test completed without panic")
}
