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
	area := wigth * heigth
	return area / 10.0, nil
}

func main() {
	amound, err := paintNeeded(4.2, -3.1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amound)
	}
}
