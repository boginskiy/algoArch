package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// Задание. Исправь представленную программу, чтобы как только какая-нибудь горутина ответила
// с ошибкой, то программа завершилась

// Реализация с минимальным набором доп структур.

type ExResult struct {
	Response *http.Response
	Err      error
}

func main() {
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.yandex.ru",
		"https://www.mail.ru",
		"https://www.google.com",
	}

	resCh := make(chan *ExResult, len(urls))
	errCh := make(chan error, 1)

	// ErrGroup
	g, gCtx := errgroup.WithContext(context.Background())

	// GraceFull Shatdown
	mCtx, stop := signal.NotifyContext(gCtx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Context Timeout
	tCtx, cancel := context.WithTimeout(mCtx, 3*time.Second)
	defer cancel()

	// Wait Group
	wg := &sync.WaitGroup{}
	wg.Add(len(urls))

	for _, url := range urls {
		url := url

		g.Go(func() error {
			resultCh := make(chan *ExResult, 1)

			go func() {
				resultCh <- fetchUrl(url)
			}()

			select {
			case <-tCtx.Done():
				return tCtx.Err()
			case result, ok := <-resultCh:
				if !ok {
					return nil
				}
				if result.Err != nil {
					return result.Err
				}

				// Отправка в основной канал
				select {
				case <-tCtx.Done():
					return tCtx.Err()
				case resCh <- result:
					return nil
				}
			}
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			select {
			case <-tCtx.Done():
			case errCh <- err:
			}
			close(errCh)
		}
	}()

	// Чтение данных
	for {
		select {
		case <-tCtx.Done():
			return
		case v, ok := <-errCh:
			if !ok {
				return
			}
			fmt.Println(v)

		case v, ok := <-resCh:
			if !ok {
				return
			}

			// Если произошла ошибка в какой то горутине - делаем отмену
			if v.Err != nil {
				cancel()
				return
			}

			// Вывод результата
			fmt.Println(v.Response)
		}
	}
}

func fetchUrl(url string) *ExResult {
	res, err := http.Get(url)
	return &ExResult{Response: res, Err: err}
}
