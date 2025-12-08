package day03

type Battery struct {
	Joltage uint64
	Index   int
}

type BatteryBank struct {
	Batteries  []Battery
	MaxJoltage uint64
}
