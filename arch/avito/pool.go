package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ГЕНЕРАТОР.
	in := generator(ctx, 100)

	// Pool
	now := time.Now()
	for i := range Pool(in, 3, timeConsuming2) {
		fmt.Println(i)
	}

	timePool := time.Since(now)
	fmt.Println(timePool)
}

func Pool(in chan int, numWorkers int, ffunc func(int) int) chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)

	go func() {
		for range numWorkers {
			go worker(wg, in, out, ffunc)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func worker(wg *sync.WaitGroup, in, out chan int, ffunc func(int) int) {
	defer wg.Done()
	for i := range in {
		out <- ffunc(i)
	}
}

// I/O : сетевой вызов, чтение файла.
func timeConsuming2(num int) int {
	time.Sleep(100 * time.Millisecond)
	return num * num
}

// Generator.
func generator(ctx context.Context, num int) chan int {
	gen := make(chan int)

	go func() {
		for i := range num {
			select {
			case <-ctx.Done():
				return
			case gen <- i:
			}
		}
		close(gen)
	}()

	return gen
}
