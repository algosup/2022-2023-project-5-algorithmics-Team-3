package main

import (
	"blend/csvutils"
	"blend/sort"
	"blend/tanks"
	"blend/ui"
)

func main() {

	/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
	಄ะ Part 1: Initialize ะ಄
	ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

	// :===== Open the selected CSV =====:
	records := csvutils.OpenCSV("UseCase2.csv")

	// :===== Create the initial Tank Slice =====:
	var Tanks []tanks.Tank

	// :===== Create the initial Formula Slice =====:
	var Formula []float32

	// :===== Parse the Records to get the Tanks Back, as well as the formula =====:
	Tanks, Formula = csvutils.ParseCSV(records)

	// Run the program id formula is not nil
	if Tanks != nil && Formula != nil {
		// :===== Create the slice of EmptyTanks =====:
		var EmptyTanks = tanks.GetEmptyTanks(Tanks)

		// :===== Create the slices for each wine Type =====:
		var WineTanks = tanks.GetWineTanks(Tanks, len(Formula))

		// :===== Sort both the Wine Tanks and Empty Tanks =====:
		EmptyTanks = sort.SortTanks(EmptyTanks)
		WineTanks = sort.SortTanks2D(WineTanks)

		// :###### DEBUG: display basic initalization debug info ######:
		ui.DebugInit(Tanks, Formula, EmptyTanks, WineTanks)
	} else {
		println("Program halted due to error.")
	}

	/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
	಄ะ Part 2: Combinations ะ಄
	ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/
}
