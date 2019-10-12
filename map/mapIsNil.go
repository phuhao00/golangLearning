package main

import "fmt"

type testMap struct {
	number     int64
	dictionary map[int64]int
}
//
func main()  {

	var testEg =make(map[int64]testMap)

	//make后的 map 为nil
	if testEg ==nil {
		fmt.Println("hhfhfh")
	}
}
