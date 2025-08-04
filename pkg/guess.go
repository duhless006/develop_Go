// guess - игра, в которой игрок должен угадать случайное число.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	i := 0
	for i < 10 {
		fmt.Println("Ваше число: (осталось попыток", 10-i, "):")
		var input string
		fmt.Scanln(&input)
		i++
		hisDigit := rand.Intn(100) + 1
		input = strings.TrimSpace((input))
		answer, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}
		if answer == float64(hisDigit) {
			fmt.Println("Поздравляем, вы угадали!")
			break
		}
		if answer < float64(hisDigit) {
			fmt.Println("LOW", "Рандомное число: ", hisDigit)
		} else {
			fmt.Println("HIGT", "Рандомное число: ", hisDigit)
		}
		if i == 10 {
			fmt.Println("Sorry. You didn't guess my number.")
		}
	}

}
