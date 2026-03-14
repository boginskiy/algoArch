package service

import (
	"context"
	"goCnter/internal/model"
	"log"
	"runtime"
)

type TaskService struct {
	TaskCh chan *model.Task
}

func NewTaskService(ctx context.Context) *TaskService {
	tmp := &TaskService{
		TaskCh: make(chan *model.Task, 10),
	}

	tmp.Consumer(ctx)
}

func (s *TaskService) Consumer(ctx context.Context) {
	numWorkers := runtime.NumCPU()

	for i := 0; i < numWorkers; i++ {
		go s.worker(i)
	}
}

func (s *TaskService) worker(workerID int) {
	for task := range s.TaskCh {

		// Запись в базу данных
		err := db.InsertOrUpdateCounter(task.ID, task.Count)

		if err != nil {
			log.Printf("Worker #%d failed to update DB: %v", workerID, err)
		} else {
			log.Printf("Worker #%d updated counter for %s with value %d", workerID, task.ID, task.Count)
		}
	}
}
