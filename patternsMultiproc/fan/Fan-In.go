// Паттерн Fan-In объединяет несколько результатов в один канал. Этот процесс также называют мультиплексированием

package main

import (
	"fmt"
	"sync"
)

// fanIn объединяет несколько каналов resultChs в один.
func fanIn(doneCh chan struct{}, resultChs ...chan int) chan int {
	// конечный выходной канал в который отправляем данные из всех каналов из слайса, назовём его результирующим
	finalCh := make(chan int)

	// понадобится для ожидания всех горутин
	var wg sync.WaitGroup

	// перебираем все входящие каналы
	for _, ch := range resultChs {
		// в горутину передавать переменную цикла нельзя, поэтому делаем так
		chClosure := ch

		// инкрементируем счётчик горутин, которые нужно подождать
		wg.Add(1)

		go func() {
			// откладываем сообщение о том, что горутина завершилась
			defer wg.Done()

			// получаем данные из канала
			for data := range chClosure {
				select {
				// выходим из горутины, если канал закрылся
				case <-doneCh:
					return
				// если не закрылся, отправляем данные в конечный выходной канал
				case finalCh <- data:
				}
			}
		}()
	}

	go func() {
		// ждём завершения всех горутин
		wg.Wait()
		// когда все горутины завершились, закрываем результирующий канал
		close(finalCh)
	}()

	// возвращаем результирующий канал
	return finalCh
}

func main() {
	// слайс данных
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// сигнальный канал для завершения горутин
	doneCh := make(chan struct{})
	// закрываем его при завершении программы
	defer close(doneCh)

	// канал с данными
	inputCh := generator(doneCh, input)

	// получаем слайс каналов из 10 рабочих add
	channels := fanOut(doneCh, inputCh)

	// а теперь объединяем десять каналов в один
	addResultCh := fanIn(doneCh, channels...)

	// передаём тот один канал в следующий этап обработки
	resultCh := multiply(doneCh, addResultCh)

	// выводим результаты расчетов из канала
	for res := range resultCh {
		fmt.Println(res)
	}
}
