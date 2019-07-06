package main

import (
	"../Message"
	"github.com/golang/protobuf/proto"
)

import "fmt"

type jkl struct {
	s interface{}
}
var jj interface{}
func main()  {
	hh:=&jkl{}
	if jj==nil {
		fmt.Println("hh")
	}
	if hh==nil {
		fmt.Println("[[[[[[[[[")
	}
	kkk(&pb.Pet{})
}
func kkk(mm interface{})  {
	lll:=mm.(*pb.Pet)
	proto.Marshal(lll)
}
