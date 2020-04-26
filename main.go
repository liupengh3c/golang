package main

import (
	"golang/ch5"
)

func main() {
	// ch1.GetStdCmd()
	// ch1.Dup2()
	// ch1.Dup3()
	// ch1.Lissajous()
	//*************爬虫*************//
	// ch := make(chan string)
	// for _, url := range os.Args[1:] {
	// 	go ch1.ParallelFetch(url, ch)
	// }
	// for range os.Args[1:] {
	// 	fmt.Println(<-ch)
	// }
	//**************************//
	// ch1.Fetch("http://www.baidu.com")
	//*************网络服务器*************//
	// http.HandleFunc("/test", ch1.Handle)
	// http.ListenAndServe("localhost:8000", nil)

	//--------------ch2示例函数--------------//
	// ch2.Echo4()
	// ch2.Gcd(9800, 400)
	// ch2.Fib(1)
	// ch2.Area()

	//--------------ch3示例函数--------------//
	// ch3.Unicode()
	// ch3.Basename("/home/work/leart/tool.go")

	// ch3.Basename2("/home/work/leart/tool.go")

	// ch3.Comma("123433433433")
	// ch3.Strings("iam have a dream iam", "iam")
	// s := []int{}
	// for i := 0; i < 10; i++ {
	// 	s = append(s, i)
	// }
	// fmt.Println(ch3.IntsToString(s))
	// ch3.Translate()
	// ch3.Const()
	// if ch3.IsSame("banlace", "banlace") {
	// 	fmt.Println("is same")
	// } else {
	// 	fmt.Println("is not same")
	// }
	// ch3.Comma2("123455712")

	//--------------ch4示例函数--------------//
	// ch4.Ch4()

	//--------------ch5示例函数--------------//
	ch5.Ch5()
	// append函数，s后的省略号代表将整个slice追加到work
	// var work = []int{1, 2}
	// var s = []int{3, 4}
	// work = append(work, s...)
	// fmt.Println(work)
}
