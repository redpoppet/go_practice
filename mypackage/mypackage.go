package mypackage

import "fmt"

func New() {
	fmt.Println("mypakcage load")
}

func Old(x, y int) int {
	return x + y
}
