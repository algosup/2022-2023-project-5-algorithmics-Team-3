package main

import (
	"blend/csvutils"
	"blend/tanks"
	"fmt"
)

func main() {

	/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
	಄ะ Part 1: Initialize ะ಄
	ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

	// :===== Open the selected CSV =====:
	records := csvutils.OpenCSV("UseCase2.csv")

	// :===== Parse the CSV =====:
	csvutils.ParseCSV(records)

	// :===== Create the initial Tank SLice =====:
	var Tanks []tanks.Tank

	// :===== Create the initial Formula Slice =====:
	var Formula []float32

	// :===== Parse the CSV to get the Tanks Back, as well as the formula =====:
	Tanks, Formula = csvutils.ParseCSV(records)

	// :===== Create the slice of EmptyTanks =====:
	var EmptyTanks = tanks.GetEmptyTanks(Tanks)

	// :===== Create the slices for each wine Type =====:
	var WineTanks = tanks.GetWineTanks(Tanks)

	// DEBUG :: Print the results //////////////////////////////////
	fmt.Printf("Tanks: %d\n", Tanks)
	fmt.Printf("Formula: %f\n", Formula)
	fmt.Printf("Number of wines used: %d\n", len(Formula))
	fmt.Printf("Empty Tanks: %d\n", EmptyTanks)
	fmt.Printf("Wine Tanks: %d\n", WineTanks)
	// DEBUG :: Print the results /////////////////////////////////

	// ////////////////////////////////////////////////////////////////////////::

	// // :===== Getting the Formula from the user =====:

	// // Get the actual number of unique wines found
	// var wineNumberCount = uint16(len(WineCounter) - 1)

	// // Tell the user
	// fmt.Printf("\nThe program has found %d types of wine in the warehouse.\n", wineNumberCount)

	// // Create the formula array
	// var formula []float32

	// // Use this variable to make sure the percentages add up
	// var totalPercentage float32

	// // Get the percentage for each wine
	// for totalPercentage != 100 {
	// 	for i := uint16(0); i < wineNumberCount; i++ {
	// 		var percentage float32

	// 		// Tell the user to input the percentage
	// 		fmt.Printf("↪ Please enter the percentage of wine n°%d:\n", i+1)

	// 		// Take the users inut (throw error if invalid)
	// 		_, err := fmt.Scan(&percentage)
	// 		if err != nil {
	// 			fmt.Println("Invalid Input:", err)
	// 			return
	// 		}

	// 		// Add the percentage to the formula
	// 		formula = (append(formula, percentage))
	// 		totalPercentage += percentage
	// 	}
	// 	// Check if criterion has been met, restart if not
	// 	if totalPercentage != 100 {
	// 		fmt.Println("=====================================\n")
	// 		fmt.Println("❗ The input percentages do not add up to 100, please try again")
	// 		totalPercentage = 0
	// 		formula = nil
	// 	}
	// }

	// // Inform the user of the formula they have chosen
	// fmt.Printf("The formula you have chosen is %d\n", formula)

	// /*
	// 	for _, element := range formula {
	// 		fmt.Println(element)
	// 	}
	// */

	// fmt.Println("\n=====================================")

	// // Print the sorted slice
	// fmt.Println("\nEmpty Tanks (sorted by capacity):")
	// for _, tank := range EmptyTanks {
	// 	fmt.Printf("TankID: %d, Capacity: %d, Wine Number: %d\n", tank.tankID, tank.capacity, tank.wineNumber)
	// }

	// /*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
	// ಄ะ Combinatorics ะ಄
	// ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

}
