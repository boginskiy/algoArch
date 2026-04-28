package main

// Ограниченные параллельные HTTP-запросы

// Условие:
// Нужно реализовать функцию, которая выполняет HTTP-запросы к переданным URL с ограничением на количество одновременно выполняемых запросов.

// Требования:
// 1. Не более K активных горутин одновременно
// 2. Возвращать все успешные ответы (в случае ошибок - пропускать)
// 3. Сохранять порядок ответов не обязательно

// Входные данные:
// - urls - список URL для запросов
// - K - максимальное количество параллельных запросов

// Выходные данные:
// - Слайс указателей на http.Response

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	// Для начала делаем GraceFull Shatdown
	mCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Context
	ctx, cancel := context.WithTimeout(mCtx, 5*time.Second)
	defer cancel()

	// Данные
	urls := []string{"https://example.com", "https://example.org", "https://example.net"}

	// Бизнес логика
	fmt.Println(callRequestsForURLs(ctx, urls, 3))

	select {
	case <-mCtx.Done():
		fmt.Println("Do Shatdown")
	case <-ctx.Done():
		fmt.Println("Do timeOut")
	}

}

// Выполнить N URL-запросов с лимитом K параллельных горутин
func callRequestsForURLs(ctx context.Context, urls []string, K int) []*http.Response {
	semaphore := make(chan struct{}, K)
	newResCh := make(chan *http.Response, len(urls))
	result := make([]*http.Response, 0, len(urls))

	wg := &sync.WaitGroup{}
	wg.Add(len(urls))

	for _, url := range urls {

		semaphore <- struct{}{}

		go func(semaphore chan struct{}, url string) {
			defer func(wg *sync.WaitGroup) {
				<-semaphore
				wg.Done()

				if r := recover(); r != nil { // Recover
					log.Printf("exception recover: %s", "callRequestsForURLs")
				}
			}(wg)

			newReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				log.Printf("error create request: %v", err)
				return
			}

			res, err := http.DefaultClient.Do(newReq)

			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					log.Println("timeOut send request")
					return
				}
				log.Printf("error send request: %v", err)
				return
			}

			defer res.Body.Close()

			select {
			case <-ctx.Done():
				log.Println("timeOut send request")
				return
			case newResCh <- res:
			}

		}(semaphore, url)
	}

	go func() {
		wg.Wait()
		close(newResCh)
	}()

	for res := range newResCh {
		result = append(result, res)
	}

	return result
}
