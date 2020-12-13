

在本专栏开始:先看一段代码
- 这是这个接口的controller层
![](http://showdoc.siwei.com/server/index.php?s=/api/attachment/visitFile/sign/288ee42c60c68405011e05495bbf6ed7&showdoc=.jpg)
- 这是这个接口的service层
![](http://showdoc.siwei.com/server/index.php?s=/api/attachment/visitFile/sign/33c0e04ca0a01913ac019e89b7c653e2&showdoc=.jpg)
-这个接口没有dao层，但是在数据请求的时候我看到这一句
![](http://showdoc.siwei.com/server/index.php?s=/api/attachment/visitFile/sign/622140c492d913054ebbc3dbfe193ad7&showdoc=.jpg)


###错误处理应该是怎么样一个姿势

1.了解golang error 的官方建议

如果一个函数返回了 value, error，你不能对这个 value 做任何假设，必须先判定 error。唯一可以忽略 error 的是，如果你连 value 也不关心。
Go 中有 panic 的机制，如果你认为和其他语言的 exception 一样，那你就错了。当我们抛出异常的时候，相当于你把 exception 扔给了调用者来处理。

you only need to check the error value if you care about the result.  -- Dave
This blog post from Microsoft’s engineering blog in 2005 still holds true today, namely:
My point isn’t that exceptions are bad. My point is that exceptions are too hard and I’m not smart enough to handle them.
简单。
考虑失败，而不是成功(Plan for failure, not success)。
没有隐藏的控制流。
完全交给你来控制 error。
Error are values。




###处理API错误

1.分清楚服务对象：我们到底应该是TOC还是对我们自己
2.理解客户/操作者 到底什么能看，什么不能看

-比如说现在有个sql错误，肯定是给我们开发人员看，不过什么运营人员还是用户都不知道这个干什么的，他们只需要知道怎么了：服务器错误，还是网络错误，还是什么参数没有填写导致的错误，解决办法就是- 错误map

3.如何用好日志

- 学会分层，将错误都返回到controller层进行处理，打印好错误堆栈（一般是程序顶部或者工作的地方进行记录）
- 记录根本：1.错误原因，2，包裹错误，3.向上返回 4.统一记录
- 发生错误就要把所有的参数记录下来，header ，body，query等

想一想下面错误几时出来
```
1.sql error 
2. param error 
3. change param string to int error

```
###我们应该




痛点一 :底层的错误到底应该在哪里处理

```
func main(){
	_,err:= B()
	if err !=nil {
		//todo record lo
	}
}
func A()(int,error) {
	return 0,errors.New("this is a Error")
}

func B()(int,error) {
	_,err:= A()
	if err !=nil {
		//todo record log
		return 0,err
	}
	// todo other
	return 1,err
}

func C()(int,error) {
	_,err:= B()
	if err !=nil {
		//todo record log
		return 0,err
	}
	return 1,err
}
```


痛点二：到处都err怎么办？

```
func CountLine(r io.Reader)(int,error){
	
	var (
		br =bufio.NewReader(r)
		line int
		err error
	)
	for {
		_,err =br.ReadString('\n')
		line++
		if err!=nil{
			break
		}
	}
	
	if err !=io.EOF{
		return 0,err
	}
	return line,nil
}
```

入手：从调用方入手*我们作为调用方的时候也可以这么处理
```
func CountLines(r io.Reader)(int,error){

	var (
		sc =bufio.NewScanner(r)
		lines = 0
	)
	for sc.Scan(){
		lines ++
	}

	return lines,sc.Err()
}
```


3.第三方包没有封装错误，我使用起来都是错误该不该忽略？

```
func WriteResponse(w io.Writer,st Status,header []Header,body io.Reader)error{
	_,err := fmt.Fprintf(w,"HTTP/1.1 %d %s\r\n",st.Code,st.Reason)
	if err !=nil {
		return err
	}
	
	for _,h :=range header{
		_,err :=fmt.Fprintf(w,"%s: %s\r\n",h.Key,h.Value)
		if err!=nil{
			return err
		}
	}
	
	if _,err := fmt.Fprintf(w,"\r\n");err !=nil{
		return err
	}
	
	_,err = io.Copy(w,body)
	
	return err
}

```

解决：从结构体封装开始

```
func (e *errWriter) Write(buf []byte)(int,error){
	if e.err !=nil {
		return 0,e.err
	}
	var n int
	n ,e.err =e.Writer.Write(buf)
	return n,nil
}

func WriteResponse(w io.Writer,st Status,header []Header,body io.Reader)error{
	ew := &errWriter{Writer:w}
	fmt.Fprintf(w,"HTTP/1.1 %d %s\r\n",st.Code,st.Reason)

	for _,h :=range header{
		fmt.Fprintf(w,"%s: %s\r\n",h.Key,h.Value)
	}
	fmt.Fprintf(w,"\r\n")
	io.Copy(w,body)
	return ew.err

}
```



error的类型
- sentinal

- errorType 