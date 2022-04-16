package main

import (
	"fmt"

	Fr "practice/fileReaders"
	Fw "practice/fileWriters"
)

func main() {
	// Return two array from file data
	srcCities, dstCities := Fr.Get_Test_Data("test.json")
	fmt.Println(srcCities, dstCities)
	
	// Loop through the elements for the arrays
	for i, _ := range srcCities {
		// Perform Bellman-Ford on each pair of cities in the arrays
		weight, path := BellmanFord(srcCities[i], dstCities[i])
		// Write data to a result file
		Fw.Write_Result_Data(weight, path)
	}
}
