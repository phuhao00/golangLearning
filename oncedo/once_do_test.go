package oncedo

import (
	"fmt"
	"sync"
	"testing"
)

var once sync.Once

func TestOnceDo(t *testing.T) {
	for {
		OneDot()
	}
}

func OneDot() {

	once.Do(func() {
		fmt.Println("888")
	})
}
