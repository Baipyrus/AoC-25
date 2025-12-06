package day01

type Direction int

const (
	LEFT  Direction = -1
	RIGHT Direction = 1
)

type Instruction struct {
	Dir   Direction
	Steps uint
}
