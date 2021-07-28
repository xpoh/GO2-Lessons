package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type data struct {
	counter int
}

func (d *data) increment(ctx context.Context, ch chan int, chsig chan os.Signal) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-ch:
			if ok {
				ch <- v + 1
				ctx.Done()
			}
			fmt.Println(v)
		case v, _ := <-chsig:
			if v == syscall.SIGTERM {
				ctx.Err()
			}
		default:

		}
	}
}

func main() {
	chsig := make(chan os.Signal, 1)
	signal.Notify(chsig, syscall.SIGTERM)

	ctx := context.Background()
	cancelctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var ch = make(chan int)
	defer close(ch)

	d := data{counter: 0}
	for i := 0; i < 1000; i++ {
		go d.increment(cancelctx, ch, chsig)
	}
	ch <- 1
	//v, _ := <-ch
	fmt.Println(d.counter)
}
