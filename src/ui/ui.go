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
	fmt.Printf("🧪  Formula: %f\n", formula)
	fmt.Println("===========================================\n\n")
	fmt.Printf("🍷  Number of wines used: %d\n\n", len(formula))
	fmt.Printf("🕳️🛢️  Sorted Empty Tanks: %d\n\n", emptyTanks)
	fmt.Printf("🍷🛢️  Sorted Wine Tanks: %d\n\n", wineTanks)
}

// Use https://github.com/Nexidian/gocliselect to build the UI
