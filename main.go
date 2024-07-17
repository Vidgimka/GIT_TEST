package main

import "fmt"

func main() {
	a := "Работа с GIT"
	fmt.Println(a)
	b := 222
	fmt.Println(b)
	if b > 50 {
		b = b + 1
	}
	fmt.Println(b)
}
