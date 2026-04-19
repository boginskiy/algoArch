package main

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

type MockResponseWriter struct{}

func (m *MockResponseWriter) Header() http.Header        { return http.Header{} }
func (m *MockResponseWriter) Write([]byte) (int, error)  { return 0, nil }
func (m *MockResponseWriter) WriteHeader(statusCode int) {}

var errorResponse SomeStruct

type SomeStruct struct {
	CntErr atomic.Int64
	CntAll atomic.Int64
	Res    http.ResponseWriter
	Err    error
}

func someHTTPRequest(ctx context.Context) (http.ResponseWriter, error) {
	// Иммитация работы
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(100 * time.Millisecond):
		// Иммитация успеха или ошибки
		if time.Now().Unix()%3 == 0 { // Каждый третить запрос будет с ошибкой
			return nil, fmt.Errorf("request failed")
		}
		return &MockResponseWriter{}, nil
	}
}

func handle(ctx context.Context, res *SomeStruct) {
	res.CntAll.Add(1)

	resExtra, err := someHTTPRequest(ctx)

	if err != nil {
		res.Err = err
		res.CntErr.Add(1)
		return
	}
	res.Res = resExtra
}

func main() {
	response := &SomeStruct{
		CntErr: atomic.Int64{},
		CntAll: atomic.Int64{},
		Res:    nil,
		Err:    nil}

	// Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Канал для сигнала завершения
	done := make(chan bool)

	// Вывод данные раз в 10 секунд
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				done <- true
				return

			case <-ticker.C:
				// Атомарно читаем значения
				cntErr := response.CntErr.Load()
				cntAll := response.CntAll.Load()

				if cntAll == 0 {
					continue
				}

				percentage := float64(cntErr) / float64(cntAll) * 100
				fmt.Printf("Error rate: %.2f\n", percentage)
			}
		}

	}(ctx)

	// Симуляция множества запросов
	for i := 0; i < 20; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			handle(ctx, response)
			time.Sleep(200 * time.Millisecond) // Задержки между запросами
		}
	}

	<-done
}
