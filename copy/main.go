package main

import "fmt"

type abd struct {
	gg1 int64
	gg2 int64
	gg3 int64
}

func main()  {
	ff:=abd{
		1,2,3,
	}
	ff1:=DeepCopy(ff)
	fmt.Println(ff1.(abd).gg1)
}
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}
		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}
		return newSlice
	}
	return value
}