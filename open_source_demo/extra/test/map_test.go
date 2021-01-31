package test

import (
	"fmt"
	"testing"
)

type Gname struct {
	GG int
}

type Gname2 struct {
	HH string
}

func (g *Gname2) UU() {
	g.HH = "yy"
}

func TestHH(t *testing.T) {
	var G = Gname2{
		"iii",
	}
	G.UU()
	fmt.Println(G.HH)
}

var (
	M1 = make(map[string]Gname)
	M2 = make(map[string]Gname2)
)

//
//func TestMap(t *testing.T) {
//	NewCorner := make(map[string]Gname2, 1)
//	corMap := make(map[string]int, 0)
//
//	for k, _ := range NewCorner {
//		if _, ok := corMap["v.HomeId"]; ok {
//			NewCorner[k].HH = corMap["v.HomeId"]
//		}
//		if _, ok := corMap["v.HomeId"]; ok {
//			NewCorner[k].HH = corMap["v.AwayId"]
//		}
//	}
//}

//type AA struct {
//	Gi int
//	Gs string
//}
//func TestHH(t *testing.T) {
//	aa := make(map[string]AA)
//	aa["1"] = AA{1, "1"}
//	if v, ok := aa["1"]; ok {
//		v.Gi = 2
//		aa["1"] = v
//	}
//
//	fmt.Printf("%+v\n", aa)
//	bb := make(map[string]*AA)
//	bb["2"] = &AA{2, "2"}
//	bb["2"].Gi = 3
//	for k, v := range bb {
//		fmt.Printf("%v:%+v\n", k, v)
//	}
//}
