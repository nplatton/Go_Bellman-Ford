package main

import (
	"fmt"

	Fr "practice/fileReaders"
	Fw "practice/fileWriters"
)

func main() {
	// Return two array from file data
	srcCities, dstCities := Fr.Get_Test_Data("test.txt")
	fmt.Println(srcCities, dstCities)
	
	// Loop through the elements for the arrays
	for i:=0; i<len(srcCities); i++ {
		srcCity, dstCity := srcCities[i], dstCities[i]
		// Perform Bellman-Ford on each pair of cities in the arrays
		weight, path := BellmanFord(srcCity, dstCity)
		// Write data to a result file
		Fw.Write_Result_Data(srcCity, dstCity, weight, path)
	}

	// Read results file data and log it to the console
	results := Fr.Read_Results_Data("results.txt")
	for _, result := range results {
		fmt.Printf("The shortest path from %v to %v has length %v and is defined as follows:\n", result.src, result.dst, result.weight)
		fmt.Println(result.path)
		fmt.Println("\n")

	}
}
