package main

import (
	"fmt"
	"time"
)

// Задача:  Что выведет данный код?
// Вывод:   panic: send on closed channel

// Корректировка:
// Использование канала с буфером ch := make(chan int, 1)
// или
// Закрытие канала в горутине, которая пишет в канал close(ch)

func main() {
	ch := make(chan int)

	go func() {
		// Операция записи блокируется, пока другая горутина не будет готова читать.
		// Когда другая горутина готова читать, текущая горутина начинает писать
		// но мы уже закрыли канал.
		ch <- 1
	}()

	time.Sleep(time.Millisecond * 500)
	close(ch) // ОШИБКА! Нельзя закрыть канал, если есть заблокированный писатель

	for i := range ch {
		fmt.Println(i)
	}

	time.Sleep(time.Millisecond * 100)
}
