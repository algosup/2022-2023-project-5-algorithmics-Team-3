package main

import (
	"blend/csvutils"
	"blend/sort"
	"blend/tanks"
	"blend/treegen"
	"blend/ui"
	"fmt"
)

func main() {

	/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
	಄ะ Part 1: Initialize ะ಄
	ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

	// :===== Open the selected CSV =====:
	records := csvutils.OpenCSV("UseCase1.csv")

	// :===== Create the initial Tank Slice =====:
	var Tanks []tanks.Tank

	// :===== Create the initial Formula Slice =====:
	var Formula []float32

	// :===== Parse the Records to get the Tanks Back, as well as the formula =====:
	Tanks, Formula = csvutils.ParseCSV(records)

	// :===== Run the program if formula is not nil =====:
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

		// CALL THE SOLVER 📞👨‍🔬
		steps := treegen.Solve(EmptyTanks, WineTanks, Formula)
		fmt.Printf("Steps: %v", steps)

		/*
			instructions := ui.PrintInstructions(steps)
			// Print the steps to the User
			fmt.Printf("Instructions: %v", instructions)
		*/

	} else {
		println("Program halted due to error reading/parsing the tanks.")
	}
}
