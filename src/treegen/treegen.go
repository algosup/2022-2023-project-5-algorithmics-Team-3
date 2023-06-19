package treegen

import (
	"blend/tanks"
	"fmt"
)

// A substep represents a unidirectional transfer between two tanks
type Substep struct {
	SourceID      uint16
	DestinationID uint16
	Volume        float64
}

type Step struct {
	Substeps []Substep
}

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ Solving Algorithm ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

// :===== The main solving function =====:
func Solve(emptyTanks []tanks.Tank, wineTanks [][]tanks.Tank, formula []float32) []Step {
	steps := []Step{}

	// ===>> Create the list of active tanks
	// ===>>  Try to pour the active tanks in the first elements of empty tanks

	// Go through the full Tanks
	for i, tankGroup := range wineTanks {
		fmt.Printf("tankgroup: %+v\n", tankGroup)
		// Go through the first tank of a winetype
		for _, sourceTank := range tankGroup {
			fmt.Printf("sourceTank: %+v\n", sourceTank)
			// Try to pour in the first empty Tank
			for _, destinationTank := range emptyTanks {
				fmt.Printf("destinationTank: %+v\n", destinationTank)
				// Here we calculate the volume we want to transfer based on the formula
				volume := float64(sourceTank.Capacity) * float64(formula[i])

				// Create a substep for the pour
				substep := Substep{
					SourceID:      sourceTank.TankID,
					DestinationID: destinationTank.TankID,
					Volume:        volume,
				}

				// Now we'll want to add this to the list of steps.
				// For simplicity, we'll assume each pour is its own step
				step := Step{Substeps: []Substep{substep}}
				steps = append(steps, step)
			}
		}
	}

	return steps
}
