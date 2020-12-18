package main

import (
	"fmt"
	"reflect"
	"testing"
)

func getinterfaceName() (arr map[int]*name) {
	//arr = make([]*name, 0)
	return arr
}

type name struct {
	int
}

func returnInterName(interf interface{}) map[string]interface{} {
	if interf == nil {
		fmt.Println("im real nil")
	}
	switch reflect.TypeOf(interf).Kind() {
	case reflect.Slice, reflect.Map:
		l := reflect.ValueOf(interf).Len()
		fmt.Println(l)
	case reflect.Uintptr:
		if reflect.ValueOf(interf).IsNil() {
			fmt.Println("fl")
			//interf = make([]int, 1)
		}
	case reflect.Array:

	}
	reflect.TypeOf(interf).Kind()

	return map[string]interface{}{
		"demo":      1,
		"interface": interf,
	}
}

func TestNil(t *testing.T) {
	var arr [1]int
	//val := getinterfaceName()
	returnInterName(arr)
}
