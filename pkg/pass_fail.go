package main

import (
	"bufio"   // создает буфер в памяти обьемом 4096 байт
	"fmt"     // форматирование ввод\вывод
	"log"     // работа с ошибками
	"os"      // работа с ос доступ к файлам
	"strconv" // преобразование строк
	"strings" //операции со строками
)

func main() {
	fmt.Print("Enter a grade: ")          // запрашиваем у пользователя значение
	reader := bufio.NewReader(os.Stdin)   // буферизированный считыватель
	input, err := reader.ReadString('\n') // получаем данные от пользователя и завершаем чтение после нажатия enter
	if err != nil {                       //если ошибка не равна nil
		log.Fatal(err) // сохраняем ее в переменную err
	}
	input = strings.TrimSpace((input))          // удаляем пробельные символы, пропуски
	grade, err := strconv.ParseFloat(input, 64) // преобразуем из строки в формат float64
	if err != nil {                             // проверка если не равно nil
		log.Fatal(err) // сохраняем ее в переменную
	}
	var status string // обьявляем переменную статус внутри блока для видимости
	if grade >= 60 {  // если значение от пользователя < или = 60
		status = "passing" //статус получает наименование
	} else { //иначе
		status = "failing" // статус получает наименование элса
	}
	fmt.Println("A grade of", grade, "is", status) // вывод полученных результатов
}
