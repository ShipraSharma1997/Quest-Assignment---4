package main

import (
	"fmt"
)

func main() {

	var x int
	var y int

	var temp int
	fmt.Printf("Enter the value of x and y\n")
	fmt.Scanf("%d%d", &x, &y)

	fmt.Printf("Before Swapping\nx = %d\ny = %d\n", x, y)

	a := &x
	b := &y

	temp = *b
	*b = *a
	*a = temp

	fmt.Printf("After Swapping\nx = %d\ny = %d\n", x, y)

	return
}

//To check Escape Ananlysis run below command
//go build -gcflags '-m -m  -l' main.go
