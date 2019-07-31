package main

import (
	"fmt"
	"sync"
)

type jjj struct {
	 jh	int64
}
type hh struct {
	clientWg sync.WaitGroup
	bbb func(int642 int64) jjj
}
func main()  {
	 //clientWg :=&hh{
	 //}
	 ff:=map[int64]struct {}{}
	 ff[1]=struct {}{}
	 ff[2]=struct {}{}
	 ff[3]=struct {}{}
	 ff[4]=struct {}{}
	for key := range ff {
		fmt.Println(key)
	}
	//
	// j_:=clientWg.bbb(5)
	// fmt.Println(j_)
	// clientWg.clientWg=sync.WaitGroup{}
	// //clientWg.clientWg.Add(2)
	// clientWg.aaa()
	//// clientWg.clientWg.Wait()
	////defer clientWg.Done()
}


func(self *hh) aaa()  {
	writeChan:= make(chan []byte, 9)
	go func() {
		fmt.Println("lllll")
		for  b := range writeChan {
			if  b == nil {
				fmt.Println("nil")
			}

		}
	}()

}