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
func Solve(emptyTanks []tanks.Tank, wineTanks [][]tanks.Tank, formula []float32, fillingRatios map[float32][]float32) []Step {

	// :===== Create the list of active tanks =====:
	selectedTanks := make([][]tanks.Tank, 0)
	for _, sublist := range wineTanks {
		if len(sublist) > 0 {
			selectedTanks = append(selectedTanks, []tanks.Tank{sublist[0]})
		}
	}
	fmt.Println("selectedTanks:", selectedTanks)

	/*
		// Try to pour the tanks
		for _, destinationTank := range emptyTanks {
			for i, sourceTank := range selectedTanks {

			}

		}
	*/

	return nil
}

// :===== This function is designed to precompute the tank filling ratios =====:
func TankFillingRatio(tanks []tanks.Tank, formula []float32) map[float32][]float32 {
	fillingRatios := make(map[float32][]float32)

	// Iterate on tanks to get their capacities
	for _, tank := range tanks {
		capacity := float32(tank.Capacity)
		// If it hasn't already been added as a key to the map, add it
		if _, exists := fillingRatios[capacity]; !exists {
			// Calculate the proratas for each capacity
			var proRatas []float32
			for _, element := range formula {
				var proRata float32 = element * float32(tank.Capacity) / 100
				// Add them to the list of proRatas
				proRatas = append(proRatas, proRata)
			}
			// Which itself becomes the value of the current key of the map
			fillingRatios[capacity] = proRatas
		}
	}

	return fillingRatios
}

/* OLD VERSION
// :===== The main solving function =====:
func Solve(emptyTanks []tanks.Tank, wineTanks [][]tanks.Tank, formula []float32) []Step {
	steps := []Step{}


	// ===>> Create the list of active tanks
	// ===>>  Try to pour the active tanks in the first elements of empty tanks

	// Go through the full Tanks
	for i, tankGroup := range wineTanks {
		// fmt.Printf("tankgroup: %+v\n", tankGroup)
		// Go through the first tank of a winetype
		for _, sourceTank := range tankGroup {
			// fmt.Printf("sourceTank: %+v\n", sourceTank)
			// Try to pour in the first empty Tank
			for _, destinationTank := range emptyTanks {
				// fmt.Printf("destinationTank: %+v\n", destinationTank)
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
*/
