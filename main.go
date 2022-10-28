package main

import (
	"fmt"
	"log"

	"github.com/sirzzang/learngo/accounts"
)

func main() {

	account := accounts.NewAccount("sirzzang") // make a new struct
	fmt.Println(account)                       // string representation
	fmt.Println(account.Owner(), account.Balance())
	account.Deposit(100)
	fmt.Println(account.Balance()) // 100

	// account.Withdraw(1000)         // error returned, nothing happens
	err := account.Withdraw(10) // check error
	if err != nil {
		log.Fatalln(err) // 2022/10/28 22:44:05 can't withdraw
	}
	fmt.Println(account.Balance())

	account.ChangeOwner("eraser")
	fmt.Println(account.Owner(), account.Balance())

}
