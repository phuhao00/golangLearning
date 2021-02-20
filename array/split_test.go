package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tmpStr := "#14'#点球"
	arr := strings.Split(tmpStr, "#")
	fmt.Println(len(arr), arr)
	fmt.Println(arr[1])
	fmt.Println(strings.Trim(arr[1], "'"))

}
