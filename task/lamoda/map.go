package main

import (
	"fmt"
	"sync"
	"time"
)

var m = map[string]int{"a": 1}

var mu sync.RWMutex

func main() {
	go read()
	time.Sleep(1 * time.Second)

	go write()
	time.Sleep(1 * time.Second)

}

func read() {
	for {
		mu.RLock()
		fmt.Println(m["a"]) // 1. Одну секунду будет выводиться значение "1"
		mu.RUnlock()
	}
}

func write() {
	for {
		mu.Lock()
		m["a"] = 2 // В какой то момент "1" поменяется на "2"
		mu.Unlock()
	}
}

// Может быть какое то время программа поработает
// Но в какой то момент будет гонка, а вот будет ли ошибка ? Да. fatal error: concurrent map read and map write
// Программа будет работать 2 секунды.
