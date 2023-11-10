package main

import (
	"fmt"
	"reflect"
)

func main() {
	// one variable in one line
	var age int
	var height float32
	var open_eyes string = "Hello World"

	// multiple variable in one line
	var a, b, c int
	var d, e, f = 5, "Good Morning", 0.10

	// short declaration one variable in one line
	g := "Good Afternoon"
	h := "Good Evening"
	i := 0.15

	// short declaration multiple variable in one line
	j, k, l := 101, "Good Night", 1.01

	// different data types
	m, n, o, p := "Hello again", 101, 1.01, false

	fmt.Println(age, height, open_eyes)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Println(g, h, i)
	fmt.Println(j, k, l)
	fmt.Println(m, reflect.TypeOf(m), n, reflect.TypeOf(n), o, reflect.TypeOf(o), p, reflect.TypeOf(p))

}
