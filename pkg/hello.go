package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	ficed := replacer.Replace(broken)
	fmt.Println(ficed)
}
