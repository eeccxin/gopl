package main

import (
	"fmt"
)

/*
练习 9.1： 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。其返回结果应
该要表明事务是成功了还是因为没有足够资金失败了。这条消息会被发送给monitor的
goroutine，且消息需要包含取款的额度和一个新的channel，这个新channel会被monitor
goroutine来把boolean结果发回给Withdraw。
*/

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan int)
var withdrawResult = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) {
	withdraws <- amount
	ok := <-withdrawResult
	if !ok {
		fmt.Printf("取钱失败:取款%d,但余额只有%d\n", amount, Balance())
	} else {
		fmt.Printf("取钱成功:取款%d,当前余额%d\n", amount, Balance())
	}

}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Printf("存款%d,当前余额%d\n", amount, balance)
		case price := <-withdraws:
			if balance >= price {
				balance -= price
				withdrawResult <- true
				break
			}
			withdrawResult <- false
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {
	Deposit(100)
	Withdraw(50)
	Withdraw(70)
	Deposit(200)

	fmt.Println("Final balance:", Balance())
}

//!-
