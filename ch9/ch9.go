package ch9

import (
	"fmt"
)

// Ch9 第九章入口测试函数
func Ch9() {
	withdraw()
}

type money struct {
	cnt int32
	ok  bool
}

var deposits = make(chan money)
var balances = make(chan int32)
var hua = make(chan money)

func deposit(amount int32) {
	var mon money
	mon.cnt = amount
	deposits <- mon
}

func balance() int32 {
	cur := <-balances
	return cur
}
func teller() {
	var balance int32 = 200
	for {
		select {
		case dep := <-deposits:
			balance += dep.cnt
		case balances <- balance:
		case cnt := <-hua:
			var draw money
			if cnt.cnt > balance {
				draw.cnt = balance
				draw.ok = false
			} else {
				balance = balance - cnt.cnt
				draw.cnt = balance
				draw.ok = true
			}
			hua <- draw
		}
	}
}

func init() {
	go teller()
}
func withdraw() {
	var consume money
	consume.cnt = 300
	hua <- consume

	cur := <-hua
	if cur.ok {
		fmt.Println("money is enough,the balance=" + fmt.Sprintf("%v", cur.cnt))
	} else {
		fmt.Println("money is not enough,the balance=" + fmt.Sprintf("%v", cur.cnt) + fmt.Sprintf(",draw money=%v", 300))
	}
}
