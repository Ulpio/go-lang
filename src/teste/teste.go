package main

import "fmt"

var C float64 = 100.0
var K float64 = 373.2

func main() {
	C = K - 273
	fmt.Printf("O valor é: %g", C)
	fmt.Println("")
}
