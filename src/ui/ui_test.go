package ui

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ UI Testing Suite ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

/*
// :===== This test function is designed to test the PrintInstructions function =====:
func TestPrintInstructions(t *testing.T) {
	// :===== This subtest is designed to Print a simple instructions set =====:
	t.Run("Simple Print", func(t *testing.T) {
		expected := "Step 1:\n=============\nTank 1 ==(50.00hL)==> Tank 3\nTank 2 ==(50.00hL)==> Tank 3\nTank 1 ==(50.00hL)==> Tank 4\nTank 2 ==(50.00hL)==> Tank 4\n"

		steps := []treegen.Step{
			{
				Substeps: []treegen.Substep{
					{SourceID: 1, DestinationID: 3, Volume: 50.00},
					{SourceID: 2, DestinationID: 3, Volume: 50.00},
					{SourceID: 1, DestinationID: 4, Volume: 50.00},
					{SourceID: 2, DestinationID: 4, Volume: 50.00},
				},
			},
		}

		got := PrintInstructions(steps)

		// Flatten the got output for proper comparison
		gotFlat := strings.Join(got[0], "")

		if gotFlat != expected {
			t.Errorf("Expected %v, but got %v", expected, gotFlat)
		}
	})

}
*/
