package main

import (
	"fmt"
	"sync"
	"time"
)

const NumWorkers = 3 // Количество ворекров в пуле.
const MaxTasks = 10  // Сколько задач обработаем.

// Worker
func worker(id int, taskCh <-chan string, resultCh chan<- string) {
	for task := range taskCh {
		// Выполнение задачи.
		time.Sleep(1 * time.Second)

		// Result
		result := fmt.Sprintf("task: %v  id: %d\n\r", task, id)
		resultCh <- result
	}
}

func main() {
	taskCh := make(chan string, MaxTasks)   // Канал для передачи задач
	resultCh := make(chan string, MaxTasks) // Канал для возвращения результатов
	var wg sync.WaitGroup                   // Контроллер завершения всех воркеров

	// Пул воркеров.
	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			worker(id, taskCh, resultCh)
		}(i)
	}

	// Задачи.
	for i := 0; i < MaxTasks; i++ {
		taskCh <- fmt.Sprintf("t-№%d", i)
	}
	close(taskCh)

	// Результаты.
	for i := 0; i < MaxTasks; i++ {
		fmt.Println(<-resultCh)
	}

	wg.Wait()
	close(resultCh)

}
