package ch2

import (
	"flag"
	"fmt"
)

// Echo4 打印程序
func Echo4() {
	var n = flag.Bool("n", false, "omit trailing newline")
	sep := flag.String("s", " ", "separator")
	flag.Parse()
	fmt.Println("n=")
	fmt.Println(*n)
	fmt.Println("sep=")
	fmt.Println(*sep)

	// fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		// fmt.Println()
	}
}

// 变量的生命周期是写出高效程序所必需清楚的，例如，在长生命周期对象中保持短生命中期对象不必要的指针，特别是在全局变量中，会阻止垃圾回收器回收短生命周期的对象空间

// 定理：两个整数的最大公约数等于其中较小的那个数和两数相除余数的最大公约数。
// Gcd 求最大公约数
func Gcd(big, small int) {
	for small != 0 {
		big, small = small, big%small
	}
	fmt.Println(big)
}

// Fib 斐波那契数列
func Fib(n int) {
	x, y := 0, 1
	for i := 0; i < n-1; i++ {
		x, y = y, x+y
	}
	fmt.Println(x)
}
