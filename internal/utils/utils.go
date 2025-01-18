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

func UniqueElementsByPosition(arrays [][]string, posInArray int) []string {
	uniqueMap := make(map[string]bool)
	var uniqueArray []string

	for _, array := range arrays {
		var element = array[posInArray]

		if !uniqueMap[element] {
			uniqueMap[element] = true
			uniqueArray = append(uniqueArray, element)
		}
	}

	return uniqueArray
}
