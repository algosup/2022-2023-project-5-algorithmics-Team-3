package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Tank struct {
	tankID     uint16
	capacity   uint32
	wineNumber uint32
}

func main() {

	/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
	಄ะ Part 1: Initialize ะ಄
	ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

	// :===== Input File Handling =====:

	// Open input file and check for errors
	fd, error := os.Open("UseCase2.csv")
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
	records, error := fileReader.ReadAll()

	// Check if error while reading
	if error != nil {
		fmt.Println(error)
	}

	// :===== Generating the main objects =====:

	// Create a slice to store the tank objects
	var Tanks []Tank

	// Create a map to store unique wine number values (used for inputting the formula later)
	wineCounter := make(map[uint32]bool)

	// Iterate over the CSV and parse the values
	for index, record := range records {

		// Parse the Tank Capacity
		capacity, err := strconv.ParseUint(record[0], 10, 32)
		if err != nil {
			fmt.Println("Failed to parse Tank Capacity", err)
		}

		// Parse the Wine Number
		wineNumber, err := strconv.ParseUint(record[1], 10, 32)
		if err != nil {
			fmt.Println("Failed to parse Wine Number")
		}

		// Create the tank object
		tank := Tank{
			tankID:     uint16(index),
			capacity:   uint32(capacity),
			wineNumber: uint32(wineNumber),
		}

		// Append it to a dynamic array
		Tanks = append(Tanks, tank)

		// Add the wineNumber to the map if it's not already present
		wineCounter[uint32(wineNumber)] = true

	}

	/*
		// Print the Tanks to the User
			fmt.Println("Tanks:")
			for _, tank := range Tanks {
				fmt.Printf("TankID: %d, Capacity: %d, Wine Number: %d\n", tank.tankID, tank.capacity, tank.wineNumber)
			}
	*/

	// :===== Create the Empty Tanks Object =====:

	// Initialize the Slice
	var EmptyTanks []Tank

	// Add all the empty tanks to the newly created slice
	for _, tank := range Tanks {
		if tank.wineNumber == 0 {
			EmptyTanks = (append(EmptyTanks, tank))
		}
	}

	// Sort the empty tanks by tank capacity in decreasing order
	sort.Slice(EmptyTanks, func(i, j int) bool {
		return EmptyTanks[i].capacity > EmptyTanks[j].capacity
	})

	/*
		// Print the sorted slice
		fmt.Println("Empty Tanks (sorted by capacity):")
		for _, tank := range EmptyTanks {
			fmt.Printf("TankID: %d, Capacity: %d, Wine Number: %d\n", tank.tankID, tank.capacity, tank.wineNumber)
		}
	*/

	// :===== User Interface =====:

	// Inform the user of the number of Tanks found
	fmt.Printf("🛢️  There are currently %d tanks in the warehouse, of which:\n", len(Tanks))

	// Inform the user of the number of Empty Tanks
	fmt.Printf("\t - 🕳️  %d Empty Tanks\n", len(EmptyTanks))

	// Inform the user of the number of Wine Tanks
	fmt.Printf("\t - 🍷 %d Tanks full of wine\n", len(Tanks)-len(EmptyTanks))

	// :===== Getting the Formula from the user =====:

	// Get the actual number of unique wines found
	var wineNumberCount = uint16(len(wineCounter) - 1)

	// Tell the user
	fmt.Printf("\nThe program has found %d types of wine in the warehouse.\n", wineNumberCount)

	// Create the formula array
	var formula []float32

	// Use this variable to make sure the percentages add up
	var totalPercentage float32

	// Get the percentage for each wine
	for totalPercentage != 100 {
		for i := uint16(0); i < wineNumberCount; i++ {
			var percentage float32

			// Tell the user to input the percentage
			fmt.Printf("↪ Please enter the percentage of wine n°%d:\n", i+1)

			// Take the users inut (throw error if invalid)
			_, err := fmt.Scan(&percentage)
			if err != nil {
				fmt.Println("Invalid Input:", err)
				return
			}

			// Add the percentage to the formula
			formula = (append(formula, percentage))
			totalPercentage += percentage
		}
		// Check if criterion has been met, restart if not
		if totalPercentage != 100 {
			fmt.Println("=====================================\n")
			fmt.Println("❗ The input percentages do not add up to 100, please try again")
			totalPercentage = 0
			formula = nil
		}
	}

	// Inform the user of the formula they have chosen
	fmt.Printf("The formula you have chosen is %d\n", formula)

	/*
		for _, element := range formula {
			fmt.Println(element)
		}
	*/

	// :===== Create the Object corresponding to each wine type in the formula =====:

	// Create the Slice of wine tanks (initialize it to avoid out of index error)
	wineTanks := make([][]Tank, wineNumberCount)

	// Iterate on the Tanks list to regroup tanks by their wine type
	for _, element := range Tanks {
		if element.wineNumber != 0 {
			wineTanks[element.wineNumber-1] = append(wineTanks[element.wineNumber-1], element)
		}
	}

	// Print Tanks
	for i, element := range wineTanks {
		fmt.Printf("\nWine n°%d tanks: %d", i+1, element)
	}

	fmt.Println("\n=====================================")

	// Print the sorted slice
	fmt.Println("\nEmpty Tanks (sorted by capacity):")
	for _, tank := range EmptyTanks {
		fmt.Printf("TankID: %d, Capacity: %d, Wine Number: %d\n", tank.tankID, tank.capacity, tank.wineNumber)
	}

	/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
	಄ะ Combinatorics ะ಄
	ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

}
