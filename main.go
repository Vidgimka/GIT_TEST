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
	Text.WriteString("Первая строка\n в новой ветке")
}
