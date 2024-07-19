package main

import (
	"fmt"
	"os"
)

func main() {

	b := 222
	c := 333
	fmt.Println(b + c)

	Text, err := os.Create("Text_File")
	if err != nil {
		fmt.Println("Файл не удалось создать")
		os.Exit(1)
	}
	defer Text.Close()
	Text.WriteString("Первая строка")

	fmt.Println(b)
	if b > 50 {
		b = b + 1
	}
	fmt.Println(b)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
