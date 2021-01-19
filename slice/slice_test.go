package slice

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	a := []int{1, 2}
	//b := append(a[:0:0], a...)
	b := append([]int(nil), a...)
	fmt.Println(b)
}
