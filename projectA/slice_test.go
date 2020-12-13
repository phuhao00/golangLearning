package main

import (
	"bytes"
	"fmt"
	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/bradleyjkemp/memviz"
	"testing"
)

type  hhh struct {

}

func TestSlice(t *testing.T)  {
	var ss1 *[]int
	s2:=[]int{1,2,4,5}
	ss1=&s2
	ss2 := (*ss1)[2:]
	ss3 := ss2[2:]
	b := &bytes.Buffer{}
	memviz.Map(b,ss3)
	fmt.Println(b.String())
	cupaloy.SnapshotT(t, b.Bytes())

}

