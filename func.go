package main

import (
	"fmt"
)

func main() {
	var a int = 3
	var b int = 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(num1 *int, num2 *int) int {
	var temp int

	temp = *num1
	*num1 = *num2
	*num2 = temp
	return temp
}
