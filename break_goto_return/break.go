package break_goto_return

import "fmt"
//在没有使用loop标签的时候break只是跳出了第一层for循环
//使用标签后跳出到指定的标签,break只能跳出到之前，如果将Loop标签放在后边则会报错
//break标签只能用于for循环，跳出后不再执行标签对应的for循环

func TestBreak() {
Loop:
	for j:=0;j<3;j++{
		fmt.Println(j)
		for a:=0;a<5;a++{
			fmt.Println(a)
			if a>3{
				break Loop
			}
		}
	}
}

