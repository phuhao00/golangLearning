package main

import "fmt"

type ErrCode int

//go:generate stringer -type ErrCode
const (
	ERR_CODE_OK             ErrCode = 0 // OK
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 无效参数
	ERR_CODE_TIMEOUT        ErrCode = 2 // 超时
)

func main() {
	fmt.Println(ERR_CODE_INVALID_PARAMS)
}
