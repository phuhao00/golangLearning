package main

import "fmt"

type iface interface {
	a ()int
	b ()
}

type im struct {
	real iface
}

func (*im)a()int  {
	return 1
}
func (*im)b()  {

}

func test(iface2 iface)  {

	fmt.Println(iface2.a())
}

func main() {
	gg:=&im{}
	test(gg)
}
