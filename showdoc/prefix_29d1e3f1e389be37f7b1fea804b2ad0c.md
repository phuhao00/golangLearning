# golang并发编程
###1.理解计算机的空间和时间的概念
空间：cup的算力，内存空间，io读写操作
时间：处理程序消费的时间



###2.并发/并行程序

计算机程序需要区分进程、线程(内核级线程)、协程(用户级线程)三个概念。

对于 进程、线程，都是有内核进行调度，有 CPU 时间片的概念，进行 抢占式调度（有多种调度算法）

对于 协程(用户级线程)，这是对内核透明的，也就是系统并不知道有协程的存在，是完全由用户自己的程序进行调度的，因为是由用户程序自己控制，那么就很难像抢占式调度那样做到强制的 CPU 控制权切换到其他进程/线程，通常只能进行 协作式调度，需要协程自己主动把控制权转让出去之后，其他协程才能被执行到。

实际的并行运行的程序只是通过电脑cpu核心数量来区分的，比如8核，那么最多同时运行的程序就是8个


#### 3.1 线程的分类
线程的实现曾有3种模型：
1.多对一(M:1)的用户级线程模型 ：协程
2.一对一(1:1)的内核级线程模型 ：线程
3.多对多(M:N)的两级线程模型 	：goroutine

上面的x对y(x:y)即x个用户线程对应y个内核调度实体(Kernel Scheduling Entity，这个是内核分配CPU的对象单位)。
！

#####一.多对一用户线级程模型

多对一线程模型中，线程的创建、调度、同步的所有细节全部由进程的用户空间线程库来处理。用户态线程的很多操作对内核来说都是透明的，因为不需要内核来接管，这意味不需要内核态和用户态频繁切换。线程的创建、调度、同步处理速度非常快。当然线程的一些其他操作还是要经过内核，如IO读写。这样导致了一个问题：当多线程并发执行时，如果其中一个线程执行IO操作时，内核接管这个操作，如果IO阻塞，用户态的其他线程都会被阻塞，因为这些线程都对应同一个内核调度实体。在多处理器机器上，内核不知道用户态有这些线程，无法把它们调度到其他处理器，也无法通过优先级来调度。这对线程的使用是没有意义的！


#####二.一对一内核极线程模型

一对一模型中，每个用户线程都对应各自的内核调度实体。内核会对每个线程进行调度，可以调度到其他处理器上面。当然由内核来调度的结果就是：线程的每次操作会在用户态和内核态切换。另外，内核为每个线程都映射调度实体，如果系统出现大量线程，会对系统性能有影响。但该模型的实用性还是高于多对一的线程模型。

#####三.多对多两极线程模型

多对多模型中，结合了1：1和M：1的优点，避免了它们的缺点。每个线程可以拥有多个调度实体，也可以多个线程对应一个调度实体。听起来好像非常完美，但线程的调度需要由内核态和用户态一起来实现。可想而知，多个对象操作一个东西时，肯定要一些其他的同步机制。用户态和内核态的分工合作导致实现该模型非常复杂。NPTL曾经也想使用该模型，但它太复杂，要对内核进行大范围改动，所以还是采用了一对一的模型！！！

#### 3.2goroutine 到底是携程吗？
Goroutine，Go语言基于并发（并行）编程给出的自家的解决方案。goroutine是什么？通常goroutine会被当做coroutine（协程）的 golang实现，从比较粗浅的层面来看，这种认知也算是合理，但实际上，goroutine并非传统意义上的协程，现在主流的线程模型分三种：内核级线程模型、用户级线程模型和两级线程模型（也称混合型线程模型），传统的协程库属于用户级线程模型，而goroutine和它的Go Scheduler在底层实现上其实是属于两级线程模型，因此，有时候为了方便理解可以简单把goroutine类比成协程，但心里一定要有个清晰的认知 — goroutine并不等同于协程


###说了这么多，你理解了什么是并行/并发了吗

![](http://showdoc.siwei.com/server/index.php?s=/api/attachment/visitFile/sign/b4b1e6be83f0e9acc21f5dbf83da61e0&showdoc=.jpg)

![](http://showdoc.siwei.com/server/index.php?s=/api/attachment/visitFile/sign/c9c41a5c6c1867fa0f4b761d807df5b5&showdoc=.jpg)

总结一下：
并行：
程序上单台机器是不可能做到真正意义上的并行，我理解的并行就是两台资源：网络带宽完全不影响的两台机器做同一件事情这就是并行

并发：
百度释义：个时间段中有几个程序都处于已启动运行到运行完毕之间，且这几个程序都是在同一个处理机上运行，但任一个时刻点上只有一个程序在处理机上运行。

我的理解：针对程序而言，在一段时间内处理了多少请求/数据，然后再讲这个时间除以以秒计算的单位时间的总值，也就是所谓的并发量（放在api层面就称为qps）



##golang 如何实现并发编程的


###关键字 GO
goroutine可以认为是轻量级的线程，与创建线程相比，创建成本和开销都很小，每个goroutine的堆栈只有几kb，并且堆栈可根据程序的需要增长和缩小(线程的堆栈需指明和固定)，所以go程序从语言层面支持了高并发,windows下创建是4kb，linux下创建是2kb，mac下是3kb

程序执行的背后：当一个程序启动的时候，只有一个goroutine来调用main函数，称它为主goroutine，新的goroutine通过go语句进行创建。

下面几个实例需要好好考虑一下输出，并提出为什么？
示例一
```
package main

import (
	"fmt"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello world goroutine")
}

func main() {
	go HelloWorld()      // 开启一个新的并发运行
	time.Sleep(1*time.Second)
	fmt.Println("hello world main")
}

```
示例二

```
package main

import (
	"fmt"
	"time"
)

func DelayPrint() {
	for i := 1; i <= 4; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(i)
	}
}

func HelloWorld() {
	fmt.Println("Hello world goroutine")
}

func main() {
	go DelayPrint()    // 开启第一个goroutine
	go HelloWorld()    // 开启第二个goroutine
	time.Sleep(2*time.Second)
	fmt.Println("main function")
}

```
###并发等待模型

#####不通过锁实现 feature 并发模式

```
func main() {
    ch := &chSql{
        in:  make(chan string, 3),
        out: make(chan interface{}, 3),
    }
    go dowork(ch)
    sql1 := "select * from "
    ch.in <- sql1
    //do something else
    fmt.Println("我吃了")
    //get result
    fmt.Println(<-ch.out)
}
type chSql struct {
    in  chan string
    out chan interface{}
}
func dowork(sql *chSql) {
    s := <-sql.in
    fmt.Println(s)
    sql.out <- "hahha"
}
```

#####通过waitgroup 实现并发模式


```
type wg struct {
    *sync.WaitGroup
}
func main() {
    Tem := wg{
        &sync.WaitGroup{},
    }
    Tem.Add(10)
    for i := 0; i < 10; i++ {
        go func(i int) {
            defer Tem.Done()
            time.Sleep(5 * time.Second)
            fmt.Println(i, "finished")
        }(i)
    }
    Tem.Wait()
}

```


###Chan 不用通过共享内存来通讯，通过通讯来共享内存（重点）

chan 是goroutine进行通讯的方式，如何理解不要通过共享内存来通讯，而是通过通讯来共享内存：假设，abc是三个方法，同时需要对变量age+1操作，那么goroutine之间没有先后顺序，没有层级主次之分，如何保证不出错呢，那么就可以使用锁。此时age就是一个共享内存的变量，被三个goroutine占用，这就是共享内存来通讯大家都是抢占式。其实这也没错，在并发编程模型中，这是非常常用的一种手段，通过给资源加锁就可以实现共享又保证数据正确，但是换个场景：我要a+1后b+1,b+1后c+1,那么这种模式可能就不行了



什么是channel
goroutine运行在相同的地址空间，因此访问共享内存必须做好同步，处理好线程安全问题
goroutine奉行通过通信来共享内存，而不是共享内存来通信
channel是一个引用类型，用于多个goroutine通讯，其内部实现了同步，确保并发安全


####channel基本使用

channel可以用内置make()函数创建
定义一个channel时，也需要定义发送到channel的值的类型
```
 make(chan 类型)   //无缓冲的通道
  make(chan 类型, 容量) //有缓冲的通道
```
当 capacity= 0 时，channel 是无缓冲阻塞读写的，当capacity> 0 时，channel 有缓冲、是非阻塞的，直到写满 capacity个元素才阻塞写入

channel通过操作符<-来接收和发送数据，发送和接收数据语法：

```
  package main
  import "fmt"

  func main() {
     //创建存放int类型的通道
     c := make(chan int)
     //子协程
     go func() {
        defer fmt.Println("子协程结束")
        fmt.Println("子协程正在运行...")
        //将666发送到通道c
        c <- 666
     }()
     //若已取出数据，下面再取会报错
     //<-c
     //主协程取数据
     //从c中取数据
     num := <-c
     fmt.Println("num = ", num)
     fmt.Println("主协程结束")
  }
```


###go带来的问题和挑战'
想象一下下面的代码会发生怎么样的问题

```
func main() {
    userCount := math.MaxInt64
    for i := 0; i < userCount; i++ {
        go func(i int) {
            // 做一些各种各样的业务逻辑处理
            fmt.Printf("go func: %d\n", i)
            time.Sleep(time.Second)
        }(i)
    }
}
```

goroutine过多
在前面花了大量篇幅，渲染了在存在大量并发 goroutine 数量时，不控制的话会出现 “严重” 的问题，接下来一起思考下解决方案。如下：

控制/限制 goroutine 同时并发运行的数量
改变应用程序的逻辑写法（避免大规模的使用系统资源和等待）
调整服务的硬件配置、最大打开数、内存等阈值



###问题？
1.假设我们有一个消息队列，其中有1000w条数据，我们如何来进行设计程序来进行消耗？假设每条数据我们需要1秒钟的时间

2.如何退出goroutine 中有死循环的go？


3.什么样的场景下使用chan的广播


2.如何合理的关闭掉chan

```
func MonitorNode(ctx *gin.Context, node *Node, token string) {
	go sendProc(node,token)        //消息发送
	go recvProc(ctx, node, token) //消息发送
}


func sendProc(node *Node,token string) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				//public.GetLogger().Error("连接写入数据错误", zap.String("error", err.Error()))
				closeConnection(node,token)
				break loop
}}}}


func recvProc(ctx *gin.Context, node *Node, token string) {
	loop:
	for {
		select {
		default:
			_, data, err := node.Conn.ReadMessage()
			if err != nil {
				//public.GetLogger().Error("读取连接错误", zap.String("error", err.Error()))
				closeConnection(node,token)
				break loop
			}
			//消息推送
			Dispatch(ctx, data, token,node)
			}}}

```

3.并发编程下我们如何去尽量提高cup的利用率？




###加餐：我们常说的高并发量究竟是什么

我通常站立的角度就是抛开场景谈性能就是耍流氓，谈极端场景的状况也是耍流氓：比如十万用户同时修改头像，这是万万不可能发生的咱就可以不用考虑成并发场景。

在服务器上我们通常按照服务器的一个cup，内存的一个指标去判断一个程序的好坏，拿实际公司的例子，在我之前在的公司服务器报警指标在：70%左右，那么在API服务中，我们的峰值时刻的cup,内存占比应该就在70以下你就算一个稳定的程序，如果超过这个值，那么服务器在稍稍极端情况下就会宕机。我们如何去模拟这个数值呢：这里经供参考：

按照二八定律：
用户总量：10w
日均pv：100w
用户活跃时间段；：9点-21点
总时间段：12小时


峰值qps：100w*0.8/(12*0.2*3600)
说明：百分之八十的流量会在百分之二十的时间到达
大概在92/s左右

其实大家可以看出所谓的高并发峰值实际上是根据实际业务的活跃时间来判断的，所有的并发和业务属性相关的



合理关闭chan

```
func main(){
    root:=context.Background()

    son,cancel:=context.WithCancel(root)

    for i:=0;i<10;i++{
    fmt.Println(<-CreateInt(son))
    }
    cancel()

}

//  简单的随机数生成器
func CreateInt(ctx context.Context)chan int{

    ch:=make(chan int)

    go func() {
        label:
            for{
                select {
                case <-ctx.Done():
                    break label
                case ch<-rand.Int():
                }
            }
            close(ch)
    }()

    return ch
}

```