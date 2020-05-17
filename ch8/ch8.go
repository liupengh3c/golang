package ch8

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

// Ch8 第八章入口函数
func Ch8() {
	// first()
	// clock()
	// netcat2()
	// pipe()
	// Noname()
	// ch8Mul()
	// ch8Tick2()
	// ch8DirMul2()
	ch8Chat()
}

func timeCost() func() {
	st := time.Now()
	return func() {
		fmt.Printf("time cost=%v\n", time.Since(st))
	}
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

func handConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		fmt.Fprintln(c, "\t", strings.ToUpper(input.Text()))
		time.Sleep(1 * time.Second)
		fmt.Fprintln(c, "\t", input.Text())
		time.Sleep(1 * time.Second)
		fmt.Fprintln(c, "\t", strings.ToLower(input.Text()))
	}
	c.Close()
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
		fmt.Println("a connecting is coming")
		// go handConn(conn)
		go handConn2(conn)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("io copy err,err=" + err.Error())
	}
	return
}

// netcat go版本netcat，可以替代nc命令
func netcat() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("net dial error,err=" + err.Error())
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
	return
}

func mustCopy2(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("io copy err,err=" + err.Error())
	}
	return
}

// netcat go版本netcat，可以替代nc命令
func netcat2() {
	conn, err := net.Dial("tcp", "106.13.105.231:8000")
	if err != nil {
		fmt.Println("net dial error,err=" + err.Error())
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdout)
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

func pipe() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for i := 0; ; i++ {
			naturals <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	// for {
	// 	fmt.Println(<-squares)
	// }

	// channel上没有数据后，循环会自动结束
	for x := range squares {
		fmt.Println(x)
	}
}

// Noname 并发匿名函数测试
func Noname() {
	ch := make(chan int)
	var wg sync.WaitGroup
	da := []string{"数据结构", "语文", "数学", "C语言"}
	for _, v := range da {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			ch <- len(s)
			fmt.Println(s)
		}(v)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 在主协程中wait会阻塞
	// wg.Wait()
	// close(ch)

	// 循环结束的条件是channel关闭，所以wait的goroutine在wait结束后必须显式关闭channel
	for v := range ch {
		fmt.Println(v)
	}
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// 计数信号量，最多创建20个goroutine
var token = make(chan struct{}, 20)

func crawl(url string) []string {
	list := []string{}
	fmt.Println(url)
	// struct{}是数据类型，struct{}{}是结构体类型空值
	token <- struct{}{}
	list, _ = extract(url)
	<-token
	return list
}
func ch8Mul() {
	worklist := make(chan []string)
	var n int
	n++
	// 匿名函数
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		lists := <-worklist
		for _, link := range lists {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
	fmt.Printf("n=%v", n)
}

// 8.7

func ch8Tick() {
	tick := time.Tick(1 * time.Second)
	for i := 10; i >= 0; i-- {
		fmt.Printf("%d\n", i)
		fmt.Println(<-tick)
	}
}

func ch8Tick2() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func(ch chan struct{}) {
		fmt.Println("enter read input")
		// 会阻塞，一直等待输入
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("read user input")
		ch <- struct{}{}
	}(abort)
	for i := 10; i >= 0; i-- {
		fmt.Printf("%d\n", i)
		select {
		case <-abort:
			fmt.Println("stop---------------")
			return
		case <-tick:

		case <-time.After(10 * time.Second):
			fmt.Println("begin")
			// default:
			// 	fmt.Println("default")
		}
	}
}
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	return entries
}
func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			walkDir(filepath.Join(dir, entry.Name()), fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func ch8Dir() {
	defer timeCost()()
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSize)
		}
		close(fileSize)
	}()
	cnt := 1
	for size := range fileSize {
		cnt++
		// fmt.Printf("the %d file,size=%[1]v\n", size)
		fmt.Printf("the %d file,size=%.1f GB\n", cnt, float64(size)/1e9)
	}
}

// walkDir2,并发递归调用
func walkDir2(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents2(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir2(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}
func ch8DirMul2() {
	defer timeCost()()
	flag.Parse()
	var done = make(chan struct{})
	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	go func() {
		os.Stdin.Read(make([]byte, 1))
		done <- struct{}{}
	}()
	//!+
	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir2(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	cnt := 0
	// for size := range fileSizes {
	// 	cnt++
	// 	fmt.Printf("the %d file,size=%.1f GB\n", cnt, float64(size)/1e9)
	// }
loop:
	//!+3
	for {
		select {
		case <-done:
			// for range fileSizes {
			// 	// Do nothing.
			// }
			fmt.Println("game over")
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			cnt++
			fmt.Printf("the %d file,size=%.1f GB\n", cnt, float64(size)/1e9)
		}
	}
}
func ch8DirMul() {
	defer timeCost()()
	flag.Parse()
	wg := sync.WaitGroup{}
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSize := make(chan int64)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, root := range roots {
			wg.Add(1)
			walkDir2(root, &wg, fileSize)
		}
		// close(fileSize)
	}()

	go func() {
		wg.Wait()
		close(fileSize)
	}()
	cnt := 0
	for size := range fileSize {
		cnt++
		fmt.Printf("the %d file,size=%.1f GB\n", cnt, float64(size)/1e9)
	}
}
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents2(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
func ch8Chat() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}
func clientWriter(conn net.Conn, ch <-chan string) {
	// for循环结束的条件是channel关闭
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
