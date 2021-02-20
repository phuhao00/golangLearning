package main

import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	mI := make(map[int64]string)
	if mI != nil {
		fmt.Println("not nil" + "")
	}

}
