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

	// ОБРАБОТКА.
	res := fanOut(ctx, in, 3, timeConsuming)
	// timeConsuming   -> 1.806154385s
	// timeConsuming2 - > 3.410208012s

	// ЧТЕНИЕ.
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		for _, ch := range res {
			ch := ch
			go func() {
				defer wg.Done()
				for num := range ch {
					fmt.Println(num)
				}
			}()
		}
	}()

	// fanIn
	now := time.Now()
	wg.Wait()

	timefanIn := time.Since(now)
	fmt.Println(timefanIn)

}

func fanOut(ctx context.Context, in chan int, w int, ffunc func(int) int) []chan int {
	res := make([]chan int, w)

	wg := &sync.WaitGroup{}
	wg.Add(w)

	for i := range res {
		res[i] = pipline(ctx, wg, in, ffunc)
	}

	go func() {
		wg.Wait()
	}()

	return res
}

func pipline(ctx context.Context, wg *sync.WaitGroup, in chan int, ffunc func(int) int) chan int {
	out := make(chan int)

	go func() {
		defer wg.Done()
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					return
				case out <- ffunc(val):
				}
			}
		}
	}()

	return out
}

// CPU intensive
func timeConsuming(num int) int {
	cnt := 0
	for range 100000000 {
		cnt++
	}
	return num * num
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
