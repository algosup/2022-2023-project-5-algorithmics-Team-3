package csvutils

import (
	"blend/tanks"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// :===== Input File Handling =====:
func OpenCSV(path string) [][]string {
	// Open input file and check for errors
	fd, error := os.Open(path)
	if error != nil {
		fmt.Println(error)
	}

	// Inform The user that the file has been processed
	fmt.Println("✅ Successfully opened the CSV file")
	fmt.Println("=====================================\n")

	// Close it to prevent memory Leak
	defer fd.Close()

	// Read csv file
	fileReader := csv.NewReader(fd)
	// This makes sure that no error is thrown due to the first line having a different number of columns than the rest of the file
	fileReader.FieldsPerRecord = -1
	records, error := fileReader.ReadAll()

	// Check if error while reading
	if error != nil {
		fmt.Println(error)
	}

	return records

}

// :===== Generating the main objects =====:
func ParseCSV(records [][]string) ([]tanks.Tank, []float32) {
	var Tanks []tanks.Tank
	var formula []float32

	// Create a map to store unique wine number values (used for inputting the formula later)
	// wineCounter := make(map[uint32]bool)

	// Iterate over the CSV and parse the values
	for index, record := range records {
		// Parse the formula which is the first line of the CSV (by design)
		if index < 1 {
			for _, wineProportion := range record {
				proportion, err := strconv.ParseFloat(wineProportion, 32)
				if err != nil {
					fmt.Println("Failed to parse wine proportion", err)
				}
				formula = append(formula, float32(proportion))
			}
		}

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

		// Add the wineNumber to the map if it's not already present
		// wineCounter[uint32(wineNumber)] = true

	}

	// Return the Tanks slice and formula to the main logic
	return Tanks, formula

	/* DEBUG
	// Print the Tanks to the User
		fmt.Println("Tanks:")
		for _, tank := range Tanks {
			fmt.Printf("TankID: %d, Capacity: %d, Wine Number: %d\n", tank.tankID, tank.capacity, tank.wineNumber)
		}
	*/
}
