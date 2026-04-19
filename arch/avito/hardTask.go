package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Задача:
// реализовать функцию 'processParallel' и прокинуть 'context'

type outVal struct {
	val int
	err error
}

var errTimeOut = errors.New("time over")

func processData(ctx context.Context, v int) chan outVal {
	chRes := make(chan outVal)
	chTime := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(chTime)
	}()

	go func() {
		select {
		case <-ctx.Done():
			chRes <- outVal{err: errTimeOut}
		case <-chTime:
			chRes <- outVal{val: v * 2}
		}
	}()

	return chRes
}

func main() {
	in := make(chan int)
	out := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		defer close(in)

		for i := range 10 {
			select {
			case <-ctx.Done():
				return
			case in <- i:
			}
		}
	}()

	start := time.Now()

	processParallel(ctx, in, out, 5)

	for v := range out {
		fmt.Println("v =", v)
	}

	fmt.Println("main duration: ", time.Since(start))
}

func processParallel(ctx context.Context, in, out chan int, numWorkers int) {
	wg := &sync.WaitGroup{}

	// Запускаем воркеры в количестве 'numWorkers'
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(ctx, wg, in, out)
	}

	// Во избежании блокировок помещаем в отдельную горутину
	go func() {
		wg.Wait()
		close(out)
	}()
}

func worker(ctx context.Context, wg *sync.WaitGroup, in, out chan int) {
	defer wg.Done()

	for {
		select {

		case num, ok := <-in:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			default:
				res := <-processData(ctx, num)
				if res.err != nil {
					return
				}
				out <- res.val
			}

		case <-ctx.Done():
			return
		}
	}
}
