#PProf之性能剖析
## 一、前言
应用程序在运行时，总会出现一些意想不到的问题，比如突然报警、监控系统提示进程CPU使用率过高、内存占用不断增大（内存泄漏）、临时内存大量申请后长时间不下降，或是goroutine泄漏、goroutine数量暴涨等。
## 二、PProf
在Go语言中，PProf是分析性能、分析数据的工具，PProf用profile.proto读取分析样本的集合，并生成可视化，以帮助分析数据。
### 1.使用模式
- `Report Generation`: 报告生成
- `Interactive Terminal Use`: 交互式终端使用。
- `Web Interface`: Web界面

### 2.作用
- `CPU Profiling`: CPU分析。按照一定频率采集所监听的应用程序CPU（含寄存器）的使用情况，确定应用程序在主动消耗CPU周期时花费时间的位置
- `Memory Profiling`: 内存分析。在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以及检查内存泄漏。
- `Block Profiling`：阻塞分析。记录goroutine阻塞等待同步（包括定时器通道）的位置，默认不开启，需要调用runtime.SetBlockProfileRate进行设置。
- `Mutex Profiling`：互斥锁分析。报告互斥锁的竞争情况，默认不开启，需要调用runtime.SetMutexProfileFraction进行设置。
- `Goroutine Profiling`：goroutine分析，可以对当前应用程序正在运行的goroutine进行堆栈跟踪和分析。这项功能在实际排查中会经常用到，因为很多问题出现时的表象就是goroutine暴增，比如go出去的方法，因为channel、互斥锁的不正确使用导致goroutine堵塞，无法结束，从而导致goroutine的数量随着时间越来越多，直到占满内存，程序崩溃。而这时候，我们可以查看应用程序中的goroutine正在做什么事情，因为什么造成的堵塞，才能更好的进行下一步排错。

## 三、PProf的使用
### 1.通过浏览器访问
`package pprof`
```
// To use pprof, link this package into your program: 
// 要使用pprof，请将这个包导入到你的项目中。
//	import _ "net/http/pprof" 
//
// If your application is not already running an http server, you
// need to start one. Add "net/http" and "log" to your imports and
// the following code to your main function:
// 如果您的应用程序还没有运行http服务器，你
// 需要启动一个。添加“net/http”和“log”到您的导入和
// 下面的代码到您的主函数
//
// 	go func() {
// 		log.Println(http.ListenAndServe("localhost:6060", nil))
// 	}()
//
// If you are not using DefaultServeMux, you will have to register handlers
// with the mux you are using.
//
// Then use the pprof tool to look at the heap profile:
//
//	go tool pprof http://localhost:6060/debug/pprof/heap
//
// Or to look at a 30-second CPU profile:
//
//	go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
//
// Or to look at the goroutine blocking profile, after calling
// runtime.SetBlockProfileRate in your program:
//
//	go tool pprof http://localhost:6060/debug/pprof/block
//
// Or to collect a 5-second execution trace:
//
//	wget http://localhost:6060/debug/pprof/trace?seconds=5
//
// Or to look at the holders of contended mutexes, after calling
// runtime.SetMutexProfileFraction in your program:
//
//	go tool pprof http://localhost:6060/debug/pprof/mutex
//
// To view all available profiles, open http://localhost:6060/debug/pprof/
// in your browser.
//
// For a study of the facility in action, visit
//
//	https://blog.golang.org/2011/06/profiling-go-programs.html
//
package pprof

func init() {
	http.HandleFunc("/debug/pprof/", Index)
	http.HandleFunc("/debug/pprof/cmdline", Cmdline)
	http.HandleFunc("/debug/pprof/profile", Profile)
	http.HandleFunc("/debug/pprof/symbol", Symbol)
	http.HandleFunc("/debug/pprof/trace", Trace)
}
```

`packge main`
```
package main

import (
	"net/http"
	_ "net/http/pprof" // 导入pprof
)
func main(){
	http.ListenAndServe("0.0.0.0:6061", nil) // 开启端口监听
}

```
	访问：http://localhost:6061/debug/pprof/


| 类型 | 描述 |
| :----------: | :----------: |
|  allocs | 查看过去所有内存分配的样本。  |
| block  | 查看导致阻塞同步的堆栈跟踪  |
| cmdline  |  当前程序命令行的完整调用路径 |
| goroutine  | 查看当前所有运行的goroutine堆栈跟踪。  |
| heap  | 查看活动对象的内存分配情况。  |
| mutex  | 查看导致互斥锁的竞争持有者的堆栈跟踪。  |
| profile  | 默认进行30s的CPU Profiling，会得到一个分析用的profile文件。  |
| threadcreate  | 查看创建新OS线程的堆栈跟踪。  |

	 若不增加debug参数，那么会直接下载对应的profile文件。

### 2.通过交互式终端使用

#### (1)CPU Profiling
`go tool pprof http://localhost:6061/debug/pprof/profile?seconds=10`

	在执行该命令后，等待10s(可修改seconds的值)，PProf会进行Profiling，然后进入PProf的命令行交互模式。通过去掉参数debug=1下载下来的profile文件，也可以通过命令行，进入PProf命令交互模式。
	使用top10可查看对应的资源开销排名前十的函数。

![](http://showdoc.siwei.com/server/index.php?s=/api/attachment/visitFile/sign/a2ffe6dd15a2f4d1a908a309ece9d54f&showdoc=.jpg)

- falt: 函数自身的运行耗时
- falt%: 函数自身占CPU运行总耗时的比例。
- sum%: 函数自身累计使用占CPU运行总耗时比例。
- cum: 函数自身及其调用函数的运行总耗时。
- cum%: 函数自身及其调用函数占CPU运行总耗时的比例。
- name: 函数名


 在大多数情况下，我们可以得出一个应用程序的运行情况，知道当前是什么函数，正在做什么事情，占用了多少资源等等。
 
```
//以下代码有什么问题：
var ch = make(chan int,10)

func main(){

	go func() {
		var i int
		for {
			ch <- i
			i++
			time.Sleep(1 * time.Second)
		}
	}()

	for i:=0;i<3;i++{
		go TestCPUProfiling(ch)
	}

	http.ListenAndServe("0.0.0.0:6061", nil)
}


func TestCPUProfiling(ch chan int){
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		default:
		}
	}
}
```

#### (2)Heap Profiling

![](http://showdoc.siwei.com/server/index.php?s=/api/attachment/visitFile/sign/a9f236333f8857e811843f5b4ef699e3&showdoc=.jpg)
在执行该命令后，能够很快地拉取结果，因为它不像CPU Profiling那样需要做采样等待。需要注意的是Type选项的默认类型是insue_space，实际上，它可以对多种内存概况进行分析，常用类别如下：
- inuse_space: 分析应用程序常驻内存的占用情况。
- alloc_objects:分析应用程序的内存临时分配情况。

```
package main

import (
"fmt"
"net/http"
_ "net/http/pprof"
"os"
"time"
)

// 运行一段时间：fatal error: runtime: out of memory
func main() {
	// 开启pprof
	go func() {
		ip := "0.0.0.0:6061"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
		time.Sleep( 3*time.Second)
	}
}
```


- heap能帮助我们发现内存问题，但不一定能发现内存泄露问题。heap记录了内存分配的情况，我们能通过heap观察内存的变化，增长与减少，内存主要被哪些代码占用了，程序存在内存问题，这只能说明内存有使用不合理的地方，但并不能说明这是内存泄露。
- heap在帮助定位内存泄露原因上贡献的力量微乎其微。如第一条所言，能通过heap找到占用内存多的位置，但这个位置通常不一定是内存泄露，就算是内存泄露，也只是内存泄露的结果，并不是真正导致内存泄露的根源。
- 对于goroutine泄露，这是通过heap无法发现的，所以heap在定位内存泄露这件事上，发挥的作用不大。


#### (3)Goroutine Profiling

##### 1.什么是goroutine泄露
如果你启动了1个goroutine，但并没有符合预期的退出，直到程序结束，此goroutine才退出，这种情况就是goroutine泄露。
##### 2.goroutine泄露怎么导致内存泄露
每个goroutine占用2KB内存，泄露1百万goroutine至少泄露2KB * 1000000 = 2GB内存，为什么说至少呢？

goroutine执行过程中还存在一些变量，如果这些变量指向堆内存中的内存，GC会认为这些内存仍在使用，不会对其进行回收，这些内存谁都无法使用，造成了内存泄露。

所以goroutine泄露有2种方式造成内存泄露：
1. goroutine本身的栈所占用的空间造成内存泄露。
1. goroutine中的变量所占用的堆内存导致堆内存泄露，这一部分是能通过heap profile体现出来的。

##### 3.怎么确定是goroutine泄露引发的内存泄露
判断依据：在节点正常运行的情况下，隔一段时间获取goroutine的数量，如果后面获取的那次，某些goroutine比前一次多，如果多获取几次，是持续增长的，就极有可能是goroutine泄露。


```
// goroutine泄露导致内存泄露
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	// 开启pprof
	go func() {
		ip := "0.0.0.0:6061"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	outCh := make(chan int)
	// 死代码，永不读取
	go func() {
		if false {
			<-outCh
		}
		select {}
	}()

	//每s起100个goroutine，goroutine会阻塞，不释放内存
	tick := time.Tick(time.Second / 100)
	i := 0
	for range tick {
		i++
		fmt.Println(i)
		go alloc(outCh)
	}
}

func alloc(outCh chan<- int) {
	defer fmt.Println("alloc-fm exit")
	// 分配内存，假用一下
	buf := make([]byte, 1024*1024*10)
	_ = len(buf)
	fmt.Println("alloc done")

	outCh <- 0
}

```

#### (4)Mutex Profiling
一般来说，在调用chan、sync.Mutex或time.Sleep()时，会造成阻塞。
Go内建的锁冲突分析支持可以方便地发现这样的锁冲突。在程序中，可以通过runtime.SetMutexProfileFraction()来设置采样频率，通过将其设为一个高于0的值来开启锁冲突分析的数据收集。

#### (5)Block Profiling
与Mutex的runtime.SetMutexProfileFraction语句相似，Block也需要调用runtime.SetBlockProfileRate语句进行设置。


```
package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

func init() {
	runtime.SetMutexProfileFraction(1)
	//runtime.SetBlockProfileRate(1)
}

func main() {
	var m sync.Mutex
	var datas = make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	_ = http.ListenAndServe(":6061", nil)
}
```

### 3.查看可视化界面
通过`http://localhost:6061/debug/pprof/profiel`，等待30s执行完毕后，会采集profile文件。
1. 该命令将在指定的端口号运行一个PProf分析用的站点。
`go tool pprof -http=:6001 profile`

2. 通过Web命令将profile文件以svg的文件格式写入图形，然后再Web浏览器中将其打开。
`go tool pprof profile`

如果出现错误提示“Could not execut dot；may need to install graphviz.”则表示需要安装graphviz组件。

### 4.总结
在程序中提前埋入PProf后，遇到问题的基本排查步骤：
1. 通过访问埋点的端口，查看异常情况，并采集对应的异常profile。
2. 使用top命令查看问题函数。
3. traces和web命令，查看调用过程。
4. list命令查看问题函数，具体出现问题的代码行数。
