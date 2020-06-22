package system

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestFor(t *testing.T)  {
	for ;true; {//等价while
		fmt.Println("dd")
	}
	for true {
		fmt.Println("b")
	}
}
//浮点数
//浮点数字面量被自动类型推断为float64类型
//计算机很难进行浮点数的精确表示和存储，因此两浮点数之间不应该使用==或！=进行比较操作，高精度科学计算应该使用math标准库
var cc = 10.0
func TestCType(t *testing.T)  {
	fmt.Println(reflect.TypeOf(cc))
}
//=== RUN   TestCType
//float64
//--- PASS: TestCType (0.00s)

//复数
var valueC complex64 =5.1 +5i
var vC =complex(3.3,5) //构造一个复数
var aaa= real(vC)
var bbb= imag(vC)

func TestComplex(t *testing.T)  {
	fmt.Println(valueC)
	fmt.Println(aaa)
	fmt.Println(bbb)
}
//=== RUN   TestComplex
//(5.1+5i)
//3.3
//5
//--- PASS: TestComplex (0.00s)
//PASS


//字符串是常量，可以访问，不可修改,
func TestStringUse(t *testing.T)  {

	var str="huhao"
	fmt.Println(str[1])
	fmt.Println(str[0:2])
	fmt.Println([]rune(str)) //ascii 码，占用四个字节，golang 默认使用utf-8编码，如果需要转，使用Unicode/utf-8标准包
	//str[1]='1'//error
}
//字符串转换为切片[]bytes要慎用，尤其是当数据量较大时(每转换一次都需要复制内容)
func TestStr2Bytes(t *testing.T)  {
	str:="huhs;asfaf"
	b:=[]byte(str)
	fmt.Println(b)
}
//go 不支持指针的运算符，那样会给垃圾回收的实现带来很多的不便，在c和c++里面指针的运算很容易出现问题，因此go直接在语言层面禁止指针运算
func TestPoint(t *testing.T)  {
	a:=1234
	p:=&a
	fmt.Println(p)
	//p++//invalid operation: p++ (non-numeric type *int)
}
//数组
func TestArray(t *testing.T)  {
	a:=[3]int64{4,56,7}
	ab:=[...]int64{1324324,244}//不指定长度
	abv:=[...]int64{1:1324324,244}//不指定长度,通过索引进行初始化
	fmt.Println(a)
	fmt.Println(ab)
	fmt.Println(abv)
}
//go内置的map不是并发安全的，并发安全的map可以使用标准包的sync中的map

//defer,主动调用os.Exit(int)推出进程时，defer将不再被执行(即使defer已经提前注册)
//defer 的好处是可以在一定成都上避免资源的泄露，特别是有很多return 语句，有多个需要关闭的场景中，很容易漏掉资源的关闭操作
func TestDefer(t *testing.T){
	defer func() {
		fmt.Println("os")
	}()
	fmt.Println("adf")
	os.Exit(9)
}
//API server listening at: 127.0.0.1:53118
//=== RUN   TestDefer
//adf
//
var val1 =0
func fa()  func(val int)int64{
	return func(val int) int64 {
		fmt.Println(val1,&val1)
		val1=val1+val
		return int64(val1)
	}
}
func TestInternal(t *testing.T)  {

	f1:=fa()
	f2:=fa()
	fmt.Println(&f1)
	fmt.Println(&f2)

	fmt.Println(f1(4))
	fmt.Println(f2(4))
	fmt.Println(f2(4))
	fmt.Println(f1(4))

}
//
func TestDeferPlus(t *testing.T){
	defer recover() //这个会捕获失败
	defer fmt.Println(recover())//捕获失败
	defer func() {
		func(){
			fmt.Println(";;;")
			recover()//无效
		}()
	}()
//如下场景会捕获成功

	defer func() {
		fmt.Println("success")
		recover()
	}()

	defer except()

}
func except()  {

	recover()
}
//
func TestDeferCatch(t *testing.T)  {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	//只有最后一次panic调用能够被捕获
	defer func() {
		panic("first defer panic !!")
	}()
	defer func() {
		panic("second defer panic !!")
	}()

	panic("main body panic !!")

}
//init 函数引发的panic 只能在init 函数中捕获，在main 函数中无法捕获，原因是init函数先于main 执行
//函数不能捕获新启动的goroutine抛出的panic
func TestCacthPlus(t *testing.T)  {
defer func() {
	if err:=recover();err!=nil {

		fmt.Println(err)
	}
	go panicTest()
}()

}
func panicTest()  {
	panic("goroutine")
}
//panic 使用场景
//1.无法正常执行下去(不能容错继续执行)
//2在调试时，panic快速退出，可以更快的定位错误
//为了保证程序的健壮性，，需要使用recover拦截运行时错误

//error通常作为最后一个返回值
//defer语句应该放到err判断的后面，不然可能产生panic

//panic 使用原则
//
//1.不能容错继续执行，主动panic
//2.虽然发生错误，允许继续，使用recover 捕获异常



//context
//核心功能是多个goroutine之间的退出机制，传递数据只是一个辅助共功能，应该慎用
//坏处
//1.传递的都是interface{}类型的值，编译器不能进行严格的类型校验
//2.从interface{}到具体类型需要使用类型断言和接口查询，有一定的运行时期的开销和性能损失
//3.值在传递的过程中有可能被后续的服务覆盖，且不易被发现
//4.传递信息不简明，较晦涩；不能通过代码或者文档一眼看到传递的是什么，不利于后续的维护
//应该传递
//1.日志信息
//2.调试信息
//3.不影响业务主逻辑的可选数据

