package sort

import (
	"blend/tanks"
	"sort"
)

// :===== This function is designed to sort lists of tanks by capacity =====:
func SortTanks(tanks []tanks.Tank) []tanks.Tank {
	sort.Slice(tanks, func(i, j int) bool {
		return tanks[i].TankID < uint16(tanks[j].Capacity)
	})
	return tanks

}

// :===== This function is designed to sort lists of lists of tanks, using the previous function =====:
func SortTanks2D(tanks [][]tanks.Tank) [][]tanks.Tank {
	for i := range tanks {
		tanks[i] = SortTanks(tanks[i])
	}
	return tanks
}
