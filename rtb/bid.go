package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), bidTimeout)
	defer cancel()

	bid := bidOn(ctx, "https://redhat.com")
	fmt.Println(bid)

	bid = bidOn(ctx, "https://ibm.com")
	fmt.Println(bid)
}

var (
	bidTimeout = 10 * time.Millisecond

	defaultBid = Bid{
		Price: 1,
		AdURL: "http://example.com/default",
	}
)

func bidOn(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // here we are using a buffered channel because we do not want to leak a goroutine

	go func() {
		ch <- findBestBid(url)
	}()

	select {
	case bid := <- ch:
		return bid
	//case <- time.After(bidTimeout):
	case <- ctx.Done():
		return defaultBid
	}
}


type Bid struct {
	Price int // in US cents
	AdURL string
}

var bidState = 0

func findBestBid(url string) Bid {
	bidState = 1 - bidState // toggle state
	if bidState == 1 {
		time.Sleep(bidTimeout / 2)
		return Bid{
			Price: 2,
			AdURL: "https://example.com/ad1",
		}
	}

	time.Sleep(bidTimeout * 2)
	return Bid{
		Price: 3,
		AdURL: "https://example.com/ad2",
	}

}
