





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