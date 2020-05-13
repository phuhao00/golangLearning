package main
/*
*/

import (
"fmt"
)

type People struct {
	Name string
	Age  int
	Sex  string
}

func (p *People) String() string {
	return fmt.Sprintf("people info: %+v", p)
}

func main() {
	p := &People{Name: "Lucy", Age: 18, Sex: "M"}
	fmt.Println(p)
}