package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

// without anything this is not thread safe
// a) use a sync.Mutex when incrementing the counter
// b) (better) use atomics

// Note: option b does not seem to be threadsafe, when running with -race it bails

var (
	counter int64
	//m sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	//m.Lock()
	//defer m.Unlock()
	atomic.AddInt64(&counter, 1)
	fmt.Fprintf(w,"count=%d\n", counter)
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
