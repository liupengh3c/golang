package ch8

import (
	"fmt"
	"io"
	"net"
	"time"
)

// Ch8 第八章入口函数
func Ch8() {
	// first()
	clock()
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
func clock() {
	listener, _ := net.Listen("tcp", "localhost:8000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handConn(conn)
	}
}
