package main

import (
	"errors"
	"log"
	"time"
)

func main() {
	input := []int{1, 2, 3, 4}

	// генератор возвращает канал, который потом передаётся получателю
	inputCh := generator(input)

	// получатель, в который передаём канал
	go consumer(inputCh)

	// добавим секунду сна, чтобы выводились сообщения ошибки
	time.Sleep(time.Second)
}

// generator генерирует данные и отправляет их в канал inputCh, который потом закрывает, потому что он отправитель
func generator(input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			inputCh <- data
		}
	}()
	return inputCh
}

// consumer читает данные из канала, отправляет в функцию, которая возвращает ошибку
func consumer(ch chan int) {
	// читаем данные из канала ch, пока он открыт
	for data := range ch {
		// получаем ошибку
		err := callDatabase(data)
		if err != nil {
			// не самый лучший способ обработки ошибок — вывод на экран
			log.Println(err)
		}
	}
}

// callDatabase функция обращения к базе данных, которая всегда возвращает ошибку
func callDatabase(data int) error {
	return errors.New("ошибка запроса к базе данных")
}
