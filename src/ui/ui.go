package ui

import (
	"blend/tanks"
	"fmt"
)

func ConfirmCSVOpen() {
	fmt.Println("✅ Successfully opened the CSV file")
	fmt.Println("=====================================\n")
}

// :###### DEBUG: Function to display basic initalization debug info ######:
func DebugInit(tanks []tanks.Tank, formula []float32, emptyTanks []tanks.Tank, wineTanks [][]tanks.Tank) {
	fmt.Printf("🛢️  Tanks: %d\n\n", tanks)
	fmt.Println("🏗️  Tank Struct: ID, Capacity, WineNumber\n")
	fmt.Printf("🧪  Formula: %.5f\n", formula)
	fmt.Println("===========================================\n\n")
	fmt.Printf("🍷  Number of wines used: %d\n\n", len(formula))
	fmt.Printf("🕳️🛢️  Sorted Empty Tanks: %d\n\n", emptyTanks)
	fmt.Printf("🍷🛢️  Sorted Wine Tanks: %d\n\n", wineTanks)
}

func Formula100(formulaSum float32) {
	fmt.Printf("❌ Formula adds up to: %.5f\n", formulaSum)
	fmt.Println("Formula should add up to 100%")
	fmt.Println("Please check again")

}

func ProgramHalt() {
	fmt.Println("Program halted due to error.")
}

// Use https://github.com/Nexidian/gocliselect to build the UI
