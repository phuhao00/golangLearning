package main

import (
	"encoding/json"
	"fmt"
)

type Pill int
//go:generate jsonenums -type=Pill -suffix=huhao  -prefix=abc
const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)


func main()  {
			//re:= &ret{
			//	Msg :"参数校验不通过",
			//	Code: 1,
			//	Data: "",}

			ret_tt:=&ret{
				"参数校验不通过",
				 1,
				 "",}
			data,err:=json.Marshal(&ret_tt)
			fmt.Println(err)
			fmt.Println(data)

			tr:=&ret{}
			json.Unmarshal(data,tr)
			fmt.Println(tr)
}
//
type ret struct {
	Msg string `json:"Msg"`
	Code int `json:"Code"`
	Data interface{} `json:",Msg"`
}
type ret_t struct {
	msg string
	code int
	data interface{}
}
