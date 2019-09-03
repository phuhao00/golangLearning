package main

import (
	"fmt"
)

func main() {
	//var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	////for key, value := range balance {
	////	fmt.Println(key,value)
	////
	////}
	//
	//fmt.Println(reflect.TypeOf(balance))

	//b()

	fmt.Println(c())
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}