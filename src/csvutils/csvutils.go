package csvutils

import (
	"blend/tanks"
	"blend/ui"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ CSV Utilities ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

// :===== Input File Handling =====:
func OpenCSV(path string) [][]string {
	// Open input file and check for errors
	fd, error := os.Open(path)
	if error != nil {
		fmt.Println(error)
	}

	// Inform The user that the file has been processed
	ui.ConfirmCSVOpen()

	// Close it to prevent memory Leak
	defer fd.Close()

	// Read csv file
	fileReader := csv.NewReader(fd)
	fileReader.FieldsPerRecord = -1 // This makes sure that no error is thrown due to the first line having a different number of columns than the rest of the file
	records, error := fileReader.ReadAll()

	// Check if error while reading
	if error != nil {
		fmt.Println(error)
	}

	return records
}

// :===== Generating the main objects =====:
func ParseCSV(records [][]string) ([]tanks.Tank, []float32) {
	// Create the basic tanks list and formula
	var Tanks []tanks.Tank
	var formula []float32
	var formulaSum float32

	// Iterate over the CSV and parse the values
	for index, record := range records {
		// Parse the formula which is the first line of the CSV (by design)
		if index == 0 {
			for i, wineProportion := range record {
				proportion, err := strconv.ParseFloat(wineProportion, 32)
				if err != nil {
					fmt.Println("Failed to parse wine proportion", err)
				}
				formula = append(formula, float32(proportion))
				formulaSum += formula[i]
			}
		}

		if formulaSum == 100 {
			// Parse the Tank Capacity
			capacity, err := strconv.ParseUint(record[0], 10, 32)
			if err != nil {
				fmt.Println("Failed to parse Tank Capacity", err)
			}

			// Parse the Wine Number
			wineNumber, err := strconv.ParseUint(record[1], 10, 32)
			if err != nil {
				fmt.Println("Failed to parse Wine Number", err)
			}

			// Create the tank object
			tank := tanks.Tank{
				TankID:     uint16(index),
				Capacity:   uint32(capacity),
				WineNumber: uint32(wineNumber),
			}

			// Append it to a dynamic array
			Tanks = append(Tanks, tank)

		}
	}

	// Check if formula adds up to 100% before continuing
	if formulaSum == 100 {
		// Remove the intruder from the list
		Tanks = Tanks[1:]
		// Return the Tanks slice and formula to the main logic
		return Tanks, formula
	} else {
		// Return the nothing if the formula doesn't add up
		ui.Formula100(formulaSum)
		return nil, nil
	}
}
