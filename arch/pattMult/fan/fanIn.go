package main

import (
	"context"
	"sync"
)

func fanIn(ctx context.Context, chans []chan int) chan int {
	res := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(chans))

		for _, ch := range chans {
			ch := ch
			go func() {
				defer wg.Done()

				for {

					select {
					case <-ctx.Done():
						return
					case v, ok := <-ch:
						if !ok {
							return
						}

						select {
						case <-ctx.Done():
							return
						case res <- v:
						}
					}
				}
			}()
		}
		wg.Wait()
		close(res)
	}()

	return res
}
