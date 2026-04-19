package main

import (
	"fmt"
	"sync"
)

/*
	Broadcaster - структура, которая публикует сообщения для множества подписчиков.
	TODO... работает некорреткно. Проверить!

*/

type Broadcaster struct {
	ch          chan interface{}
	subscribers []chan<- interface{} // Канал типа send-only
	mu          sync.Mutex           // Мутация списка подписчиков
	wg          sync.WaitGroup       // Ожидание закрытия всех подписков
}

func NewBroadcaster(bufferSize int) *Broadcaster {
	return &Broadcaster{
		ch:          make(chan interface{}, bufferSize),
		subscribers: nil,
	}
}

// Подписываемся на получение событий
func (b *Broadcaster) Subscribe() <-chan interface{} {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan interface{})
	b.subscribers = append(b.subscribers, ch)
	b.wg.Add(1) // добавляем одну ожидаемую горутину

	go func() {
		defer func() {
			close(ch)
			b.Unsubscribe(ch) // удаляем подписку при выходе
			b.wg.Done()       // сообщаем, что горутина завершилась
		}()

		for {
			v, ok := <-b.ch
			if !ok {
				break
			}
			ch <- v
		}
	}()

	return ch
}

// Отменяем подписку
func (b *Broadcaster) Unsubscribe(ch chan<- interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for i, sub := range b.subscribers {
		if sub == ch {
			copy(b.subscribers[i:], b.subscribers[i+1:])
			b.subscribers[len(b.subscribers)-1] = nil
			b.subscribers = b.subscribers[:len(b.subscribers)-1]
			break
		}
	}
}

// Публикуем событие для всех подписчиков
func (b *Broadcaster) Broadcast(v interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, sub := range b.subscribers {
		select {
		case sub <- v:
		default:
			fmt.Println("Subscriber channel is full")
		}
	}
}

// Ждём завершения всех подписки
func (b *Broadcaster) WaitForAllSubscribers() {
	b.wg.Wait()
}

func main() {
	broadcaster := NewBroadcaster(10)

	// Два получателя подписываются на канал
	ch1 := broadcaster.Subscribe()
	ch2 := broadcaster.Subscribe()

	// Читаем полученные сообщения
	go func() {
		for msg := range ch1 {
			fmt.Printf("Ch1 получил %v\n", msg)
		}
	}()

	go func() {
		for msg := range ch2 {
			fmt.Printf("Ch2 получил %v\n", msg)
		}
	}()

	// Посылаем несколько сообщений
	for i := 0; i < 1; i++ {
		broadcaster.Broadcast(i)
	}

	// Закрываем основной канал, сигнализируя конец передачи сообщений
	close(broadcaster.ch)

	// Дожидаемся завершения получения сообщений всеми подписчиками
	broadcaster.WaitForAllSubscribers()

}
