package main

import (
	"fmt"
	"time"
)

func test(a string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(a + " good!")
}

func main() {
	go test("dog", 5)
	go test("cat", 2)
	fmt.Println("Now is ready")
	time.Sleep(time.Duration(10) * time.Second)
}
