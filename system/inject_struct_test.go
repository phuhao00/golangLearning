package system

import (
	"fmt"
	"github.com/codegangsta/inject"
	"testing"
)


type H1 struct {}
type H2 struct {}


//
type Staff struct {

	Name string 	`inject`
	Company H1		`inject`
	Level H2 		`inject`
	Age int 		`inject`

}

func TestInjectStruct(t *testing.T)  {

	//创建被注入实例
	s:=Staff{}
	inj:=inject.New()

	//初始化注入值
	inj.Map("tom")
	inj.MapTo("tencent",(*H1)(nil))
	inj.MapTo("t4",(*H2)(nil))
	inj.Map(23)

	//实现struct 注入
	inj.Apply(&s)

	fmt.Println(s)

}
