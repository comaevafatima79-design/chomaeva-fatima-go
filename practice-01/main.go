package main

import (
	"fmt"
	"runtime"
)

func main() {
	student := "Чомаева Фатима"
	group := "ПрИ-251"
	platform := runtime.GOOS
	if platform == "windows" {
		platform = "Windows"
	}
	year := "2026-2027"
	practice := "1"

	sep := "================================"

	fmt.Println(sep)
	fmt.Println("ЯЗЫК ПРОГРАММИРОВАНИЯ GO")
	fmt.Println(sep)
	fmt.Println("Студент: " + student)
	fmt.Println("Группа: " + group)
	fmt.Println("Платформа: " + platform)
	fmt.Println("Учебный год: " + year)
	fmt.Println("Практическая работа: " + practice)
	fmt.Println(sep)
	fmt.Println("Среда готова к работе")
}
