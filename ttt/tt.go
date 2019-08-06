package main

import "fmt"

func main()  {
	//attach:=[]int64{11,22,33,44,55}
	//gg:=[]int64{1,2,3,4,5,6}
	//for _,val:=range gg{
	//	petattach:=val
	//	tm:=attach
	//	fmt.Printf("attach:%p ;\n",attach)
	//	fmt.Printf("tm:%p ;\n",tm)
	//	tm=append(tm,petattach)
	//	//fmt.Println(tm)
	//	fmt.Printf("tm:--%p ;\n",tm)
	//}
	test()
}

func  test()  {

		arr := []int32{1,2,4,5,7,8}

		arr1 := arr
		fmt.Printf("arr1:%p ;\n",arr1)
		arr1 = append(arr1, 3)
		fmt.Printf("arr:%p ;\n",arr)
		arr2 := arr
		fmt.Printf("arr2:%p ;\n",arr2)
		arr2 = append(arr2, 4)
		fmt.Printf("arr:%p ;\n",arr)
		fmt.Println(arr1)
		fmt.Println(arr2)

}