package main

import (
	"fmt"

	F "./fileReaders"
)

func main() {
	srcCities, dstCities := F.Get_Test_Data("test.json")
	fmt.Println(srcCities, dstCities)
	
}
