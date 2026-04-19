package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// 1. Иногда приходят нули. В чем проблема? Исправь ее
// 2. Если функция bank_network_call выполняется 5 секунд,
// то за сколько выполнится balance()? Как исправить проблему? // 5 секунд, можно добавить контекст с таймаутом
// 3. Представим, что bank_network_call возвращает ошибку дополнительно.
// Если хотя бы один вызов завершился с ошибкой, то balance должен вернуть ошибку.

func balance(ctx context.Context) (int, error) {
	x := make(map[int]int, 1)
	var m sync.Mutex

	errCh := make(chan error, 1)

	wg := &sync.WaitGroup{}
	wg.Add(5)

	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		for i := 0; i < 5; i++ {
			i := i
			go func() {
				defer wg.Done()

				if ctx.Err() != nil {
					return
				}

				b, err := bank_network_call(ctx, i)
				if err != nil {
					select {
					case <-ctx.Done():
					case errCh <- err:
						cancel()

					}
					return
				}

				m.Lock()
				x[i] = b
				m.Unlock()
			}()
		}
	}()

	wg.Wait()
	close(errCh)

	select {
	case <-ctx.Done():
		return 0, errors.New("Time out")
	case err, ok := <-errCh:
		if ok {
			return 0, err
		}
	}
	// Как-то считается сумма значений в мапе и возвращается
	sumOfMap := 19
	return sumOfMap, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := balance(ctx)
	fmt.Println(res, err)
}
