package main

import (
	"fmt"
)

func resetToZero(ptr *float64) {
	*ptr = 0.0

}
func main() {
	zero := 12.2
	resetToZero(&zero)
	fmt.Println(zero)
}
