package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)
func main()  {
	TestSliceSort()

}
//测试随机
func test(old int) int {
	randSeed:=rand.New(rand.NewSource(time.Now().UnixNano()))
	poolIndex:= randSeed.Intn(10)
	if poolIndex==old {
		return test(old)
	}
	return poolIndex
}
//测试排序
func TestSliceSort()  {
	var arr =[]int32{
		5,
		6,
		9,
	}
	sort.Slice(arr, func(i, j int) bool {
		return i>j
	})
	fmt.Println(arr)

}
