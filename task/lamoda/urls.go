package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Задание. Сделать ревью кода и провести рефакторинг.

// Не забывать про:
// 1. Обработка паник в горутинах
// 2. ! Лимит на количество одновременных запросов - семафор
// 3. Более детальная обработка ошибок
// 4. ! Закрывать response body
// 5. Логирование с уровнями (debug/info/error)
// 6. Метрики и observability. Время выполнения запроса, счетчики успехов/ошибок

// Дополнительные улучшения:
// Retry logic с exponential backoff
// Circuit breaker для проблемных хостов
// Rate limiting на стороне клиента
// Tracing (например, OpenTelemetry)
// Проверка DNS до запроса

func fetchUrl(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	return res, err
}

type Response struct {
	Url        string
	Body       []byte
	Err        error
	StatusCode int
	TimeWorked time.Duration
}

func main() {
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.yandex.ru",
		"https://www.mail.ru",
		"https://www.google.com",
	}

	// GraceFull Shatdown
	mainCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Context common
	comCtx, cancel := context.WithTimeout(mainCtx, 5*time.Second)
	defer cancel()

	// WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(len(urls))

	// Channel for Data
	dataCh := make(chan Response, len(urls))

	// Semaphore
	semaphore := make(chan struct{}, 3)

	for _, url := range urls {

		go func(ctx context.Context, ch chan Response, url string) {
			semaphore <- struct{}{}

			defer func() {
				if r := recover(); r != nil {
					log.Printf("panic request: %v", r)
				}
				wg.Done()
				<-semaphore
			}()

			// Save start time
			start := time.Now()

			res, err := fetchUrl(comCtx, url)

			response := Response{
				Url:        url,
				TimeWorked: time.Since(start),
				Err:        err,
			}

			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					log.Printf("timeout: %s", url)
				} else {
					log.Printf("error: %s - %v", url, err)
				}
				ch <- response
				return
			}

			// Обязательно закрываем body
			defer res.Body.Close()
			response.StatusCode = res.StatusCode

			if res.StatusCode >= 400 {
				response.Err = fmt.Errorf("bad request status %d\n", res.StatusCode)
			} else {
				response.Body, response.Err = io.ReadAll(res.Body)
			}

			select {
			case <-ctx.Done():
				return
			case ch <- response:
			}

		}(comCtx, dataCh, url)

	}

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	for {
		select {
		case <-comCtx.Done():
			fmt.Println("TimeOut done")
			return

			// Не просто завершаем, а ждем когда горутины закончат работу.
		case <-mainCtx.Done():

			cancel()
			done := make(chan struct{})

			go func() {
				wg.Wait()
				close(done)
				fmt.Println("Shatdown done")
			}()

			select {
			case <-done:
			case <-time.After(2 * time.Second):
			}

		case response, ok := <-dataCh:
			if !ok {
				fmt.Println("Works done")
				return
			}
			fmt.Println(response)
		}
	}
}

// // Код до исправлений.
// func main() {
// 	urls := []string{
// 		"https://www.lamoda.ru",
// 		"https://www.yandex.ru",
// 		"https://www.mail.ru",
// 		"https://www.google.com",
// 	}
// 	for _, url := range urls {
// 		go func(url string) {
// 			fmt.Printf("Fetching %s...\n", url)
// 			err := fetchUrl(url)
// 			if err != nil {
// 				fmt.Printf("Error fetching %s: %v\n", url, err)
// 				return
// 			}
// 			fmt.Printf("Fetched %s\n", url)
// 		}(url)
// 	}
// 	fmt.Println("All requests ...")
// 	time.Sleep(400 * time.Millisecond)
// 	fmt.Println("Programm finished")
// }
// func fetchUrl(url string) error {
// 	_, err := http.Get(url)
// 	return err
// }
