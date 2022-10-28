package main

import (
	"fmt"

	"github.com/sirzzang/learngo/accounts"
	"github.com/sirzzang/learngo/mydict"
)

func main() {

	/** Account Example **/
	account := accounts.NewAccount("sirzzang") // make a new struct
	fmt.Println(account)                       // string representation
	fmt.Println(account.Owner(), account.Balance())
	account.Deposit(100)
	fmt.Println(account.Balance()) // 100
	// account.Withdraw(1000)         // error returned, nothing happens
	// err := account.Withdraw(10) // check error
	// if err != nil {
	// 	log.Fatalln(err) // 2022/10/28 22:44:05 can't withdraw
	// }
	// fmt.Println(account.Balance())

	account.ChangeOwner("eraser")
	fmt.Println(account.Owner(), account.Balance())

	/** Dictionary Example **/
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
	// definition, err := dictionary.Search("hello")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(definition)
	err := dictionary.Add("hi", "greeting")
	if err != nil {
		fmt.Println("err: ", err)
	}
	err2 := dictionary.Update("hi", "newGreeting")
	if err2 != nil {
		fmt.Println("err2: ", err2)
	}
	err3 := dictionary.Delete("hi")
	if err3 != nil {
		fmt.Println("err3: ", err3)
	}
	defintion, err4 := dictionary.Search("hi")
	if err4 != nil {
		fmt.Println("err4: ", err4)
	}
	fmt.Println(defintion)

}
