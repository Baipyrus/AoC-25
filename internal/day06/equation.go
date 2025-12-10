package day06

import "fmt"

type Equation struct {
	Numbers  []int64
	Operator string
}

func (e *Equation) Evaluate() (result int64) {
	// This is requires a different starting value from 0
	// because obviously starting at 0 would negate all
	// incoming values ...
	if e.Operator == "*" {
		result = 1
	}

	for _, num := range e.Numbers {
		switch e.Operator {
		case "*":
			result *= num
		case "+":
			result += num
		}
	}

	return result
}

func (e Equation) String() (out string) {
	for _, num := range e.Numbers {
		out += fmt.Sprintf("%d %s ", num, e.Operator)
	}

	return out[:len(out)-3] + " = ?"
}
