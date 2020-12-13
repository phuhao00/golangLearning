package gin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"testing"
)

//ret 0 正常  msg   data
type Response struct {
	Ret  int64       `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}


func TestResponse(t *testing.T)  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//c.JSON(200, gin.H{a
		//	"message": "pong",
		//})
		resp := &Response{Ret: 0, Msg: "成功", Data: "kkk"}
		c.JSON(200, resp)
		response, _ := json.Marshal(resp)
		c.Set("response", string(response))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}