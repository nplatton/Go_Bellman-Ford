package main

import (
	"fmt"

	Fr "practice/fileReaders"
	Fw "practice/fileWriters"
)

func main() {
	// Return two array from file data
	srcNodes, dstNodes := Fr.Get_Test_Data("test.txt")
	fmt.Println(srcNodes, dstNodes)
	
	// Loop through the elements for the arrays
	for i:=0; i<len(srcNodes); i++ {
		srcNode, dstNode := srcNodes[i], dstNodes[i]
		// Perform Bellman-Ford on each pair of nodes in the arrays
		weight, path := BellmanFord(srcNode, dstNode)
		// Write data to a result file
		Fw.Write_Result_Data(srcNode, dstNode, weight, path)
	}

	// Read results file data and log it to the console
	results := Fr.Read_Results_Data("results.txt")
	for _, result := range results {
		fmt.Printf("The shortest path from %v to %v has length %v and is defined as follows:\n", result.Src, result.Dst, result.Weight)
		fmt.Println(result.Path)
		fmt.Println("\n")
	}
}
