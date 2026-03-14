package main

import (
	"fmt"
	"sync"
	"time"
)

// Oграничить количество одновременно работающих горутин, например, до 20.
// Это означает, что запросы будут обрабатываться "пакетами": 20 запросов
// выполняются одновременно, остальные ждут своей очереди.

func main() {
	fmt.Println("Отправляем 100 HTTP-запросов с лимитом 20 одновременных")

	// Семафор - ограничиваем до 20 одновременных запросов
	semaphore := make(chan struct{}, 20)
	var wg sync.WaitGroup

	// Запускаем 100 запросов
	for i := 1; i <= 100; i++ {
		semaphore <- struct{}{} // Ждем свободный слот
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			defer func() { <-semaphore }() // Освобождаем слот

			sendRequest(id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Все запросы завершены!")
}

func sendRequest(id int) {
	fmt.Printf("Запрос %d: начался\n", id)
	time.Sleep(2 * time.Second) // Имитация HTTP-запроса
	fmt.Printf("Запрос %d: завершен\n", id)
}
