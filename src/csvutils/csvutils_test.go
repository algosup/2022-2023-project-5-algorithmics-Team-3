package csvutils

import (
	"blend/tanks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// :===== Enable/Disable specific subtest suites =====:
var testFloating bool = false

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ CSV Utilities Test Suite ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

// :===== This function is designed to test the opening of the input CSV file =====:
func TestOpenCSV(t *testing.T) {
	// :===== File not found subtest =====:
	t.Run("File opens and reads correctly", func(t *testing.T) {
		want := [][]string{
			{"7967", "5255850"},
			{"762289", "529347"},
			{"692095508", "8413"},
			{"2646949789", "3297"},
			{"9349248858", "004"},
			{"233", "9323"},
			{"271437796", "8116"},
			{"320", "380642"},
		}

		got := OpenCSV("csvTests/openTest.csv")

		assert.Equal(t, got, want, "File not found/open...")
	})
}

// :===== This function is designed to test the parsing of the input CSV into readable data =====:
func TestParseCSV(t *testing.T) {
	// :===== Few tanks & small formula subtest=====:
	t.Run("Few tanks & small formula", func(t *testing.T) {
		// The expected outcome
		wantedTanks, wantedFormula :=
			// The expected tanks
			[]tanks.Tank{
				{TankID: 1, Capacity: 100, WineNumber: 1},
				{TankID: 2, Capacity: 50, WineNumber: 1},
				{TankID: 3, Capacity: 25, WineNumber: 2},
				{TankID: 4, Capacity: 100, WineNumber: 2},
				{TankID: 5, Capacity: 200, WineNumber: 3},
				{TankID: 6, Capacity: 100, WineNumber: 0},
				{TankID: 7, Capacity: 200, WineNumber: 0},
				{TankID: 8, Capacity: 50, WineNumber: 0},
				{TankID: 9, Capacity: 25, WineNumber: 0},
				{TankID: 10, Capacity: 50, WineNumber: 0},
			},
			// The expected formula
			[]float32{
				40.000,
				45.000,
				15.000,
			}

		//
		gotOpenedCSV := OpenCSV("csvTests/ftsf.csv")

		gotTanks, gotFormula := ParseCSV(gotOpenedCSV)

		assert.Equal(t, gotTanks, wantedTanks, "The tanks do not match.")
		assert.Equal(t, gotFormula, wantedFormula, "The formulas do not match.")
	})

	// :===== Many tanks & big formula subtest =====:
	t.Run("Many tanks & big formula", func(t *testing.T) {
		// As the csv is to big we will cherry pick specific values to compare and assume the rest is OK
		// The expected formula

		wantedTanks, wantedFormula :=
			[]tanks.Tank{
				{TankID: 1, Capacity: 73453868, WineNumber: 699860},
				{TankID: 501, Capacity: 234, WineNumber: 81726},
				{TankID: 1000, Capacity: 1395711, WineNumber: 28464},
			},
			[]float32{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}

		gotOpenedCSV := OpenCSV("csvTests/mtbf.csv")

		gotTanks, gotFormula := ParseCSV(gotOpenedCSV)

		fmt.Printf("gotTanks: %+v", gotTanks)
		fmt.Printf("wantedTanks: %+v", wantedTanks)

		assert.Equal(t, gotTanks[0], wantedTanks[0], "The tanks n°1 do not match.")
		assert.Equal(t, gotTanks[500], wantedTanks[1], "The tanks n°501 do not match.")
		assert.Equal(t, gotTanks[999], wantedTanks[2], "The tanks n°1000 doe not match.")
		assert.Equal(t, gotFormula, wantedFormula, "The formulas do not match.")
	})

	// :===== DISABLED Few tanks & small formula with floating points subtest =====:
	t.Run("Few tanks & small formula with floating points", func(t *testing.T) {
		// Enable/Disable test
		if !testFloating {
			t.Skip()
		}

		// The expected outcome
		wantedTanks, wantedFormula :=
			// The expected tanks
			[]tanks.Tank{
				{TankID: 1, Capacity: 100, WineNumber: 1},
				{TankID: 2, Capacity: 50, WineNumber: 1},
				{TankID: 3, Capacity: 25, WineNumber: 2},
				{TankID: 4, Capacity: 100, WineNumber: 2},
				{TankID: 5, Capacity: 200, WineNumber: 3},
				{TankID: 6, Capacity: 100, WineNumber: 0},
				{TankID: 7, Capacity: 200, WineNumber: 0},
				{TankID: 8, Capacity: 50, WineNumber: 0},
				{TankID: 9, Capacity: 25, WineNumber: 0},
				{TankID: 10, Capacity: 50, WineNumber: 0},
			},
			// The expected formula
			[]float32{
				40.000,
				45.000,
				15.000,
			}

		//
		gotOpenedCSV := OpenCSV("csvTests/ftsf.csv")

		gotTanks, gotFormula := ParseCSV(gotOpenedCSV)

		assert.Equal(t, gotTanks, wantedTanks, "The tanks do not match.")
		assert.Equal(t, gotFormula, wantedFormula, "The formulas do not match.")

	})

	// :===== DISABLED Many tanks &  big formula with floating points subtest =====:
	t.Run("Many tanks & big formula", func(t *testing.T) {
		// As the csv is to big we will cherry pick specific values to compare and assume the rest is OK
		// The expected formula

		// Disable/enable test
		if !testFloating {
			t.Skip()
		}

		wantedTanks, wantedFormula :=
			[]tanks.Tank{
				{TankID: 1, Capacity: 73453868, WineNumber: 699860},
				{TankID: 501, Capacity: 234, WineNumber: 81726},
				{TankID: 1000, Capacity: 1395711, WineNumber: 28464},
			},
			[]float32{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}

		gotOpenedCSV := OpenCSV("csvTests/mtbf.csv")

		gotTanks, gotFormula := ParseCSV(gotOpenedCSV)

		fmt.Printf("gotTanks: %+v", gotTanks)
		fmt.Printf("wantedTanks: %+v", wantedTanks)

		assert.Equal(t, gotTanks[0], wantedTanks[0], "The tanks n°1 do not match.")
		assert.Equal(t, gotTanks[500], wantedTanks[1], "The tanks n°501 do not match.")
		assert.Equal(t, gotTanks[999], wantedTanks[2], "The tanks n°1000 doe not match.")
		assert.Equal(t, gotFormula, wantedFormula, "The formulas do not match.")
	})

	// :===== Return error if there is a problem with the formula =====:
	t.Run("Formula incorrect (doesn't add up to 100%)", func(t *testing.T) {
		gotOpenedCSV := OpenCSV("csvTests/formula100.csv")

		gotTanks, gotFormula := ParseCSV(gotOpenedCSV)

		// Check if gotTanks or gotFormula are not nil, and report a failure if so
		if gotTanks != nil || gotFormula != nil {
			t.Errorf("Unexpected non-nil values. gotTanks: %v, gotFormula: %v", gotTanks, gotFormula)
		}

	})
}
