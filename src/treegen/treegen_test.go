package treegen

import (
	"blend/tanks"
	"reflect"
	"testing"
)

// :===== This test function is designed to test the Solving function =====:
func TestSolve(t *testing.T) {
	// :===== This subtest is designed to test a simple 3 tank solving =====:
	t.Run("Simple Solve", func(t *testing.T) {
		// The initial tanks
		gotEmptyTanks, gotWineTanks, gotFormula :=
			[]tanks.Tank{
				{TankID: 3, Capacity: 100, WineNumber: 0},
				{TankID: 4, Capacity: 100, WineNumber: 0},
			},
			[][]tanks.Tank{
				{{TankID: 1, Capacity: 100, WineNumber: 1}},
				{{TankID: 2, Capacity: 100, WineNumber: 2}},
			},
			[]float64{50.00, 50.00}

		// The expected list of steps after processing the previous input
		expected :=
			[]Step{
				{
					Substeps: []Substep{
						{SourceID: 1, DestinationID: 3, Volume: 50.00},
						{SourceID: 2, DestinationID: 3, Volume: 50.00},
						{SourceID: 1, DestinationID: 4, Volume: 50.00},
						{SourceID: 2, DestinationID: 4, Volume: 50.00},
					},
				},
			}

		// Calling the function
		got := Solve(gotEmptyTanks, gotWineTanks, gotFormula, nil)

		// Assert the results
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}

	})
}

// :===== This test function is designed to test the orecomputation of the TankFillingRatios =====:
func TestTankFillingRatios(t *testing.T) {
	// :===== This subtest is designed to try out mapping simple filling ratios =====:
	t.Run("Test simple filling mapping", func(t *testing.T) {
		// The initial tanks and formula
		gotTanks, gotFormula :=
			[]tanks.Tank{
				{TankID: 1, Capacity: 100, WineNumber: 1},
				{TankID: 2, Capacity: 100, WineNumber: 2},
				{TankID: 3, Capacity: 100, WineNumber: 0},
				{TankID: 4, Capacity: 100, WineNumber: 0},
			}, []float64{
				50.00,
				50.00,
			}

		// Calling the function
		got := TankFillingRatio(gotTanks, gotFormula)

		// Assert the results
		expected := []float64{50.0, 50.0}
		if !reflect.DeepEqual(got[100.0], expected) {
			t.Errorf("got %v, expected %v", got[100.0], expected)
		}

	})
}
