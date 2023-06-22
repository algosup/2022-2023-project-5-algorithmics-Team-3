package sort

import (
	"blend/tanks"
	"testing"
)

// :===== Enable/Disable specific subtest suites =====:
var simpleTests bool = true
var alreadySortedTests bool = true
var nullTanksTests bool = true
var testFloating bool = false // disabled because floating point tank capacities have not been implemented yet

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ Sorting Test Suite ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

// :===== This test function is designed to test the sorting capabilities of SortTanks (decreasing order) =====:
func TestSortTanks(t *testing.T) {
	// :===== Simple 1D sort subtest =====:
	t.Run("Simple 1D sort subtest", func(t *testing.T) {
		if !simpleTests {
			t.SkipNow()
		}
		tanks, expected :=
			// The starting case
			[]tanks.Tank{
				{TankID: 5, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 6, Capacity: 300, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 4, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 2, Capacity: 1750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 1, Capacity: 1898, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 3, Capacity: 1500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
			},
			// The expected result
			[]tanks.Tank{
				{TankID: 6, Capacity: 300, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 5, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 4, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 3, Capacity: 1500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 2, Capacity: 1750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 1, Capacity: 1898, BlendNewField: []float64{0, 0, 0}, Volume: 0},
			}

		// Call the function to be tested
		sortedTanks := SortTanks(tanks)

		// The ID of the tanks is used to check if they have effectively sorted
		if sortedTanks[0].TankID != expected[0].TankID || sortedTanks[1].TankID != expected[1].TankID || sortedTanks[2].TankID != expected[2].TankID || sortedTanks[3].TankID != expected[3].TankID || sortedTanks[4].TankID != expected[4].TankID || sortedTanks[5].TankID != expected[5].TankID {
			t.Errorf("Tanks were not sorted correctly:\n Want: %+v \nGot: %+v", sortedTanks, expected)
		}
	})

	// :===== 1D sort subtest when some tanks are empty =====:
	t.Run("1D sort with null tanks mixed in subtest", func(t *testing.T) {
		if !simpleTests {
			t.SkipNow()
		}
		tanks, expected :=
			// The starting case
			[]tanks.Tank{
				{},
				{TankID: 6, Capacity: 300, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{},
				{TankID: 2, Capacity: 1750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 1, Capacity: 1898, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 3, Capacity: 1500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
			},
			// The expected result
			[]tanks.Tank{
				{},
				{},
				{TankID: 6, Capacity: 300, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 3, Capacity: 1500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 2, Capacity: 1750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 1, Capacity: 1898, BlendNewField: []float64{0, 0, 0}, Volume: 0},
			}

		// Call the function to be tested
		sortedTanks := SortTanks(tanks)

		// The ID of the tanks is used to check if they have effectively sorted
		if sortedTanks[0].TankID != expected[0].TankID || sortedTanks[1].TankID != expected[1].TankID || sortedTanks[2].TankID != expected[2].TankID || sortedTanks[3].TankID != expected[3].TankID || sortedTanks[4].TankID != expected[4].TankID || sortedTanks[5].TankID != expected[5].TankID {
			t.Errorf("Tanks were not sorted correctly:\n Want: %+v \nGot: %+v", sortedTanks, expected)
		}
	})

	// :===== 1D sort subtest when tanks are already in order =====:
	t.Run("Simple 1D sort subtest", func(t *testing.T) {
		if !simpleTests {
			t.SkipNow()
		}
		tanks, expected :=
			// The starting case
			[]tanks.Tank{
				{TankID: 6, Capacity: 300, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 5, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 4, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 3, Capacity: 1500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 2, Capacity: 1750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 1, Capacity: 1898, BlendNewField: []float64{0, 0, 0}, Volume: 0},
			},
			// The expected result
			[]tanks.Tank{
				{TankID: 6, Capacity: 300, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 5, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 4, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 3, Capacity: 1500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 2, Capacity: 1750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 1, Capacity: 1898, BlendNewField: []float64{0, 0, 0}, Volume: 0},
			}

		// Call the function to be tested
		sortedTanks := SortTanks(tanks)

		// The ID of the tanks is used to check if they have effectively sorted
		if sortedTanks[0].TankID != expected[0].TankID || sortedTanks[1].TankID != expected[1].TankID || sortedTanks[2].TankID != expected[2].TankID || sortedTanks[3].TankID != expected[3].TankID || sortedTanks[4].TankID != expected[4].TankID || sortedTanks[5].TankID != expected[5].TankID {
			t.Errorf("Tanks were not sorted correctly:\n Want: %+v \nGot: %+v", sortedTanks, expected)
		}
	})

	// :===== DISABLED 1D sort subtest with floating point Capacities =====:
	t.Run("1D Sort for floating point tank capacities subtest", func(t *testing.T) {
		// Disable test if set up to do so
		if !testFloating {
			t.SkipNow()
		}
		tanks, expected :=
			// The starting case
			[]tanks.Tank{
				{TankID: 2, Capacity: 100, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 1, Capacity: 300, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 3, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 5, Capacity: 1750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 6, Capacity: 1898, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 4, Capacity: 1500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
			},
			// The expected result
			[]tanks.Tank{
				{TankID: 1, Capacity: 300, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 2, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 3, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 4, Capacity: 1500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 5, Capacity: 1750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				{TankID: 6, Capacity: 1898, BlendNewField: []float64{0, 0, 0}, Volume: 0},
			}

		// Call the function to be tested
		sortedTanks := SortTanks(tanks)

		// The ID of the tanks is used to check if they have effectively sorted
		if sortedTanks[0].TankID != expected[0].TankID ||
			sortedTanks[1].TankID != expected[1].TankID ||
			sortedTanks[2].TankID != expected[2].TankID ||
			sortedTanks[3].TankID != expected[3].TankID ||
			sortedTanks[4].TankID != expected[4].TankID ||
			sortedTanks[5].TankID != expected[5].TankID {
			t.Errorf("Tanks were not sorted correctly: %+v", sortedTanks)
		}
	})
}

// :===== This test function is desifgned to test the sorting capabilities of SortTanks 2D (decreasing order) =====:
func TestSortTanks2D(t *testing.T) {
	// :===== Simple 2D sort subtest =====:
	t.Run("Simple 2D sort subtest", func(t *testing.T) {
		if !simpleTests {
			t.SkipNow()
		}
		tankGroups, expected :=
			// The starting case
			[][]tanks.Tank{
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
			},
			// The expected result
			[][]tanks.Tank{
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
			}

		// Call the function to be tested
		sortedTankGroups := SortTanks2D(tankGroups)

		// The ID of the tanks is used to check if the tanks have effectively been sorted
		if sortedTankGroups[0][0].TankID != expected[0][0].TankID ||
			sortedTankGroups[0][1].TankID != expected[0][1].TankID ||
			sortedTankGroups[0][2].TankID != expected[0][2].TankID ||
			sortedTankGroups[1][0].TankID != expected[1][0].TankID ||
			sortedTankGroups[1][1].TankID != expected[1][1].TankID ||
			sortedTankGroups[1][2].TankID != expected[1][2].TankID ||
			sortedTankGroups[2][0].TankID != expected[2][0].TankID ||
			sortedTankGroups[2][1].TankID != expected[2][1].TankID ||
			sortedTankGroups[2][2].TankID != expected[2][2].TankID {
			t.Errorf("Tanks were not sorted correctly: %+v", sortedTankGroups)
		}
	})

	// :===== 2D sort subtest for when tanks are already sorted =====:
	t.Run("Already sorted 2D sort subtest", func(t *testing.T) {
		// Disable test if set up to do so
		if !alreadySortedTests {
			t.SkipNow()
		}
		tankGroups, expected :=
			// The starting case
			[][]tanks.Tank{
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
			},
			// The expected result
			[][]tanks.Tank{
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
			}

		// Call the function to be tested
		sortedTankGroups := SortTanks2D(tankGroups)

		// The ID of the tanks is used to check if the tanks have effectively been sorted
		if sortedTankGroups[0][0].TankID != expected[0][0].TankID ||
			sortedTankGroups[0][1].TankID != expected[0][1].TankID ||
			sortedTankGroups[0][2].TankID != expected[0][2].TankID ||
			sortedTankGroups[1][0].TankID != expected[1][0].TankID ||
			sortedTankGroups[1][1].TankID != expected[1][1].TankID ||
			sortedTankGroups[1][2].TankID != expected[1][2].TankID ||
			sortedTankGroups[2][0].TankID != expected[2][0].TankID ||
			sortedTankGroups[2][1].TankID != expected[2][1].TankID ||
			sortedTankGroups[2][2].TankID != expected[2][2].TankID {
			t.Errorf("Tanks were not sorted correctly: %+v", sortedTankGroups)
		}
	})

	// :===== 2D sort subtest for when some tanks are null =====:
	t.Run("2D sort with null tanks mixed in subtest", func(t *testing.T) {
		// Disable test if set up to do so
		if !nullTanksTests {
			t.SkipNow()
		}
		tankGroups, expected :=
			// The starting case
			[][]tanks.Tank{
				{
					{},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
			},
			// The expected result
			[][]tanks.Tank{
				{
					{},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
			}

		// Call the function to be tested
		sortedTankGroups := SortTanks2D(tankGroups)

		// The ID of the tanks is used to check if the tanks have effectively been sorted
		if sortedTankGroups[0][0].TankID != expected[0][0].TankID ||
			sortedTankGroups[0][1].TankID != expected[0][1].TankID ||
			sortedTankGroups[0][2].TankID != expected[0][2].TankID ||
			sortedTankGroups[1][0].TankID != expected[1][0].TankID ||
			sortedTankGroups[1][1].TankID != expected[1][1].TankID ||
			sortedTankGroups[1][2].TankID != expected[1][2].TankID ||
			sortedTankGroups[2][0].TankID != expected[2][0].TankID ||
			sortedTankGroups[2][1].TankID != expected[2][1].TankID ||
			sortedTankGroups[2][2].TankID != expected[2][2].TankID {
			t.Errorf("Tanks were not sorted correctly: %+v", sortedTankGroups)
		}
	})

	// :===== DISABLED 2D sort subtest for floating point capacities =====:
	t.Run("2D sort for floating point tank capacities subtest", func(t *testing.T) {
		if !testFloating {
			t.SkipNow()
		}
		tankGroups, expected :=
			// The starting case
			[][]tanks.Tank{
				{
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
			},
			// The expected result
			[][]tanks.Tank{
				{
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
				{
					{TankID: 1, Capacity: 1000, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 2, Capacity: 750, BlendNewField: []float64{0, 0, 0}, Volume: 0},
					{TankID: 3, Capacity: 500, BlendNewField: []float64{0, 0, 0}, Volume: 0},
				},
			}

		// Call the function to be tested
		sortedTankGroups := SortTanks2D(tankGroups)

		// The ID of the tanks is used to check if the tanks have effectively been sorted
		if sortedTankGroups[0][0].TankID != expected[0][0].TankID ||
			sortedTankGroups[0][1].TankID != expected[0][1].TankID ||
			sortedTankGroups[0][2].TankID != expected[0][2].TankID ||
			sortedTankGroups[1][0].TankID != expected[1][0].TankID ||
			sortedTankGroups[1][1].TankID != expected[1][1].TankID ||
			sortedTankGroups[1][2].TankID != expected[1][2].TankID ||
			sortedTankGroups[2][0].TankID != expected[2][0].TankID ||
			sortedTankGroups[2][1].TankID != expected[2][1].TankID ||
			sortedTankGroups[2][2].TankID != expected[2][2].TankID {
			t.Errorf("Tanks were not sorted correctly: %+v", sortedTankGroups)
		}
	})
}
