package main

import "fmt"

type iface_eg interface {
	a ()int
	b ()
}

type im struct {
	real iface_eg
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
