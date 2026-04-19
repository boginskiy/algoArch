package main

import (
	"context"
	"fmt"
	"time"
)

func callRPC() (string, error) {
	return externalService.Call()
}

// Решение. Горутина + select
func callRPCWithTimeOut(ctx context.Context, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	type result struct {
		data string
		err  error
	}

	ch := make(chan result, 1) // Буфер 1 - чтобы горутина не зависла.

	go func() {
		data, err := callRPC()
		ch <- result{data, err}
	}()

	select {
	case res := <-ch:
		return res.data, res.err

	case <-ctx.Done():
		return "", fmt.Errorf("time out: %w", ctx.Err())
	}
}

func main() {

}
