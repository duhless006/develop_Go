package main

import (
	"fmt"
	"math"
)

func floatPaint(number float64) (integerPaint int, fractionalPaint float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber

}

func main() {
	cans, remainder := floatPaint(1.26)
	fmt.Println(cans, remainder)
}
