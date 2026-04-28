package main

import "fmt"

// Задача: Что выведет код?
// Результат: Функция foo(N) всегда будет возвращать значение != nil

// Почему так ?
// err != nil → true, потому что тип интерфейса не nil (он содержит информацию о типе *MyError)

type MyError struct {
	data string
}

func (e MyError) Error() string {
	return e.data
}

func main() {
	err := foo(4)
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok")
	}
}

func foo2(i int) error {
	var err error

	if i > 5 {
		err = &MyError{data: "i > 5"}

	}
	return err
}

func foo(i int) error {
	var err *MyError

	if i > 5 {
		err = &MyError{data: "i > 5"}
		// return &MyError{data: "i > 5"}
	}

	return err
	// return nil
}
