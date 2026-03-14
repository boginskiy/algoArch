package main

import (
	"fmt"
	"sync"
	"time"
)

// Semaphore структура семафора
type Semaphore struct {
	semaCh chan struct{}
}

// NewSemaphore создает семафор с буферизованным каналом емкостью maxReq
func NewSemaphore(maxReq int) *Semaphore {
	return &Semaphore{
		semaCh: make(chan struct{}, maxReq),
	}
}

// когда горутина запускается, отправляем пустую структуру в канал semaCh
func (s *Semaphore) Acquire() {
	s.semaCh <- struct{}{}
}

// когда горутина завершается, из канала semaCh убирается пустая структура
func (s *Semaphore) Release() {
	<-s.semaCh
}

func main() {
	// Чтобы дождаться всех горутин
	var wg sync.WaitGroup

	// Создаем семафор емкостью 2: он будет пропускать только 2 горутины
	// для работы с общим ресурсом.
	semaphore := NewSemaphore(2)

	// Создаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)

		// Горутина в которую помещаем ее порядковый номер
		go func(taskID int) {
			// Отправляем в канал семафора пустую структуру, тем самым мы заняли место
			semaphore.Acquire()

			//
			defer wg.Done()

			// Забираем из канала семафора пустую структуру, освобождаем место
			defer semaphore.Release()

			fmt.Printf("Запущен рабочий %d", taskID)

			// Полезная работа, например обращения к БД
			time.Sleep(1 * time.Second)
		}(i)
	}
	wg.Wait()
}
