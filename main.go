package main

import (
	"fmt"
	"os"
)

func main() {
	a := "Работа с GIT"
	fmt.Println(a)
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

	Text2, err := os.Create("Text_File2")
	if err != nil {
		fmt.Println("Файл не удалось создать")
		os.Exit(1)
	}
	defer Text2.Close()
	Text2.WriteString("Первая строка2")
}
