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

	stop := time.Since(start).Microseconds()
	total := time.Since(entry).Milliseconds()

	fmt.Println("\nExecution time recorded at:")
	fmt.Printf("Main: %d ms\n", total)
	fmt.Printf("Challenge: %d μs\n", stop)
}
