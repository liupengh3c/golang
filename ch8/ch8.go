package ch8

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// Ch8 第八章入口函数
func Ch8() {
	// first()
	clock()
	netcat()
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r计算中：%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// first 第一个简单的并发测试程序
// 主进程结束，其他所有goroutine暴力的直接终结
func first() {
	go spinner(100 * time.Millisecond)
	fmt.Printf("\nfib(45)=%d\n", fib(45))
	return
}

func handConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// clock 并发时钟，第一个示例程序
// nc localhost 8000命令行可以链接验证
func clock() {
	listener, _ := net.Listen("tcp", "localhost:8000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handConn(conn)
	}
}

// netcat go版本netcat，可以替代nc命令
func netcat() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("net dial error,err=" + err.Error())
	}
	defer conn.Close()
	io.Copy(os.Stdout, conn)
	return
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("io copy err,err=" + err.Error())
	}
	return
}

var ch = make(chan int)

func test(index int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1 + index
	}
}
func tests() {
	for i := 0; i < 2; i++ {
		go test(i * 10)
	}
}
