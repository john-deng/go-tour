package primary

import (
	"log"
	"testing"
)

func TestLogInfo(t *testing.T)  {
	log.Printf("http://www.flysnow.org")
	log.SetFlags(log.Ldate|log.Ltime |log.LUTC)
}

func init()  {
	//包的名字    行数
	log.SetFlags(log.Ldate|log.Lshortfile)
}

const (
	Ldate = 1<<iota  //日期示例 2009/01/23
	Ltime  // 时间示例 01：23；23
	Lmicroseconds //毫秒
	Llongfile
	Lshortfile  //文件行数
	LUTC   //文件
	LstdFlags = Ldate|Ltime //抬头信息

)




