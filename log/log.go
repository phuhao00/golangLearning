package log
import (
	"io"
	"log"
	"os"
)
var (
	Info *log.Logger
	Warning *log.Logger
	Error * log.Logger
)
func init() {
	errFile,err1:=os.OpenFile("log/errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err1!=nil{
		log.Fatalln("打开日志文件失败：",err1)
	}
	warningFile,err2:=os.OpenFile("log/warnings.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err2!=nil{
		log.Fatalln("打开日志文件失败：",err2)
	}
	InfoFile,err3:=os.OpenFile("log/Infos.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err3!=nil{
		log.Fatalln("打开日志文件失败：",err3)
	}

	Info	 	= log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	Warning 	= log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile)
	Error 		= log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)
	Warning 	= log.New(io.MultiWriter(os.Stderr,warningFile),"Warning:",log.Ldate | log.Ltime | log.Lshortfile)
	Info 		= log.New(io.MultiWriter(os.Stderr,InfoFile),"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	//Info.Printf("<<<<<<<<< Task_State Check >>>>>>>>>>>>>>\n")
}
