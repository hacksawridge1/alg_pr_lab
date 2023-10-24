package main

import (
	"fmt"
)

func main() {
	var y string = "Да"
	var n string = "Нет"
	var a, b, c string

	fmt.Scan(&a, &b, &c)

	if a == n && b == n && c == n {
		fmt.Println("Кот")
	} else if a == n && b == n && c == y {
		fmt.Println("Жираф")
	} else if a == n && b == y && c == n {
		fmt.Println("Курица")
	} else if a == n && b == y && c == y {
		fmt.Println("Страус")
	} else if a == y && b == n && c == n {
		fmt.Println("Дельфин")
	} else if a == y && b == n && c == y {
		fmt.Println("Плезиозавры")
	} else if a == y && b == y && c == n {
		fmt.Println("Пингвин")
	} else if a == y && b == y && c == y {
		fmt.Println("Утка")
	}
}
