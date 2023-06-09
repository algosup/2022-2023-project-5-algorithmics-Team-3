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

	// :===== Use Case Selection =====:
	var ucNumber int
	fmt.Printf("Please select a use case number: ")
	// Scan user input & throw error if needed
	_, err := fmt.Scan(&ucNumber)
	if err != nil {
		fmt.Println("Error Reading input:", err)
		return
	}
	fmt.Println("Use case selected:", ucNumber)

	// :===== Open the selected CSV =====:
	records := csvutils.OpenCSV(fmt.Sprintf("UseCase%d.csv", ucNumber))
	// records := csvutils.OpenCSV("UseCase1.csv")

	// :===== Create the initial Tank Slice =====:
	var Tanks []tanks.Tank

	// :===== Create the initial Formula Slice =====:
	var Formula []float64

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

		// :!!!!!! DEBUG: display basic initalization debug info !!!!!!:
		ui.DebugInit(Tanks, Formula, EmptyTanks, WineTanks)

		// :===== Precompute the tank filling ratios =====:
		fillingRatios := treegen.TankFillingRatio(Tanks, Formula)

		// :!!!!!! DEBUG: print the filling Ratios !!!!!!:
		// ui.DebugTankFillingRatios(fillingRatios)

		// :===== CALL THE SOLVER 📞👨‍🔬allo?? =====:
		steps := treegen.Solve(EmptyTanks, WineTanks, Formula, fillingRatios)

		// :!!!!!!  DEBUG: print the final instructions !!!!!!:
		ui.PrintInstructions(steps)

		// fmt.Printf("Steps: %v", steps)

		/*
			instructions := ui.PrintInstructions(steps)
			// Print the steps to the User
			fmt.Printf("Instructions: %v", instructions)
		*/

	} else {
		println("Program halted due to error reading/parsing the tanks.")
	}
}
