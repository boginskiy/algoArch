package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	// Исходящие запросы по http с таймаутом.
	FetchData("hhp://www.google.com")

	// Проверка выставления таймаута в существующем таймауте.
	CheckOfTimeOut()

	// Клиент получил ответ и закрыл соединение, а мы хотим статистику собрать.
	// Handler(w, r)

	// After func
	// Worker(ctx)

}

func Worker(ctx context.Context) {
	conn := openConnection()

	// Регистрируем callback - горутина создается только при отмене
	// Если контекст не отменен, то и context.AfterFunc не сработает
	// Сработает только когда ctx.Done()

	stop := context.AfterFunc(ctx, func() {
		conn.Close()
		log.Println("Connection closed")
	})
	defer stop()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Отправляем ответ
	w.Write([]byte("OK"))

	// Клиент получил ответ -> закрыл соединение (это норма для HTTP)
	// Go видит:соединение закрыто -> отменяет r.Context()

	// Решение.Добавление обертки контекста без отмены.
	ctx := context.WithoutCancel(r.Context())

	go func() {
		// Контекст тут может быть отменен!
		// Если sendMetrics проверяет ctx.Done() - операция прервется.

		// // Было: r.Context()
		// sendMetrics(r.Context())

		// Стало: ctx
		sendMetrics(ctx)
	}()
}

func sendMetrics(ctx context.Context) {
	// TODO some logic.
}

func CheckOfTimeOut() {
	// Нет timeout.
	ctx1 := context.Background()
	deadline, ok := ctx1.Deadline()
	fmt.Printf("ctx1: Remains of time: %v. Is there timeout: %v", deadline, ok)

	// Есть timeout.
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	deadline, ok = ctx2.Deadline()
	fmt.Printf("ctx2: Remains of time: %v. Is there timeout: %v", deadline, ok)

}

func FetchData(url string) ([]byte, error) {
	//Таймаут 5 секунд - если не успели, то abort
	ctx, cancel := context.WithTimeout((context.Background()), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, fmt.Errorf("Abort: таймаут превысил %d секунд", 5)
		}
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
