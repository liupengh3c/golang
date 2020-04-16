package main

import (
	"golang/ch2"
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

	//*************网络服务器*************//
	// http.HandleFunc("/test", ch1.Handle)
	// http.ListenAndServe("localhost:8000", nil)

	//--------------ch2示例函数--------------//
	ch2.Echo4()
	ch2.Gcd(9800, 400)
	ch2.Fib(1)
}
