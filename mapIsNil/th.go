package main

import "fmt"
type gg struct {
	g int64
	jj map[int64]int
}

func main()  {
	var gg=make(map[int64]gg)
	if gg==nil {
		fmt.Println("hhfhfh")
	}
}
