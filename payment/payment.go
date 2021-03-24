package main

import (
	"fmt"
	"sync"
	"time"
)


type Payment struct {
	From string
	To string
	Amount float64 // EUR
	once sync.Once
}

func (p *Payment) Pay() {
	time := time.Now()
	p.once.Do(func() {
		p.pay(time)
	})
}

func (p *Payment) pay(time time.Time) {
	fmt.Printf("%v: %s -> [EUR %.2f] -> %s\n", time, p.From, p.Amount, p.To)
}

func main() {
	payment := Payment{
		From:   "Wile E. Coyote",
		To:     "ACME",
		Amount: 127.3,
	}
	payment.Pay()
	payment.Pay()
}
