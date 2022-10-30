package main

import (
	"fmt"
	"time"
)

func main() {
	go coolCount("sirzzang")
	go coolCount("eraser")
	time.Sleep(time.Second * 5)
}

func coolCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is cool", i)
		time.Sleep(time.Second)
	}
}
