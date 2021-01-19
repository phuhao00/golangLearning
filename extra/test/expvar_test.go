package test

import (
	"expvar"
	"fmt"
	"net/http"
	"runtime"
	"testing"
	"time"
)

var visits = expvar.NewInt("visits")

func handler(w http.ResponseWriter, r *http.Request) {
	visits.Add(1)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func TestExpvar(t *testing.T) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1818", nil)
}

func TestU(t *testing.T) {
	runtime.GOMAXPROCS(1)
	var c = make(chan int)
	time.Sleep(time.Second)
	fu := func() {
		c <- 1
	}
	go fu()

	go func() {
		<-c
	}()

}
