package tanks

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ Tanks Logic & Structureะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

/* DEPRECATED
type OldTank struct {
	TankID     uint16
	Capacity   uint32
	WineNumber uint32
}
*/

// :===== The main Tank data structure =====:
type Tank struct {
	TankID   uint16
	Capacity uint32
	Blend    []float64
	Volume   float64
}

// :===== This function is meant to retrieve all the empty tanks from the list of tanks =====:
func GetEmptyTanks(Tanks []Tank) []Tank {
	// Create the return slice
	var EmptyTanks []Tank

	// Add all the empty tanks to a slice
	for _, tank := range Tanks {
		var wineProportionSum float64 = 0
		for _, wineProportion := range tank.Blend {
			wineProportionSum += wineProportion
		}
		if wineProportionSum == 0 {
			EmptyTanks = append(EmptyTanks, tank)
		}
	}
	return EmptyTanks
}

// :===== This function is meant to retrieve all theWine Tanks from the list of tanks  =====:
func GetWineTanks(Tanks []Tank, len int) [][]Tank {
	// Create the List of list of wines; initalize it to have as many empty spaces as the number of wines in the formula to avoid out of index error
	WineTanks := make([][]Tank, len)

	// Iterate on the Tanks list to regroup tanks by their wine number
	for _, tank := range Tanks {
		for i, wineProportion := range tank.Blend {
			if wineProportion != 0.0 {
				WineTanks[i] = append(WineTanks[i], tank)
			}
		}

		/*
			if tank.Blend != 0 {
				WineTanks[tank.Blend-1] = append(WineTanks[tank.Blend-1], tank)
			}
		*/
	}
	return WineTanks
}
