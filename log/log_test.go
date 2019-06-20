package log

import "testing"

func TestLog(t *testing.T) {
	Info.Println("888")
	Warning.Printf("999")
	Error.Println("000")
}
