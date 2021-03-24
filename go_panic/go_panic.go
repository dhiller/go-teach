package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR", err)
		}
	}()

	// as we are using a goroutine it's still panicking even if we have the recover
	go div(1,0)

	time.Sleep(time.Millisecond)
	fmt.Println("DONE")
}

func div(a, b int) int {
	return a / b
}
