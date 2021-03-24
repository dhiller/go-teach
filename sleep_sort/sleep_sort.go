package main

import (
	"fmt"
	"time"
)

func sleepSort(nums []int) []int {
	c := make(chan int)
	defer close(c)
	for _, i := range nums {
		go func(n int) {
			time.Sleep(time.Duration(n)*10*time.Millisecond)
			c <- n
		}(i)
	}
	result := []int{}
	for range nums {
		r := <- c
		result = append(result, r)
	}
	return result
}

func main() {
	nums := []int{15,23,8,4,42,16}
	fmt.Println(sleepSort(nums))
}


