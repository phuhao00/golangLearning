package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main()  {
	old:=4
	new:=test(old)
	fmt.Println(new)

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