





```go
package main
import "fmt"
type abc struct {
   name interface{}
}
func main()  {
var mm map[int]bool=map[int]bool{
   7:true,
}
   ABC:=&abc{
      "huhao",
   }
   gg(ABC)
   fmt.Println(mm[900])
}
func gg(i interface{})  {
ll:=i.(*abc).name.(string)
fmt.Println(ll)
}

---
huhao
false
```









```go
package main

import "fmt"

var ttt =map[int32]bool{	//不需要这样子 var ttt map[int32]bool=map[int32]bool
   8:true,
}
func main()  {
   fmt.Println(ttt[9])
}
```