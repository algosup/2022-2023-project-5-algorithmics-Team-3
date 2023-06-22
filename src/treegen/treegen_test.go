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
		// DISABLE

		// The initial tanks
		gotEmptyTanks, gotWineTanks, gotTanks, gotFormula :=
			[]tanks.Tank{
				{TankID: 3, Capacity: 100, Blend: []float64{0, 0}, Volume: 0},
				{TankID: 4, Capacity: 100, Blend: []float64{0, 0}, Volume: 0},
			},
			[][]tanks.Tank{
				{{TankID: 1, Capacity: 100, Blend: []float64{100, 0}, Volume: 100}},
				{{TankID: 2, Capacity: 100, Blend: []float64{0, 100}, Volume: 100}},
			},
			[]tanks.Tank{
				{TankID: 1, Capacity: 100, Blend: []float64{100, 0}, Volume: 100},
				{TankID: 2, Capacity: 100, Blend: []float64{0, 100}, Volume: 100},
				{TankID: 3, Capacity: 100, Blend: []float64{0, 0}, Volume: 0},
				{TankID: 4, Capacity: 100, Blend: []float64{0, 0}, Volume: 0},
			},
			[]float64{50.00, 50.00}

		gotTankFillingRatios := TankFillingRatio(gotTanks, gotFormula)

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
		got := Solve(gotEmptyTanks, gotWineTanks, gotFormula, gotTankFillingRatios)

		// Assert the results
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("\ngot %v\n, get %v", got, expected)
		}

	})
}

// :===== This test function is designed to test the precomputation of the TankFillingRatios =====:
func TestTankFillingRatios(t *testing.T) {

	// :===== This subtest is designed to try out mapping simple filling ratios =====:
	t.Run("Test simple filling mapping", func(t *testing.T) {

		// The initial tanks and formula
		gotTanks, gotFormula :=
			[]tanks.Tank{
				{TankID: 1, Capacity: 100, Blend: []float64{100, 0}, Volume: 0},
				{TankID: 2, Capacity: 100, Blend: []float64{0, 100}, Volume: 0},
				{TankID: 3, Capacity: 100, Blend: []float64{0, 0}, Volume: 0},
				{TankID: 4, Capacity: 100, Blend: []float64{0, 0}, Volume: 0},
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
