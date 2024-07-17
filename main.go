package main

import "fmt"

func main() {

	b := 222
	fmt.Println(b)
	if b > 50 {
		b = b + 1
	}
	fmt.Println(b)
}
