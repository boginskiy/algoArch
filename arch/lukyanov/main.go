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
// то за сколько выполнится balance()? Как исправить проблему?
// 3. Представим, что bank_network_call возвращает ошибку дополнительно.
// Если хотя бы один вызов завершился с ошибкой, то balance должен вернуть ошибку.

func bank_network_call(num int) (int, error) {
	time.Sleep(1 * time.Second)

	if time.Now().UnixNano()%10 == 0 {
		return num * num, errors.New("error")
	}
	return num * num, nil
}

func balance(live int) (int, error) {
	x := make(map[int]int, 5)

	var (
		mu        sync.Mutex
		wg        sync.WaitGroup
		sOnce     sync.Once
		errReturn error
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(5)

	for i := 0; i < 5; i++ {
		i := i

		go func() {
			defer wg.Done()

			val, err := bank_network_call(i)
			if err != nil {
				sOnce.Do(func() {
					errReturn = err
					cancel()
				})
				return
			}

			select {
			case <-ctx.Done():
				cancel()
				return
			default:
				mu.Lock()
				x[i] = val
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	if errReturn != nil {
		return 0, errReturn
	}

	return len(x), nil
}

func main() {
	res, err := balance(4)
	fmt.Println(res, err)
}
