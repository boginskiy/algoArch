package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Иммитация разных серверов с разной скоростью
var Servers = []struct {
	name  string
	delay time.Duration
}{
	{"server1", 1 * time.Second},
	{"server2", 2 * time.Second},
	{"server3", 500 * time.Millisecond},
}

func fastSearch(ctx context.Context, query string, results chan<- string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	resultCh := make(chan string, len(Servers))
	wg := &sync.WaitGroup{}
	wg.Add(len(Servers))

	go func() {
		// Запускаем запросы ко всем серверам
		for _, s := range Servers {
			go func(srv struct {
				name  string
				delay time.Duration
			}) {
				defer wg.Done()

				select {
				case <-time.After(srv.delay):
					resultCh <- fmt.Sprintf("%s: %s", srv.name, query)
				case <-ctx.Done():
					return
				}
			}(s)
		}
		wg.Wait()
		close(resultCh)
	}()

	// Берем первый успешный ответ
	select {
	case result, ok := <-resultCh:
		if !ok {
			return
		}

		select {
		case <-ctx.Done():
		case results <- result:
		}

	case <-ctx.Done():
		return
	}

}

func main() {
	ctx := context.Background()
	results := make(chan string)

	go fastSearch(ctx, "golang", results)
	fmt.Println("Первый ответ: ", <-results)
}
