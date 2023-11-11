package main

import (
	"fmt"
)

func flowControl() {
	// condition and operators
	p, q := 100, 101
	if p%2 == 1 {
		fmt.Println(p, " is odd")
	} else {
		fmt.Println(p, " is even")
	}
	if p%2 == 0 && q%2 == 0 {
		fmt.Println(p, " ", q, " both are even")
	}
	if p%2 == 0 || q%2 == 0 {
		fmt.Println("One of them is even")
	}
	if p%2 == 0 {
		fmt.Println(p, " is devisible by 2")
	} else if p%3 == 0 {
		fmt.Println(p, " is devisible by 3")
	}

	//for loop
	arr := []int{50, 100, 300, 450, 525}
	for i := 0; i < 5; i++ {
		fmt.Println(i, " is index", arr[i], " is value")
	}

}
