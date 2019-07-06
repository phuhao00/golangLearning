package main

import (
	"../Message"
	"github.com/golang/protobuf/proto"
)

import "fmt"

var jj interface{}
func main()  {
	hh:=&pb.Jj{}
	if jj==nil {
		fmt.Println("hh")
	}
	if hh==nil {
		fmt.Println("[[[[[[[[[")
	}
	hh.Sd=make([]*pb.Pet,0)

	kkk(hh)
}
func kkk(mm interface{})  {
	lll:=mm.(*pb.Jj)
	proto.Marshal(lll)
}
