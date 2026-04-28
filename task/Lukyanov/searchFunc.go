package main

import (
	"context"
	"fmt"
	"time"
)

// Нужно реализовать функцию, которая выполняет
// поиск query во всех переданных SearchFunc
// Когда получаем первый успешный результат -
// отдаем его сразу. Если все SearchFunc отработали
// с ошибкой - отдаем последнюю полученную ошибку

type Result struct{}

type SearchFunc func(ctx context.Context, query string) (Result, error)

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
	// errs := make([]error, 0, len(sfs))
	// result := make(chan Result)

	resultCh := make(chan Result, 1)
	errCh := make(chan error, len(sfs))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, sf := range sfs {
		go func(sf SearchFunc) {
			// Если контекст уже отменен - выходим
			if ctx.Err() != nil {
				return
			}

			res, err := sf(ctx, query)
			if err != nil {
				select {
				case errCh <- err:
				case <-ctx.Done():
				}
				return
			}

			// Первый успешный результат отправляем в resultCh и делаем cancel()
			select {
			case <-ctx.Done():
				return
			case resultCh <- res:

				// Закрываем контекст → остальные горутины
				cancel()
			}

		}(sf)
	}

	var lastErr error

	for i := 0; i < len(sfs); i++ {
		select {
		case <-ctx.Done():
			return Result{}, ctx.Err()
		case res := <-resultCh:
			return res, nil
		case err := <-errCh:
			lastErr = err
		}
	}

	return Result{}, lastErr
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ff := func(context.Context, string) (Result, error) { return Result{}, nil }

	res, err := MultiSearch(ctx, "qwery", []SearchFunc{ff})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}
