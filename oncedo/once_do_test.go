package oncedo

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnceDo(t *testing.T)  {
	for {
		OneDot()
	}
}

func OneDot()  {
	var once sync.Once
	once.Do(func() {
		fmt.Println("888")
	})
}