package main

import (
	"fmt"
)

const (
	ab = 1
	bb = 2
)

func main() {

	var a = "GOlong"
	fmt.Print("var a=", a)
	var b int
	fmt.Print("var b=", b)

	var c bool
	fmt.Print("var c=", c)

	var d *int
	fmt.Println("var d=", d)

	var e []int
	fmt.Print("var e=", e)

	var f error
	fmt.Print("var f=", f)

	g := 111
	fmt.Println("var g=", g)

	h, j, k := 1, 2, "hh"
	fmt.Println("var h,j,k=", h, j, k)

	const l string = "aaa"

	fmt.Println(l)
	fmt.Println(ab, "\t", bb)
}
