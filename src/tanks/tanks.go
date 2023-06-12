package tanks

// :===== The main Tank data structure =====:
type Tank struct {
	TankID     uint16
	Capacity   uint32
	WineNumber uint32
}

// :===== This function is meant to retrieve all the empty tanks from the list of tanks =====:
func GetEmptyTanks(Tanks []Tank) []Tank {
	// Create the return slice
	var EmptyTanks []Tank

	// Add all the empty tanks to a slice
	for _, tank := range Tanks {
		if tank.WineNumber == 0 {
			EmptyTanks = append(EmptyTanks, tank)
		}
	}
	return EmptyTanks
}
func GetWineTanks(Tanks []Tank) [][]Tank {
	// Create the return List of list of wines
	var WineTanks [][]Tank

	// Iterate on the Tanks list to regroup tanks by their wine number !-!-!-!/!\BUG/!\!-!-!-!
	for _, element := range Tanks {
		if element.WineNumber != 0 {
			WineTanks[element.WineNumber-1] = append(WineTanks[element.WineNumber-1], element)
		}
	}
	return WineTanks
}

/*

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

*/

/*
// :===== Create the Empty Tanks Object =====:
// Initialize the Slice
var EmptyTanks []Tank

// Add all the empty tanks to the newly created slice
for _, tank := range Tanks {
	if tanks.Tank.WineNumber == 0 {
		EmptyTanks = (append(EmptyTanks, tank))
	}
}

// Sort the empty tanks by tank capacity in decreasing order
sort.Slice(EmptyTanks, func(i, j int) bool {
	return EmptyTanks[i].Capacity > EmptyTanks[j].Capacity
})

/*
	// Print the sorted slice
	fmt.Println("Empty Tanks (sorted by capacity):")
	for _, tank := range EmptyTanks {
		fmt.Printf("TankID: %d, Capacity: %d, Wine Number: %d\n", tank.tankID, tank.capacity, tank.wineNumber)
	}
*/
