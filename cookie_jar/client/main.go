package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	jar, _ := cookiejar.New(nil)
	fmt.Println("Start Request Server")
	client := http.Client{
		Jar: jar,
	}
	url := "http://127.0.0.1:8889/test"
	req, _ := http.NewRequest("GET", url, nil)
	//第一次发请求
	client.Do(req)
	fmt.Printf("第一次 %s \n", req.Cookies())

	//第二次发请求
	client.Do(req)
	fmt.Printf("第二次 %s \n", req.Cookies())

	//第三次发请求
	client.Do(req)
	fmt.Printf("第三次 %s \n", req.Cookies())

	//第四次发请求
	client.Do(req)
	fmt.Printf("第四次 %s \n", req.Cookies())

	//第五次发请求
	client.Do(req)
	fmt.Printf("第五次 %s \n", req.Cookies())

}

////didTimeout is non-nil only if err != nil.
//func (c *Client) send(req *Request, deadline time.Time) (resp *Response, didTimeout func() bool, err error) {
//	if c.Jar != nil {
//		for _, cookie := range c.Jar.Cookies(req.URL) {
//			req.AddCookie(cookie)  // 这里
//		}
//	}
//	resp, didTimeout, err = send(req, c.transport(), deadline)
//	if err != nil {
//		return nil, didTimeout, err
//	}
//	if c.Jar != nil {
//		if rc := resp.Cookies(); len(rc) &gt; 0 {
//			c.Jar.SetCookies(req.URL, rc) // 这里
//		}
//	}
//	return resp, nil, nil
//}
