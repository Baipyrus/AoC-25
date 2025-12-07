package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Baipyrus/AoC-25/internal/inputs"
	"github.com/Baipyrus/AoC-25/internal/registry"
)

func main() {
	entry := time.Now()
	challenges := registry.Get()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	exec := inputs.GetChallenge(challenges)
	input := inputs.GetInput(cwd)

	start := time.Now()
	exec(input)

	stop := time.Since(start)
	total := time.Since(entry)

	fmt.Println("\nExecution time recorded at:")
	fmt.Printf("Main: %s\n", formatDuration(total))
	fmt.Printf("Challenge: %s\n", formatDuration(stop))
}

// Converts into the most appropriate unit string for human readability.
func formatDuration(d time.Duration) string {
	if d >= time.Second {
		// Use d.Seconds() which returns float64
		return fmt.Sprintf("%.4fs", d.Seconds())
	}

	if d >= time.Millisecond {
		// Divide by the Millisecond constant to get a precise float value
		ms := float64(d) / float64(time.Millisecond)
		return fmt.Sprintf("%.4fms", ms)
	}

	if d >= time.Microsecond {
		// Divide by the Microsecond constant
		us := float64(d) / float64(time.Microsecond)
		return fmt.Sprintf("%.4fµs", us)
	}

	// Default to Nanoseconds
	return fmt.Sprintf("%dns", d.Nanoseconds())
}
