package main

import "fmt"

type iface interface {
	a () int64
	b () bool
}

type cba struct {

	real iface

}

type abc struct {
	cba
}

func (*abc)a()int64  {

	return  9
}

func (*abc)b()bool  {

	return  false
}

func main() {
	abc_sub:=&abc{	}
	abc_sub.real=abc_sub
	tt:=abc_sub.real.(iface).a()//
	fmt.Println(tt)
}
