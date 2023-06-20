package ui

import (
	"blend/tanks"
	"blend/treegen"
	"fmt"
)

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ UI Functions: Success ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

// :===== Confirm CSV Opening =====:
func ConfirmCSVOpen() {
	fmt.Println("✅ Successfully opened the CSV file")
	fmt.Println("=====================================")
}

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ UI Functions: Failure ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

// :===== Print error for Formula not equal to 100 =====:
func Formula100(formulaSum float32) {
	fmt.Printf("❌ Formula adds up to: %.5f\n", formulaSum)
	fmt.Println("Formula should add up to 100%")
	fmt.Println("Please check again")
}

// :===== Print error for program halt =====:
func ProgramHalt() {
	fmt.Println("Program halted due to error.")
}

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ UI Functions: Main Logic ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

// :===== Print the full list of instructions to the user =====:
func PrintInstructions(steps []treegen.Step) {
	// The final list of steps to be printed to the user
	var instructions [][]string

	// Iterate over the steps and substeps to populate the instructions array
	for stepIndex, step := range steps {
		stepInstructions := []string{}
		for _, substep := range step.Substeps {
			substepInstruction := fmt.Sprintf("Tank %d ==(%.2fhL)==> Tank %d\n", substep.SourceID, substep.Volume, substep.DestinationID)
			stepInstructions = append(stepInstructions, substepInstruction)
			fmt.Println(substepInstruction)
		}
		instructions = append(instructions, stepInstructions)

		// Add step number to the beginning of stepInstructions
		instructions[stepIndex] = append([]string{fmt.Sprintf("Step %d:\n=============\n", stepIndex+1)}, instructions[stepIndex]...)
	}
	// fmt.Println("Instructions: \n", instructions)
}

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ UI Function: DEBUG ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

// :###### DEBUG: Function to display basic initalization debug info ######:
func DebugInit(tanks []tanks.Tank, formula []float32, emptyTanks []tanks.Tank, wineTanks [][]tanks.Tank) {
	fmt.Printf("\n🛢️  Tanks: %d\n\n", tanks)
	fmt.Println("🏗️  Tank Struct: ID, Capacity, WineNumber")
	fmt.Printf("\n🧪  Formula: %.5f\n\n", formula)
	fmt.Println("=====================================")
	fmt.Printf("\n🍷  Number of wines used: %d\n\n", len(formula))
	fmt.Printf("🕳️ 🛢️  Sorted Empty Tanks: %d\n\n", emptyTanks)
	fmt.Printf("🍷🛢️  Sorted Wine Tanks: %d\n\n", wineTanks)
	fmt.Println("=====================================")
	fmt.Println("")
}

func DebugTankFillingRatios(fillingRatios map[float32][]float32) {
	fmt.Println("Filling Ratios:")
	for capacity, ratios := range fillingRatios {
		fmt.Printf("\nCapacity: %.2f\n", capacity)
		fmt.Println("Ratios:", ratios)
		fmt.Println("------------------------")
	}
}

// Use https://github.com/Nexidian/gocliselect to build the UI
