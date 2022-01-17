package main

import (
	"fmt"
)

//For int
func main() {
	var a *int
	b := 4
	a = &b
	fmt.Printf("%v", *a)
}

//for string
// func main() {
// 	var a *string
// 	b := "4"
// 	a = &b
// 	fmt.Printf("%s", *a)
// }

//for float
// func main() {
// 	var a *float64
// 	b := 4.00
// 	a = &b
// 	fmt.Printf("%f", *a)
// }
