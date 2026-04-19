package main

//Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
//Её тело нельзя изменять (представим, что внутри сетевой запрос).
//Нужно написать обертку predictableFunc, которая будет работать с заданным фиксированным
// таймаутом (например, 1 секунду).
//Нужно изменить функцию обертку, которая будет работать с заданным таймаутом (например, 1
// секунду).
//Если "длинная" функция отработала за это время - отлично, возвращаем результат.
//Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
//Сигнатуру функцию обёртки менять можно.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

func predictableFunc(ctx context.Context) (int64, error) {
	result := make(chan int64, 1)
	tresult := make(chan time.Duration, 1)

	go func() {
		start := time.Now()
		result <- unpredictableFunc()
		tresult <- time.Since(start)
		close(result)
		close(tresult)
	}()

	select {
	case <-ctx.Done():
		return 0, errors.New("Time is over!")
	case num, ok := <-result:
		if !ok {
			return 0, errors.New("Value not valid!")
		}

		select {
		case <-ctx.Done():
		case t, ok := <-tresult:
			if !ok {
				return num, errors.New("Need to check timer's work")
			}
			fmt.Println("Time Work: ", t)
		}
		return num, nil
	}
}

func main() {
	fmt.Println("started")

	// Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := predictableFunc(ctx)
	fmt.Println(result, err)
}
