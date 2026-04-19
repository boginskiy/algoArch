// merge N каналов 2

package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 1. Merge n channels
// 2. Если один из входных каналов закрывается,
// то нужно закрыть все остальные каналы

// Аномалия в работе. Непонятно в чем дело.

func case3(ctx context.Context, channels ...chan int) chan int {
	resCh := make(chan int)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {

		go func(ch chan int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						// Если канал закрыт, останавливаем горутины
						cancel()
						return
					}

					select {
					case <-ctx.Done():
						return
					case resCh <- v:
					}
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	return resCh
}

func main() {
	// GraceFull Shatdown
	mCtx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM)

	defer stop()

	// Context
	ctx, cancel := context.WithTimeout(mCtx, 3*time.Second)
	defer cancel()

	ch1 := make(chan int, 1000)
	ch2 := make(chan int, 1000)

	for i := 1; i <= 1000; i++ {
		ch1 <- i
		ch2 <- 1000 + i
	}

	close(ch1)
	close(ch2)

	for res := range case3(ctx, ch1, ch2) {
		fmt.Println(res)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Context worked")
	case <-mCtx.Done():
		fmt.Println("Shatdown worked")
	}

	fmt.Println("Ok stop")
}
