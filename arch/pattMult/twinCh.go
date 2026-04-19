package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func twinChan(ctx context.Context, in chan int) (chan int, chan int) {
	resCh1 := make(chan int, 100)
	resCh2 := make(chan int, 100)

	go func() {
		defer close(resCh1)
		defer close(resCh2)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					return
				case resCh1 <- v:
				}

				select {
				case <-ctx.Done():
					return
				case resCh2 <- v:
				}
			}
		}
	}()
	return resCh1, resCh2
}

func main() {
	in := make(chan int, 100)

	for i := 1; i <= 100; i++ {
		in <- i
	}
	close(in)

	// Graceful Shatdown
	mCtx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// Context
	ctx, cancel := context.WithTimeout(mCtx, 3*time.Second)
	defer cancel()

	ch1, ch2 := twinChan(ctx, in)

	// Чтение из каналов.
	go func() {
		for {
			select {
			case <-ctx.Done():
			case v, ok := <-ch1:
				if !ok {
					return
				}
				fmt.Println(v)
			}

			select {
			case <-ctx.Done():
			case v, ok := <-ch2:
				if !ok {
					return
				}
				fmt.Println(v)
			}
		}
	}()

	select {
	case <-mCtx.Done():
		fmt.Println("Shatdown...")
	case <-ctx.Done():
		fmt.Println("TimeOut...")
	}
}
