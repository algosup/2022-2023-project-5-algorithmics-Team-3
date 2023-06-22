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
func Solve(emptyTanks []tanks.Tank, wineTanks [][]tanks.Tank, formula []float64, fillingRatios map[float64][]float64) []Step {
	var steps []Step
	var step Step
	var fullTanks []tanks.Tank

	/*
		// :===== Create the list of active tanks =====:
		selectedTanks := make([][]tanks.Tank, 0)
		for _, sublist := range wineTanks {
			if len(sublist) > 0 {
				selectedTanks = append(selectedTanks, []tanks.Tank{sublist[len(CHANGEME)-1]})
			}
		}
	*/

	// WAIT A MINUTE, WHAT IF I TRY TO EXTERNALIZE THIS FUNCTION AND CALL IT AGAIN TO REPOPULATE SELECTED TANKS
	// :===== Create the list of active tanks =====:
	selectedTanks := make([][]tanks.Tank, 0)
	selectedTanks = TankSelector(wineTanks, selectedTanks)

	// :===== Create the list of leftover tanks =====:
	leftoverTanks := make([][]tanks.Tank, len(formula))

	fmt.Println(leftoverTanks)

	// Variables to count when to create a new step
	var sourceTanksInStep uint8
	var destinationTanksInStep uint8

	// :===== Try  =====:
	for sourceTanksInStep < 6 && destinationTanksInStep < 6 {

		// Select an empty (destination) Tank
		for _, destinationTank := range emptyTanks {
			fmt.Println(" ⭕ Destination Tank ⭕ ")
			fmt.Println(destinationTank)

			destinationTanksInStep++
			fmt.Println("destinationTanksInStep: ", destinationTanksInStep)

			// If the tank is not filled yet, do:
			if destinationTank.Volume < float64(destinationTank.Capacity) {

				// If there is enough of each wine type (in the sourceTanksByWine), proceed, otherwise go get another tank

				// Reset the substeps for each destination tank
				step.Substeps = nil

				// Select a source Tank from the tank
				for i, sourceTanksByWine := range selectedTanks {

					// Create the source tank
					sourceTank := &sourceTanksByWine[len(sourceTanksByWine)-1]

					sourceTanksInStep++
					fmt.Println(" 🛢️  #️⃣  sourceTanksInStep: ", sourceTanksInStep)
					fmt.Println(" 🛢️  ↩️  sourceTankBeforePour: ", sourceTank)

					// Verify that source has the capacity to pour in destination tank
					if float64(sourceTank.Volume) >= fillingRatios[float64(destinationTank.Capacity)][i] {

						// If maximum reached, remove this tank from emptyTanks and put it in fullTanks
						if destinationTank.Volume == float64(destinationTank.Capacity) {
							fullTanks = append(fullTanks, destinationTank)
							emptyTanks = removeLastTank(emptyTanks)
						}

						fmt.Println("sourceTank.Volume: ")
						fmt.Println(sourceTank.Volume)
						// If the sourceTank has the capacity to pour into destination tank:
						if sourceTank.Volume >= fillingRatios[float64(destinationTank.Capacity)][i] {
							// fmt.Println("ratio:")
							// fmt.Println(fillingRatios[float64(destinationTank.Capacity)][i])
							fmt.Println("Sufficient tank found")

							// Remove the volume from the sourceTank
							fmt.Println("Is about to be removed: ", fillingRatios[float64(destinationTank.Capacity)][i])
							sourceTanksByWine[len(sourceTanksByWine)-1].Volume -= fillingRatios[float64(destinationTank.Capacity)][i]

							// Add the volume to the destination tank HERE HERE
							fmt.Println(destinationTank.Volume)
							destinationTank.Volume += fillingRatios[float64(destinationTank.Capacity)][i]
							fmt.Println(destinationTank.Volume)

						}

						// Add the instruction the substep to the list of substeps
						substep := Substep{SourceID: sourceTank.TankID, DestinationID: destinationTank.TankID, Volume: float64(fillingRatios[float64(destinationTank.Capacity)][i])}
						step.Substeps = append(step.Substeps, substep)

						// Otherwise, if it is empty, add it to emptyTanks and remove it from sourceTanksByWine, then replace it with a fresh one if possible
					} else if sourceTank.Volume == 0.0 {
						fmt.Println("Empty tank found")
						fmt.Println("sourceTanksByWine Before: ")
						fmt.Println(sourceTanksByWine)
						fmt.Println("emptyTanks before: ")
						fmt.Println(emptyTanks)
						fmt.Println("selectedTanks before: ")
						fmt.Println(selectedTanks)
						fmt.Println("++++++++++++++++++++++++++")

						emptyTanks = append(emptyTanks, *sourceTank)
						sourceTanksByWine = removeLastTank(sourceTanksByWine)
						selectedTanks[i] = removeLastTank(selectedTanks[i])
						selectedTanks[i] = append(selectedTanks[i], wineTanks[i][0]) // len(wineTanks)-1

						fmt.Println("sourceTanksByWine After: ")
						fmt.Println(sourceTanksByWine)
						fmt.Println("emptyTanks After: ")
						fmt.Println(emptyTanks)
						fmt.Println("selectedTanks after: ")
						fmt.Println(selectedTanks)

						// Otherwise, add it to leftover tanks where it belongs (in terms of wineType), then replace it with a fresh one if possible
					} else {
						fmt.Println("Leftover tank found")
						fmt.Println("sourceTanksByWine Before: ")
						fmt.Println(sourceTanksByWine)
						fmt.Println("leftoverTanks before: ")
						fmt.Println(leftoverTanks[i])
						fmt.Println("selectedTanks before: ")
						fmt.Println(selectedTanks)
						fmt.Println("++++++++++++++++++++++++++")

						leftoverTanks[i] = append(leftoverTanks[i], *sourceTank)
						sourceTanksByWine = removeLastTank(sourceTanksByWine)
						selectedTanks[i] = append(selectedTanks[i], wineTanks[i][len(wineTanks)-1])

						fmt.Println("sourceTanksByWine After: ")
						fmt.Println(sourceTanksByWine)
						fmt.Println("leftoverTanks After: ")
						fmt.Println(leftoverTanks[i])
						fmt.Println("selectedTanks after: ")
						fmt.Println(selectedTanks)
					}

					fmt.Println(" 🛢️  ↪️  sourceTankAfterPour: ", sourceTank)
					fmt.Println("====================================================")
				}
			} else {
				fullTanks = append(fullTanks, destinationTank)
				fmt.Println(fullTanks)
				emptyTanks = removeLastTank(emptyTanks)
				fmt.Println(emptyTanks)
			}
		}
	}

	steps = append(steps, step)

	// :===== Update the tanks with their new contents =====:
	return steps
}

// :===== This function is designed to precompute the tank filling ratios =====:
func TankFillingRatio(tanks []tanks.Tank, formula []float64) map[float64][]float64 {
	fillingRatios := make(map[float64][]float64)

	// Iterate on tanks to get their capacities
	for _, tank := range tanks {
		capacity := float64(tank.Capacity)

		// If it hasn't already been added as a key to the map, add it
		if _, exists := fillingRatios[capacity]; !exists {

			// Calculate the proratas for each capacity
			var proRatas []float64
			for _, element := range formula {
				var proRata float64 = element * float64(tank.Capacity) / 100

				// Add them to the list of proRatas
				proRatas = append(proRatas, proRata)
			}

			// Which itself becomes the value of the current key of the map
			fillingRatios[capacity] = proRatas
		}
	}

	return fillingRatios
}

// :===== This function is designed to remove the last element of a slice =====:
func removeLastTank(slice []tanks.Tank) []tanks.Tank {
	if len(slice) == 0 {
		return slice // Return the same slice if it's empty
	}
	return slice[:len(slice)-1]
}

// :===== This function is here to select a tanks to be put in the selector =====:
func TankSelector(wineTanks [][]tanks.Tank, selectedTanks [][]tanks.Tank) [][]tanks.Tank {
	for _, sublist := range wineTanks {
		if len(sublist) > 0 {

			// creating new Tank based on the Tank
			activeTank := tanks.Tank{
				TankID:        sublist[len(sublist)-1].TankID,
				Capacity:      sublist[len(sublist)-1].Capacity,
				BlendNewField: sublist[len(sublist)-1].BlendNewField,
				Volume:        float64(sublist[len(sublist)-1].Capacity),
			}
			selectedTanks = append(selectedTanks, []tanks.Tank{activeTank})
		}
	}
	return selectedTanks
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

/* OLDEST VERSION
// :===== The main solving function =====:
func Solve(emptyTanks []tanks.Tank, wineTanks [][]tanks.Tank, formula []float64) []Step {
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

/* OLDER VERSION
// :===== The main solving function =====:
func Solve(emptyTanks []tanks.Tank, wineTanks [][]tanks.Tank, formula []float64, fillingRatios map[float64][]float64) []Step {
	var steps []Step
	var step Step

	// :===== Create the list of active tanks =====:
	selectedTanks := make([][]tanks.Tank, 0)
	for _, sublist := range wineTanks {
		if len(sublist) > 0 {
			selectedTanks = append(selectedTanks, []tanks.Tank{sublist[len(CHANGEME)-1]})
		}
	}
	fmt.Println("")
	fmt.Println("ALL selectedTanks:", selectedTanks)
	fmt.Println("")

	// :===== Try pouring =====:
	// Select an empty (destination) Tank
	for j, destinationTank := range emptyTanks {
		fmt.Println("destinationTank:", destinationTank)
		// Select a source Tank
		for i, sourceTanksByWine := range selectedTanks {
			fmt.Println("sourceTanksByWine:", sourceTanksByWine)
			// Verify if it has the capacity to pour into the destination tank
			fmt.Println("selectedTanks[i][len(CHANGEME)-1]:", selectedTanks[i][len(CHANGEME)-1])
			sourceTank := selectedTanks[i][len(CHANGEME)-1]
			if float64(sourceTank.Capacity) >= fillingRatios[float64(sourceTank.Capacity)][i] {
				substep := Substep{SourceID: sourceTank.TankID, DestinationID: destinationTank.TankID, Volume: float64(fillingRatios[float64(sourceTank.Capacity)][i])}

			}
		}
	}

	return nil
}
*/
