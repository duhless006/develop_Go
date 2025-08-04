package main

import (
	"fmt"
	"log"
)

func paintNeeded(wigth float64, heigth float64) (float64, error) {
	if wigth < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", wigth)
	}
	if heigth < 0 {
		return 0, fmt.Errorf("a heigth of %0.2f is invalid", heigth)
	}
	reverse := wigth * heigth
	return reverse / 10.0, nil

}

func main() {
	var amor, total float64
	amor, err := paintNeeded(4.2, 5.0)
	if err != nil {
		log.Fatal(err)
	}
	amor, err = paintNeeded(4.2, 3.0)
	total += amor
	fmt.Printf("%0.2f first sten\n", amor)
	amor, err = paintNeeded(5.2, 2.5)
	total += amor
	fmt.Printf("%0.2f two sten\n", amor)
	fmt.Printf("%0.2f all litres", total)
}
