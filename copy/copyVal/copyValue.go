package main

import "fmt"

type Struct1 struct {
	A []int32
}

var AA map[int32]Struct1

func main()  {
	AA= map[int32]Struct1{
		1:Struct1{A:[]int32{1,2,3,4}},
	}
	var B=make([]int32,0)
	B=AA[1].A				//会有问题
	//B=append(B,AA[1].A...)	//这样子不会有问题
	fmt.Println("old:",B)
	AA[1].A[0]=4
	fmt.Println("new:",B)
}