package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	ch chan struct{}
}

// New Semaphore.
func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, limit),
	}
}

// Acquire - захват ресурсаю.
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

// Release - освобождение ресурса.
func (s *Semaphore) Release() {
	<-s.ch
}

// Worker имитирует выполнение некоторой длительной операции
func worker(id int, semaphore *Semaphore) {
	defer semaphore.Release()

	semaphore.Acquire()

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second) // Имитация длительной операции
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	maxWorker := 3 // Лимит параллельного исполнения.
	numJobs := 10  // Количество заданий.

	semaphore := NewSemaphore(maxWorker)

	var wg sync.WaitGroup // Синхронизируем завершение всех задач

	wg.Add(numJobs)

	for i := 0; i < numJobs; i++ {
		go func(id int) {
			defer wg.Done()
			worker(id+1, semaphore)
		}(i)
	}

	wg.Wait()
}
