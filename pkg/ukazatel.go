// программа с указателями, передает значение в Float64
package main

import (
	"fmt"
)

func createPointer() *float64 { // указатель на тип данных
	var myFloat = 98.5
	return &myFloat // возвращаем указатель заданного типа
}

func main() {
	var myFloatPointer *float64 = createPointer() // назначаем возвращенный указатель
	fmt.Println(*myFloatPointer)                  //вывод значения который ссылается на указатель
}
