package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"unsafe"
)
var json = jsoniter.ConfigCompatibleWithStandardLibrary
func main()  {
	gg:=make(map[int64]bool)
	gg[9]=true
	testUnmarsalMap(gg)
}

func testUnmarsalMap(gg map[int64]bool)  {
	bytes,_ := json.Marshal(gg)
	result:= Bytes2str(bytes)
	hh:=make(map[int64]bool)
	err:= json.Unmarshal(Str2bytes(result), &hh)
	if err==nil {
		fmt.Println(hh)
	}

}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
