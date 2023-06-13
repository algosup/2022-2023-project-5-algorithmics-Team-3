package sort

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ Sorting Functions ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

import (
	"blend/tanks"
	"sort"
)

// :===== This function is designed to sort lists of tanks by capacity (decreasing order) =====:
func SortTanks(tanks []tanks.Tank) []tanks.Tank {
	sort.Slice(tanks, func(i, j int) bool {
		return uint16(tanks[i].Capacity) > uint16(tanks[j].Capacity)
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
