package day10

import (
	"log"
	"strings"
)

func ParseInput(input string) (machines []*Machine) {
	lines := strings.SplitSeq(input, "\n")

	for line := range lines {
		bline := strings.TrimSpace(line)
		if bline == "" {
			continue
		}

		var machine Machine

		// NOTE: I know, I know ... I don't usually do error
		//       handling in AoC entries like this, but I was
		//       already writing all this "clean" looking
		//       boilerplate and so I just had to keep passing
		//       the error upward! :D
		err := machine.Deserialize(bline)
		if err != nil {
			log.Fatal(err)
		}

		machines = append(machines, &machine)
	}

	return machines
}
