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

// :===== This function is meant to retrieve all theWine Tanks from the list of tanks  =====:
func GetWineTanks(Tanks []Tank, len int) [][]Tank {
	// Create the return List of list of wines initalize it to have as many empty spaces as the number of wines in the formula to avoid out of index
	WineTanks := make([][]Tank, len)

	// Iterate on the Tanks list to regroup tanks by their wine number
	for _, tank := range Tanks {
		if tank.WineNumber != 0 {
			WineTanks[tank.WineNumber-1] = append(WineTanks[tank.WineNumber-1], tank)
		}
	}
	return WineTanks
}
