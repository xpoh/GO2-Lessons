package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type data struct {
	counter int
}

func (d *data) increment(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-ch:
			fmt.Println(v)

			if v<1000 {
				ch <- v + 1
			} else {
				ch <-v
			}
		default:

		}
	}
}

func main() {
	chsig := make(chan os.Signal, 1)
	defer close(chsig)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	ctx, cancel = signal.NotifyContext(ctx, syscall.SIGTERM)
	defer cancel()

	var ch = make(chan int)
	defer close(ch)

	d := data{counter: 0}
	for i := 0; i < 1000; i++ {
		go d.increment(ctx, ch)
	}
	ch <- 1

	//	time.Sleep(1*time.Second)
	//	fmt.Println("Sleep done")

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
			if v == 1000 {
				cancel()
				fmt.Println("Total count", v)
				return
			}
			ch <- v

		default:
		}
	}
}