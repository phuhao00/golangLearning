package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr1(t *testing.T) {
	еrr := errors.New("foo")
	var err error
	if еrr != nil {
		fmt.Printf("%T %v", err, err)
	}
}
