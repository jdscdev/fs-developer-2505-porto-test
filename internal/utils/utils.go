package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ReadCSVFile(csvFile string) ([][]string, error) {
	// open CSV file
	fd, error := os.Open(csvFile)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Successfully opened the CSV file")
	defer fd.Close()

	// read CSV file
	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()
	if error != nil {
		fmt.Println(error)
	}
	return records, error
}

func ConvertToInt(stringValue string) int {
	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		fmt.Println("ConvertToInf Error:", err, " String:", stringValue)
		return 0
	}
	return intValue
}
