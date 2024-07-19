package main

import "fmt"

func main() {

	b := 222
	fmt.Println(b)
	if b > 50 {
		b = b + 1
	}
	fmt.Println(b)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
