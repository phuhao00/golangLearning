package main

import "fmt"

//匿名结构体
func main()  {

	map2Anonymous_struct := map[int64]struct { int}{}

	map2Anonymous_struct[1] = struct {int}{1}
	map2Anonymous_struct[2] = struct {int}{2}
	map2Anonymous_struct[3] = struct {int}{3}
	map2Anonymous_struct[4] = struct {int}{4}

	for key,v := range map2Anonymous_struct {
		fmt.Println(key,v)
	}
}
