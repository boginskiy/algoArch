package main

import (
	"fmt"
	"sync"
	"time"
)

func startWorkers(c *sync.Cond, val *int) {
	workerCount := 3

	for i := 0; i < workerCount; i++ {
		go func(workerId int) {
			c.L.Lock()
			for {
				c.Wait()
				// получили сигнал
				fmt.Printf("val %v processed by worker %v\n", *val, workerId)
			}
		}(i)
	}
}

func main() {
	var m sync.Mutex
	c := sync.NewCond(&m)
	val := 0
	startWorkers(c, &val)
	// ждём, чтобы стартанули все воркеры
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 4; i++ {
		m.Lock()
		val = i
		fmt.Printf("set val to %v\n", val)
		// отправляем сигнал всем воркерам
		c.Broadcast()
		m.Unlock()
		time.Sleep(time.Millisecond)
	}
}
