package day03

import (
	"strconv"
	"strings"
)

func ParseInput(input string) (batteryBanks []BatteryBank) {
	lines := strings.SplitSeq(input, "\n")

	for line := range lines {
		bline := strings.TrimSpace(line)
		if bline == "" {
			continue
		}

		var row BatteryBank
		for i, battery := range bline {
			joltage, _ := strconv.ParseUint(string(battery), 10, 64)
			row.Batteries = append(row.Batteries, Battery{Joltage: joltage, Index: i})
		}

		batteryBanks = append(batteryBanks, row)
	}

	return batteryBanks
}
