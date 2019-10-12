package main
import "fmt"

//map 的range是乱序的,slice是有序的
func main() {

	map2slice :=map[int32][]int64{
		1:[]int64{1,2,3,4,5,6,7},
		23:[]int64{1,2,10,4,5,6,7},
		2:[]int64{1,2,3,4,5,6,7},
		3:[]int64{1,2,3,4,5,6,7},
		5:[]int64{1,2,3,4,5,6,7},
	}

	for key, value := range map2slice {
		for _, val := range value {
			if val==10 {
				fmt.Println("-------------")
				break
			}
		}
		fmt.Println(key)
	}
}
