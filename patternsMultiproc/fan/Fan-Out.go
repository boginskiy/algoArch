// Паттерн Fan-Out позволяет увеличить количество рабочих на этапах с наибольшей нагрузкой.

package main

import "time"

// generator функция из предыдущего примера, делает то же, что и делала
func generator(doneCh chan struct{}, input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			case <-doneCh:
				return
			case inputCh <- data:
			}
		}
	}()

	return inputCh
}

// multiply функция из предыдущего примера, делает то же, что и делала
func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	multiplyRes := make(chan int)

	go func() {
		defer close(multiplyRes)

		for data := range inputCh {
			result := data * 2

			select {
			case <-doneCh:
				return
			case multiplyRes <- result:
			}
		}
	}()
	return multiplyRes
}

// add функция из предыдущего примера, делает то же, что и делала
func add(doneCh chan struct{}, inputCh chan int) chan int {
	addRes := make(chan int)

	go func() {
		defer close(addRes)

		for data := range inputCh {
			// замедлим вычисление, как будто функция add требует больше вычислительных ресурсов
			time.Sleep(time.Second)

			result := data + 1

			select {
			case <-doneCh:
				return
			case addRes <- result:
			}
		}
	}()
	return addRes
}

// fanOut принимает канал данных, порождает 10 горутин
func fanOut(doneCh chan struct{}, inputCh chan int) []chan int {
	// количество горутин add
	numWorkers := 10
	// каналы, в которые отправляются результаты
	channels := make([]chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		// получаем канал из горутины add
		addResultCh := add(doneCh, inputCh)
		// отправляем его в слайс каналов
		channels[i] = addResultCh
	}

	// возвращаем слайс каналов
	return channels
}
