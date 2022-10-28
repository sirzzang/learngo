package something

import "fmt"

// private function
func sayBye() {
	fmt.Println("Bye")
}

// exported function
func SayHello() {
	fmt.Println("Hello")
}
