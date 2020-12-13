package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSelect(t *testing.T)  {
	go Go1()
	time.Sleep(time.Second*10)
}

func Go1()  {

	select {

	}
	fmt.Println("sdfsdf")
}

func BB()  {

	var ss sync.Cond
	var sy sync.Pool
}

