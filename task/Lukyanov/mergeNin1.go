package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// merge N каналов в 1

func joinChannels(ctx context.Context, chs ...<-chan int) <-chan int {
	result := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(len(chs))

	semafore := make(chan struct{}, 5)

	// Чтобы не было блокировки в следствии использования семафора
	// помещаем блок кода в отдельную горутину.

	go func() {
		for _, ch := range chs {
			// Занимаем место в семафоре.
			semafore <- struct{}{}

			go func(ch <-chan int, semafore chan struct{}) {
				defer func() {
					<-semafore
					wg.Done()
				}()

				for {
					select {
					case <-ctx.Done():
						return
					case num, ok := <-ch:
						if !ok {
							return
						}

						select {
						case <-ctx.Done():
							return
						case result <- num:
						}
					}
				}
			}(ch, semafore)
		}

		go func() {
			wg.Wait()
			close(result)
		}()
	}()

	return result
}

func main() {
	// Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Создаем список каналов
	tmpCh1 := make(chan int, 100)
	tmpCh2 := make(chan int, 100)
	tmpCh3 := make(chan int, 100)

	cnt := 1

	// Пишем данные в канал
	for i, ch := range []chan int{tmpCh1, tmpCh2, tmpCh3} {
		cnt *= 10

		go func() {
			for j := 1; i < 100; j++ {
				select {
				case <-ctx.Done():
					return
				case ch <- j * cnt:
				}
			}
			close(ch)
		}()
	}

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bad work!")
		}
	}()

	result := joinChannels(ctx, tmpCh1, tmpCh2, tmpCh3)

	for num := range result {
		fmt.Println(num)
	}
}
