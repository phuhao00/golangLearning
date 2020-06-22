package system

import (
	"fmt"
	"github.com/codegangsta/inject"
	"testing"
)

type S1 interface {

}

type S2 interface {

}
//
func Format8(name string ,company S1,level S2,age int)  {
	fmt.Println(name,company,level,age)
}

func TestInject(t *testing.T)  {
	//控制实例的创建
	inj:=inject.New()

	//实参注入
	inj.Map("tom")
	inj.MapTo("tencent",(*S1)(nil))
	inj.MapTo("t4",(*S2)(nil))
	inj.Map(23)

	//函数反转调用
	inj.Invoke(Format8)
}
