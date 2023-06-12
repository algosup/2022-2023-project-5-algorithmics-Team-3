package csvutils

import (
	"encoding/csv"
	"fmt"
	"os"
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
	records, error := fileReader.ReadAll()

	// Check if error while reading
	if error != nil {
		fmt.Println(error)
	}

	return records

}
