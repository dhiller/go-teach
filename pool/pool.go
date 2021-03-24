package main

import (
	"log"
	"runtime"
	"sync"
)

func cpuHungry(n float64) {
	val := n / 2.3
	log.Printf("%f -> %f", n, val)
}

func poolWorker(ch <-chan float64, wg *sync.WaitGroup) {
	for val := range ch {
		cpuHungry(val)
		wg.Done()
	}
}

// the waitgroup here is signaling when all the jobs are done so we cam safely close the channel

func multiWork(nums []float64) {
	ch := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for i := 0; i < runtime.NumCPU(); i++ {
		go poolWorker(ch, &wg)
	}

	for _, i := range nums {
		ch <- i
	}

	wg.Wait()
	close(ch)
}

func main() {
	multiWork([]float64{1.3, 2.7, 3.6, 4.9})
}
