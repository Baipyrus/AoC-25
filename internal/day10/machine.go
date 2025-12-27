package day10

import (
	"fmt"
	"strings"
)

type Machine struct {
	State    MachineState
	Buttons  []Button
	Joltages string
}

func (m *Machine) Deserialize(serialized string) error {
	parts := strings.Split(serialized, " ")

	for i, p := range parts {
		startEnd := string(p[0]) + string(p[len(p)-1])
		switch startEnd {
		case "[]":
			// NOTE: Machine state *should* only appear once
			//       per serialized machine string!!!
			machineStateGoal, err := NewMachineState(p)
			if err != nil {
				return fmt.Errorf("Failed to parse state for machine at %d ('%s'): %w", i, p, err)
			}

			m.State = machineStateGoal
		case "()":
			b, err := NewButton(p)
			if err != nil {
				return fmt.Errorf("Failed to parse button for machine at %d ('%s'): %w", i, p, err)
			}

			m.Buttons = append(m.Buttons, b)
		case "{}":
			// NOTE: Machine joltages *should* only appear once
			//       per serialized machine string!!!
			m.Joltages = p
		default:
			return fmt.Errorf("Unknown machine syntax!")
		}
	}

	// NOTE: Technically speaking, you'd have to check if all the
	//       buttons actually reference existing indices for state
	//       indicator lights here ... D:

	return nil
}

func (m Machine) String() string {
	var buttons string
	for _, b := range m.Buttons {
		buttons += fmt.Sprintf("%s ", b)
	}
	buttons = strings.TrimSpace(buttons)

	return fmt.Sprintf("%s %s %s", m.State, buttons, m.Joltages)
}
