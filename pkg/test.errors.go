// программа теста ошибок
package main

import (
	"errors" // подключаем пакет ошибок
	"fmt"
)

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 { // если divisor равен 0
		return 0, errors.New("can't divide by 0") //возвращаем нулевое значение с сообщением
	}
	return dividend / divisor, nil // если divisor > 0 делим его на дивиденды и возвращаем
	// вместе с nil
}
func main() {
	quotient, err := divide(6.0, 0.0) //создаем 2 переменные и передаем параметры функции
	if err != nil {                   //если ерр не равен нил выводим ошибку
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient) //иначе пишем результат
	}
}
