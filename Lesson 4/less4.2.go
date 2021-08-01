package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			time.Sleep(1*time.Second)
			fmt.Println("Get SIGTERM signal.")
			return
		default:
		}
	}
}