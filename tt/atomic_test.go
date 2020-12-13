package main

import (
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T)  {

	atomic.CompareAndSwapInt64()
}
